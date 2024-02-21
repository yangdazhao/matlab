package mclmcr

import (
	"fmt"
	"git.zuihuayin.work/ben/log"
	"golang.org/x/sys/windows"
	"golang.org/x/sys/windows/registry"
	"os"
	"strings"
	"syscall"
)

type MclFunction int

//go:generate stringer -type MclFunction -linecomment
const (
	mclmcrInitialize         MclFunction = iota // mclmcrInitialize_proxy
	mclInitializeApplication                    // mclInitializeApplication_800_proxy
	mclTerminateApplication                     // mclTerminateApplication_proxy
	mclGetLastErrorMessage                      // mclGetLastErrorMessage_proxy
	mclGetLogFileName                           // mclGetLogFileName_proxy
)

const (
	MATLAB          = "MATLAB"
	Runtime         = "MATLAB Runtime"
	MathWorks       = `SOFTWARE\MathWorks`
	MathWorksFormat = `SOFTWARE\MathWorks\%s`
)

func GetMatLabRootByRegistry(name string) (string, string, error) {
	openKey, err := registry.OpenKey(registry.LOCAL_MACHINE, fmt.Sprintf(MathWorksFormat, name), registry.ALL_ACCESS)
	if err != nil {
		log.Println(err)
		return `nil`, "", err
	}
	RuntimeNames, err := openKey.ReadSubKeyNames(2)
	if err != nil {
		log.Println(err)
		return `nil`, "", err
	}
	log.Println(RuntimeNames)
	for _, runtimeName := range RuntimeNames {
		log.Println(runtimeName)
		runtimeKey, err := registry.OpenKey(registry.LOCAL_MACHINE, fmt.Sprintf(MathWorksFormat, name)+`\`+runtimeName, registry.ALL_ACCESS)
		if err != nil {
			log.Println(err)
			break
		}
		MATLABROOT, _, err := runtimeKey.GetStringValue("MATLABROOT")
		if err != nil {
			log.Println(err)
			break
		}

		MATLABROOT1 := MATLABROOT + `\VersionInfo.xml`

		VersionInfo, err := os.ReadFile(MATLABROOT1)
		if err != nil {
			log.Println(err)
			continue
		}

		log.Debug(string(VersionInfo))
		return MATLABROOT, runtimeName, nil
	}
	return `RuntimeNames`, "", nil
}

func MatlabRoot() (string, string, error) {
	dir, name, err := GetMatLabRootByRegistry(Runtime)
	if err != nil {
		return "", "", err
	}
	//log.Println(dir, name)
	return dir, name, nil
}

var lazyDLL *syscall.LazyDLL

func init() {
	root, version, err := MatlabRoot()
	if err != nil {
		log.Error(err)
		return
	}
	//log.Info(root, version)
	mclmcrrtDllFile := fmt.Sprintf(`%s\runtime\win64\mclmcrrt%s.dll`, root, strings.ReplaceAll(version, `.`, `_`))
	log.Info(mclmcrrtDllFile)
	lazyDLL = syscall.NewLazyDLL(mclmcrrtDllFile)
	if err := lazyDLL.Load(); err != nil {
		log.Println(err)
	}
}

func Initialize() error {
	if r1, _, err := lazyDLL.NewProc(mclmcrInitialize.String()).Call(); r1 != 0 {
		//if r1, r2, err := lazyDLL.NewProc(mclmcrInitialize.String()).Call(uintptr(5)); CheckError(err) && r1 != 0 {
		//log.Error(err)
		return err
	} else {
		//log.Debugf("%d %d %s", r1, r2, err)
		return nil
	}
}

func InitializeApplication() error {
	if r1, _, err := lazyDLL.NewProc(mclInitializeApplication.String()).Call(uintptr(0), uintptr(0)); r1 != 0 && windows.DS_S_SUCCESS != err {
		log.Println(r1, err)
		return err
	} else {
		//log.Debugf("%d %d %s", r1, r2, err)
		return nil
	}
}

func TerminateApplication() error {
	if r1, _, err := lazyDLL.NewProc(mclTerminateApplication.String()).Call(); r1 != 0 && windows.DS_S_SUCCESS != err {
		//windows.DS_S_SUCCESS
		log.Println(err)
		return err
	} else {
		//log.Debugf("%d %d %s", r1, r2, err)
		return nil
	}
}

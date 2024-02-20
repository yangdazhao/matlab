package mclmcr

import (
	"log"
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

var lazyDLL *syscall.LazyDLL

func init() {
	lazyDLL = syscall.NewLazyDLL(`MCLMCRRT9_13.DLL`)
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
	if r1, _, err := lazyDLL.NewProc(mclInitializeApplication.String()).Call(uintptr(0), uintptr(0)); r1 != 0 {
		log.Println(err)
		return err
	} else {
		//log.Debugf("%d %d %s", r1, r2, err)
		return nil
	}
}

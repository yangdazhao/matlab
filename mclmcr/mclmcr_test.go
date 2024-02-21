package mclmcr

import (
	"fmt"
	"golang.org/x/sys/windows/registry"
	"log"
	"testing"
)

func TestInitializeApplication(t *testing.T) {
	key, err := registry.OpenKey(registry.LOCAL_MACHINE, MathWorks, registry.ALL_ACCESS)
	if err != nil {
		t.Log(err)
		return
	}

	names, err := key.ReadSubKeyNames(2)
	if err != nil {
		t.Log(err)
		return
	}
	for i, name := range names {
		t.Log(i, name)
		switch name {
		case Runtime:
			log.Println(name)
			if RuntimeNames, _, err := GetMatLabRootByRegistry(name); err == nil {
				return
			} else {
				log.Println(RuntimeNames)
			}

		case MATLAB:
			log.Println(name)
			openKey, err := registry.OpenKey(registry.LOCAL_MACHINE, fmt.Sprintf(MathWorksFormat, name), registry.ALL_ACCESS)
			if err != nil {
				log.Println(err)
				break
			}
			RuntimeNames, err := openKey.ReadSubKeyNames(1)
			if err != nil {
				t.Log(err)
				return
			}
			stat, err := openKey.Stat()
			if err != nil {
				return
			}
			log.Println(stat.SubKeyCount)
			log.Println(RuntimeNames)
		}
	}

	//		if err := InitializeApplication(); (err != nil) != tt.wantErr {
	//			t.Errorf("InitializeApplication() error = %v, wantErr %v", err, tt.wantErr)
	//		}
	//	})
	//}
}

func Test_funcName(t *testing.T) {
	GetMatLabRootByRegistry(Runtime)

}

func TestMatlabRoot(t *testing.T) {
	t.Log(MatlabRoot())
}

func TestInitializeApplication1(t *testing.T) {
	defer TerminateApplication()
	InitializeApplication()
}

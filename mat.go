package goMatlab

import "C"

/*
#cgo windows CPPFLAGS: -I"C:/Program Files/MATLAB/MATLAB Runtime/R2022b/extern/include"
#cgo windows LDFLAGS: -L"C:/Program Files/MATLAB/MATLAB Runtime/R2022b/extern/lib/win64/mingw64" -lmat -lmclmcrrt
#include <mat.h>
*/
import "C"
import (
	"errors"
	"unsafe"
)

type MatMode int

const (
	r MatMode = iota
	u
	w
	w4
	w6
	wL
	w7
	wz
	w7_3
)

type Mat struct {
	MATFile *C.MATFile
}

func (m *Mat) Open(filename, mode string) error {
	_file := C.CString(filename)
	_mode := C.CString(mode)
	m.MATFile = C.matOpen(_file, _mode)
	if nil == m.MATFile {
		return m.GetErrno()
	}

	return nil
}

func (m *Mat) Close() error {
	matError := C.matClose(m.MATFile)
	if matNoError == MatError(matError) {
		return nil
	}
	return errors.New(MatError(matError).String())
}

func (m *Mat) GetDir() []string {
	var num C.int
	cDirs := C.matGetDir(m.MATFile, &num)
	var goStringArray []string

	for i := C.int(0); i < num; i++ {
		// 获取char*
		//cStr := *(*C.char)(unsafe.Pointer(uintptr(unsafe.Pointer(cDirs)) + uintptr(len(goStringArray)*int(unsafe.Sizeof(uintptr(0))))))
		cStr := (*C.char)(unsafe.Pointer(uintptr(unsafe.Pointer(cDirs)) + uintptr(len(goStringArray)*int(unsafe.Sizeof(uintptr(0))))))
		//if cStr == nil {
		//	break // 遇到nil结束
		//}

		// 将C的char*转换为Go的string
		goStringArray = append(goStringArray, C.GoString((*C.char)(unsafe.Pointer(uintptr(unsafe.Pointer(cStr))+1))))
	}
	return goStringArray
}

func (m *Mat) GetErrno() error {
	err := C.matGetErrno(m.MATFile)
	if matNoError == MatError(err) {
		return nil
	}
	return errors.New(MatError(err).String())
}

func (m *Mat) GetVariable(name string) MxArray {
	_name := C.CString(name)
	var mx MxArray
	mx.mxArray = C.matGetVariable(m.MATFile, _name)
	return mx
}

func (m *Mat) GetVariableInfo(name string) MxArray {
	_name := C.CString(name)
	var mx MxArray
	mx.mxArray = C.matGetVariableInfo(m.MATFile, _name)
	return mx
}

func (m *Mat) GetNextVariable(name string) MxArray {
	_name := C.CString(name)
	var mx MxArray
	mx.mxArray = C.matGetNextVariable(m.MATFile, &_name)
	return mx
}

func (m *Mat) GetNextVariableInfo(name string) MxArray {
	_name := C.CString(name)
	var mx MxArray
	mx.mxArray = C.matGetNextVariableInfo(m.MATFile, &_name)
	return mx
}

func (m *Mat) DeleteVariable(name string) error {
	_name := C.CString(name)
	matError := C.matDeleteVariable(m.MATFile, _name)
	if matNoError == MatError(matError) {
		return nil
	}
	return errors.New(MatError(matError).String())
}

//func (m *Mat) DestroyArray(name string, mx MxArray) int {
//	_name := C.CString(name)
//	err := C.matGetVariable(m.MATFile, _name, mx.mxArray)
//	return int(err)
//}

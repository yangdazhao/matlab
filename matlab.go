package goMatlab

/*
#cgo windows CPPFLAGS: -I"C:/Program Files/MATLAB/MATLAB Runtime/R2022b/extern/include"
#cgo windows LDFLAGS: -L"C:/Program Files/MATLAB/MATLAB Runtime/R2022b/extern/lib/win64/mingw64" -lmclbase -lmclmcrrt
#include <mclmcrrt.h>
#include <mclcppclass.h>
*/
import "C"

// MclmcrInitialize Initialize the MATLAB Runtime proxy library
// mclmcrInitialize is called by mclInitializeApplication. Therefore, in most cases, you should not explicitly call this function in your application code.
func mclmcrInitialize() {
	C.mclmcrInitialize()
}

// MclInitializeApplication Set up application state shared by all MATLAB Runtime instances created in current process
// MclInitializeApplication must be called once only per process.
func MclInitializeApplication() bool {
	return bool(C.mclInitializeApplication(nil, 0))
}

// MclTerminateApplication Close MATLAB Runtime-internal application state
func MclTerminateApplication() bool {
	return bool(C.mclTerminateApplication())
}

// MclWaitForFiguresToDie Enable deployed applications to process graphics events so that figure windows remain displayed
func MclWaitForFiguresToDie() {

}

//func mclRunMain() string {
//	Msg := C.mclGetLastErrorMessage()
//	return C.GoString(Msg)
//}

func MclGetLastErrorMessage() string {
	Msg := C.mclGetLastErrorMessage()
	return C.GoString(Msg)
}

// MclGetLogFileName Retrieve name of log file used by MATLAB Runtime
func MclGetLogFileName() string {
	Msg := C.mclGetLogFileName()
	return C.GoString(Msg)
}

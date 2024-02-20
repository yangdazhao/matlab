package matlab

/*
#cgo windows CPPFLAGS: -I"C:/Program Files/MATLAB/MATLAB Runtime/R2022b/extern/include"
#cgo windows LDFLAGS: -L"C:/Program Files/MATLAB/MATLAB Runtime/R2022b/extern/lib/win64/mingw64" -lmat -lmclmcrrt
#include <matrix.h>
*/
import "C"
import "unsafe"

type MxArray struct {
	mxArray *C.mxArray
}

func (a MxArray) ToString() {

}

func (a MxArray) IsNumeric() bool {
	return bool(C.mxIsNumeric(a.mxArray))
}

func (a MxArray) IsComplex() bool {
	return bool(C.mxIsComplex(a.mxArray))
}

func (a MxArray) IsEmpty() bool {
	return bool(C.mxIsEmpty(a.mxArray))
}

func (a MxArray) IsFromGlobalWS() bool {
	return bool(C.mxIsFromGlobalWS(a.mxArray))
}

func (a MxArray) DestroyArray() {
	C.mxDestroyArray(a.mxArray)
}

func (a MxArray) GetNumberOfDimensions() int {
	mwSize := C.mxGetNumberOfDimensions(a.mxArray)
	return int(mwSize)
}

func (a MxArray) GetNumberOfElements() int {
	return int(C.mxGetNumberOfElements(a.mxArray))
}

func (a MxArray) GetElementSize() int {
	return int(C.mxGetElementSize(a.mxArray))
}

func (a MxArray) GetM() int {
	return int(C.mxGetM(a.mxArray))
}

func (a MxArray) SetM(n int) {
	C.mxSetM(a.mxArray, C.size_t(n))
}

func (a MxArray) GetN() int {
	return int(C.mxGetN(a.mxArray))
}

func (a MxArray) SetN(n int) {
	C.mxSetN(a.mxArray, C.size_t(n))
}

func (a MxArray) GetPr() unsafe.Pointer {
	return unsafe.Pointer(C.mxGetPr(a.mxArray))
}

package matlab

/*
#cgo windows CPPFLAGS: -I"C:/Program Files/MATLAB/MATLAB Runtime/R2022b/extern/include"
#cgo windows LDFLAGS: -L"C:/Program Files/MATLAB/MATLAB Runtime/R2022b/extern/lib/win64/mingw64" -leng
#include <engine.h>
*/
import "C"

type Engine struct {
	engine *C.Engine
}

func (e *Engine) Open(startcmd string) {
	_cmd := C.CString(startcmd)
	e.engine = C.engOpen(_cmd)
}

// EvalString engEvalString evaluates the expression contained in string for the MATLAB® engine session
// Status, returned as int. The function returns 1 if the engine session is no longer running or the engine pointer is invalid or NULL.
// Otherwise, returns 0 even if the MATLAB engine session cannot evaluate the command.
func (e *Engine) EvalString(cmd string) {
	_cmd := C.CString(cmd)
	C.engEvalString(e.engine, _cmd)
}

func (e *Engine) engOpenSingleUse(startcmd string) {
	_cmd := C.CString(startcmd)
	e.engine = C.engOpen(_cmd)
}

func (e *Engine) Close() {
	C.engClose(e.engine)
}

// GetVariable Pointer to a newly allocated mxArray structure, returned as mxArray *. Returns NULL if the attempt fails. engGetVariable fails if the named variable does not exist.
func (e *Engine) GetVariable(name string) MxArray {
	_name := C.CString(name)
	var mx MxArray
	mx.mxArray = C.engGetVariable(e.engine, _name)
	return mx
}

func (e *Engine) PutVariable(name string, pm MxArray) {
	_name := C.CString(name)
	C.engPutVariable(e.engine, _name, pm.mxArray)
}

package v8engine

// #include <stdlib.h>
// #include "v8engine.h"
import "C"

import (
	"runtime"
	"unsafe"
)

// Value represents a JavaScript value
type Value struct {
	ptr C.ValuePtr
}

// String returns the string representation of the value
func (v *Value) String() string {
	s := C.ValueToString(v.ptr)
	defer C.free(unsafe.Pointer(s))

	return C.GoString(s)
}

func (v *Value) finalizer() {
	C.DisposeValue(v.ptr)
	v.ptr = nil
	runtime.SetFinalizer(v, nil)
}
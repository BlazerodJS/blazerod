package v8engine

// #include "v8engine.h"
import "C"

// Version returns the version of the V8 engine
func Version() string {
	return C.GoString(C.Version())
}

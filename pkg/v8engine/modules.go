package v8engine

// #include <stdlib.h>
import "C"

import (
	"sync"
)

var (
	resolverTableLock sync.Mutex
	nextResolverToken int
	resolverFuncs     = make(map[int]ModuleResolverCallback)
)

// ModuleResolverCallback is a callback function type used to resolve modules
type ModuleResolverCallback func(moduleName, referrerName string) (string, int)

// ResolveModule resolves module requests to source contents
//export ResolveModule
func ResolveModule(moduleSpecifier *C.char, referrerSpecifier *C.char, resolverToken int) (*C.char, C.int) {
	moduleName := C.GoString(moduleSpecifier)
	referrerName := C.GoString(referrerSpecifier)

	resolverTableLock.Lock()
	resolve := resolverFuncs[resolverToken]
	resolverTableLock.Unlock()

	if resolve == nil {
		return nil, C.int(1)
	}
	canon, ret := resolve(moduleName, referrerName)

	return C.CString(canon), C.int(ret)
}

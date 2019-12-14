package v8engine

// #cgo CXXFLAGS: -fno-rtti -fpic -Ideps/include -std=c++11
// #cgo LDFLAGS: -pthread -lv8
// #cgo darwin LDFLAGS: -Ldeps/darwin-x86_64
// #cgo linux LDFLAGS: -Ldeps/linux-x86_64
import "C"

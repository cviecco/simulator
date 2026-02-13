package ossl

/*
 This file is just to try to force the use fo cgo within the package
 code coming primarily from: https://go.dev/blog/cgo
*/

// #include <stdio.h>
// #include <stdlib.h>
import "C"
import "unsafe"

func Print(s string) {
	cs := C.CString(s)
	defer C.free(unsafe.Pointer(cs))
	C.fputs(cs, (*C.FILE)(C.stdout))
}

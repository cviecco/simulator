package src

/*
 This file is just to try to force the use fo cgo within the package
 code coming primarily from: https://go.dev/blog/cgo
*/

// #cgo CFLAGS: -I ../
//
// // Silence known warnings from the reference code and CGO code.
// #cgo CFLAGS: -Wno-missing-braces -Wno-empty-body -Wno-unused-variable -Wno-uninitialized
// // Silence openssl deprecation warnings for ms-tpm-20-ref
// #cgo CFLAGS: -Wno-deprecated-declarations
//
// #cgo CFLAGS: -DCERTIFYX509_DEBUG=NO
//
// #include <stdio.h>
// #include <stdlib.h>
import "C"
import "unsafe"

func Print(s string) {
	cs := C.CString(s)
	defer C.free(unsafe.Pointer(cs))
	C.fputs(cs, (*C.FILE)(C.stdout))
}

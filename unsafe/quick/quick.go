// Copyright 2015 The GOMACRO Authors. All rights reserved.
// Use of this source code is governed by a GPLv2-style
// license that can be found in the LICENSE file.

package quick

import (
	"reflect"
	"unsafe"

	quick32 "github.com/gomacro/sort/32/quick"
	quick64 "github.com/gomacro/sort/64/quick"
	quick8 "github.com/gomacro/sort/8/quick"
)

////////////////////////////////////////////////////////////////////////////////
func elemsize(slice interface{}) uintptr {
	return uintptr(reflect.TypeOf(slice).Elem().Size())
}
func mvetype(dst, src *interface{}) {
	*(*uintptr)(unsafe.Pointer(dst)) = *(*uintptr)(unsafe.Pointer(src))
}
func arg8(fun interface{}) (dst func(*uint8, *uint8) int) {
	var ction interface{}
	ction = dst
	mvetype(&fun, &ction)
	return fun.(func(*uint8, *uint8) int)
}
func arg32(fun interface{}) (dst func(*uint32, *uint32) int) {
	var ction interface{}
	ction = dst
	mvetype(&fun, &ction)
	return fun.(func(*uint32, *uint32) int)
}
func arg64(fun interface{}) (dst func(*uint64, *uint64) int) {
	var ction interface{}
	ction = dst
	mvetype(&fun, &ction)
	return fun.(func(*uint64, *uint64) int)
}
func u8(slice interface{}, size uintptr) (src []uint8) {
	var dst interface{}
	dst = src
	mvetype(&slice, &dst)
	src = slice.([]uint8)
	h := (*reflect.SliceHeader)(unsafe.Pointer(&src))
	h.Len *= int(size)
	h.Cap *= int(size)
	return src
}
func u32(slice interface{}, size uintptr) (src []uint32) {
	var dst interface{}
	dst = src
	mvetype(&slice, &dst)
	src = slice.([]uint32)
	h := (*reflect.SliceHeader)(unsafe.Pointer(&src))
	h.Len *= int(size)
	h.Cap *= int(size)
	return src
}
func u64(slice interface{}, size uintptr) (src []uint64) {
	var dst interface{}
	dst = src
	mvetype(&slice, &dst)
	src = slice.([]uint64)
	h := (*reflect.SliceHeader)(unsafe.Pointer(&src))
	h.Len *= int(size)
	h.Cap *= int(size)
	return src
}

////////////////////////////////////////////////////////////////////////////////

func Sort(compar interface{}, s interface{}) {
	size := elemsize(s) //8,4,1

	if (size & 7) == 0 { // use 8 (64bit)
		var m = [1]uintptr{size / 8}
		quick64.Sort(&m, arg64(compar), u64(s, m[0]))
		return
	}

	if (size & 3) == 0 { // use 4 (32bit)
		var m = [1]uintptr{size / 4}
		quick32.Sort(&m, arg32(compar), u32(s, m[0]))
		return
	}

	// use 1 (8bit)
	var m = [1]uintptr{size}
	quick8.Sort(&m, arg8(compar), u8(s, m[0]))
}

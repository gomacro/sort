// Copyright 2009 The Go Authors. All rights reserved.

// Package quick is a 64-bit macro function, callable from a third-party macros.
package quick

import (
	insertion64 "github.com/gomacro/sort/64/insertion"
)

type Quick struct {
	heapDepth int
}

func Sort(ts0 *[1]uintptr, compar func(*uint64, *uint64) int, s []uint64) {
	q := Make(ts0, compar, s)
	q.Sort(ts0, compar, s)
}

func Make(ts0 *[1]uintptr, compar func(*uint64, *uint64) int, s []uint64) (q Quick) {
	incr := int((*ts0)[0])

	for i := (len(s) / incr); i > 0; i >>= 1 {
		q.heapDepth++
	}
	q.heapDepth *= 2

	return q
}

// Sort macro .... SAFE!
func (q Quick) Sort(ts0 *[1]uintptr, compar func(*uint64, *uint64) int, s []uint64) {
	incr := int((*ts0)[0])

	for len(s) > 7*incr {
		if q.heapDepth <= 0 {

			insertion64.Sort(ts0, compar, s)

			// 				heapSort(data, a, b)
			return
		}
		q.heapDepth--
		mlo, mhi := doPivot(ts0, compar, s)
		// Avoiding recursion on the larger subproblem guarantees
		// a stack depth of at most lg(b-a).
		if mlo < (len(s)/incr)-mhi {
			q.Sort(ts0, compar, s[:mlo*incr])
			s = s[mhi*incr:]

		} else {
			q.Sort(ts0, compar, s[mhi*incr:])

			s = s[:mlo*incr]
		}
	}
	if len(s) > incr {
		insertion64.Sort(ts0, compar, s)
	}
}

// Copyright 2009 The Go Authors. All rights reserved.

// Package insertion is a 64-bit macro function, callable from a third-party macros.
package insertion

func Sort(ts0 *[1]uintptr, compar func(*uint64, *uint64) int, s []uint64) {
	incr := int((*ts0)[0])

	for i := len(s) - incr; i > 0; i -= incr {
		j := i
		sjj := s[j-incr:]
		sj := s[j:]
		for j < len(s) && compar(&sj[0], &sjj[0]) < 0 {

			for q := 0; q < incr; q++ {
				x := sjj[q]
				sjj[q] = sj[q]
				sj[q] = x
			}

			j += incr
			sjj = sj
			sj = sj[incr:]
		}
	}

}

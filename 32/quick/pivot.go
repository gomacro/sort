package quick

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func medianOfThree(ts0 *[1]uintptr, data []uint32, a, b, c int, compar func(*uint32, *uint32) int) {
	incr := int((*ts0)[0])

	m0 := b
	m1 := a
	m2 := c

	s0 := data[incr*m0:]
	s1 := data[incr*m1:]
	s2 := data[incr*m2:]

	// bubble sort on 3 elements
	if compar(&s1[0], &s0[0]) < 0 {
		for q := 0; q < incr; q++ {
			x := s0[q]
			s0[q] = s1[q]
			s1[q] = x
		}
	}
	if compar(&s2[0], &s1[0]) < 0 {
		for q := 0; q < incr; q++ {
			x := s2[q]
			s2[q] = s1[q]
			s1[q] = x
		}
	}
	if compar(&s1[0], &s0[0]) < 0 {
		for q := 0; q < incr; q++ {
			x := s0[q]
			s0[q] = s1[q]
			s1[q] = x
		}
	}
	// now data[m0] <= data[m1] <= data[m2]
}

func swapRange(ts0 *[1]uintptr, data []uint32, a, b, n int) {
	incr := int((*ts0)[0])

	sl := data[incr*a:]
	sr := data[incr*b:]

	l := incr * n

	for q := 0; q < l; q++ {
		x := sl[q]
		sl[q] = sr[q]
		sr[q] = x
	}
}

func doPivot(ts0 *[1]uintptr, data []uint32, compar func(*uint32, *uint32) int) (midlo, midhi int) {
	incr := int((*ts0)[0])

	lo := 0
	hi := len(data) / incr

	m := lo + (hi-lo)/2 // Written like this to avoid integer overflow.

	if hi-lo > 40 {
		// Tukey's ``Ninther,'' median of three medians of three.
		s := (hi - lo) / 8
		medianOfThree(ts0, data, lo, lo+s, lo+2*s, compar)
		medianOfThree(ts0, data, m, m-s, m+s, compar)
		medianOfThree(ts0, data, hi-1, hi-1-s, hi-1-2*s, compar)
	}

	medianOfThree(ts0, data, lo, m, hi-1, compar)

	// Invariants are:
	//	data[lo] = pivot (set up by ChoosePivot)
	//	data[lo <= i < a] = pivot
	//	data[a <= i < b] < pivot
	//	data[b <= i < c] is unexamined
	//	data[c <= i < d] > pivot
	//	data[d <= i < hi] = pivot
	//
	// Once b meets c, can swap the "= pivot" sections
	// into the middle of the slice.

	pivot := lo
	a, b, c, d := lo+1, lo+1, hi, hi

	pidata := &data[incr*pivot]

	for {
		for b < c {
			comp := compar(&data[incr*b], pidata)

			if comp < 0 { // data[b] < pivot
				b++
			} else if comp == 0 { // data[b] = pivot

				for q := 0; q < incr; q++ {
					x := data[a*incr+q]
					data[a*incr+q] = data[b*incr+q]
					data[b*incr+q] = x
				}

				a++
				b++

			} else {
				break
			}
		}
		for b < c {
			comp := compar(pidata, &data[incr*(c-1)])

			if comp < 0 { // data[c-1] > pivot
				c--
			} else if comp == 0 { // data[c-1] = pivot

				c--
				d--
				for q := 0; q < incr; q++ {
					x := data[c*incr+q]
					data[c*incr+q] = data[d*incr+q]
					data[d*incr+q] = x
				}

			} else {
				break
			}
		}
		if b >= c {
			break
		}
		// data[b] > pivot; data[c-1] < pivot
		c--

		for q := 0; q < incr; q++ {
			x := data[c*incr+q]
			data[c*incr+q] = data[b*incr+q]
			data[b*incr+q] = x
		}

		b++

	}

	n := min(b-a, a-lo)
	swapRange(ts0, data, lo, b-n, n)

	n = min(hi-d, d-c)
	swapRange(ts0, data, c, hi-n, n)

	return lo + b - a, hi - (d - c)
}

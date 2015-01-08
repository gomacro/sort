package quick

import (
	insertion64 "example.com/repo.git/sort/64/insertion"
)

type Quick struct {
	heapDepth int
}

func Sort(ts0 *[1]uintptr, s []uint64, compar func(*uint64, *uint64) int) {
	q := Make(ts0, s, compar)
	q.Sort(ts0, s, compar)
}

func Make(ts0 *[1]uintptr, s []uint64, compar func(*uint64, *uint64) int) (q Quick) {
	incr := int((*ts0)[0])

	for i := (len(s) / incr); i > 0; i >>= 1 {
		q.heapDepth++
	}
	q.heapDepth *= 2

	return q
}

// Sort macro .... SAFE!
func (q Quick) Sort(ts0 *[1]uintptr, s []uint64, compar func(*uint64, *uint64) int) {
	incr := int((*ts0)[0])

	for len(s) > 7*incr {
		if q.heapDepth <= 0 {

			insertion64.Sort(ts0, s, compar)

			// 				heapSort(data, a, b)
			return
		}
		q.heapDepth--
		mlo, mhi := doPivot(ts0, s, compar)
		// Avoiding recursion on the larger subproblem guarantees
		// a stack depth of at most lg(b-a).
		if mlo < (len(s)/incr)-mhi {
			q.Sort(ts0, s[:mlo*incr], compar)
			s = s[mhi*incr:]

		} else {
			q.Sort(ts0, s[mhi*incr:], compar)

			s = s[:mlo*incr]
		}
	}
	if len(s) > incr {
		insertion64.Sort(ts0, s, compar)
	}
}

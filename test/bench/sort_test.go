// Copyright 2015 The GOMACRO Authors. All rights reserved.
// Use of this source code is governed by a GPLv2-style
// license that can be found in the LICENSE file.

package sort_test

import (
	"fmt"
	"github.com/gomacro/compare"
	"github.com/gomacro/sort/unsafe/quick"
	"sort"
	"testing"
)

////////////////////////////////////////////////////////////////////////////////

type int64slice []int64

func (p int64slice) Len() int           { return len(p) }
func (p int64slice) Less(i, j int) bool { return p[i] < p[j] }
func (p int64slice) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }

////////////////////////////////////////////////////////////////////////////////

func random(prng *[2]uint64) uint64 {
	s1 := prng[0]
	s0 := prng[1]
	prng[0] = s0
	s1 ^= s1 << 23 // a
	prng[1] = (s1 ^ s0 ^ (s1 >> 17) ^ (s0 >> 26))
	return prng[1] + s0 // b, c
}

func fillu(data []int64, seed *[2]uint64) {
	for i := 0; i < len(data); i++ {
		data[i] = int64(random(seed))
	}
}

const debug = true
const shortdatasize = 100
const datasize = 10000000

func BenchmarkSortLarge_Random(b *testing.B) {
	b.StopTimer()

	fmt.Printf("")

	var seed = [2]uint64{0x13371337, 0x1337beef}
	n := datasize
	if testing.Short() {
		n /= shortdatasize
	}
	data := make([]int64, n)
	fillu(data, &seed)
	if sort.IsSorted(int64slice(data)) {
		b.Fatalf("terrible rand.rand")
	}

	for n := 0; n < b.N; n++ {

		b.StartTimer()
		sort.Sort(int64slice(data))
		b.StopTimer()

		if debug {

			if !sort.IsSorted(int64slice(data)) {
				b.Errorf("sort didn't sort - 1M ints")
			}

		}

		fillu(data, &seed)

	}
}

func BenchmarkMySortLarge_Random(b *testing.B) {
	b.StopTimer()

	fmt.Printf("")

	var seed = [2]uint64{0x13371337, 0x1337beef}
	n := datasize
	if testing.Short() {
		n /= shortdatasize
	}
	data := make([]int64, n)
	fillu(data, &seed)
	if sort.IsSorted(int64slice(data)) {
		b.Fatalf("terrible rand.rand")
	}

	for n := 0; n < b.N; n++ {

		b.StartTimer()

		quick.Sort(compare.Int64, data)

		b.StopTimer()

		if debug {

			if !sort.IsSorted(int64slice(data)) {
				b.Errorf("sort didn't sort - 1M ints")
			}
		}
		fillu(data, &seed)

	}
}

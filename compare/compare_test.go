package compare

import (
	"fmt"
	"testing"
)

func test(do func(i int) (int, string), sign []int, t *testing.T) {
	for i := 0; i < len(sign); i++ {

		r, xy := do(i)
		s := sign[i]
		if r*s < 0 {
			t.Errorf("i=%v|XY=%v| should=%v result=%v\n", i, xy, s, r)
		} else {
			//			fmt.Printf("i=%v|XY=%v| should=%v result=%v\n",i,xy,s, r)
		}

	}
}

func TestFloat32s(t *testing.T) {
	d := []float32{654, -56, 146, -44, 896, 352, 123, 687, 688, 689}
	sign := []int{1, -1, 1, -1, 1, 1, -1, -1, -1}

	test(func(i int) (int, string) {
		x := &d[i]
		y := &d[i+1]
		return Float32(x, y), fmt.Sprintf("%v %v %v", *x, *y, *x-*y)
	}, sign, t)
}

func TestFloat64s(t *testing.T) {
	d := []float64{654, -56, 146, -44, 896, 352, 123, 687, 688, 689}
	sign := []int{1, -1, 1, -1, 1, 1, -1, -1, -1}

	test(func(i int) (int, string) {
		x := &d[i]
		y := &d[i+1]
		return Float64(x, y), fmt.Sprintf("%v %v %v", *x, *y, *x-*y)
	}, sign, t)
}

func TestInt64s(t *testing.T) {
	d := []int64{654, -56, 146, -44, 896, 352, 123, 687, 688, 689}
	sign := []int{1, -1, 1, -1, 1, 1, -1, -1, -1}

	test(func(i int) (int, string) {
		x := &d[i]
		y := &d[i+1]
		return Int64(x, y), fmt.Sprintf("%v %v %v", *x, *y, *x-*y)
	}, sign, t)
}

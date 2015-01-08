package compare

import (
	"math"
)

func Float32(a, b *float32) int {
	r := *a - *b
	rr := int32(math.Float32bits(r))
	return int(rr)
}

func Int64(a, b *int64) int {
	r := int(*a>>32) - int(*b>>32)
	if r != 0 {
		return r
	}
	return int(*a) - int(*b)
}

func Float64(a, b *float64) int {
	p := *a - *b
	q := int64(math.Float64bits(p))
	r := int(q >> 32)
	if r != 0 {
		return r
	}
	return int(r)
}

# quick.Sort(compar, slice)

A sorting algorithm for slices.

# Benchmark

N        | Std sort | Gomacro sort 
-------- | -------- | -------
10 | 992 | 1015 
100 | 16662 | 11287 
1000 | 232382 | 156376 
10000 | 3108161 | 2025253 
100000 | 40125309 | 24780259 
1000000 | 482527863 | 293896621 
10000000 | 5697779309 | 3452470132 

# Go get

  go get github.com/gomacro/sort/unsafe/quick

# Import

	"github.com/gomacro/sort/compare"
	"github.com/gomacro/sort/unsafe/quick"

# Example

N/A

# Docs

N/A

# License

Adapted from go std sort.

The unsafe/, test/, compare/ is GPLv2

Version: v0.1

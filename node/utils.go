package node

import "math"

func log2(n int) int {
	if n < 1 {
		return math.MinInt64
	}

	count := 0
	for ; n >= 2; n >>= 1 {
		count++
	}
	return count
}

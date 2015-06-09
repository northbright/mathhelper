package mathhelper

// MinInt64 is the min(a, b) for int64 type.
func MinInt64(a int64, b int64) int64 {
	if a <= b {
		return a
	} else {
		return b
	}
}

// MaxInt64 is the max(a, b) for int64 type.
func MaxInt64(a int64, b int64) int64 {
	if a >= b {
		return a
	} else {
		return b
	}
}

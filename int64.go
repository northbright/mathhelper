package mathhelper

func MinInt64(a int64, b int64) int64 {
	if a <= b {
		return a
	} else {
		return b
	}
}

func MaxInt64(a int64, b int64) int64 {
	if a >= b {
		return a
	} else {
		return b
	}
}

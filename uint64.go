package mathhelper

func MinUInt64(a uint64, b uint64) uint64 {
	if a <= b {
		return a
	} else {
		return b
	}
}

func MaxUInt64(a uint64, b uint64) uint64 {
	if a >= b {
		return a
	} else {
		return b
	}
}

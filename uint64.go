package mathhelper

// MinUInt64 is the min(a, b) for uint64 type.
func MinUInt64(a uint64, b uint64) uint64 {
	if a <= b {
		return a
	}
	return b
}

// MaxUInt64 is the max(a, b) for uint64 type.
func MaxUInt64(a uint64, b uint64) uint64 {
	if a >= b {
		return a
	}
	return b
}

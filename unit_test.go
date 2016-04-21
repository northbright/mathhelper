package mathhelper_test

import (
	"fmt"

	"github.com/northbright/mathhelper"
)

func Example() {

	var a int64 = -32455
	var b int64 = 32923
	var a1 uint64 = 1002993
	var b1 uint64 = 2222222

	c := mathhelper.MinInt64(a, b)
	fmt.Printf("MinInt64(%v, %v) = %v\n", a, b, c)

	c = mathhelper.MaxInt64(a, b)
	fmt.Printf("MaxInt64(%v, %v) = %v\n", a, b, c)

	d := mathhelper.MinUInt64(a1, b1)
	fmt.Printf("MinUInt64(%v, %v) = %v\n", a1, b1, d)

	d = mathhelper.MaxUInt64(a1, b1)
	fmt.Printf("MaxUInt64(%v, %v) = %v\n", a1, b1, d)

	// Output:
	// MinInt64(-32455, 32923) = -32455
	// MaxInt64(-32455, 32923) = 32923
	// MinUInt64(1002993, 2222222) = 1002993
	// MaxUInt64(1002993, 2222222) = 2222222
}

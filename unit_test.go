package mathhelper_test

import (
	"github.com/northbright/fnlog"
	"github.com/northbright/mathhelper"
	"log"
)

var (
	logger *log.Logger
)

func Example() {

	var a int64 = -32455
	var b int64 = 32923
	var a1 uint64 = 1002993
	var b1 uint64 = 2222222

	c := mathhelper.MinInt64(a, b)
	logger.Printf("MinInt64(%v, %v) = %v", a, b, c)

	c = mathhelper.MaxInt64(a, b)
	logger.Printf("MaxInt64(%v, %v) = %v", a, b, c)

	d := mathhelper.MinUInt64(a1, b1)
	logger.Printf("MinUInt64(%v, %v) = %v", a1, b1, d)

	d = mathhelper.MaxUInt64(a1, b1)
	logger.Printf("MaxUInt64(%v, %v) = %v", a1, b1, d)

	// Output:
}

func init() {
	logger = fnlog.New("")
}

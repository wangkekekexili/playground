package main

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"math"
)

func main() {
	nans := []float64{
		math.NaN(),
		math.Copysign(math.NaN(), -1),
		math.Float64frombits(0xFFFF000000000000),
	}
	for _, n := range nans {
		fmt.Printf("%v %v\n", math.IsNaN(n), n)
	}

	nums := []float64{
		0,
		math.Copysign(0, -1),
		math.NaN(),
		math.Copysign(math.NaN(), -1),
		math.Inf(1),
		math.Inf(-1),
		1,
		2,
		3,
	}

	for _, n := range nums {
		var b bytes.Buffer
		binary.Write(&b, binary.BigEndian, n)
		fmt.Printf("%08b %v\n", b.Bytes(), n)
	}
}

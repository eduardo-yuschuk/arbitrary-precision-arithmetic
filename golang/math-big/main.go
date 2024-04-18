package main

import (
	"fmt"
	"math/big"
)

func main() {

	// add #1
	{
		var x, y, z, expected big.Int
		x.SetUint64(10)
		y.SetUint64(10)
		z.Add(&x, &y)
		expected.SetUint64(20)
		if z.Cmp(&expected) != 0 {
			fmt.Println("add #1 failed")
		}
	}

	// add #2
	{
		var x, y, z, expected big.Int
		x.SetString("FFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFF", 16)
		y.SetUint64(1)
		z.Add(&x, &y)
		fmt.Println(z.Text(16))
		expected.SetUint64(0)
		if z.Cmp(&expected) != 0 {
			fmt.Println("add #2 failed")
		}
	}

	// var value big.Int

	// value.SetUint64(1)
	// //value.SetString("0010000000000000000000000000000000000000000000000000", 2)
	// value.Lsh(&value, 200)
	// fmt.Println(value.Bytes())
}

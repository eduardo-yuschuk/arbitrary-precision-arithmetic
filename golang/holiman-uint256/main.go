package main

import (
	"fmt"

	"github.com/holiman/uint256"
)

func main() {

	var z uint256.Int

	// ADD #1
	if z.Add(uint256.MustFromDecimal("10"), uint256.MustFromDecimal("10")).Cmp(uint256.MustFromDecimal("20")) != 0 {
		fmt.Println("ADD #1 failed")
	}

	// ADD #2
	if z.Add(
		uint256.MustFromHex("0xFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFF"),
		uint256.MustFromHex("0x1")).Cmp(
		uint256.MustFromHex("0x0")) != 0 {
		fmt.Println("ADD #2 failed")
	}

	// MUL #1
	if z.Mul(uint256.MustFromDecimal("10"), uint256.MustFromDecimal("10")).Cmp(uint256.MustFromDecimal("100")) != 0 {
		fmt.Println("MUL #1 failed")
	}

	// MUL #2
	if z.Mul(
		uint256.MustFromHex("0xFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFF"),
		uint256.MustFromDecimal("2")).Cmp(
		uint256.MustFromHex("0xFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFE")) != 0 {
		fmt.Println("MUL #2 failed")
	}

	// SUB #1
	if z.Sub(uint256.MustFromDecimal("10"), uint256.MustFromDecimal("10")).Cmp(uint256.MustFromDecimal("0")) != 0 {
		fmt.Println("SUB #1 failed")
	}

	// SUB #2
	if z.Sub(
		uint256.MustFromDecimal("0"),
		uint256.MustFromDecimal("1")).Cmp(
		uint256.MustFromHex("0xFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFF")) != 0 {
		fmt.Println("SUB #2 failed")
	}

	// DIV #1
	if z.Div(uint256.MustFromDecimal("10"), uint256.MustFromDecimal("10")).Cmp(uint256.MustFromDecimal("1")) != 0 {
		fmt.Println("DIV #1 failed")
	}

	// DIV #2
	if z.Div(
		uint256.MustFromDecimal("1"),
		uint256.MustFromDecimal("2")).Cmp(
		uint256.MustFromDecimal("0")) != 0 {
		fmt.Println("DIV #2 failed")
	}

	// SDIV #1
	if z.SDiv(uint256.MustFromDecimal("10"), uint256.MustFromDecimal("10")).Cmp(uint256.MustFromDecimal("1")) != 0 {
		fmt.Println("SDIV #1 failed")
	}

	// SDIV #2
	if z.SDiv(
		uint256.MustFromHex("0xFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFE"),
		uint256.MustFromHex("0xFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFF")).Cmp(
		uint256.MustFromDecimal("2")) != 0 {
		fmt.Println("SDIV #2 failed")
	}
}

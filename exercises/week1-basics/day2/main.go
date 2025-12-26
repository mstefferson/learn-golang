package main

import (
	"fmt"
	"strconv"
)

func main() {

	// Exercise 1
	var temperature float64 = 98.6
	isValid := true
	var errorCount int

	fmt.Println("== Exercise 1 ==")
	fmt.Println("Temperature:", temperature)
	fmt.Println("Is Active:", isValid)
	fmt.Println("Error Count", errorCount)

	// Exercise 2
	var zeroInt int
	var zeroFloat float64
	var zeroBool bool
	var zeroString string

	fmt.Println("== Exercise 2 ==")
	fmt.Println("int zero", zeroInt)
	fmt.Println("float64 zero", zeroFloat)
	fmt.Println("bool zero", zeroBool)
	fmt.Printf("string zero: %q\n", zeroString)

	// Exercise 3
	intVal := 42
	floatVal := float64(intVal)
	floatPi := 3.14158
	intPi := floatPi
	stringVal := string(42)
	stringValConvert := strconv.Itoa(42)

	fmt.Println("int", intVal, " -> float64:", floatVal)
	fmt.Println("float64", floatPi, " -> int:", intPi)
	// Danger, 42 is an ASCII asterisk
	fmt.Println("stingInt", stringVal)
	fmt.Println("stingInt convert", stringValConvert)

	// Exercise 4
	const (
		Pending  = iota // 0
		Running         // 1
		Complete        // 2
		Failed          // 3
	)
	const MaxRetries = 3

	fmt.Printf("Status values - Pending: %d, Running: %d, Complete: %d, Failed: %d\n", Pending, Running, Complete, Failed)
	fmt.Printf("Max retries: %d\n", MaxRetries)

	// Exercise 5
	fmt.Println("== Exercise 5 ==")
	fmt.Printf("42 has type: %T\n", 42)
	fmt.Printf("3.14 has type: %T\n", 3.14)
	fmt.Printf("hello has type: %T\n", "hello")
	fmt.Printf("true has type: %T\n", true)
	// Rune literal
	fmt.Printf("'A' has type: %T\n", 'A')
	fmt.Println('A')
}

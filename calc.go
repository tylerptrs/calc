package main

import (
	"fmt"
	"log"
	"strconv"
)

func add(x, y float64) float64 {
	return x + y
}

func subtract(x, y float64) float64 {
	return x - y
}

func multiply(x, y float64) float64 {
	return x * y
}

func divide(x, y float64) float64 {
	return x / y
}

func second() float64 {
	var ss string
	fmt.Printf("What is the second number?")
	fmt.Scanln(&ss)
	s, err := strconv.ParseFloat(ss, 64)
	if err != nil {
		log.Fatal(err)
	}
	return s
}

func main() {
	var fs string
	fmt.Printf("What is the first number?")
	fmt.Scanln(&fs)
	f, err := strconv.ParseFloat(fs, 64)
	if err != nil {
		log.Fatal(err)
	}

	var o string
	fmt.Printf("What is the operator? (+, -, *, /)")
	fmt.Scanln(&o)

	var a float64
	switch o {
	case "+":
		secondNumber := second()
		a = add(f, secondNumber)
		fmt.Printf("%v %s %v = %v\n", f, o, secondNumber, a)
	case "-":
		secondNumber := second()
		a = subtract(f, secondNumber)
		fmt.Printf("%v %s %v = %v\n", f, o, secondNumber, a)

	case "*":
		secondNumber := second()
		a = multiply(f, secondNumber)
		fmt.Printf("%v %s %v = %v\n", f, o, secondNumber, a)

	case "/":
		secondNumber := second()
		a = divide(f, secondNumber)
		fmt.Printf("%v %s %v = %v\n", f, o, secondNumber, a)
	default:
		fmt.Printf("%v not recognized \n", o)
	}
}

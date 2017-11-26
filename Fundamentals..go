//Every Go program starts with a package decleration which provides a way for
//use to reuse
package main

import "fmt"

//import enables us to use code from other packages
// The format package provides formatting for input and output

// A comment

/*
Multi line comment.
 */

//The main function starts with 'func' and is opened and closed with {}
//When executing a program main is the backbone for everything.

func main() {
	println("Hello world")
	// You can get a description of a function by typing godoc fmt Println
	// for example in the terminal

	// You execute Go programs like this in the terminal go run herewego.go

	// Variables are statically typed, which means their type can't change
	// Variable names must start with a letter and may contain letters, numbers
	// or the _

	// An int is a positive or negative number without decimals
	// Versions
	// uint8 : unsigned  8-bit integers (0 to 255)
	// uint16 : unsigned 16-bit integers (0 to 65535)
	// uint32 : unsigned 32-bit integers (0 to 4294967295)
	// uint64 : unsigned 64-bit integers (0 to 18446744073709551615)
	// int8 : signed  8-bit integers (-128 to 127)
	// int16 : signed 16-bit integers (-32768 to 32767)
	// int32 : signed 32-bit integers (-2147483648 to 2147483647)
	// int64 : signed 64-bit integers (-9223372036854775808 to
	// 9223372036854775807)

	var age int64 = 40

	//A float is a decimal number

	var favNum float64 = 0.69

	oldAge := 56

	fmt.Println(oldAge)
	//DO NOT HAVE TO DEFINE THE DATA TYPE AS SHOWN 'randNum'
	//You cannot assign a non-compatible type later.
	//Example randNum = "Hello"

	fmt.Println(age, " ", favNum)

}

package main

import (
	"fmt"

)

func main(){
	//constant is a variable that cannot be changed.
	const pi float64  = 3.14159265359

	// Declearing multiple variables.
	var(
		student = "Peter"
		studentId = 21344
		studentGpa = 3.3
	)
	//Prints all the variables.
	fmt.Println("Student name:",student, "\nIdentification No.",studentId,"\nStudent GPA:", studentGpa)

	//Prints out how many characters are inside the variable.
	fmt.Println(len(student))

}
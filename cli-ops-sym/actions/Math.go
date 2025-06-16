package actions

import (
	"fmt"
)

func Arithmeticops(){
	fmt.Println("--- Arithmetic operations ---")

	// What math do u want?
	fmt.Println("What math do u want?")
	fmt.Println("1. Addition")
	fmt.Println("2. Subtraction")
	fmt.Println("3. Multiplication")
	fmt.Println("4. Division")
	fmt.Println("5. Exit")

	var choice int
	fmt.Scan(&choice)

	// Key in the values 
	var value1, value2 int

	switch choice {
	case 1:
		fmt.Println("Enter the first number")
		fmt.Scan(&value1)
		fmt.Println("Enter the second number")
		fmt.Scan(&value2)
		fmt.Println(value1 + value2)
	case 2:
		fmt.Println("Enter the first number")
		fmt.Scan(&value1)
		fmt.Println("Enter the second number")
		fmt.Scan(&value2)
		fmt.Println(value1 - value2)
	case 3:
		fmt.Println("Enter the first number")
		fmt.Scan(&value1)
		fmt.Println("Enter the second number")
		fmt.Scan(&value2)
		if value2 == 0 {
			fmt.Println("Division by zero is not allowed")
		} else {
			fmt.Println(value1 * value2)
		}
	case 4:
		fmt.Println("Enter the first number")
		fmt.Scan(&value1)
		fmt.Println("Enter the second number")
		fmt.Scan(&value2)
		if value2 == 0 {
			fmt.Println("Division by zero is not allowed")
		} else {
			fmt.Println(value1 / value2)
		}
	case 5:
		fmt.Println("Thanks for using this app")
		fmt.Println("Goodbye")
	default:
		fmt.Println("Invalid choice")
		fmt.Println("Try again")
		Arithmeticops()
	}
}
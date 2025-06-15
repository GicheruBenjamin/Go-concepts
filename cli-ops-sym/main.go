package main

import (
	"fmt"
	"ops-cli/actions"
)

func main(){
	// make a choice what action u wanna do

	fmt.Println(" +@#--==   The app Starts from here on out+@#--==")
	fmt .Println(" Make a choice what action u wanna do..") 
	fmt.Println(" 1. Access stuff \n. Arithmetic operations \n 3. Binary data \n 4. Compare \n 5. Desion making \n 6. Spread stuff \n 7. Store stuff \n 8. Working with routines \n")
	fmt.Println("Enter your choice: ")
	var choice int
	fmt.Scan(&choice)

	switch choice {
	case 1:
		actions.Access()
	case 2:
		actions.Arithmeticops()
	case 3:
		actions.Bin()
	case 4:
		actions.Compare()
	case 5:
		actions.Desionmaking()
	case 6:
		actions.Spread()
	case 7:
		actions.Store()
	case 8:
		actions.Routines()
	default:
		fmt.Println("Invalid choice")
	}
}
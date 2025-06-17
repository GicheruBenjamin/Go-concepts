package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

// A simple struct to demonstrate field access
type Data struct {
	Value int
}

func main() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Go Operator Categories CLI Demo")
	fmt.Println("--------------------------------")

	// 1. Input: User gives a number
	fmt.Print("Enter a number: ")
	inputStr, _ := reader.ReadString('\n')
	inputStr = strings.TrimSpace(inputStr)
	num, _ := strconv.Atoi(inputStr)

	// 2. Storing & accessing struct field
	data := Data{Value: num}
	fmt.Printf("\n[Accessing] Struct Value: data.Value = %d\n", data.Value)

	// 3. Bitwise operations
	bitwise := data.Value ^ 0xF // XOR with 15
	fmt.Printf("[Bitwise] data.Value ^ 15 = %d\n", bitwise)

	// 4. Comparisons
	isEven := data.Value%2 == 0
	fmt.Printf("[Comparison] Is %d even? %v\n", data.Value, isEven)

	// 5. Decision making
	if data.Value > 10 {
		fmt.Println("[Decision] It's greater than 10")
	} else {
		fmt.Println("[Decision] It's 10 or less")
	}

	// 6. Arithmetic operation
	data.Value += 5
	fmt.Printf("[Arithmetic] Value after += 5: %d\n", data.Value)

	// 7. Assignment (storing)
	memory := data.Value
	fmt.Printf("[Assignment] Stored in memory variable: %d\n", memory)

	// 8. Type conversion
	var asFloat float64 = float64(data.Value)
	fmt.Printf("[Type Conversion] As float64: %.2f\n", asFloat)

	// 9. Variadic function usage
	sum := sumAll(1, 2, 3, data.Value)
	fmt.Printf("[Variadic] Sum of 1,2,3,data.Value: %d\n", sum)

	// 10. Goroutine usage (concurrent print)
	fmt.Println("[Goroutine] Running in background...")
	go showLater(data.Value)

	// Wait for goroutine
	time.Sleep(1 * time.Second)
}

// Variadic function
func sumAll(nums ...int) int {
	total := 0
	for _, n := range nums {
		total += n
	}
	return total
}

// Background goroutine function
func showLater(x int) {
	time.Sleep(500 * time.Millisecond)
	fmt.Printf(">>> [Goroutine Output] Final value is %d\n", x)
}

package main

import (
	"gcwork/pkg/object"
	"gcwork/pkg/utils"
	"time"
)

func main() {
	manager := object.NewObjectManager()

	// Simulate object creation and destruction over time
	for i := 0; i < 1000; i++ {
		obj := manager.CreateObject(1024 * 10) // 10 KB per object

		// Simulate random destruction
		if i%3 == 0 {
			manager.DestroyObject(obj)
		}

		// Print memory usage every 100 iterations
		if i%100 == 0 {
			utils.PrintMemoryUsage()
		}

		// Introduce a small delay to simulate real-time operations
		time.Sleep(100 * time.Millisecond) // 10ms delay per iteration
	}

	// Force garbage collection and print final memory stats
	utils.PrintMemoryUsage()
	utils.ForceGC()
	utils.PrintMemoryUsage()
}

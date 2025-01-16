
package utils

import (
	"fmt"
	"runtime"
)

// PrintMemoryUsage prints the current memory usage
func PrintMemoryUsage() {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	fmt.Printf("Alloc: %v KB, TotalAlloc: %v KB, Sys: %v KB, NumGC: %v\n",
		m.Alloc/1024, m.TotalAlloc/1024, m.Sys/1024, m.NumGC)
}

// ForceGC triggers garbage collection
func ForceGC() {
	fmt.Println("Forcing Garbage Collection...")
	runtime.GC()
}

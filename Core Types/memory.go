package main

import(
	"fmt"
	"runtime"
)
// To check the memory used by a type
func main() {
	var list []int
	for i := 0; i < 10000000; i++ {
		list = append(list, 100)
	}
	fmt.Println(list)
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	fmt.Printf("TotalAlloc (Heap) = %v MiB\n", m.TotalAlloc/1024/1024)
}
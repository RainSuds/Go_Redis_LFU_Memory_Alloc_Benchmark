package main

import (
	"encoding/csv"
	"fmt"
	"runtime"
)

func bToMb(b uint64) uint64 {
	return b / 1024 / 1024
}

// Function to log memory usage at different stages
func logMemUsage(stage string, writer *csv.Writer) {
	var memStats runtime.MemStats
	runtime.ReadMemStats(&memStats)

	// Log the memory usage in megabytes
	allocatedMB := bToMb(memStats.Alloc) // The number of bytes of allocated heap ojects, the amount of memory currently in use by the application
	totalAllocatedMB := bToMb(memStats.TotalAlloc) // Cumulative total number of bytes allocated for heap objects
	systemMemoryMB := bToMb(memStats.Sys) // Total bytes of memory obtained from thhe system, memory that is not in use but allocated by Go runtime
	numGC := memStats.NumGC
	heapIdleMB := bToMb(memStats.HeapIdle)
	// freed := memStats.Frees // Frees is the cumulative count of heap objects freed in this size class.

	// Write memory statistics to the CSV file
	writer.Write([]string{
		stage,
		fmt.Sprintf("%d", allocatedMB),
		fmt.Sprintf("%d", totalAllocatedMB),
		fmt.Sprintf("%d", systemMemoryMB),
		fmt.Sprintf("%d", numGC),
		fmt.Sprintf("%v", heapIdleMB),
	})

	// Flush the writer to ensure data is written
	writer.Flush()
}

// Function to force garbage collection and log memory usage
func gcAndLog(stage string, writer *csv.Writer) {
	// Force garbage collection
	runtime.GC()
	
	// Log memory usage after garbage collection
	logMemUsage(stage, writer)
}

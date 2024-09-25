package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

var (
	redisCache  = make(map[int][]byte)
	keyOrder    = []int{}
	CACHE_LIMIT int
)

func generateRandomValue(size int, r *rand.Rand) []byte {
	value := make([]byte, size)
	for i := range value {
		value[i] = byte(r.Intn(256))
	}
	return value
}

func main() {

	// Read input from console
	fmt.Print("Input a file: ")
	reader := bufio.NewReader(os.Stdin)
	filename, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("An error occurred while reading input. Please try again", err)
		return
	}
	filename = strings.TrimSpace(filename)

	// Read file content
	fileContent, err := os.ReadFile(filename)
	if err != nil {
		fmt.Println("An error occurred while reading the file. Please try again", err)
		return
	}
	scanner := bufio.NewScanner(strings.NewReader(string(fileContent)))

	// Read the first line for entry limit
	scanner.Scan()
	CACHE_LIMIT, err = strconv.Atoi(scanner.Text())
	if err != nil {
		fmt.Println("Invalid entry limit in the file. Please check the file format.", err)
		return
	}

	operations := make([]Operation, 0)
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Fields(line)
		if len(parts) == 0 {
			continue
		}

		currentOp := Operation{Type: parts[0]}
		if len(parts) > 1 {
			currentOp.NumEntries, _ = strconv.Atoi(parts[1])
		}
		if len(parts) > 2 {
			currentOp.ByteSize, _ = strconv.Atoi(parts[2])
		}
		operations = append(operations, currentOp)
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("An error occurred while reading the file. Please try again", err)
		return
	}

	csvfile, err := os.Create("redis_LRC_benchmark_memory_usage.csv")
	if err != nil {
		fmt.Println("Error creating CSV file:", err)
		return
	}
	defer csvfile.Close()

	writer := csv.NewWriter(csvfile)
	defer writer.Flush()

	writer.Write([]string{"Stage", "Allocated Memory (MiB)", "Total Allocated Memory (MiB)", "System Memory (MiB)",
		"Number of Garbage Collections", "Idle Heap Memory (MiB)"})

	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for _, op := range operations { // _ is index, op is value
		switch op.Type {
		case "READ":
			readOperation(op.NumEntries, writer)
		case "WRITE":
			writeOperation(op.NumEntries, op.ByteSize, r, writer)
		case "DELETE":
			deleteOperation(op.NumEntries, writer)
		case "UPDATE":
			updateOperation(op.NumEntries, op.ByteSize, r, writer)
		case "GC":
			gcAndLog("After manual GC", writer)
		default:
			fmt.Println("Unknown operation type:", op.Type)
		}
	}

	// Clear the cache and run garbage collection
	redisCache = nil
	keyOrder = nil
	gcAndLog("Final GC", writer)

	fmt.Println("Benchmark complete, memory usage logged to CSV.")
}

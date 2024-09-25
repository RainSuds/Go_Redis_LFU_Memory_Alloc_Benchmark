package main

import (
	"encoding/csv"
	"fmt"
	"math/rand"
)

type Operation struct {
	Type         string
	NumEntries int
	ByteSize     int
}

/*
 * This function takes in a string and a value then add  it to the cache
 * if the cache is full, it deletes the earliest entry and adds the new entry
*/
func addToCache(key int, value []byte) {
	if len(keyOrder) >= CACHE_LIMIT {
		deleteEarliestCache()
	}
	redisCache[key] = value
	keyOrder = append(keyOrder, key)
}

/*
 * This function deletes the earliest entry in the cache
*/
func deleteEarliestCache() {
	if len(keyOrder) > 0 {
		earliest := keyOrder[0]
		delete(redisCache, earliest)
		keyOrder = keyOrder[1:]
	}
}

/*
 * This function reads in the first NumEntries entries in the cache (no return)
*/
func readOperation(NumEntries int, writer *csv.Writer) {
	for i := 0; i < NumEntries; i++ {
		key := keyOrder[i]
		_ = redisCache[key]
	}
	logMemUsage(fmt.Sprintf("After reading %d entries", NumEntries), writer)
}

/*
 * This function takes in NumEntries and ByteSize and writes NumEntries entries with ByteSize byte values randomly generated
 * values to the cache
*/
func writeOperation(NumEntries, ByteSize int, r *rand.Rand, writer *csv.Writer) {
	for i := 0; i < NumEntries; i++ {
		addToCache(i, generateRandomValue(ByteSize, r))
	}
	logMemUsage(fmt.Sprintf("After writing %d entries with %d-byte values", NumEntries, ByteSize), writer)
}

/*
 * This function takes in NumEntries and deletes the earliest entry in the cache NumEntries times
*/
func deleteOperation(NumEntries int, writer *csv.Writer) {
	for i := 0; i < NumEntries; i++ {
		deleteEarliestCache()
	}
	logMemUsage(fmt.Sprintf("After deleting %d entries", NumEntries), writer)
}

/*
 * This function reads in total number of NumEntries and change the entries to another random value of same size
*/
func updateOperation(NumEntries, ByteSize int, r *rand.Rand, writer *csv.Writer) {
	for i := 0; i < NumEntries; i++ {
		key := r.Intn(len(redisCache))
		redisCache[key] = generateRandomValue(ByteSize, r)
	}
	logMemUsage(fmt.Sprintf("After updating %d entries with %d-byte values", NumEntries, ByteSize), writer)
}
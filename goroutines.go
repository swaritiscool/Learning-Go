package main

import (
	"fmt"
	"time"
	"sync"
)

var m = sync.Mutex{}
var wg = sync.WaitGroup{}
var dbData = []string{"id1", "id2", "id3", "id4", "id5", "id6"}
var results = []string{}

func main() {
	t0 := time.Now()
	for i := 0; i < len(dbData); i++ {
		wg.Add(1)
		go get_dbData(i) // On its own, this won't work... It will call the tasks but won't wait... hence we need to use waitgroups
	}
	wg.Wait()
	fmt.Printf("\n\nResults are: %v", results)
	fmt.Printf("\n\nTime Taken: %v\n", time.Since(t0)) // Without go routines: 12 seconds | With go routines: 2 seconds
}

func get_dbData(i int) {
	var delay float32 = 2000 // Delay for experimentation
	time.Sleep(time.Duration(delay)*time.Millisecond)
	fmt.Printf("DB returned %v\n", dbData[i])
	m.Lock() // We are locking i.e. allowing only one process to edit the results slice to prevent corrupt memory caused by multiple programs saving the same thing to one memory location
	results = append(results, dbData[i])
	m.Unlock()
	wg.Done()
}

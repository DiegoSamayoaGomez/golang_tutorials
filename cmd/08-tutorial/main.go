package main

// Just a quick exercise in order to understand concurrency
// concurrency ins´t equal to parallelism

import (
	"fmt"
	"sync"
	"time"
)

var m = sync.Mutex{}
var wg = sync.WaitGroup{}
var dbData = []string{"id1", "id2", "id3", "id4", "id5"}
var results = []string{}

func main() {
	t0 := time.Now()
	for i := 0; i < len(dbData); i++ {
		wg.Add(1)
		go dbCall(i)
	}
	wg.Wait()
	fmt.Printf("\nTotal execution time: %v ", time.Since(t0))
	fmt.Printf("\nThe results are %v ", results)
}
func dbCall(i int) {
	//Simulates DB call delay
	var delay float32 = 2000
	time.Sleep(time.Duration(delay) * time.Millisecond)
	/*
		fmt.Println("The result from the database is: ", dbData[i])
		m.Lock()
		results = append(results, dbData[i])
		m.Unlock()
	*/
	save(dbData[i])
	log()

	wg.Done()
}

func save(result string) {
	m.Lock()
	results = append(results, result)
	m.Unlock()

}

func log() {
	m.Lock()
	fmt.Printf("The results are %v \n", results)
	m.Unlock()

}

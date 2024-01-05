package main
// Счетчик синхронизированных горутин

import (
	"fmt"
	"sync"
	"time"
)

var counter int
var mutex sync.Mutex


func incrementCounter(wg *sync.WaitGroup) {
	defer wg.Done()

	mutex.Lock()
	defer mutex.Unlock()

	time.Sleep(500 * time.Millisecond)

	counter++
	fmt.Printf("Counter: %d\n", counter)
}

func main() {
	var wg sync.WaitGroup

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go incrementCounter(&wg)
	}

	wg.Wait()

	fmt.Printf("Final Counter: %d\n", counter)
}

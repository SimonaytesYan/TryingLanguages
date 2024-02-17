//==============================================================================
//  								USING GORUTINES
//==============================================================================

package main

import (
	"fmt"
	"time"
	"sync"
)

var a int = 0
var mutex sync.Mutex

func start() {
	a = a + 1
	
	fmt.Println("No mutex:", a);
}

func normal_start() {
	mutex.Lock()
	
	a = a + 1
	fmt.Println("With mutex:", a);
	
	mutex.Unlock()
}

func main() {

	for i := 0; i < 5; i++ {
		go start()
	}

	for i := 0; i < 5; i++ {
		go normal_start()
	}

	time.Sleep(1 * time.Second)

	fmt.Println("==========\nres =", a);
}

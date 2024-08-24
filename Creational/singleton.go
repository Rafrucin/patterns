package patterns

import (
	"math/rand"
	"sync"
	"time"
)

// SINGLETON

type singletonStruct struct {
	CreatorID int
}

var singleton *singletonStruct
var mu sync.Mutex
var wg sync.WaitGroup

func Singleton() {
	wg = sync.WaitGroup{}

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go initSingleton(i)
	}

	wg.Wait()
	println("-----------", singleton.CreatorID)
}

func initSingleton(i int) {
	time.Sleep(time.Duration(rand.Intn(100)) * time.Millisecond)
	println(i)
	defer wg.Done()
	if singleton == nil {
		mu.Lock()
		defer mu.Unlock()
		if singleton != nil {
			return
		}
		singleton = &singletonStruct{CreatorID: i}
	}
}

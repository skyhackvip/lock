package main

import (
	"fmt"
	"github.com/go-redis/redis"
	"sync"
	"time"
)

var client = redis.NewClient(&redis.Options{
	Addr:     "localhost:6379",
	Password: "",
	DB:       0,
})

var counter int64
var wg sync.WaitGroup
var lockKey = "myrslock"

func main() {
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func() {
			lock(incr)
		}()
	}
	wg.Wait()
	fmt.Printf("final counter is %d\n", counter)
}

func incr() {
	counter++
	fmt.Printf("after incr is %d\n", counter)
}

func lock(myfunc func()) {
	defer wg.Done()

	//lock
	lockSuccess, err := client.SetNX(lockKey, 1, time.Second*5).Result()
	if err != nil || !lockSuccess {
		fmt.Println("get lock fail")
		return
	} else {
		fmt.Println("get lock")
	}

	//run func
	myfunc()

	//unlock
	_, err = client.Del(lockKey).Result()
	if err != nil {
		fmt.Println("unlock fail")
	} else {
		fmt.Println("unlock")
	}
}

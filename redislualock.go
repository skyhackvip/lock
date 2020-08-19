package main

import (
	"fmt"
	"github.com/go-redis/redis"
	"os/exec"
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
	uuid := getUuid()

	//lock
	lockSuccess, err := client.SetNX(lockKey, uuid, time.Second*5).Result()
	if err != nil || !lockSuccess {
		fmt.Println("get lock fail")
		return
	} else {
		fmt.Println("get lock")
	}

	//run func
	myfunc()

	//unlock
	var luaScript = redis.NewScript(`
		if redis.call("get", KEYS[1]) == ARGV[1]
			then
				return redis.call("del", KEYS[1])
			else
				return 0
		end
	`)
	rs, _ := luaScript.Run(client, []string{lockKey}, uuid).Result()
	if rs == 0 {
		fmt.Println("unlock fail")
	} else {
		fmt.Println("unlock")
	}

}

func getUuid() string {
	out, err := exec.Command("uuidgen").Output()
	if err != nil {
		panic(err)
	}
	return string(out)
}

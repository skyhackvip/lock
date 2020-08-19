package main

import (
	"fmt"
	"github.com/samuel/go-zookeeper/zk"
	"sync"
	"time"
)

var wg sync.WaitGroup
var counter int64

func main() {
	var conn *zk.Conn
	conn, _, err := zk.Connect([]string{"localhost:2181"}, time.Second)
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	for i := 0; i < 5; i++ {
		wg.Add(1)
		go zklock(conn, incr)
	}
	wg.Wait()
	fmt.Printf("final counter is %d\n", counter)
}

func incr() {
	counter++
	fmt.Printf("after incr is %d\n", counter)
}

func zklock(conn *zk.Conn, myfunc func()) {
	defer wg.Done()
	lock := zk.NewLock(conn, "/mylock", zk.WorldACL(zk.PermAll))
	err := lock.Lock()
	if err != nil {
		panic(err)
	}
	fmt.Println("get lock")

	myfunc()

	lock.Unlock()
	fmt.Println("unlock")
}

func testchild(conn *zk.Conn) {
	ch, _, err := conn.Children("/test")
	if err != nil {
		panic(err)
	}
	fmt.Printf("%v \n", ch)
}

func createNode(conn *zk.Conn) {
	//create node
	nodeName, err := conn.Create("/testlock", nil, zk.FlagSequence, zk.WorldACL(zk.PermAll))
	if err != nil {
		panic(err)
	}
	fmt.Println(nodeName)
}

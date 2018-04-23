package main

import (
	"fmt"
	"math/rand"
	"time"
)

func raceDemo() {
	start := time.Now()
	var t *time.Timer
	t = time.AfterFunc(randDuration(), func() {
		fmt.Println(time.Since(start))
		t.Reset(randDuration()) // 无同步读写t导致race。 go run/test/build -race 检查
	})

	time.Sleep(time.Second * 5)
}

func randDuration() time.Duration {
	return time.Duration(rand.Int63n(1e9))
}

func noRace() {
	start := time.Now()
	var t *time.Timer
	reset := make(chan bool)
	t = time.AfterFunc(randDuration(), func() {
		fmt.Println(time.Since(start))
		reset <- true
	})

	for time.Since(start) < 5*time.Second {
		<-reset
		t.Reset(randDuration()) // 写t操作移出读t操作，解决race
	}
}

func main() {
	//raceDemo()
	var foo bool
	if foo==false && foo==true {
		noRace()
	}
}

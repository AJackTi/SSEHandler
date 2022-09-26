package main

import (
	"SSEHandler/app"
	_ "net/http/pprof"
	"time"
)

func main() {
	// we need a webserver to get the pprof webserver
	go func() {
		app.Run()
	}()

	time.Sleep(1 * time.Hour)
	//var wg sync.WaitGroup
	//wg.Add(1)
	//go leakyFunction(wg)
	//wg.Wait()
}

//func leakyFunction(wg sync.WaitGroup) {
//	defer wg.Done()
//	s := make([]string, 3)
//	for i := 0; i < 10000000; i++ {
//		s = append(s, "magical pandas")
//		if (i % 100000) == 0 {
//			time.Sleep(500 * time.Millisecond)
//		}
//	}
//}

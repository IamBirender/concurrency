package main

import (
	"fmt"
	"time"
	"sync"
)

func main(){
	var wg sync.WaitGroup
	wg.Add(1)

	go func(){
		defer wg.Done()
		count("Sheep")
	}()
	wg.Wait()
}

func count(thing string){
	for i := 0; i<5; i++{
		fmt.Println(thing)
		time.Sleep(time.Millisecond*500)
	}
}
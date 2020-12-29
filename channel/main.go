package main

import (
	"fmt"
	"time"
)

func main(){
	c := make(chan string) 
	go count("Sheep", c)

	for msg := range c{
		fmt.Println(msg)
	}
}

func count(thing string, c chan string){
	for i := 0; i<5; i++{
		c <- thing 
		//sending to a channel is a blocking operation. 
		time.Sleep(time.Millisecond*500)
	}
	close(c)		//reciever should never close a channel
}
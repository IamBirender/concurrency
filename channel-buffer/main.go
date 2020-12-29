package main

import (
	"fmt"
)

func main(){
	c := make(chan string, 2)  //2 is the buffer size of the channel. 
	c <- "hello"
	c <- "world"

	// as we are using the buffer here so this will not block the main go routine.
	fmt.Println(<-c)
	fmt.Println(<-c) 
}

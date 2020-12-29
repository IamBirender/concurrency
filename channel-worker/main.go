package main

import (
	"fmt"
)

func main(){
	jobs := make(chan int, 100)
	results := make(chan int, 100)

	go worker(jobs, results) //this will be using 100 % of cpu
	//and if we copy paste it multiple times we will be using up 100*n % of cpu
	go worker(jobs, results)
	go worker(jobs, results)

	for i:=0; i<100 ; i++{
		jobs <- i
	}

	close(jobs) //close jobs because we are finished sending the messages. 

	for j := range results{
		fmt.Println(j)
	}
	
}

func worker(jobs <-chan int, results chan<- int){
	for n := range jobs{
		results <- fib(n)
	}
}

func fib(n int) int{
	if n < 2{
		return n 
	}
	return fib(n-1) + fib(n-2)
}

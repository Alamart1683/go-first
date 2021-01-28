package main

import (
	"fmt"
	"../mymath"
	"time"
)

func main() {
	//var number int
	var N int
	var channels int
	fmt.Print("Input N: ")
	fmt.Scan(&N)
	fmt.Print("Input number of channels: ")
	fmt.Scan(&channels)
	start := time.Now()
	parallelPi := mymath.GetParallelPi(N, channels)
	duration := time.Since(start)
	fmt.Println("Parallel Pi:",parallelPi,"time", duration.Milliseconds(),"мс")
	start = time.Now()
	consistPi := mymath.GetPi(N)
	duration = time.Since(start)
	fmt.Println("Consist Pi:",consistPi,"time:", duration.Milliseconds(),"мс")
	//fmt.Print("Input number: ")
	//fmt.Scan(&number)
	//fmt.Println("Fib for ",number," = ",mymath.GetFib(number))
}
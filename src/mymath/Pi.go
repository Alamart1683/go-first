package mymath

import (
	"fmt"
	"math"
	"sync"
)

var waitGroup sync.WaitGroup

// Get Pi number
func GetPi(n int) float64 {
	var sum float64 = 1.0
	for i := 2; i < n; i++ {
		sum += math.Pow(1.0 / float64(i), 2)
	}
	var pi float64 = math.Sqrt(sum * 6.0)
	return pi
}

// Get parallel Pi number
func GetParallelPi(n int, channels int) float64  {
	localSums := make(chan float64, channels)
	var globalSum float64
	for i:= 0; i < channels; i++ {
		var start int
		if i == 0 {
			start = 1
		} else {
			start = i * (n / channels)
		}
		var stop int
		if i == channels - 1 {
			stop = n
		} else {
			stop = i * (n / channels) + (n / channels)
		}
		waitGroup.Add(1)
		go GetPiRoutine(start, stop, localSums, i)
	}
	waitGroup.Wait()
	fmt.Println("LocalSums len is ", len(localSums))
	for i:= 0; i < len(localSums); i++ {
		globalSum += <-localSums
	}
	return math.Sqrt(globalSum * 6.0)
}

func GetPiRoutine(start int, stop int, localSums chan <- float64, channelNumber int) {
	defer waitGroup.Done()
	var localSum float64
	for i:= start; i < stop; i++ {
		localSum += math.Pow(1.0 / float64(i), 2)
	}
	localSums <- localSum
}
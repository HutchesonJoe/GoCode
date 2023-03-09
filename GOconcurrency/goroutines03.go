package main

import (
    "fmt"
    "sync"
)

var (
	counter int32 
	wg sync.WaitGroup
	mutex   sync.Mutex
)

func main(){
	wg.Add(4)

	go SuperPower("controlling elements", 2)
  go SuperPower("energy projection", 5)
  go SuperPower("flight", 6)
  go SuperPower("healing", 2)

	wg.Wait()
	fmt.Println("Total Mutant Powers used:", counter)
}

func SuperPower(power string, num int) {
	defer wg.Done()

	for i :=0; i < num; i ++{
		mutex.Lock()
		{
			fmt.Println(power)
			counter++
		}
		mutex.Unlock()
	}
}
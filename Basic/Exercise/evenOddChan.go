package main

import (
	"fmt"
	"sync"
)

var c = make(chan []int)

func main() {
	wg := sync.WaitGroup{}
	wg.Add(1)
	
	go server(&wg)	
	globalArray := []int{1,2,3,4,5,6,7,8,9,10}
	c <- globalArray
	wg.Wait()

}

func server(wg *sync.WaitGroup) {
	
	totalArray := <- c
	evenValue := []int{}
	oddValue := []int{}
	for _,i := range totalArray{
		if i % 2 == 0{
			evenValue = append(evenValue,i)
		}else{
			oddValue = append(oddValue,i)
		}
	} 
	fmt.Println(evenValue)
	fmt.Println(oddValue)
	wg.Done()
}

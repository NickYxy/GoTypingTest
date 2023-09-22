package main

import (
	"fmt"
	"sync"
)

func main() {
	letter, number := make(chan bool), make(chan bool)
	wg := sync.WaitGroup{}

	go func() {
		i := 1
		for {
			select {
			case <-number:
				fmt.Print(i)
				i++
				fmt.Print(i)
				i++
				letter <- true
			}
		}
	}()
	wg.Add(1)
	go func(wg1 *sync.WaitGroup) {
		j := 'A'
		for {
			select {
			case <-letter:
				fmt.Print(string(j))
				j++
				fmt.Print(string(j))
				j++
				if j >= 'Z' {
					wg.Done()
					return
				}
				number <- true
			}
		}
	}(&wg)

	number <- true
	wg.Wait()
}

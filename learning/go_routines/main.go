package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	done := make(chan struct{})
	
	wg.Add(2)
	
	go func() {
		defer wg.Done()
		for i := 0; i <= 100; i += 2 {
			select {
			case <-done:
				return
			default:
				fmt.Println("Even: ", i)
			}
		}
	}()

	go func() {
		defer wg.Done()
		for i := 1; i <= 100; i += 2 {
			select {
			case <-done:
				return
			default:
				fmt.Println("Odd: ", i)
			}
		}
	}()

	wg.Wait()
	close(done)
	fmt.Println("Done")
}
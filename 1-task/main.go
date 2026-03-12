package main

import (
	"fmt"
	"sync"
)



func main(){
	total := 1_000_000
	workers := 10

	chunk := total / workers

	results := make([]int, workers)

	var wg sync.WaitGroup

	for i := 0; i < workers; i++ {
		wg.Add(1)

		go func(id int) {
			defer wg.Done()

			start := id * chunk
			end := start + chunk

			local := 0

			for i := start; i < end; i++ {
				local++
			}
			results[id] = local

		}(i)
	}
wg.Wait()

sum := 0

for _, v := range results {
	sum+=v
}

fmt.Println("natija: ", sum)


}
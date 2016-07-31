package main

import(
	"fmt"
	"time"
)

func main() {
	jobs := make(chan int, 100)
	results := make(chan int, 100)

	// Try increase the workers or the jobs number
	jobsN := 9
	workersN := 3

	for w := 1; w <= workersN; w++ {
		go worker(w, jobs, results)
	}
	
	for j := 1; j <= jobsN; j++ {
		 jobs <- j
	}
	close(jobs)

	for a :=1; a <= jobsN; a++ {
		<-results
	}
}

func worker(id int, jobs <-chan int, results chan<- int) {
	for j := range jobs {
		fmt.Printf("Worker %d processing job %d\n", id, j)
		time.Sleep(time.Second)
		results <- j * 2
	}
}

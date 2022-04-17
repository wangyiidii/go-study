package main

import (
	"fmt"
	"time"
)

// work_pool

func main() {

	jobs := make(chan int, 100)
	results := make(chan int, 100)

	// 开启3个goroutine
	for i := 0; i < 3; i++ {
		go worker(i, jobs, results)
	}

	// 再发送5个任务
	for i := 0; i < 5; i++ {
		jobs <- i
	}
	close(jobs)

	//for ret := range results {
	for i := 0; i < 5; i++ {
		ret := <-results
		fmt.Println(ret)
	}

}

func worker(id int, jobs <-chan int, results chan<- int) {
	for job := range jobs {
		fmt.Printf("worker: %d, start job: %d\n", id, job)
		time.Sleep(time.Millisecond * 500)
		results <- job * 2
	}
}

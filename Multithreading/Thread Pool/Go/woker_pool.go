package main

import (
	"fmt"
	"math/rand/v2"
	"time"
)

type Task struct {
	id int
}

func NewTasks(id int) *Task {
	return &Task{
		id: id,
	}
}

func worker(workerID int, jobs <-chan *Task, results chan<- int) {
	for job := range jobs {
		fmt.Printf("Worker %d started job %d \n", workerID, job.id)
		time.Sleep(1 * time.Second)
		fmt.Printf("Worker %d finished job %d \n", workerID, job.id)
		results <- rand.IntN(100)
	}
}

func main() {
	const numJobs = 10
	const numWorker = 3

	jobs := make(chan *Task, numJobs)
	results := make(chan int, numJobs)

	for w := 1; w <= numWorker; w++ {
		go worker(w, jobs, results)
	}

	for job := 1; job <= numJobs; job++ {
		jobs <- NewTasks(rand.IntN(100))
	}
	close(jobs)

	for a := 1; a <= numJobs; a++ {
		<-results
	}
}

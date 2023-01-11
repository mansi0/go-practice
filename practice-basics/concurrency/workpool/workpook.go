package main

import "fmt"
import "time"

func work(workerId int, tasks <-chan int, result chan<- int ) {
	for task:= range(tasks) {
		fmt.Printf("\nWorker %d has started working on task %d",workerId,task)
		time.Sleep(1*time.Second)
		fmt.Printf("\nWorker %d has worked on task %d",workerId,task)
		result <- task *10
	}
}
func main() {
	startTime := time.Now()
	const noOfTasks = 10
	tasks := make(chan int, noOfTasks)
	result := make(chan int, noOfTasks)
	noOfWorker := 3
	for worker := 1; worker <= noOfWorker; worker++ {
		go work(worker,tasks,result)
	}
	for task := 1; task <= noOfTasks; task++ {
		tasks<-task
	}

	for r := 1; r <= noOfTasks; r++  {
		fmt.Println(<-result)
	}
	endTime := time.Now()
	diff := endTime.Sub(startTime)
	fmt.Println("total time taken ", diff.Seconds(), "seconds")
}
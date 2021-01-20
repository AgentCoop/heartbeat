package main

import (
	"fmt"
	job "github.com/AgentCoop/go-work"
)

func main() {

	j := job.NewJob(nil)
	j.AddTask(func(j job.Job) (job.Init, job.Run, job.Finalize) {
		return nil, func(task job.Task) {
			fmt.Print("Run task\n")
			task.Done()
		}, nil
	})
	<-j.Run()
	fmt.Printf("Hello, World!")
}

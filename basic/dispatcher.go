package basic

import "fmt"

var WorkerList chan chan WorkRequest

func StartDispatcher(nworkers int) {

	WorkerList = make(chan chan WorkRequest, nworkers)

	// Create Workers
	for i:=0; i<nworkers; i++ {
		fmt.Println("Starting worker", i+1)
		worker := New(i+1, WorkerList)
		worker.Start()
	}

	go distributeWork()
}

func distributeWork() {

	for {
		select {
		case work := <-WorkQueue:

			fmt.Printf("Recieved work request: %s", work.Name)

			go func() {
				worker := <-WorkerList
				fmt.Println("Dispatching work request")
				worker <- work
			}()
		}
	}
}

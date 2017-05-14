package basic

import (
	"fmt"
	"time"
)

type Worker struct {
	ID 			int
	Work 		chan WorkRequest
	// Keeps track of all Work channels
	WorkerList chan chan WorkRequest
	QuitChan 	chan bool
}

func New(id int, workerList chan chan WorkRequest) Worker {

	worker := Worker{
		ID:          id,
		Work:        make(chan WorkRequest),
		WorkerList: workerList,
		QuitChan:    make(chan bool),
	}

	return worker
}

func (w *Worker) Start() {
	go func() {
		for {
			// Keep track of channel
			w.WorkerList <- w.Work

			select {
			case work := <-w.Work:
				fmt.Printf("worker%d: Recieved work request, delaying for %f seconds\n", w.ID, work.Name)
				time.Sleep(work.Delay)
				fmt.Printf("worker%d: Hello, %s!\n", w.ID, work.Name)

			case <-w.QuitChan:
				fmt.Printf("worker%d stopping\n", w.ID)
				return
			}
		}
	}()
}

func (w *Worker) Stop() {
	go func() {
		w.QuitChan <- true
	}()
}
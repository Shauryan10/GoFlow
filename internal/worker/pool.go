package worker

import "fmt"

func StartWorkerPool(workerCount int) {

	for i := 1; i <= workerCount; i++ {

		go StartWorker(i)

		fmt.Printf("Worker %d started\n", i)

	}

}

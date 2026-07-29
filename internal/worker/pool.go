package worker

import (
	"fmt"

	"github.com/Shauryan10/GoFlow/internal/models"
)

func StartWorkerPool(workerCount int) {

	for i := 1; i <= workerCount; i++ {

		mu.Lock()

		Workers[i] = &models.Worker{
			ID:            i,
			Status:        "IDLE",
			CurrentTaskID: 0,
		}

		mu.Unlock()

		go StartWorker(i)

		fmt.Printf("Worker %d started\n", i)
	}
}

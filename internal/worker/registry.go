package worker

import (
	"sync"

	"github.com/Shauryan10/GoFlow/internal/models"
)

var (
	Workers = make(map[int]*models.Worker)

	mu sync.RWMutex
)

func GetWorkers() []*models.Worker {

	mu.RLock()
	defer mu.RUnlock()

	workers := make([]*models.Worker, 0, len(Workers))

	for _, worker := range Workers {
		workers = append(workers, worker)
	}

	return workers
}

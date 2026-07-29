package worker

import (
	"sync"

	"github.com/Shauryan10/GoFlow/internal/models"
)

var (
	Workers = make(map[int]*models.Worker)

	mu sync.RWMutex
)

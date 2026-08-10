package worker

import (
	"context"
	"fmt"
	"sync"
)

func StartWorkerPool(
	ctx context.Context,
	count int,
	wg *sync.WaitGroup,
) {

	for i := 1; i <= count; i++ {

		wg.Add(1)

		workerID := i

		go func() {

			defer wg.Done()

			fmt.Printf(
				"Worker %d started\n",
				workerID,
			)

			StartWorker(ctx, workerID)

			fmt.Printf(
				"Worker %d stopped\n",
				workerID,
			)

		}()
	}
}

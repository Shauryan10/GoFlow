package worker

import (
	"fmt"
	"time"

	"github.com/Shauryan10/GoFlow/internal/queue"
)

func StartWorker(id int) {
	for {

		task := <-queue.TaskQueue

		fmt.Printf(
			"Worker %d started task %d (%s)\n",
			id,
			task.ID,
			task.Name,
		)

		time.Sleep(5 * time.Second)

		fmt.Printf(
			"Worker %d completed task %d\n",
			id,
			task.ID,
		)

	}

}

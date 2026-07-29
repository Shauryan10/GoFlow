package worker

import (
	"fmt"
	"time"

	"github.com/Shauryan10/GoFlow/internal/queue"
	"github.com/Shauryan10/GoFlow/internal/service"
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

		err := service.CompleteTask(task.ID)

		if err != nil {
			fmt.Printf(
				"Worker %d failed updating task %d: %v\n",
				id,
				task.ID,
				err,
			)
			continue
		}
		
		fmt.Printf(
			"Worker %d completed task %d\n",
			id,
			task.ID,
		)

	}

}

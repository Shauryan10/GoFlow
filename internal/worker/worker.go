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

		err := service.StartTask(task.ID)

		if err != nil {

			fmt.Printf(
				"Worker %d failed starting task %d : %v\n",
				id,
				task.ID,
				err,
			)

			continue
		}

		mu.Lock()

		Workers[id].Status = "BUSY"
		Workers[id].CurrentTaskID = task.ID

		mu.Unlock()

		fmt.Printf(
			"Worker %d started task %d (%s)\n",
			id,
			task.ID,
			task.Name,
		)

		for progress := uint(10); progress <= 100; progress += 10 {

			err := service.UpdateTaskProgress(task.ID, progress)

			if err != nil {

				fmt.Printf("Failed updating progress\n")

				break
			}

			fmt.Printf(
				"Worker %d | Task %d | %d%% complete\n",
				id,
				task.ID,
				progress,
			)

			time.Sleep(500 * time.Millisecond)
		}

		err = service.CompleteTask(task.ID)

		if err != nil {

			fmt.Printf(
				"Worker %d failed task %d : %v\n",
				id,
				task.ID,
				err,
			)

			mu.Lock()

			Workers[id].Status = "IDLE"
			Workers[id].CurrentTaskID = 0

			mu.Unlock()

			continue
		}

		mu.Lock()

		Workers[id].Status = "IDLE"
		Workers[id].CurrentTaskID = 0

		mu.Unlock()

		fmt.Printf(
			"Worker %d completed task %d\n",
			id,
			task.ID,
		)
	}
}

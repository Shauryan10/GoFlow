package worker

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/Shauryan10/GoFlow/internal/metrics"
	"github.com/Shauryan10/GoFlow/internal/queue"
	"github.com/Shauryan10/GoFlow/internal/service"
)

const MaxRetries = 3

func StartWorker(id int) {

	for {

		task := <-queue.TaskQueue
		metrics.QueueLength.Set(float64(len(queue.TaskQueue)))

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

		metrics.WorkersBusy.Inc()   //increae by 1
		fmt.Printf(
			"Worker %d started task %d (%s)\n",
			id,
			task.ID,
			task.Name,
		)

		success := false

		for attempt := 1; attempt <= MaxRetries; attempt++ {

			fmt.Printf(
				"Worker %d | Starting Attempt %d for Task %d\n",
				id,
				attempt,
				task.ID,
			)

			// Reset progress for this attempt
			service.UpdateTaskProgress(task.ID, 0)

			failed := false

			for progress := uint(10); progress <= 100; progress += 10 {

				err := service.UpdateTaskProgress(task.ID, progress)

				if err != nil {
					fmt.Println(err)
					failed = true
					break
				}

				fmt.Printf(
					"Worker %d | Task %d | %d%% complete\n",
					id,
					task.ID,
					progress,
				)

				time.Sleep(500 * time.Millisecond)

				// 20% chance of failure
				if rand.Intn(100) < 20 {

					fmt.Printf(
						"Worker %d | Task %d failed at %d%%\n",
						id,
						task.ID,
						progress,
					)

					service.IncrementRetry(task.ID)
					metrics.TaskRetries.Inc()

					failed = true

					break
				}
			}

			if !failed {

				success = true

				break
			}

			fmt.Printf(
				"Retrying Task %d...\n",
				task.ID,
			)

			time.Sleep(1 * time.Second)
		}

		if success {

			err = service.CompleteTask(task.ID)

			if err != nil {

				fmt.Println(err)

			} else {

				metrics.TasksProcessed.Inc()
				fmt.Printf(
					"Worker %d completed task %d\n",
					id,
					task.ID,
				)
			}

		} else {

			service.FailTask(
				task.ID,
				"Maximum retry limit exceeded",
			)

			metrics.TasksFailed.Inc()

			fmt.Printf(
				"Worker %d permanently failed task %d\n",
				id,
				task.ID,
			)
		}

		mu.Lock()

		Workers[id].Status = "IDLE"
		Workers[id].CurrentTaskID = 0

		mu.Unlock()

		metrics.WorkersBusy.Dec()   // decrease by 1
		fmt.Printf(
			"Worker %d completed task %d\n",
			id,
			task.ID,
		)
	}
}

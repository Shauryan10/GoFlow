package worker

import (
	"context"
	"fmt"
	"log"
	"math/rand"
	"time"

	"github.com/Shauryan10/GoFlow/internal/metrics"
	"github.com/Shauryan10/GoFlow/internal/models"
	"github.com/Shauryan10/GoFlow/internal/queue"
	"github.com/Shauryan10/GoFlow/internal/service"
)

const MaxRetries = 3

func StartWorker(ctx context.Context, id int) {

	// Make sure this worker exists in the registry.
	mu.Lock()

	if Workers[id] == nil {
		Workers[id] = &models.Worker{
			ID:            id,
			Status:        "IDLE",
			CurrentTaskID: 0,
		}
	}

	mu.Unlock()

	for {

		select {

		case task := <-queue.TaskQueue:

			metrics.QueueLength.Set(
				float64(len(queue.TaskQueue)),
			)

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

			// Worker becomes BUSY
			mu.Lock()

			Workers[id].Status = "BUSY"
			Workers[id].CurrentTaskID = task.ID

			mu.Unlock()

			metrics.WorkersBusy.Inc()

			fmt.Printf(
				"Worker %d started task %d (%s)\n",
				id,
				task.ID,
				task.Name,
			)

			success := false

			// Retry loop
			for attempt := 1; attempt <= MaxRetries; attempt++ {

				fmt.Printf(
					"Worker %d | Starting Attempt %d for Task %d\n",
					id,
					attempt,
					task.ID,
				)

				err = service.UpdateTaskProgress(
					task.ID,
					0,
				)

				if err != nil {

					fmt.Printf(
						"Failed resetting progress for task %d: %v\n",
						task.ID,
						err,
					)

					break
				}

				failed := false

				// Progress loop
				for progress := uint(10); progress <= 100; progress += 10 {

					err := service.UpdateTaskProgress(
						task.ID,
						progress,
					)

					if err != nil {

						fmt.Printf(
							"Failed updating progress for task %d: %v\n",
							task.ID,
							err,
						)

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

						if err := service.IncrementRetry(task.ID); err != nil {
							log.Printf("Failed to increment retry count for task %d: %v", task.ID, err)
						}

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

			// Task completed successfully
			if success {

				err = service.CompleteTask(task.ID)

				if err != nil {

					fmt.Printf(
						"Worker %d failed completing task %d: %v\n",
						id,
						task.ID,
						err,
					)

				} else {

					metrics.TasksProcessed.Inc()

					fmt.Printf(
						"Worker %d completed task %d\n",
						id,
						task.ID,
					)
				}

			} else {

				// Task permanently failed
				err = service.FailTask(
					task.ID,
					"Maximum retry limit exceeded",
				)

				if err != nil {

					fmt.Printf(
						"Worker %d failed updating failed task %d: %v\n",
						id,
						task.ID,
						err,
					)
				}

				metrics.TasksFailed.Inc()

				fmt.Printf(
					"Worker %d permanently failed task %d\n",
					id,
					task.ID,
				)
			}

			// Worker becomes IDLE
			mu.Lock()

			Workers[id].Status = "IDLE"
			Workers[id].CurrentTaskID = 0

			mu.Unlock()

			metrics.WorkersBusy.Dec()

		case <-ctx.Done():

			fmt.Printf(
				"Worker %d shutting down\n",
				id,
			)

			return
		}
	}
}

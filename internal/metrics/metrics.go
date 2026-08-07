package metrics

import "github.com/prometheus/client_golang/prometheus"

var (
	TasksProcessed = prometheus.NewCounter(

		prometheus.CounterOpts{
			Name: "goflow_tasks_processed_total",
			Help: "Total completed tasks",
		},
	)

	TasksFailed = prometheus.NewCounter(

		prometheus.CounterOpts{
			Name: "goflow_tasks_failed_total",
			Help: "Total failed tasks",
		},
	)

	TaskRetries = prometheus.NewCounter(

		prometheus.CounterOpts{
			Name: "goflow_task_retries_total",
			Help: "Total retries",
		},
	)

	WorkersBusy = prometheus.NewGauge(

		prometheus.GaugeOpts{
			Name: "goflow_workers_busy",
			Help: "Current busy workers",
		},
	)

	QueueLength = prometheus.NewGauge(

		prometheus.GaugeOpts{
			Name: "goflow_queue_length",
			Help: "Current queue size",
		},
	)
)

func Register() {

	prometheus.MustRegister(
		TasksProcessed,
		TasksFailed,
		TaskRetries,
		WorkersBusy,
		QueueLength,
	)

}

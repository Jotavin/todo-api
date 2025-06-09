package metrics

import (
	"fmt"

	"github.com/prometheus/client_golang/prometheus"
)


var (

	RequestsTotal = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "todo_api_requests_total",
			Help: "Total number of requests",
		},
		[]string{"endpoint", "method", "status"},
	)

	RequestDuration = prometheus.NewHistogramVec(
		prometheus.HistogramOpts{
			Name: "todo_api_request_duration_seconds",
			Help: "Request duration in seconds",
			Buckets: []float64{0.001, 0.01, 0.1, 0.5, 1.0, 2.5, 5.0, 10.0},
		},
		[]string{"endpoint", "method"}, // ← 2 labels
	)

	TasksCreated = prometheus.NewCounter(
		prometheus.CounterOpts{
			Name: "todo_api_tasks_created_total",
			Help: "Total number of tasks created",
		},
	)

	TasksDeleted = prometheus.NewCounter(
		prometheus.CounterOpts{
			Name: "todo_api_tasks_deleted_total",
			Help: "Total number of tasks deleted",
		},
	)
	
	TasksUpdated = prometheus.NewCounter(
		prometheus.CounterOpts{
			Name: "todo_api_tasks_updated_total",
			Help: "Total number of tasks updated",
		},
	)
	
	TasksQueried = prometheus.NewCounter(
		prometheus.CounterOpts{
			Name: "todo_api_tasks_queried_total",
			Help: "Total number of task queries",
		},
	)
	
	// Métricas de erro com labels
	DatabaseErrors = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "todo_api_database_errors_total",
			Help: "Total number of database errors",
		},
		[]string{"operation", "error_type"}, // ← 2 labels
	)
	
	TasksNotFound = prometheus.NewCounter(
		prometheus.CounterOpts{
			Name: "todo_api_tasks_not_found_total",
			Help: "Total number of tasks not found (404 errors)",
		},
	)

	MigrationsExecuted = prometheus.NewCounter(
		prometheus.CounterOpts{
			Name: "todo_api_migrations_executed_total",
			Help: "Total number of migrations executed",
		},
	)
)

func init(){
	prometheus.MustRegister(RequestsTotal)
	prometheus.MustRegister(RequestDuration)
	prometheus.MustRegister(TasksCreated)
	prometheus.MustRegister(TasksDeleted)
	prometheus.MustRegister(TasksUpdated)
	prometheus.MustRegister(TasksQueried)
	prometheus.MustRegister(DatabaseErrors)
	prometheus.MustRegister(TasksNotFound)
	prometheus.MustRegister(MigrationsExecuted)

	fmt.Println("✅ Métricas registradas com sucesso!")

}
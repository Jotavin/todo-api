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

	// Migrations executadas

	MigrationsExecuted = prometheus.NewCounter(
		prometheus.CounterOpts{
			Name: "todo_api_migrations_executed_total",
			Help: "Total number of migrations executed",
		},
	)
	// Conexões com banco
	DatabaseConnectionAttempts = prometheus.NewCounter(
		prometheus.CounterOpts{
			Name: "todo_api_database_connection_attempts_total",
			Help: "Total number of database connection attempts",
		},
	)
	
	DatabaseConnectionFailures = prometheus.NewCounter(
		prometheus.CounterOpts{
			Name: "todo_api_database_connection_failures_total",
			Help: "Total number of database connection failures",
		},
	)
	
	DatabaseConnectionsActive = prometheus.NewGauge(
		prometheus.GaugeOpts{
			Name: "todo_api_database_connections_active",
			Help: "Number of active database connections",
		},
	)
	
	DatabaseConnectionDuration = prometheus.NewHistogram(
		prometheus.HistogramOpts{
			Name:    "todo_api_database_connection_duration_seconds",
			Help:    "Time spent connecting to database",
			Buckets: []float64{0.001, 0.01, 0.1, 0.5, 1.0, 2.5, 5.0},
		},
	)
	
	// Queries do banco
	DatabaseQueries = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "todo_api_database_queries_total",
			Help: "Total number of database queries",
		},
		[]string{"operation"}, // create, select, update, delete
	)
	
	DatabaseQueryDuration = prometheus.NewHistogramVec(
		prometheus.HistogramOpts{
			Name:    "todo_api_database_query_duration_seconds",
			Help:    "Database query duration",
			Buckets: []float64{0.001, 0.01, 0.1, 0.5, 1.0, 2.5, 5.0},
		},
		[]string{"operation"},
	)
	
	// Migrations
	DatabaseMigrationAttempts = prometheus.NewCounter(
		prometheus.CounterOpts{
			Name: "todo_api_database_migration_attempts_total",
			Help: "Total number of migration attempts",
		},
	)
	
	DatabaseMigrationFailures = prometheus.NewCounter(
		prometheus.CounterOpts{
			Name: "todo_api_database_migration_failures_total",
			Help: "Total number of migration failures",
		},
	)
	
	DatabaseMigrationDuration = prometheus.NewHistogram(
		prometheus.HistogramOpts{
			Name:    "todo_api_database_migration_duration_seconds",
			Help:    "Time spent on database migrations",
			Buckets: []float64{0.1, 0.5, 1.0, 5.0, 10.0, 30.0, 60.0},
		},
	)
	
	// Pool de conexões
	DatabasePoolConnectionsOpen = prometheus.NewGauge(
		prometheus.GaugeOpts{
			Name: "todo_api_database_pool_connections_open",
			Help: "Number of open connections in the pool",
		},
	)
	
	DatabasePoolConnectionsInUse = prometheus.NewGauge(
		prometheus.GaugeOpts{
			Name: "todo_api_database_pool_connections_in_use",
			Help: "Number of connections currently in use",
		},
	)
	
	DatabasePoolConnectionsIdle = prometheus.NewGauge(
		prometheus.GaugeOpts{
			Name: "todo_api_database_pool_connections_idle",
			Help: "Number of idle connections in the pool",
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
	
	// Database metrics
	prometheus.MustRegister(DatabaseConnectionAttempts)
	prometheus.MustRegister(DatabaseConnectionFailures)
	prometheus.MustRegister(DatabaseConnectionsActive)
	prometheus.MustRegister(DatabaseConnectionDuration)
	prometheus.MustRegister(DatabaseQueries)
	prometheus.MustRegister(DatabaseQueryDuration)
	prometheus.MustRegister(DatabaseMigrationAttempts)
	prometheus.MustRegister(DatabaseMigrationFailures)
	prometheus.MustRegister(DatabaseMigrationDuration)
	prometheus.MustRegister(DatabasePoolConnectionsOpen)
	prometheus.MustRegister(DatabasePoolConnectionsInUse)
	prometheus.MustRegister(DatabasePoolConnectionsIdle)
	fmt.Println("✅ Metrics registrad successfully!")

}
package main

import (
	"context"
	"flag"
	"fmt"
	"log/slog"
	"math/rand"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

// Config struct contains configuration variables for the server.
type Config struct {
	serverPort         string  // Port for server to listen on
	minSleepDurationMs int     // Minimum sleep duration in milliseconds to simulate work
	maxSleepDurationMs int     // Maximum sleep duration in milliseconds to simulate work
	errorRate          float64 // Rate of simulated errors
	name               string  // Name used in greeting
}

// Initialising default configuration values.
var config = Config{
	serverPort:         "5000",
	minSleepDurationMs: 0,
	maxSleepDurationMs: 0,
	errorRate:          0.0,
	name:               "Brave",
}

// Initialising Prometheus metrics.
var (
	requestsTotal    = promauto.NewCounterVec(prometheus.CounterOpts{Name: "requests_total", Help: "Total number of requests by status code."}, []string{"code", "method"})
	requestDuration  = promauto.NewHistogramVec(prometheus.HistogramOpts{Name: "request_duration_seconds", Help: "Duration of HTTP requests in seconds", Buckets: prometheus.DefBuckets}, []string{"method"})
	inFlightRequests = promauto.NewGauge(prometheus.GaugeOpts{Name: "in_flight_requests", Help: "Number of in-flight HTTP requests."})
)

// loadEnvConfig loads configuration from environment variables with defaults as fallback.
func loadEnvConfig() (Config, error) {
	var config Config

	name := getEnvOrDefault("NAME", "Brave")
	serverPort := getEnvOrDefault("SERVER_PORT", "5000")
	minSleep := getEnvOrDefault("MIN_SLEEP_MS", "0")
	maxSleep := getEnvOrDefault("MAX_SLEEP_MS", "0")
	errorRate := getEnvOrDefault("ERROR_RATE", "0")

	config.name = name
	config.serverPort = serverPort

	var err error
	config.minSleepDurationMs, err = parseToInt(minSleep, "MIN_SLEEP_MS")
	if err != nil {
		return Config{}, err
	}

	config.maxSleepDurationMs, err = parseToInt(maxSleep, "MAX_SLEEP_MS")
	if err != nil {
		return Config{}, err
	}

	config.errorRate, err = parseToFloat(errorRate, "ERROR_RATE")
	if err != nil {
		return Config{}, err
	}

	return config, nil
}

// getEnvOrDefault retrieves the value of an environment variable or a default value if the variable is not set.
func getEnvOrDefault(key, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultValue
}

// parseToInt converts string to integer with error handling.
func parseToInt(value string, envVar string) (int, error) {
	parsed, err := strconv.Atoi(value)
	if err != nil {
		slog.LogAttrs(context.Background(), slog.LevelError, "Error converting to integer", slog.String(envVar, value), slog.String("error", err.Error()))
		return 0, err
	}
	return parsed, nil
}

// parseToFloat converts string to float64 with error handling.
func parseToFloat(value string, envVar string) (float64, error) {
	parsed, err := strconv.ParseFloat(value, 64)
	if err != nil {
		slog.LogAttrs(context.Background(), slog.LevelError, "Error converting to float64", slog.String(envVar, value), slog.String("error", err.Error()))
		return 0.0, err
	}
	return parsed, nil
}

// main function initialises and starts the server.
func main() {
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	slog.SetDefault(logger)

	envConf, err := loadEnvConfig()
	if err != nil {
		slog.LogAttrs(context.Background(), slog.LevelError, "Error loading environment variables", slog.String("error", err.Error()))
		os.Exit(1)
	}

	config = envConf

	flag.StringVar(&config.serverPort, "port", config.serverPort, "Port on which the server runs")
	flag.IntVar(&config.minSleepDurationMs, "min-sleep", config.minSleepDurationMs, "Min sleep duration in milliseconds to simulate work")
	flag.IntVar(&config.maxSleepDurationMs, "max-sleep", config.maxSleepDurationMs, "Max sleep duration in milliseconds to simulate work")
	flag.Float64Var(&config.errorRate, "error-rate", config.errorRate, "Error simulation rate")
	flag.StringVar(&config.name, "name", config.name, "Name to be used in greeting")

	flag.Parse()

	config.serverPort = ":" + config.serverPort

	rand.Seed(time.Now().UnixNano())
	slog.LogAttrs(context.Background(), slog.LevelInfo, "Starting server", slog.String("port", config.serverPort), slog.Int("min-sleep", config.minSleepDurationMs), slog.Int("max-sleep", config.maxSleepDurationMs), slog.Float64("error-rate", config.errorRate), slog.String("name", config.name))

	http.HandleFunc("/", greet)
	http.Handle("/metrics", promhttp.Handler())
	http.HandleFunc("/healthz", healthz)

	err = http.ListenAndServe(config.serverPort, nil)
	if err != nil {
		slog.LogAttrs(context.Background(), slog.LevelError, "Error starting server", slog.String("error", err.Error()))
		os.Exit(1)
	}
}

// greet function handles incoming requests and sends responses.
func greet(w http.ResponseWriter, r *http.Request) {
	inFlightRequests.Inc()
	defer inFlightRequests.Dec()

	start := time.Now()

	delay := time.Duration(config.minSleepDurationMs) * time.Millisecond
	if config.maxSleepDurationMs > config.minSleepDurationMs {
		delay += time.Duration(rand.Intn(config.maxSleepDurationMs-config.minSleepDurationMs)) * time.Millisecond
	}
	time.Sleep(delay)

	if rand.Float64() < config.errorRate {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(w, "Internal Server Error")
		requestsTotal.WithLabelValues(fmt.Sprint(http.StatusInternalServerError), r.Method).Inc()
		slog.LogAttrs(context.Background(), slog.LevelError, "Received request", slog.String("method", r.Method), slog.String("status", fmt.Sprint(http.StatusInternalServerError)))
		return
	}

	w.Header().Set("Content-Type", "text/plain")
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Hello, World! Greetings from %s.", config.name)

	duration := time.Since(start).Seconds()
	requestDuration.WithLabelValues(r.Method).Observe(duration)
	requestsTotal.WithLabelValues(fmt.Sprint(http.StatusOK), r.Method).Inc()
	slog.LogAttrs(context.Background(), slog.LevelInfo, "Received request", slog.String("method", r.Method), slog.Float64("duration", duration), slog.String("status", fmt.Sprint(http.StatusOK)))
}

// healthz is a health check endpoint that responds with 'OK' status.
func healthz(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, "OK")
}

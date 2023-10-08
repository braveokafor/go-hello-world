package main

import (
	"net/http"
	"net/http/httptest"
	"os"
	"reflect"
	"testing"
)

func Test_loadEnvConfig(t *testing.T) {
	tests := []struct {
		envVars map[string]string
		want    Config
		wantErr bool
	}{
		{
			envVars: map[string]string{
				"SERVER_PORT":  "6000",
				"MIN_SLEEP_MS": "100",
				"MAX_SLEEP_MS": "200",
				"ERROR_RATE":   "0.5",
				"NAME":         "TestName",
			},
			want: Config{
				serverPort:         ":6000",
				minSleepDurationMs: 100,
				maxSleepDurationMs: 200,
				errorRate:          0.5,
				name:               "TestName",
			},
			wantErr: false,
		},
		{
			envVars: map[string]string{
				"SERVER_PORT":  "invalid",
				"MIN_SLEEP_MS": "invalid",
				"MAX_SLEEP_MS": "invalid",
				"ERROR_RATE":   "invalid",
				"NAME":         "TestName",
			},
			want:    Config{},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			for key, value := range tt.envVars {
				os.Setenv(key, value)
				defer os.Unsetenv(key)
			}

			got, err := loadEnvConfig()
			if (err != nil) != tt.wantErr {
				t.Errorf("loadEnvConfig() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("loadEnvConfig() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_getEnvOrDefault(t *testing.T) {
	tests := []struct {
		name string
		key  string
		def  string
		want string
	}{
		{"Var exists", "EXISTING_VAR", "default", "existing_value"},
		{"Var does not exist", "NON_EXISTING_VAR", "default", "default"},
	}

	os.Setenv("EXISTING_VAR", "existing_value")
	defer os.Unsetenv("EXISTING_VAR")

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getEnvOrDefault(tt.key, tt.def); got != tt.want {
				t.Errorf("getEnvOrDefault() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_parseToInt(t *testing.T) {
	tests := []struct {
		name    string
		value   string
		envVar  string
		want    int
		wantErr bool
	}{
		{"Valid integer", "123", "TEST_INT", 123, false},
		{"Invalid integer", "abc", "TEST_INT", 0, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := parseToInt(tt.value, tt.envVar)
			if (err != nil) != tt.wantErr {
				t.Errorf("parseToInt() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("parseToInt() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_parseToFloat(t *testing.T) {
	tests := []struct {
		name    string
		value   string
		envVar  string
		want    float64
		wantErr bool
	}{
		{"Valid float", "123.45", "TEST_FLOAT", 123.45, false},
		{"Invalid float", "abc", "TEST_FLOAT", 0, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := parseToFloat(tt.value, tt.envVar)
			if (err != nil) != tt.wantErr {
				t.Errorf("parseToFloat() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("parseToFloat() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_greet(t *testing.T) {
	tests := []struct {
		name       string
		errorRate  float64
		wantStatus int
		wantBody   string
	}{
		{"normal case", 0, http.StatusOK, "Hello, World! Greetings from Brave."},
		{"error case", 1, http.StatusInternalServerError, "Internal Server Error"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			config.errorRate = tt.errorRate
			rr := httptest.NewRecorder()
			req, _ := http.NewRequest("GET", "/", nil)

			greet(rr, req)

			if status := rr.Code; status != tt.wantStatus {
				t.Errorf("handler returned wrong status code: got %v want %v", status, tt.wantStatus)
			}

			expected := tt.wantBody
			if rr.Body.String() != expected {
				t.Errorf("handler returned unexpected body: got %v want %v", rr.Body.String(), expected)
			}
		})
	}
}

func Test_healthz(t *testing.T) {
	tests := []struct {
		name string
		want string
	}{
		{"normal case", "OK"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rr := httptest.NewRecorder()
			req, _ := http.NewRequest("GET", "/healthz", nil)

			healthz(rr, req)

			if status := rr.Code; status != http.StatusOK {
				t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
			}

			expected := tt.want
			if rr.Body.String() != expected {
				t.Errorf("handler returned unexpected body: got %v want %v", rr.Body.String(), expected)
			}
		})
	}
}

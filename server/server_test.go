package server

import (
	"fmt"
	"log"
	"net/http"
	"testing"
	"time"

	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
)

func TestNew(t *testing.T) {
	var tests = []struct {
		name  string
		input Options
		want  *server
	}{
		{
			name: "default options",
			want: &server{
				httpServer: &http.Server{
					Addr:         ":8080",
					Handler:      http.NewServeMux(),
					ReadTimeout:  15 * time.Second,
					WriteTimeout: 15 * time.Second,
					IdleTimeout:  30 * time.Second,
				},
				router: http.NewServeMux(),
				log:    log.Default(),
			},
		},
		{
			name: "with options",
			input: Options{
				Router: http.NewServeMux(),
				Log:    mockLogger{},
			},
			want: &server{
				httpServer: &http.Server{
					Addr:         ":8080",
					Handler:      http.NewServeMux(),
					ReadTimeout:  15 * time.Second,
					WriteTimeout: 15 * time.Second,
					IdleTimeout:  30 * time.Second,
				},
				router: http.NewServeMux(),
				log:    mockLogger{},
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got := New(test.input)

			if diff := cmp.Diff(test.want, got, cmp.AllowUnexported(server{}, mockLogger{}), cmpopts.IgnoreUnexported(http.Server{}, http.ServeMux{}, log.Logger{})); diff != "" {
				t.Errorf("New() = unexpected result (-want +got):\n%s\n", diff)
			}
		})
	}
}

var testLogs = []string{}

type mockLogger struct{}

func (m mockLogger) Printf(format string, v ...any) {
	testLogs = append(testLogs, fmt.Sprintf(format, v...))
}

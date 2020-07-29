package main

import (
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httptest"
	"net/url"
	"os"
	"os/signal"
	"syscall"
	"testing"
	"time"
)

func TestTodos(t *testing.T) {
	tests := []struct {
		name   string
		r      *http.Request
		method string
		path   string
		want   string
	}{
		{
			name:   "Get from /foo/bar",
			method: "GET",
			path:   "/foo/bar",
			want:   `[{"id":42,"title":"MyTask","completed":false,"url":"","order":1}]`,
		},
	}
	for i := range tests {
		tt := tests[i]
		t.Run(tt.name, func(t *testing.T) {
			r := &http.Request{
				Method: tt.method,
				URL: &url.URL{
					Path: tt.path,
				},
			}
			rr := httptest.NewRecorder()
			ListTodos(rr, r)
			if rr.Code != http.StatusOK {
				t.Errorf("Echo did not get correct status, got %d, want %d", rr.Code, http.StatusOK)
			}
			if got := rr.Body.String(); got != tt.want {
				t.Errorf("Echo response body is wrong, got %q want %q", got, tt.want)
			}
		})
	}
}

func TestWaiter(t *testing.T) {
	t.Run("Signal runserver graceful shutdown", func(t *testing.T) {
		var finished bool
		// Get the operating system process
		proc, err := os.FindProcess(os.Getpid())
		if err != nil {
			t.Fatal(err)
		}
		// Discard noisy logs
		logger := log.New(ioutil.Discard, "", log.LstdFlags)
		go func() {
			runServer(logger)
			finished = true
		}()

		if finished {
			t.Error("runServer Exit before signal sent")
		}

		// if we signal too early, Waiter isn't listening yet
		time.Sleep(10 * time.Millisecond)
		// Send the SIGQUIT
		proc.Signal(syscall.SIGQUIT)
		// if we test finished too early, finished may not have been updated yet
		time.Sleep(10 * time.Millisecond)
		// reset signal notification
		signal.Reset()
		if !finished {
			t.Error("runServer Did Not Exit")
		}
	})
}

func TestHealthCheckHandler(t *testing.T) {
	// Create a request to pass to our handler. We don't have any query parameters for now, so we'll
	// pass 'nil' as the third parameter.
	req, err := http.NewRequest("GET", "/health", nil)
	if err != nil {
		t.Fatal(err)
	}

	// We create a ResponseRecorder (which satisfies http.ResponseWriter) to record the response.
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(ListTodos)

	// Our handlers satisfy http.Handler, so we can call their ServeHTTP method
	// directly and pass in our Request and ResponseRecorder.
	handler.ServeHTTP(rr, req)

	// Check the status code is what we expect.
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	want := "application/json; charset=utf-8"
	if contentType := rr.Header().Get("Content-Type"); contentType != want {
		t.Errorf("handler returned wrong status code: got %v want %v", contentType, want)
	}
}

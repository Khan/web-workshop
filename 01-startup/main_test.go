package main

import (
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
)

func TestEcho(t *testing.T) {
	tests := []struct {
		name   string
		method string
		path   string
		want   string
	}{
		{
			name:   "No specific path",
			method: "GET",
			path:   "/",
			want:   "You asked to GET /\n",
		},
		{
			name:   "Posting to /foo/bar",
			method: "POST",
			path:   "/foo/bar",
			want:   "You asked to POST /foo/bar\n",
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
			Echo(rr, r)
			if rr.Code != http.StatusOK {
				t.Errorf("Echo did not get correct status, got %d, want %d", rr.Code, http.StatusOK)
			}
			if got := rr.Body.String(); got != tt.want {
				t.Errorf("Echo response body is wrong, got %q want %q", got, tt.want)
			}
		})
	}
}

package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"reflect"
	"strings"
	"testing"
)

type handlerFunc func(http.ResponseWriter, *http.Request)

func TestAllHandlers(t *testing.T) {
	tests := map[string]struct {
		handler handlerFunc
		method  string
		path    string
		code    int
		want    interface{}
	}{
		"root happy pass": {
			handler: helloHandler,
			method:  http.MethodGet,
			path:    "/",
			code:    http.StatusOK,
			want:    "Hello World",
		},
		"root bad method": {
			handler: helloHandler,
			method:  http.MethodPut,
			path:    "/",
			code:    http.StatusMethodNotAllowed,
			want:    "PUT is not supported for this endpoint",
		},
		"root bad path": {
			handler: helloHandler,
			method:  http.MethodGet,
			path:    "/wrongPath",
			code:    http.StatusNotFound,
			want:    "Not Found",
		},
		"health happy path": {
			handler: healthHandler,
			method:  http.MethodGet,
			path:    "/health",
			code:    http.StatusOK,
			want:    "200",
		},
		"health bad method": {
			handler: healthHandler,
			method:  http.MethodPost,
			path:    "/health",
			code:    http.StatusMethodNotAllowed,
			want:    "POST is not supported for this endpoint",
		},
		"metadata happy pass": {
			handler: metadataHandler,
			method:  http.MethodGet,
			path:    "/metadata",
			code:    http.StatusOK,
			want: MyApplication{
				[]Metadata{{
					Version:       "HEAD",
					Description:   "pre-interview technical test",
					LastCommitSha: "noData",
				}},
			},
		},
	}

	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			req, err := http.NewRequest(tt.method, tt.path, nil)
			if err != nil {
				t.Fatalf("failed to create request: %s", err)
			}

			w := httptest.NewRecorder()
			h := http.HandlerFunc(tt.handler)
			h.ServeHTTP(w, req)

			//determine the response type
			switch d := tt.want.(type) {
			case string:
				if !strings.Contains(w.Body.String(), d) {
					t.Errorf("handler returned unexpected body: got %v want %v",
						d, tt.want)
				}
			case MyApplication:
				got := MyApplication{}
				err := json.Unmarshal(w.Body.Bytes(), &got)
				if err != nil {
					t.Fatalf("unable to parse response from server %q into slice of Metadata, '%v'", got, err)
				}
				if !reflect.DeepEqual(got, d) {
					t.Errorf("got %q, want %q", got, tt.want)
				}
			default:
				t.Fatalf("missing test implementation for type: %T", tt.want)
			}
			//check for the expected response code
			if status := w.Code; status != tt.code {
				t.Errorf("handler returned wrong status code: got %v want %v",
					status, tt.code)
			}
		})
	}
}

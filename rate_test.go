package ratelmt

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

var chain = Mw(1, reply())
var testTable = []struct {
	name         string
	status       int
	responseBody []byte
}{
	{
		"pass",
		200,
		[]byte("done"),
	},
	{
		"denied",
		429,
		[]byte(http.StatusText(429)),
	},
}

func TestMw(t *testing.T) {
	for _, tt := range testTable {
		t.Run(tt.name, func(t *testing.T) {
			r := httptest.NewRequest("GET", "/", nil)
			w := httptest.NewRecorder()
			chain(w, r)
			res := w.Result()
			body, _ := ioutil.ReadAll(res.Body)
			if res.StatusCode != tt.status {
				t.Errorf("expected %d, got %d", tt.status, res.StatusCode)
			}
			if !bytes.Equal(bytes.TrimSpace(body), tt.responseBody) {
				t.Errorf("expected %s, got %s", tt.responseBody, bytes.TrimSpace(body))
			}
		})
	}
}

func reply() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("done"))
	}
}

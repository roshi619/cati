package pushsafer

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestSend(t *testing.T) {
	n := CATIfication{
		Title:   "title",
		Message: "mesg",
		Key:     "key",
		Client:  &http.Client{Timeout: 3 * time.Second},
	}
	var mockResp apiResponse
	var hitServer bool

	ts := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		hitServer = true

		if r.Method != "POST" {
			t.Error("HTTP method should be POST")
		}

		if r.FormValue("k") == "" {
			t.Error("missing private or alias key")
		}

		json.NewEncoder(rw).Encode(mockResp)
	}))
	defer ts.Close()

	API = ts.URL
	mockResp.Status = 1 // success
	if err := n.Send(); err != nil {
		t.Error(err)
	}

	if !hitServer {
		t.Error("didn't reach server")
	}

	mockResp.Status = 0 // failure
	mockResp.Errors = []string{"error fu"}
	if err := n.Send(); err == nil {
		t.Error("unexpected success")
	}

	mockResp.Status = 0 // failure
	mockResp.Info = "no active devices to send to"
	if err := n.Send(); err == nil {
		t.Error("unexpected success")
	}
}

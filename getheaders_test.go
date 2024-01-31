package getheaders

import (
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

func TestDoRequest(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Server", "TestGetHeaders")
	}))
	defer server.Close()

	expected := make(headerMap)
	expected["Server"] = "TestGetHeaders"

	result, err := DoRequest(http.MethodGet, server.URL)
	if err != nil {
		t.Fatal("Failed to get headers from DoRequest()")
		os.Exit(1)
	}

	_, ok := result["Server"]
	if !ok {
		t.Errorf("got %v\nwant %v", result, expected)
	}
}

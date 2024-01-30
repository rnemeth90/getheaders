package getheaders

import (
	"fmt"
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
		fmt.Fprintf(os.Stderr, err.Error())
		os.Exit(1)
	}

	_, ok := result["Server"]
	if !ok {
		t.Errorf("got %v\nwant %v", result, expected)
	}
}

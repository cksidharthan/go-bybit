package testdata

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"os"
)

func FakeBybitServer(resourcePath, jsonFileName string, statusCode int) *httptest.Server {
	httpServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.String() != resourcePath {
			fmt.Printf("expected to request '%s', got: %s\n", resourcePath, r.URL.String())
			w.WriteHeader(http.StatusNotFound)
		} else {
			w.WriteHeader(statusCode)
		}
		jsonResponse, err := jsonFileReader(jsonFileName)
		if err != nil {
			fmt.Printf("failed to read json file: %s\n", err)
		}
		_, _ = w.Write(jsonResponse)
	}))
	return httpServer
}

func jsonFileReader(filePath string) ([]byte, error) {
	data, err := os.ReadFile(filePath)
	if err != nil {
		return nil, err
	}
	return data, nil
}

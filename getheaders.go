package getheaders

import (
	"net/http"
	"strings"
)

type headerMap map[string]string

func DoRequest(method, url string) (headerMap, error) {
	headers := make(headerMap)

	client := &http.Client{}

	request, err := http.NewRequest(method, url, nil)
	if err != nil {
		return nil, err
	}

	response, err := client.Do(request)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	for key, val := range response.Header {
		headers[key] = strings.Join(val, ",")
	}

	return headers, err
}

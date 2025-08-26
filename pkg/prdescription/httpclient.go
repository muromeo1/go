package prdescription

import (
	"bytes"
	"encoding/json"
	"io"
	"log"
	"net/http"
)

type HttpClient struct {
	Url     string
	Headers map[string]string
}

func NewHttpClient(url string, headers map[string]string) *HttpClient {
	return &HttpClient{
		Url:     url,
		Headers: headers,
	}
}

func (client *HttpClient) Post(body map[string]string) []byte {
	result, err := json.Marshal(body)
	handleError("Marshal error: ", err)

	httpClient := &http.Client{}

	r, err := http.NewRequest("POST", client.Url, bytes.NewBuffer(result))

	for k, v := range client.Headers {
		r.Header.Add(k, v)
	}

	res, err := httpClient.Do(r)
	handleError("Do failed: ", err)

	defer res.Body.Close()

	data, err := io.ReadAll(res.Body)
	handleError("ReadAll failed: ", err)

	return data
}

func handleError(msg string, err error) {
	if err != nil {
		log.Fatal(msg, err)
	}
}

package scoro

import (
	"bytes"
	"io/ioutil"
	"log"
	"net/http"
)

// HTTPClient The API connection data
type HTTPClient struct {
	BaseURL       string
	HTTPClient    *http.Client
	customHeaders map[string]string
}

// MakeGETRequest Method to make http GET request
func (c *HTTPClient) MakeGETRequest(Path string) string {
	response, err := c.HTTPClient.Get(c.BaseURL + "/" + Path)
	if err != nil {
		panic(err)
	}

	defer response.Body.Close()
	if response.StatusCode == http.StatusOK {
		bodyBytes, err := ioutil.ReadAll(response.Body)
		if err != nil {
			log.Fatal(err)
		} else {
			return string(bodyBytes)
		}
	}

	return ""
}

// MakePOSTRequest Method to make http GET request
func (c *HTTPClient) MakePOSTRequest(Path string, Data []byte) []byte {
	req, err := http.NewRequest("POST", c.BaseURL+"/"+Path, bytes.NewBuffer(Data))
	if err != nil {
		log.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/json")

	for key, value := range c.customHeaders {
		req.Header.Set(key, value)
	}

	response, err := c.HTTPClient.Do(req)

	if err != nil {
		panic(err)
	}

	defer response.Body.Close()
	bodyBytes, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}

	return bodyBytes
}

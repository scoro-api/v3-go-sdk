package scoro

import (
	"bytes"
	"io/ioutil"
	"log"
	"net/http"
)

//HTTPClient The API connection data
type HTTPClient struct {
	BaseURL    string
	HTTPClient *http.Client
}

//MakeGETRequest Method to make http GET request
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

//MakePOSTRequest Method to make http GET request
func (c *HTTPClient) MakePOSTRequest(Path string, Data []byte) string {

	response, err := c.HTTPClient.Post(c.BaseURL+"/"+Path, "application/json", bytes.NewBuffer(Data))
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

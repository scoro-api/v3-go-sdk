package scoro

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

const apiVersion = "v3"

// APIClient The API connection data
type APIClient struct {
	config        ApiConfig
	httpClient    *http.Client
	customHeaders map[string]string
}

type apiRequestData struct {
	Request map[string]interface{} `json:"request,omitempty"`
	Filter  map[string]interface{} `json:"filter,omitempty"`
}

func (d *apiRequestData) covertToByteArray() []byte {
	bytes, _ := json.Marshal(d)
	return bytes
}

// AddCustomHeader Add custom http header to the request
func (c *APIClient) AddCustomHeader(key string, value string) {
	c.customHeaders[key] = value
}

// View object
func (c *APIClient) View(Endpoint string, id int) []byte {
	httpClient := HTTPClient{c.config.siteUrl + "/api/" + apiVersion, c.httpClient, c.customHeaders}
	request := httpClient.MakePOSTRequest(Endpoint+"/view/"+strconv.Itoa(id), nil)
	return request
}

// List request using api client
func (c *APIClient) List(endpoint string, Filters []byte) []byte {
	var filterData map[string]interface{}
	json.Unmarshal(Filters, &filterData)

	requestData := apiRequestData{
		Filter: filterData,
	}

	httpClient := HTTPClient{c.config.siteUrl + "/api/" + apiVersion, c.httpClient, c.customHeaders}
	request := httpClient.MakePOSTRequest(endpoint+"/list", requestData.covertToByteArray())
	return request
}

// Create request using api client
func (c *APIClient) Create(endpoint string, Data []byte) []byte {
	return c.Update(endpoint, 0, Data)
}

// Update request using api client
func (c *APIClient) Update(endpoint string, id int, Data []byte) []byte {
	var objectData map[string]interface{}
	json.Unmarshal(Data, &objectData)

	requestData := apiRequestData{
		Request: objectData,
	}

	httpClient := HTTPClient{c.config.siteUrl + "/api/" + apiVersion, c.httpClient, c.customHeaders}
	request := httpClient.MakePOSTRequest(endpoint+"/modify/"+strconv.Itoa(id), requestData.covertToByteArray())
	return request
}

// Delete object
func (c *APIClient) Delete(Endpoint string, id int) []byte {
	httpClient := HTTPClient{c.config.siteUrl + "/api/" + apiVersion, c.httpClient, c.customHeaders}
	request := httpClient.MakePOSTRequest(Endpoint+"/delete/"+strconv.Itoa(id), nil)
	return request
}

func (c *APIClient) makeScoroAPIRequest(name string, path string, Request map[string]string) []byte {
	emp := make(map[string]interface{})
	emp[name] = Request

	// Marshal the map into a JSON string.
	empData, err := json.Marshal(emp)
	if err != nil {
		fmt.Println(err.Error())
		return []byte{}
	}

	httpClient := HTTPClient{c.config.siteUrl + "/api/" + apiVersion, c.httpClient, c.customHeaders}
	return httpClient.MakePOSTRequest(path, empData)
}

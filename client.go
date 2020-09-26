package scoro

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"v3-go-sdk/models"
)

const apiVersion = "v3"

//APIClient The API connection data
type APIClient struct {
	config     ApiConfig
	httpClient *http.Client
}

//List request using api client
func (c *APIClient) List(Endpoint string, Filters map[string]string) string {
	return c.makeScoroAPIRequest("filter", Endpoint, Filters)
}

//Create object using api client
func (c *APIClient) Create(Endpoint string, Request map[string]string) string {
	return c.makeScoroAPIRequest("request", Endpoint+"/modify/", Request)
}

//Update request using api client
func (c *APIClient) Update(Endpoint string, id int, Request map[string]string) string {
	return c.makeScoroAPIRequest("request", Endpoint+"/modify/"+strconv.Itoa(id), Request)
}

//View object
func (c *APIClient) View(model models.ApiModel) models.ApiModel {
	httpClient := HTTPClient{c.config.siteUrl + "/api/" + apiVersion, c.httpClient}
	response := httpClient.MakePOSTRequest(model.Endpoint()+"/view/"+strconv.Itoa(model.Id()), nil)

	err := model.InitFromJSON([]byte(response))

	if err == nil {
		return model
	}
	return model
}

// Deprecated
//ViewLegacy object
func (c *APIClient) ViewLegacy(Endpoint string, id int) string {
	httpClient := HTTPClient{c.config.siteUrl + "/api/" + apiVersion, c.httpClient}
	return httpClient.MakePOSTRequest(Endpoint+"/view/"+strconv.Itoa(id), nil)
}

//Delete object
func (c *APIClient) Delete(model models.ApiModel) string {
	httpClient := HTTPClient{c.config.siteUrl + "/api/" + apiVersion, c.httpClient}
	return httpClient.MakePOSTRequest(model.Endpoint()+"/delete/"+strconv.Itoa(model.Id()), nil)
}

// Deprecated
//Delete object
func (c *APIClient) DeleteLegacy(Endpoint string, id int) string {
	httpClient := HTTPClient{c.config.siteUrl + "/api/" + apiVersion, c.httpClient}
	return httpClient.MakePOSTRequest(Endpoint+"/delete/"+strconv.Itoa(id), nil)
}

func (c *APIClient) makeScoroAPIRequest(name string, path string, Request map[string]string) string {
	emp := make(map[string]interface{})
	emp[name] = Request

	// Marshal the map into a JSON string.
	empData, err := json.Marshal(emp)
	if err != nil {
		fmt.Println(err.Error())
		return ""
	}

	httpClient := HTTPClient{c.config.siteUrl + "/api/" + apiVersion, c.httpClient}
	return httpClient.MakePOSTRequest(path, empData)
}

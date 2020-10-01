package scoro

import (
	"encoding/json"
	"fmt"
	"net/http"
)

const apiVersion = "v3"

//APIClient The API connection data
type APIClient struct {
	config     ApiConfig
	httpClient *http.Client
}

type apiRequestData struct {
	Request map[string]interface{} `json:"request,omitempty"`
	Filter  map[string]interface{} `json:"filter,omitempty"`
}

func (d *apiRequestData) covertToByteArray() []byte {
	bytes, _ := json.Marshal(d)
	return bytes
}

////Create object using api client
//func (c *APIClient) Create(Endpoint string, Request map[string]string) string {
//	return c.makeScoroAPIRequest("request", Endpoint+"/modify/", Request)
//}
//
////Update request using api client
//func (c *APIClient) Update(Endpoint string, id int, Request map[string]string) string {
//	return c.makeScoroAPIRequest("request", Endpoint+"/modify/"+strconv.Itoa(id), Request)
//}

//View object
func (c *APIClient) View(path string) []byte {
	httpClient := HTTPClient{c.config.siteUrl + "/api/" + apiVersion, c.httpClient}
	request := httpClient.MakePOSTRequest(path, nil)
	return request
}

//List request using api client
func (c *APIClient) List(path string, Filters []byte) []byte {
	var result map[string]interface{}
	json.Unmarshal(Filters, &result)

	requestData := apiRequestData{
		Filter: result,
	}

	httpClient := HTTPClient{c.config.siteUrl + "/api/" + apiVersion, c.httpClient}
	request := httpClient.MakePOSTRequest(path, requestData.covertToByteArray())
	return request
}

//// Deprecated
////ViewLegacy object
//func (c *APIClient) ViewLegacy(Endpoint string, id int) string {
//	httpClient := HTTPClient{c.config.siteUrl + "/api/" + apiVersion, c.httpClient}
//	return httpClient.MakePOSTRequest(Endpoint+"/view/"+strconv.Itoa(id), nil)
//}
//
////Delete object
//func (c *APIClient) Delete(model models.ApiModel) string {
//	httpClient := HTTPClient{c.config.siteUrl + "/api/" + apiVersion, c.httpClient}
//	return httpClient.MakePOSTRequest(model.Endpoint()+"/delete/"+strconv.Itoa(model.Id()), nil)
//}
//
//// Deprecated
////Delete object
//func (c *APIClient) DeleteLegacy(Endpoint string, id int) string {
//	httpClient := HTTPClient{c.config.siteUrl + "/api/" + apiVersion, c.httpClient}
//	return httpClient.MakePOSTRequest(Endpoint+"/delete/"+strconv.Itoa(id), nil)
//}

func (c *APIClient) makeScoroAPIRequest(name string, path string, Request map[string]string) []byte {
	emp := make(map[string]interface{})
	emp[name] = Request

	// Marshal the map into a JSON string.
	empData, err := json.Marshal(emp)
	if err != nil {
		fmt.Println(err.Error())
		return []byte{}
	}

	httpClient := HTTPClient{c.config.siteUrl + "/api/" + apiVersion, c.httpClient}
	return httpClient.MakePOSTRequest(path, empData)
}

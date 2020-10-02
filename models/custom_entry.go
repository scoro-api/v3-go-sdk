package models

import (
	"encoding/json"
	"fmt"
	"log"
	"strconv"
)

type CustomEntry struct {
	AbstractModel

	Id           int    `json:"item_id,omitempty"`
	Status       string `json:"status,omitempty"`
	OwnerID      string `json:"owner_id,omitempty"`
	ModifiedDate string `json:"modified_date,omitempty"`
	DeletedDate  string `json:"deleted_date,omitempty"`
	IsDeleted    string `json:"is_deleted,omitempty"`

	Module string
}

type CustomViewResponse struct {
	Status   string                 `json:"status,omitempty"`
	Messages map[string]interface{} `json:"messages,omitempty"`
	Errors   map[string]interface{} `json:"errors,omitempty"`
	Data     CustomEntry            `json:"data,omitempty"`
}
type CustomListResponse struct {
	Status   string                   `json:"status,omitempty"`
	Messages interface{}              `json:"messages,omitempty"`
	Data     []map[string]interface{} `json:"data,omitempty"`
}

func (model *CustomEntry) ModuleName() string {
	return model.Module
}

func (model *CustomEntry) SetModuleName(moduleName string) {
	model.Module = moduleName
}

func (model *CustomEntry) validateModule() {
	if model.ModuleName() == "" {
		log.Fatal("Module unset, use SetModuleName")
	}
}

func (model *CustomEntry) FindById(id int) CustomEntry {
	model.validateModule()

	bytes := model.client.View(model.ModuleName(), id)
	CustomViewResponse := CustomViewResponse{}
	json.Unmarshal(bytes, &CustomViewResponse)
	customEntry := CustomViewResponse.Data
	customEntry.SetModuleName(model.ModuleName())
	customEntry.SetClient(model.client)
	customEntry.SetRawDataFromBytes(bytes)
	return customEntry
}

func (model *CustomEntry) Find(filter CustomEntry) []CustomEntry {
	model.validateModule()

	filterData := filter.RawData()
	if filter.Id > 0 {
		filterData["item_id"] = filter.Id
	}
	if filter.Status != "" {
		filterData["status"] = filter.Status
	}
	outputJSON, _ := json.Marshal(filterData)

	bytes := model.client.List(model.ModuleName(), outputJSON)

	response := CustomListResponse{}
	json.Unmarshal(bytes, &response)

	customEntriesMap := response.Data

	var customEntries []CustomEntry
	for _, customEntryData := range customEntriesMap {
		id, _ := strconv.Atoi(fmt.Sprintf("%v", customEntryData["item_id"]))
		customEntry := CustomEntry{
			Id:     id,
			Status: fmt.Sprintf("%v", customEntryData["status"]),
		}
		customEntry.SetModuleName(model.ModuleName())
		customEntry.SetRawDataFromMap(customEntryData)
		customEntry.SetClient(model.client)
		customEntries = append(customEntries, customEntry)
	}

	return customEntries
}

func (model *CustomEntry) Create() {
	model.Id = 0
	model.Update()
}

func (model *CustomEntry) Update() {
	model.validateModule()

	filterData := model.RawData()
	if model.Id > 0 {
		filterData["item_id"] = model.Id
	}
	if model.Status != "" {
		filterData["status"] = model.Status
	}
	outputJSON, _ := json.Marshal(filterData)

	var bytes []byte
	if model.Id > 0 {
		bytes = model.client.Update(model.ModuleName(), model.Id, outputJSON)
	} else {
		bytes = model.client.Create(model.ModuleName(), outputJSON)
	}

	CustomViewResponse := CustomViewResponse{}
	json.Unmarshal(bytes, &CustomViewResponse)
	newEntry := CustomViewResponse.Data

	model.Id = newEntry.Id
	model.Status = newEntry.Status
	model.OwnerID = newEntry.OwnerID
	model.ModifiedDate = newEntry.ModifiedDate
	model.IsDeleted = newEntry.IsDeleted
	model.DeletedDate = newEntry.DeletedDate

	model.SetRawDataFromBytes(bytes)
}

func (model *CustomEntry) Delete() CustomViewResponse {
	model.validateModule()

	return model.DeleteById(model.Id)
}

func (model *CustomEntry) DeleteById(id int) CustomViewResponse {
	model.validateModule()

	bytes := model.client.Delete(model.ModuleName(), id)
	CustomViewResponse := CustomViewResponse{}
	json.Unmarshal(bytes, &CustomViewResponse)
	return CustomViewResponse
}

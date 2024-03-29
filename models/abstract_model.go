package models

import (
	"encoding/json"
	scoro "github.com/scoro-api/v3-go-sdk"
)

type AbstractModel struct {
	client  *scoro.APIClient
	rawData map[string]interface{}
}

func (model *AbstractModel) SetClient(client *scoro.APIClient) {
	model.client = client
}

func (model *AbstractModel) RawData() map[string]interface{} {
	if model.rawData == nil {
		model.rawData = map[string]interface{}{}
	}

	return model.rawData
}

func (model *AbstractModel) SetRawDataFromBytes(data []byte) {
	var result map[string]map[string]interface{}
	json.Unmarshal(data, &result)
	model.rawData = result["data"]
}

func (model *AbstractModel) SetRawDataFromMap(data map[string]interface{}) {
	model.rawData = data
}

func (model *AbstractModel) GetValueFor(key string) interface{} {
	return model.RawData()[key]
}

func (model *AbstractModel) SetValueFor(key string, value interface{}) {
	data := model.RawData()
	data[key] = value

	model.SetRawDataFromMap(data)
}

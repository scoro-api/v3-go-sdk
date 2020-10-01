package models

import (
	"encoding/json"
	"github.com/siimtalts/v3-go-sdk"
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

func (model *AbstractModel) SetRawData(data []byte) {
	var result map[string]map[string]interface{}
	json.Unmarshal(data, &result)
	model.rawData = result["data"]
}

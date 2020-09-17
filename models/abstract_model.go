package models

import "google.golang.org/genproto/googleapis/type/date"

type AbstractModel struct {
	id           int
	createdUser  int
	modifiedUser int
	createdDate  date.Date
	modifiedDate date.Date
	rawData      map[string]interface{}
}

func (model *AbstractModel) ModifiedDate() date.Date {
	return model.modifiedDate
}

func (model *AbstractModel) SetModifiedDate(modifiedDate date.Date) {
	model.modifiedDate = modifiedDate
}

func (model *AbstractModel) CreatedDate() date.Date {
	return model.createdDate
}

func (model *AbstractModel) SetCreatedDate(createdDate date.Date) {
	model.createdDate = createdDate
}

func (model *AbstractModel) ModifiedUser() int {
	return model.modifiedUser
}

func (model *AbstractModel) SetModifiedUser(modifiedUser int) {
	model.modifiedUser = modifiedUser
}

func (model *AbstractModel) CreatedUser() int {
	return model.createdUser
}

func (model *AbstractModel) SetCreatedUser(createdUser int) {
	model.createdUser = createdUser
}

func (model *AbstractModel) Id() int {
	return model.id
}

func (model *AbstractModel) SetId(id int) {
	model.id = id
}

func (model *AbstractModel) RawData() map[string]interface{} {
	if model.rawData == nil {
		model.rawData = map[string]interface{}{}
	}

	return model.rawData
}

func (model *AbstractModel) SetRawData(rawData map[string]interface{}) {
	model.rawData = rawData
}

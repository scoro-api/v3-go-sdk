package models

import "google.golang.org/genproto/googleapis/type/date"

type ApiModel interface {
	ModifiedDate() date.Date
	SetModifiedDate(modifiedDate date.Date)
	CreatedDate() date.Date
	SetCreatedDate(createdDate date.Date)
	ModifiedUser() int
	SetModifiedUser(modifiedUser int)
	CreatedUser() int
	SetCreatedUser(createdUser int)
	Id() int
	SetId(id int)
	RawData() map[string]interface{}
	SetRawData(rawData map[string]interface{})
	Endpoint() string
	InitFromJSON(bytes []byte) error
}

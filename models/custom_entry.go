package models

type CustomEntry struct {
	AbstractModel

	Id     int    `json:"item_id,omitempty"`
	Status string `json:"status,omitempty"`
}

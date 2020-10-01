package models

import (
	"encoding/json"
	"google.golang.org/genproto/googleapis/type/date"
	"google.golang.org/genproto/googleapis/type/datetime"
	"strconv"
)

type status string

type statusesList struct {
	Active   status
	InActive status
	Pending  status
	Awaiting status
}

var Status = &statusesList{
	Active:   "active",
	InActive: "inactive",
	Pending:  "pending",
	Awaiting: "awaiting",
}

type User struct {
	AbstractModel

	Id           int    `json:"id,omitempty"`
	UserName     string `json:"username,omitempty"`
	FirstName    string `json:"firstname,omitempty"`
	LastName     string `json:"lastname,omitempty"`
	FullName     string `json:"full_name,omitempty"`
	Initials     string `json:"initials,omitempty"`
	Email        string `json:"email,omitempty"`
	Status       string `json:"status,omitempty"`
	Birthday     string `json:"birthday,omitempty"`
	Category     string `json:"category,omitempty"`
	position     string
	userPicture  string
	roleId       int
	userGroupIds []int
	countryId    int
	gsm          string
	timezone     datetime.TimeZone

	createdUser  int
	modifiedUser int
	createdDate  date.Date
	modifiedDate date.Date
}

type UserViewResponse struct {
	Status   string      `json:"status,omitempty"`
	Messages interface{} `json:"messages,omitempty"`
	Data     User        `json:"data,omitempty"`
}
type UserListResponse struct {
	Status   string      `json:"status,omitempty"`
	Messages interface{} `json:"messages,omitempty"`
	Data     []User      `json:"data,omitempty"`
}

func (model *User) FindById(id int) User {
	bytes := model.client.View("users/view/" + strconv.Itoa(id))

	userViewResponse := UserViewResponse{}
	json.Unmarshal(bytes, &userViewResponse)
	user := userViewResponse.Data
	user.SetRawData(bytes)

	return user
}

func (model *User) Find(filter User) []User {
	filterData, _ := json.Marshal(filter)
	bytes := model.client.List("users/list/", filterData)

	response := UserListResponse{}
	json.Unmarshal(bytes, &response)

	users := response.Data
	return users
}

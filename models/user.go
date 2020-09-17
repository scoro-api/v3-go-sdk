package models

import (
	"google.golang.org/genproto/googleapis/type/date"
	"google.golang.org/genproto/googleapis/type/datetime"
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

	userName     string
	firstName    string
	lastName     string
	fullName     string
	initials     string
	email        string
	status       status
	birthday     date.Date
	category     string
	position     string
	userPicture  string
	roleId       int
	userGroupIds []int
	countryId    int
	gsm          string
	timezone     datetime.TimeZone
}

func (user *User) Timezone() datetime.TimeZone {
	return user.timezone
}

func (user *User) SetTimezone(timezone datetime.TimeZone) {
	user.timezone = timezone
}

func (user *User) Gsm() string {
	return user.gsm
}

func (user *User) SetGsm(gsm string) {
	user.gsm = gsm
}

func (user *User) CountryId() int {
	return user.countryId
}

func (user *User) SetCountryId(countryId int) {
	user.countryId = countryId
}

func (user *User) UserGroupIds() []int {
	return user.userGroupIds
}

func (user *User) SetUserGroupIds(userGroupIds []int) {
	user.userGroupIds = userGroupIds
}

func (user *User) RoleId() int {
	return user.roleId
}

func (user *User) SetRoleId(roleId int) {
	user.roleId = roleId
}

func (user *User) UserPicture() string {
	return user.userPicture
}

func (user *User) SetUserPicture(userPicture string) {
	user.userPicture = userPicture
}

func (user *User) Position() string {
	return user.position
}

func (user *User) SetPosition(position string) {
	user.position = position
}

func (user *User) Category() string {
	return user.category
}

func (user *User) SetCategory(category string) {
	user.category = category
}

func (user *User) Birthday() date.Date {
	return user.birthday
}

func (user *User) SetBirthday(birthday date.Date) {
	user.birthday = birthday
}

func (user *User) Status() status {
	return user.status
}

func (user *User) SetStatus(status status) {
	user.status = status
}

func (user *User) Email() string {
	return user.email
}

func (user *User) SetEmail(email string) {
	user.email = email
}

func (user *User) Initials() string {
	return user.initials
}

func (user *User) SetInitials(initials string) {
	user.initials = initials
}

func (user *User) FullName() string {
	return user.fullName
}

func (user *User) SetFullName(fullName string) {
	user.fullName = fullName
}

func (user *User) LastName() string {
	return user.lastName
}

func (user *User) SetLastName(lastName string) {
	user.lastName = lastName
}

func (user *User) FirstName() string {
	return user.firstName
}

func (user *User) SetFirstName(firstName string) {
	user.firstName = firstName
}

func (user *User) UserName() string {
	return user.userName
}

func (user *User) SetUserName(userName string) {
	user.userName = userName
}

func (user *User) Endpoint() string {
	return "users"
}

package services

var UserService Service

func SetUp() {
	UserService = CreateService("http://localhost:8080/api/user/")
}

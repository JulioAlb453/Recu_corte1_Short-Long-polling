package domain

type User struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
	Sex string `json:"sex"`
	Gender string `json:"gender"`
}
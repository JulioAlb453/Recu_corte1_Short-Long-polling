package domain

type User struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
	Gender string `json:"gender"`
}
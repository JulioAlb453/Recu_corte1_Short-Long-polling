package domain

type UserRepository interface{
	AddUser(user User)
	GetUser() []User
}


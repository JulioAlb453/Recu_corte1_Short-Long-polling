package infraestructure

import "recu_c1/Users/domain"

type InMemomyUserRepository struct {
	users []domain.User
}

func NewInMemomyUserRepository() *InMemomyUserRepository {
    return &InMemomyUserRepository{}
}

func (r *InMemomyUserRepository) AddUser(user domain.User) {
    r.users = append(r.users, user)
}

func (r *InMemomyUserRepository) GetUser() []domain.User {
    return r.users
}


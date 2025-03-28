package application

import (
	"recu_c1/Users/domain"
	"sync"
	"time"
)

type UserService struct {
	repo    domain.UserRepository
	counter domain.GenderCounter
	mu      sync.Mutex
}

func NewUserService(repo domain.UserRepository) *UserService {
	return &UserService{
		repo:    repo,
		counter: domain.GenderCounter{Counts: map[string]int{"Masculino": 0, "Femenino": 0}},
	}
}

func (s *UserService) AddUser(user domain.User) {
	s.repo.AddUser(user)

	if user.Gender == "Masculino" || user.Gender == "Femenino" {
		s.mu.Lock()
		s.counter.Counts[user.Gender]++
		s.mu.Unlock()
	}

}

func (s *UserService) GetUser() []domain.User {
	return s.repo.GetUser()
}

func (s *UserService) GetGenderCounts(lastCount map[string]int) map[string]int {
	timeout := time.After(1 * time.Second)
	ch := make(chan bool, 1)

	go func() {
		for {
			s.mu.Lock()
			changed := false
			for _, gender := range []string{"Masculino", "Femenino"} {
				if s.counter.Counts[gender] != lastCount[gender] {
					lastCount[gender] = s.counter.Counts[gender]
					changed = true
				}
			}
			s.mu.Unlock()

			if changed {
				ch <- true
				break
			}
			time.Sleep(100 * time.Millisecond)
		}
	}()

	select {
	case <-ch:
		return lastCount
	case <-timeout:
		return lastCount

	}

}

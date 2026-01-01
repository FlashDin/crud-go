package repository

import (
	"crud-go/model"
	"sync"
)

type UserRepository struct {
	data   map[int]model.User
	nextID int
	mu     sync.Mutex
}

func NewUserRepository() *UserRepository {
	return &UserRepository{
		data:   make(map[int]model.User),
		nextID: 1,
	}
}

func (r *UserRepository) FindAll() []model.User {
	r.mu.Lock()
	defer r.mu.Unlock()

	var users []model.User
	for _, u := range r.data {
		users = append(users, u)
	}
	return users
}

func (r *UserRepository) FindByID(id int) (model.User, bool) {
	r.mu.Lock()
	defer r.mu.Unlock()

	u, ok := r.data[id]
	return u, ok
}

func (r *UserRepository) Save(user model.User) model.User {
	r.mu.Lock()
	defer r.mu.Unlock()

	if user.ID == 0 {
		user.ID = r.nextID
		r.nextID++
	}
	r.data[user.ID] = user
	return user
}

func (r *UserRepository) Delete(id int) {
	r.mu.Lock()
	defer r.mu.Unlock()

	delete(r.data, id)
}

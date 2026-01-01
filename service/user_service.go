package service

import (
	"crud-go/dto"
	"crud-go/model"
	"crud-go/repository"
)

type UserService struct {
	repo *repository.UserRepository
}

func NewUserService(repo *repository.UserRepository) *UserService {
	return &UserService{repo: repo}
}

func (s *UserService) GetAll() []dto.UserResponse {
	users := s.repo.FindAll()
	var res []dto.UserResponse

	for _, u := range users {
		res = append(res, toResponse(u))
	}
	return res
}

func (s *UserService) Create(req dto.UserRequest) dto.UserResponse {
	user := model.User{
		Name: req.Name,
		Age:  req.Age,
	}
	saved := s.repo.Save(user)
	return toResponse(saved)
}

func toResponse(u model.User) dto.UserResponse {
	return dto.UserResponse{
		ID:   uint(u.ID),
		Name: u.Name,
		Age:  u.Age,
	}
}

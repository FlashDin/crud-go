package service

import (
	"crud-go/dto"
	"crud-go/model"
	"crud-go/repository"
	"errors"

	"gorm.io/gorm"
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

func (s *UserService) GetByID(id uint) (dto.UserResponse, error) {
	user, err := s.repo.FindByID(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return dto.UserResponse{}, errors.New("user not found")
		}
		return dto.UserResponse{}, err
	}
	return toResponse(user), nil
}

func (s *UserService) Update(id uint, req dto.UserRequest) (dto.UserResponse, error) {
	user, err := s.repo.FindByID(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return dto.UserResponse{}, errors.New("user not found")
		}
		return dto.UserResponse{}, err
	}

	user.Name = req.Name
	user.Age = req.Age
	saved := s.repo.Save(user)
	return toResponse(saved), nil
}

func (s *UserService) Delete(id uint) error {
	err := s.repo.Delete(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return errors.New("user not found")
		}
		return err
	}
	return nil
}

func toResponse(u model.User) dto.UserResponse {
	return dto.UserResponse{
		ID:   uint(u.ID),
		Name: u.Name,
		Age:  u.Age,
	}
}

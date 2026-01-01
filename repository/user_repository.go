package repository

import (
	"crud-go/model"

	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (r *UserRepository) FindAll() []model.User {
	var users []model.User
	r.db.Find(&users)
	return users
}

func (r *UserRepository) FindByID(id uint) (model.User, bool) {
	var user model.User
	result := r.db.First(&user, id)
	return user, result.Error == nil
}

func (r *UserRepository) Save(user model.User) model.User {
	r.db.Save(&user)
	return user
}

func (r *UserRepository) Delete(id uint) {
	r.db.Delete(&model.User{}, id)
}

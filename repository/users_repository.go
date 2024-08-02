package repository

import (
	"errors"
	"the_pizza_api/helper"
	"the_pizza_api/model"

	"gorm.io/gorm"
)

type UsersRepository interface {
	Save(users model.Users)
	// Update(users model.Users)
	Delete(usersId int)
	FindById(usersId int) (users model.Users, err error)
	FindAll() []model.Users
	FindByEmail(email string) (users model.Users, err error)
}

type UsersRepositoryImpl struct {
	Db *gorm.DB
}

func NewUsersRepositoryImpl(Db *gorm.DB) UsersRepository {
	return &UsersRepositoryImpl{Db: Db}
}

// Delete implements UsersRepository.
func (p *UsersRepositoryImpl) Delete(usersId int) {
	var user model.Users
	result := p.Db.Where("id = ?", usersId).Delete(&user)
	helper.ErrorPanic(result.Error)
}

// FindAll implements UsersRepository.
func (p *UsersRepositoryImpl) FindAll() []model.Users {
	var users []model.Users
	result := p.Db.Find(&users)
	helper.ErrorPanic(result.Error)
	return users
}

// FindById implements UsersRepository.
func (p *UsersRepositoryImpl) FindById(usersId int) (users model.Users, err error) {
	var user model.Users
	result := p.Db.Find(&user, usersId)
	if result != nil {
		return user, nil
	} else {
		return user, errors.New("user is not found")
	}
}

// Save implements UsersRepository.
func (p *UsersRepositoryImpl) Save(user model.Users) {
	result := p.Db.Create(&user)
	helper.ErrorPanic(result.Error)
}

func (p *UsersRepositoryImpl) FindByEmail(email string) (users model.Users, err error) {
	var user model.Users
	result := p.Db.First(&user, "email = ?", email)
	if result != nil {
		return user, nil
	} else {
		return user, errors.New("user is not found")
	}
}

// Update implements UsersRepository.
// func (p *UsersRepositoryImpl) Update(user model.Users) {
// 	var updateUser = request.UpdateUsersRequest{
// 		ID:          user.ID,
// 		Name:        user.Name,
// 		Description: user.Description,
// 	}
// 	result := p.Db.Model(&user).Updates(updateUser)
// 	helper.ErrorPanic(result.Error)
// }

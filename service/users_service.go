package service

import (
	"os"
	"strconv"
	"the_pizza_api/data/request"
	"the_pizza_api/data/response"
	"the_pizza_api/helper"
	"the_pizza_api/model"
	"the_pizza_api/repository"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

type UsersService interface {
	Create(users request.CreateUsersRequest)
	// Update(users request.UpdateUsersRequest)
	Delete(usersId int)
	FindById(usersId int) response.UsersResponse
	FindAll() []response.UsersResponse
	Login(user request.CreateUsersRequest) string
}

type UsersServiceImpl struct {
	UsersRepository repository.UsersRepository
	Validate        *validator.Validate
}

func NewUsersServiceImpl(userRepository repository.UsersRepository, validate *validator.Validate) UsersService {
	return &UsersServiceImpl{
		UsersRepository: userRepository,
		Validate:        validate,
	}
}

// Login a user
func (p *UsersServiceImpl) Login(requestBody request.CreateUsersRequest) string {
	err := p.Validate.Struct(requestBody)
	helper.ErrorPanic(err)

	// look up requested user
	user, err := p.UsersRepository.FindByEmail(requestBody.Email)
	helper.ErrorPanic(err)

	// compare sent in pass with saved user pass hash
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(requestBody.Password))
	helper.ErrorPanic(err)

	// generate a jwt token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.RegisteredClaims{
		Subject:   strconv.Itoa(int(user.ID)),
		ExpiresAt: jwt.NewNumericDate(time.Now().UTC().Add(time.Hour)),
	})

	// Sign and get the complete encoded token as a string using the secret
	secret := os.Getenv("SECRET")
	tokenString, err := token.SignedString([]byte(secret))
	helper.ErrorPanic(err)
	return tokenString
}

// Create implements UsersService
func (p *UsersServiceImpl) Create(users request.CreateUsersRequest) {
	err := p.Validate.Struct(users)
	helper.ErrorPanic(err)

	// hash password
	hash, err := bcrypt.GenerateFromPassword([]byte(users.Password), 10)
	helper.ErrorPanic(err)

	userModel := model.Users{
		Email:    users.Email,
		Password: string(hash),
	}
	p.UsersRepository.Save(userModel)
}

// Delete implements UsersService
func (p *UsersServiceImpl) Delete(usersId int) {
	p.UsersRepository.Delete(usersId)
}

// FindAll implements UsersService
func (p *UsersServiceImpl) FindAll() []response.UsersResponse {
	result := p.UsersRepository.FindAll()

	var users []response.UsersResponse
	for _, value := range result {
		user := response.UsersResponse{
			ID:    value.ID,
			Email: value.Email,
		}
		users = append(users, user)
	}

	return users
}

// FindById implements UsersService
func (p *UsersServiceImpl) FindById(usersId int) response.UsersResponse {
	userData, err := p.UsersRepository.FindById(usersId)
	helper.ErrorPanic(err)

	userResponse := response.UsersResponse{
		ID:    userData.ID,
		Email: userData.Email,
	}
	return userResponse
}

// Update implements UsersService
// func (p *UsersServiceImpl) Update(users request.UpdateUsersRequest) {
// 	userData, err := p.UsersRepository.FindById(users.ID)
// 	helper.ErrorPanic(err)
// 	userData.Name = users.Name
// 	userData.Description = users.Description
// 	p.UsersRepository.Update(userData)
// }

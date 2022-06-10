package users

import (
	"errors"

	"golang.org/x/crypto/bcrypt"
)

type Services interface {
	RegisterUser(input RegisterUserInput) (User, error)
	Login(input LoginInput) (User, error)
	IsEmailAvailable(input CheckEmailInput) (bool, error)
	GetUsersByID(ID int) (User, error)
	GetUserByToken(token string) (User, error)
	SaveToken(ID int, token string) (User, error)
	GetToken(ID int) (string, error)
}

type services struct {
	repository Repository
}

func NewServices(repository Repository) *services {
	return &services{repository}
}

func (s *services) RegisterUser(input RegisterUserInput) (User, error) {
	user := User{}
	user.Email = input.Email
	password, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)
	if err != nil {
		return user, err
	}
	user.Password = string(password)
	newUser, err := s.repository.Save(user)
	if err != nil {
		return newUser, err
	}
	return newUser, nil
}

func (s *services) Login(input LoginInput) (User, error) {
	email := input.Email
	password := input.Password

	user, err := s.repository.FindByEmail(email)

	if err != nil {
		return user, err
	}
	if user.ID == 0 {
		return user, errors.New("User Not Found on the Database")
	}
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return user, err
	}

	return user, nil
}

func (s *services) IsEmailAvailable(input CheckEmailInput) (bool, error) {
	email := input.Email
	user, err := s.repository.FindByEmail(email)
	if err != nil {
		return false, err
	}
	if user.ID == 0 {
		return true, nil
	}
	return false, nil

}

func (s *services) GetUsersByID(ID int) (User, error) {
	user, err := s.repository.FindByID(ID)
	if err != nil {
		return user, err
	}
	if user.ID == 0 {
		return user, errors.New("User Not Found on the Database")
	}
	return user, nil

}

func (s *services) GetUserByToken(token string) (User, error) {
	user, err := s.repository.FindByToken(token)
	if err != nil {
		return user, err
	}
	if user.ID == 0 {
		return user, errors.New("User Not Found on the Database")
	}
	return user, nil

}

func (s *services) SaveToken(ID int, token string) (User, error) {
	user, err := s.repository.FindByIdAndUpdateToken(ID, token)
	if err != nil {
		return user, err
	}

	return user, nil
}
func (s *services) GetToken(ID int) (string, error) {
	token, err := s.repository.FindByID(ID)
	if err != nil {
		return token.Token, err
	}
	return token.Token, nil
}

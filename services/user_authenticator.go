package services

import (
	"errors"
	"golang.org/x/crypto/bcrypt"
	"log"
	"ropc-service/model"
	"ropc-service/repositories"
)

type UserAuthenticatorContract interface {
	Authenticate(username, password string) (*model.User, error)
}

type UserAuthenticator struct {
	userRepository repositories.UserRepository
}

func InstantiateUserAuthenticator() *UserAuthenticator {
	return &UserAuthenticator{
		userRepository: repositories.NewUserRepository(),
	}
}

func (selfC UserAuthenticator) Authenticate(username, password string) (*model.User, error) {

	user, err := selfC.userRepository.GetUser(username)

	if err != nil {
		log.Println(err)
		return nil, err
	}

	if err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil || err == bcrypt.ErrMismatchedHashAndPassword {
		log.Println(err)
		return nil, errors.New("invalid user credentials")
	}

	return user, nil
}

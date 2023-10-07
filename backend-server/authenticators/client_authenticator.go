package authenticators

import (
	"errors"
	"ropc-service/model"
	"ropc-service/repositories"
	"ropc-service/utils"

	"golang.org/x/crypto/bcrypt"
)

const InvalidClientMessage = "invalid client credentials"
const ConnectionErrorMessage = "could not authenticate client"

type ClientAuthenticator interface {
	Authenticate(clientId, clientSecret string) (*model.Token, error)
}

type clientAuthenticator struct {
	repository                    repositories.ApplicationRepository
	thirdPartyClientAuthenticator ThirdPartyClientAuthenticator
}

func NewClientAuthenticator(repository repositories.ApplicationRepository) ClientAuthenticator {
	return &clientAuthenticator{
		repository: repository,
	}
}

func (selfC clientAuthenticator) Authenticate(clientId, clientSecret string) (*model.Token, error) {

	client, err := selfC.repository.Get(clientId)
	if err != nil {
		return nil, err
	}

	if err = bcrypt.CompareHashAndPassword([]byte(client.ClientSecret), []byte(clientSecret)); err != nil || errors.Is(err, bcrypt.ErrMismatchedHashAndPassword) {
		return nil, errors.New(InvalidClientMessage)
	}

	token, err := utils.GenerateToken(client, clientSecret)
	if err != nil {
		return nil, err
	}

	tokenResponse := &model.Token{
		AccessToken: token,
	}

	return tokenResponse, nil
}

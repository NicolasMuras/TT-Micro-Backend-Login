package service

import (
	"TT-Micro-Backend-Login/pkg/dto"
	"TT-Micro-Backend-Login/pkg/models"
	"TT-Micro-Backend-Login/pkg/repository"
	"fmt"

	"github.com/aws/aws-lambda-go/events"
)

// NewUserService return the user service interface
// with "user" db collection.
func NewUserService(repo repository.Interface) UserInterface {

	return &UserService{
		Repository: repo,
	}
}

func (s *UserService) Login(request events.APIGatewayProxyRequest, loginReq dto.LoginRequest) (dto.LoginResponse, error) {
	user := models.User{
		Email:    loginReq.Email,
		Password: loginReq.Password,
	}
	_, err := s.Repository.ValidateDBUserData(request, user)
	if err != nil {
		return dto.LoginResponse{}, fmt.Errorf("\n[ERROR] %v:", err)
	}

	token := dto.LoginResponse{
		Token: "",
	}

	new_token, err := s.Repository.RetrieveToken(request)
	token.Token = new_token.Token
	if err != nil {
		return dto.LoginResponse{}, err
	}
	return token, nil
}

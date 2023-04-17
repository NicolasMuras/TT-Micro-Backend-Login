package service

import (
	"TT-Micro-Backend-Login/pkg/dto"
	"TT-Micro-Backend-Login/pkg/repository"

	"github.com/aws/aws-lambda-go/events"
)

// UserService contains the attributes for the User object service.
type UserService struct {
	Repository repository.UserInterface
}

// UserInterface define the methods for the User service.
type UserInterface interface {
	Login(request events.APIGatewayProxyRequest, loginReq dto.LoginRequest) (dto.LoginResponse, error)
}

// MockUserService defines the mock User Service.
type MockUserService struct{}

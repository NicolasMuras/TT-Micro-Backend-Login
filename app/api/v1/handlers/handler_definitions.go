package handlers

import (
	"TT-Micro-Backend-Login/pkg/service"

	"github.com/aws/aws-lambda-go/events"
)

// Handler contains attributes for the object handlers.
type Handler struct {
	userService service.UserInterface
}

// UserHandler define the methods for the user object.
type UserHandler interface {
	Login(events.APIGatewayProxyRequest) (*events.APIGatewayProxyResponse, error)
	UnhandledMethod(string) (*events.APIGatewayProxyResponse, error)
}

// Interface define the methods for the handler.
type Interface interface {
	UserHandler
}

// NewHandler Return the Handlder with all the handler dependencies.
func NewHandler(userService service.UserInterface) Interface {
	return &Handler{
		userService: userService,
	}
}

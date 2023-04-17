package routers

import (
	"TT-Micro-Backend-Login/api/v1/handlers"
	"TT-Micro-Backend-Login/pkg/db"
	"TT-Micro-Backend-Login/pkg/repository"
	"TT-Micro-Backend-Login/pkg/service"

	"github.com/aws/aws-lambda-go/events"
	"github.com/jmoiron/sqlx"
)

// Server define the server dependencies & attributes.
type Server struct {
	db *sqlx.DB
}

func Router(request events.APIGatewayProxyRequest) (*events.APIGatewayProxyResponse, error) {
	s := Server{
		db: db.ConnectDB(),
	}
	userRepository := repository.NewUserRepository(s.db)
	//Initialize services
	userService := service.NewUserService(userRepository)
	// handler creation
	handler := handlers.NewHandler(userService)

	switch request.HTTPMethod {

	case "POST":
		return handler.Login(request)
	default:
		return handler.UnhandledMethod(request.HTTPMethod)
	}
}

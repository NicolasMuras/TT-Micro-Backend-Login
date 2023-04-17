package repository

import (
	"TT-Micro-Backend-Login/pkg/models"

	"github.com/aws/aws-lambda-go/events"
	"github.com/jmoiron/sqlx"
)

// Repository contains attributes for the object repository.
type UserRepository struct {
	db *sqlx.DB
}

// UserRepository define the methods for the user object.
type UserInterface interface {
	ValidateDBUserData(request events.APIGatewayProxyRequest, user models.User) (bool, error)
	RetrieveToken(request events.APIGatewayProxyRequest) (models.Token, error)
}

// Interface define the methods for the repository.
type Interface interface {
	UserInterface
}

// NewHandler Return the Handlder with all the handler dependencies.
func NewUserRepository(db *sqlx.DB) Interface {
	return &UserRepository{
		db: db,
	}
}

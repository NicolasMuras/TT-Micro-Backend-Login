package repository

import (
	"TT-Micro-Backend-Login/pkg/auth"
	"TT-Micro-Backend-Login/pkg/models"
	"fmt"

	"github.com/aws/aws-lambda-go/events"
)

// ValidateDBUserData verify that the user and password are presented in db.
// Returns and boolean value that represent: success: true | fail: false
func (r *UserRepository) ValidateDBUserData(request events.APIGatewayProxyRequest, user models.User) (bool, error) {
	response := []models.User{}

	err := r.db.Select(&response, "SELECT * FROM users WHERE email = $1 AND password = $2", user.Email, user.Password)
	if err != nil || len(response) == 0 {
		return false, err
	}
	return true, nil
}

// RetrieveToken retrieves a token for the user.
func (r *UserRepository) RetrieveToken(request events.APIGatewayProxyRequest) (models.Token, error) {
	token := models.Token{}

	tokenString, err := auth.GenerateJWT()
	token.Token = tokenString
	if err != nil {
		return models.Token{}, fmt.Errorf("\n[ERROR] Fail at token generation: %s", err)
	}
	return token, nil
}

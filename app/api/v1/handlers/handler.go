package handlers

import (
	"TT-Micro-Backend-Login/pkg/dto"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-sdk-go/aws"
)

var ErrorMethodNotAllowed = "method not allowed"

type ErrorBody struct {
	ErrorMsg *string `json:"error,omitempty"`
}

func (h *Handler) Login(request events.APIGatewayProxyRequest) (*events.APIGatewayProxyResponse, error) {
	var newLoginRequest dto.LoginRequest

	if err := json.Unmarshal([]byte(request.Body), &newLoginRequest); err != nil {
		return apiResponse(http.StatusBadRequest, ErrorBody{aws.String(err.Error())})
	}

	newLoginResponse, err := h.userService.Login(request, newLoginRequest)
	if err != nil {
		return apiResponse(http.StatusInternalServerError, ErrorBody{aws.String(err.Error())})
	}
	return apiResponse(http.StatusOK, newLoginResponse)
}

func (h *Handler) UnhandledMethod(method string) (*events.APIGatewayProxyResponse, error) {
	return apiResponse(http.StatusMethodNotAllowed, fmt.Sprintf("[ERROR] Method not allowed: "+method))
}

package main

import (
	routers "TT-Micro-Backend-Login/api/v1/router"

	"github.com/aws/aws-lambda-go/lambda"
)

func main() {
	lambda.Start(routers.Router)
}

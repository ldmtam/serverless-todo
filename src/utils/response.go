package utils

import (
	"encoding/json"

	"github.com/aws/aws-lambda-go/events"
)

// ResponseOK ...
func ResponseOK(data interface{}) (events.APIGatewayProxyResponse, error) {
	result := map[string]interface{}{
		"result": data,
	}

	body, err := json.Marshal(result)
	if err != nil {
		return events.APIGatewayProxyResponse{
			StatusCode: 400,
			Body:       err.Error(),
		}, nil
	}

	return events.APIGatewayProxyResponse{
		StatusCode: 200,
		Headers: map[string]string{
			"Content-Type": "application/json",
		},
		Body: string(body),
	}, nil
}

// ResponseError ...
func ResponseError(err error) (events.APIGatewayProxyResponse, error) {
	result := map[string]string{
		"error": err.Error(),
	}

	body, err := json.Marshal(result)
	if err != nil {
		return events.APIGatewayProxyResponse{
			StatusCode: 400,
			Body:       err.Error(),
		}, nil
	}

	return events.APIGatewayProxyResponse{
		StatusCode: 200,
		Headers: map[string]string{
			"Content-Type": "application/json",
		},
		Body: string(body),
	}, nil
}

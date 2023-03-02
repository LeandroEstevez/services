package services

import (
	"io"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Service struct {
	endpoint string
}

func CreateService(endpoint string) Service {
	return Service{
		endpoint: endpoint,
	}
}

func (service *Service) Request(method string, path string, header []string, body io.Reader) (int, []byte, error) {
	req, err := http.NewRequest(method, service.endpoint+path, body)
	if err != nil {
		return -1, nil, err
	}
	req.Header.Add("Content-Type", "application/json")

	if header != nil {
		req.Header.Add(header[0], header[1])
	}

	client := &http.Client{}
	response, err := client.Do(req)
	if err != nil {
		return -1, nil, err
	}

	bodyBytes, err := io.ReadAll(response.Body)
	if response.StatusCode != 200 {
		if err != nil {
			return http.StatusInternalServerError, nil, err
		}

		statusCode := ReturnStatusCode(response)
		// ctx.JSON(statusCode, string(bodyBytes))
		return statusCode, bodyBytes, nil
	}

	return response.StatusCode, bodyBytes, nil
}

func ErrorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}

func ReturnStatusCode(response *http.Response) int {
	switch response.StatusCode {
	case 500:
		return http.StatusInternalServerError
	case 400:
		return http.StatusBadRequest
	case 404:
		return http.StatusNotFound
	case 401:
		return http.StatusUnauthorized
	case 403:
		return http.StatusForbidden
	default:
		return http.StatusBadRequest
	}
}

package response

import (
	"net/http"
	"user-service/constants"
	errorConstant "user-service/constants/error"

	"github.com/gin-gonic/gin"
)

type Response struct {
	Status  string      `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
	Token   *string     `json:"token,omitempty"`
}

type ParamHTTPResp struct {
	Code    int
	Err     error
	Message *string
	Gin     *gin.Context
	Data    *interface{}
	Token   *string
}

func HttpResponse(param ParamHTTPResp) {
	// Success case
	if param.Err == nil {
		var data interface{}
		if param.Data != nil {
			data = *param.Data
		}

		param.Gin.JSON(param.Code, Response{
			Status:  constants.Success,
			Message: http.StatusText(param.Code),
			Data:    data,
			Token:   param.Token,
		})
		return
	}

	// Error case
	message := errorConstant.ErrInternalServerError.Error() // Default message

	// Check if custom message provided
	if param.Message != nil {
		message = *param.Message
	} else if param.Err != nil {
		// Check if it's a known error
		if errorConstant.ErrMapping(param.Err) {
			message = param.Err.Error()
		}
	}

	var data interface{}
	if param.Data != nil {
		data = *param.Data
	}

	param.Gin.JSON(param.Code, Response{
		Status:  constants.Error,
		Message: message,
		Data:    data,
	})
	return
}

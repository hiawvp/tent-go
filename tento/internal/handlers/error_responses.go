package handlers

import (
	"errors"
	"fmt"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type CustomErrorResponse struct {
	httpStatusCode int
	body           ErrorResponseBody
}

type ErrorResponseBody struct {
	code    string
	message string
}

type handledError struct {
	httpStatusCode int
	code           string
	baseMessage    string
}

var handledErrors = map[string]handledError{
	"kek": handledError{code: "INVALID_BARCODE", baseMessage: "Could not find product"},
}

func newErrorResponse(err error) *CustomErrorResponse {
	switch true {
	case errors.Is(err, gorm.ErrRecordNotFound):
		return &CustomErrorResponse{
			httpStatusCode: 404,
			body: ErrorResponseBody{code: "UNKNOWN_BARCODE",
				message: "Could not find product",
			},
		}
	}
	return &CustomErrorResponse{
		httpStatusCode: 400,
		body: ErrorResponseBody{code: "UNKNOWN_ERROR",
			message: "Could not handle your request",
		},
	}
}

func NewCustomErrorResponse(err error, v ...interface{}) *CustomErrorResponse {
	resp := newErrorResponse(err)
	extraInfo := " Aditional Info: " + fmt.Sprint(v...)
	resp.body.message = resp.body.message + extraInfo
	return resp
}

//func (errResponse CustomErrorResponse) AsGinH()  {
//return gin.H
//}

func abortRequest(c *gin.Context, httpStatusCode int, errorCode, errorMsg string) {
	c.JSON(httpStatusCode, gin.H{"code": errorCode, "message": errorMsg})
}

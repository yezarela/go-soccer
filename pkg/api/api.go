package api

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

// ResponseMeta represents api response meta
type ResponseMeta struct {
	Code  int    `json:"code"`
	Error string `json:"error,omitempty"`
	Type  string `json:"type,omitempty"`
}

// Response represents api response
type Response struct {
	Meta ResponseMeta `json:"meta"`
	Data interface{}  `json:"data"`
}

// ResponseOK returns 200
func ResponseOK(c echo.Context, data interface{}) error {

	body := &Response{
		Meta: ResponseMeta{
			Code: http.StatusOK,
		},
		Data: data,
	}

	return c.JSON(http.StatusOK, body)
}

func responseError(c echo.Context, statusCode int, errType, msg string) error {

	body := &Response{
		Meta: ResponseMeta{
			Code:  statusCode,
			Type:  errType,
			Error: msg,
		},
	}

	return c.JSON(statusCode, body)
}

// ResponseBadRequest returns 400
func ResponseBadRequest(c echo.Context, msg string) error {
	return responseError(c, http.StatusBadRequest, "Bad Request", msg)
}

// ResponseUnprocessableEntity returns 422
func ResponseUnprocessableEntity(c echo.Context, msg string) error {
	return responseError(c, http.StatusUnprocessableEntity, "Unprocessable Entity", msg)
}

// ResponseNotFound returns 404
func ResponseNotFound(c echo.Context, msg string) error {
	return responseError(c, http.StatusNotFound, "Resource Not Found", msg)
}

// ResponseError returns 500
func ResponseError(c echo.Context, err error) error {
	return responseError(c, http.StatusInternalServerError, "Internal Server Error", err.Error())
}

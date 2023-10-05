package response

import (
	"fmt"
	"net/http"
	"test-lion-superindo/lib/log"

	"github.com/labstack/echo/v4"
)

type httpContext struct {
	ctx echo.Context
}

type response struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func ToJson(ctx echo.Context) *httpContext {
	return &httpContext{ctx}
}

func (hc httpContext) builder(code int, msg string, data interface{}) error {
	return hc.ctx.JSON(code, &response{
		Code:    code,
		Message: msg,
		Data:    data,
	})
}

func (hc httpContext) OK(data interface{}, msg string) error {
	log.Info(fmt.Sprintf("response: [%v] %v", http.StatusOK, msg))
	return hc.builder(http.StatusOK, msg, data)
}

func (hc httpContext) InternalServerError(msg string) error {
	log.Info(fmt.Sprintf("response: [%v] %v", http.StatusInternalServerError, msg))
	return hc.builder(http.StatusInternalServerError, "Internal Service Error", nil)
}

func (hc httpContext) BadRequest(msg string) error {
	log.Info(fmt.Sprintf("response: [%v] %v", http.StatusBadRequest, msg))
	return hc.builder(http.StatusBadRequest, "Bad Request", nil)
}

func (hc httpContext) NotFound(msg string) error {
	log.Info(fmt.Sprintf("response: [%v] %v", http.StatusNotFound, msg))
	return hc.builder(http.StatusNotFound, msg, nil)
}

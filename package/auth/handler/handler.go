package handler

import (
	"encoding/json"
	"fmt"
	"io"
	"strings"
	"test-lion-superindo/lib/response"
	"test-lion-superindo/package/auth/model"
	"test-lion-superindo/package/auth/usecase"

	"github.com/labstack/echo/v4"
)

type authHandler struct {
	uc usecase.AuthUsecase
}

func NewAuthHandler(uc usecase.AuthUsecase) *authHandler {
	return &authHandler{uc}
}

func (h *authHandler) Mount(g *echo.Group) {
	g.POST("/registration", h.Registration)
	g.POST("/login", h.Login)
}

func (h *authHandler) Registration(e echo.Context) error {
	body, err := io.ReadAll(e.Request().Body)
	if err != nil {
		return response.ToJson(e).BadRequest("Failed get body")
	}

	user := model.UserAuth{}
	if err = json.Unmarshal(body, &user); err != nil {
		return response.ToJson(e).BadRequest("Failed unmarshal")
	}

	if err = h.uc.Registration(user); err != nil {
		switch {
		case strings.Contains(err.Error(), "(SQLSTATE 23502)"):
			column := strings.Split(err.Error(), "\"")[1]
			return response.ToJson(e).BadRequest(fmt.Sprintf("Can't Null value on column \"%v\"", column))
		case strings.Contains(err.Error(), "(SQLSTATE 23505)"):
			column := strings.Split(err.Error(), "\"")[1]
			return response.ToJson(e).BadRequest(fmt.Sprintf("Duplicate value on column \"%v\"", column))
		default:
			return response.ToJson(e).InternalServerError(err.Error())
		}
	}

	return response.ToJson(e).OK(nil, "Registration Success")
}

func (h *authHandler) Login(e echo.Context) error {
	body, err := io.ReadAll(e.Request().Body)
	if err != nil {
		return response.ToJson(e).BadRequest("Failed get body")
	}

	user := model.UserAuth{}
	if err = json.Unmarshal(body, &user); err != nil {
		return response.ToJson(e).BadRequest("Failed unmarshal")
	}

	token, isMatch, err := h.uc.Login(user)

	if err != nil {
		return response.ToJson(e).InternalServerError(err.Error())
	}

	if !isMatch {
		return response.ToJson(e).NotFound("Username & Password not found!")
	}

	data := model.ResponseLogin{Token: token}

	return response.ToJson(e).OK(data, "Login Success")
}

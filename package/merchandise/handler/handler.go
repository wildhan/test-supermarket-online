package handler

import (
	"encoding/json"
	"io"
	"strconv"
	"test-lion-superindo/lib/helper"
	"test-lion-superindo/lib/response"
	"test-lion-superindo/package/merchandise/model"
	"test-lion-superindo/package/merchandise/usecase"

	"github.com/labstack/echo/v4"
)

type merchandiseHandler struct {
	uc usecase.MerchandiseUsecase
}

func NewMerchandiseHandler(uc usecase.MerchandiseUsecase) *merchandiseHandler {
	return &merchandiseHandler{uc}
}

func (h *merchandiseHandler) Mount(g *echo.Group) {
	g.GET("/categories", h.GetCategories)
	g.GET("/listByCategory", h.GetListByCategory)
	g.GET("/detail", h.GetDetailMerchandise)
	g.POST("/addToCart", h.AddToCart)
	g.GET("/getCart", h.GetCart)
}

func (h *merchandiseHandler) GetCategories(e echo.Context) error {
	data, err := h.uc.GetCategories()

	if err != nil {
		return response.ToJson(e).InternalServerError(err.Error())
	}

	return response.ToJson(e).OK(data, "Get Categories Merchandise Success")
}

func (h *merchandiseHandler) GetListByCategory(e echo.Context) error {
	categoryId, err := strconv.Atoi(e.QueryParam("category_id"))
	if err != nil {
		return response.ToJson(e).BadRequest("Failed covert param to integer")
	}

	data, err := h.uc.GetMerchandise(categoryId)

	if err != nil {
		return response.ToJson(e).InternalServerError(err.Error())
	}
	return response.ToJson(e).OK(data, "Get List Merchandise Success")
}

func (h *merchandiseHandler) GetDetailMerchandise(e echo.Context) error {
	merchandiseId, err := strconv.Atoi(e.QueryParam("merchandise_id"))
	if err != nil {
		return response.ToJson(e).BadRequest("Failed covert param to integer")
	}

	username := helper.GetUsername(e)

	data, err := h.uc.GetDetailMerchandise(merchandiseId, username)

	if err != nil {
		return response.ToJson(e).InternalServerError(err.Error())
	}
	return response.ToJson(e).OK(data, "Get Detail Merchandise Success")
}

func (h *merchandiseHandler) AddToCart(e echo.Context) error {
	body, err := io.ReadAll(e.Request().Body)
	if err != nil {
		return response.ToJson(e).BadRequest("Failed get body")
	}

	merchandise := model.MerchandiseAddCart{}
	if err = json.Unmarshal(body, &merchandise); err != nil {
		return response.ToJson(e).BadRequest("Failed unmarshal")
	}

	username := helper.GetUsername(e)

	if err := h.uc.AddToCart(merchandise, username); err != nil {
		return response.ToJson(e).InternalServerError(err.Error())
	}

	return response.ToJson(e).OK(nil, "Add to Cart Success")
}

func (h *merchandiseHandler) GetCart(e echo.Context) error {
	username := helper.GetUsername(e)
	data, err := h.uc.GetCart(username)

	if err != nil {
		return response.ToJson(e).InternalServerError(err.Error())
	}

	return response.ToJson(e).OK(data, "Get Cart Success")
}

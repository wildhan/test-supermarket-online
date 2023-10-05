package usecase

import (
	"test-lion-superindo/package/merchandise/model"
	"test-lion-superindo/package/merchandise/repository"
)

type merchandiseUsecase struct {
	repo repository.MerchandiseRepo
}

func NewMerchandiseUsecase(repo repository.MerchandiseRepo) MerchandiseUsecase {
	return &merchandiseUsecase{repo}
}

type MerchandiseUsecase interface {
	GetCategories() ([]model.Category, error)
	GetMerchandise(categoryId int) ([]model.Merchandise, error)
	GetDetailMerchandise(merchandiseId int, username string) (model.DetailMerchandise, error)
	AddToCart(merchandise model.MerchandiseAddCart, username string) error
	GetCart(username string) ([]model.MerchandiseInCart, error)
}

func (uc *merchandiseUsecase) GetCategories() ([]model.Category, error) {
	return uc.repo.GetCategories()
}
func (uc *merchandiseUsecase) GetMerchandise(categoryId int) ([]model.Merchandise, error) {
	return uc.repo.GetMerchandise(categoryId)
}
func (uc *merchandiseUsecase) GetDetailMerchandise(merchandiseId int, username string) (model.DetailMerchandise, error) {
	return uc.repo.GetDetailMerchandise(merchandiseId, username)
}

func (uc *merchandiseUsecase) AddToCart(merchandise model.MerchandiseAddCart, username string) error {
	return uc.repo.AddToCart(merchandise, username)
}

func (uc *merchandiseUsecase) GetCart(username string) ([]model.MerchandiseInCart, error) {
	return uc.repo.GetCart(username)
}

package repository

import (
	"test-lion-superindo/config/database"
	"test-lion-superindo/package/merchandise/model"
)

type merchandiseRepo struct {
	dbConn *database.DbConnection
}

func NewMerchandiseRepo(dbConn *database.DbConnection) MerchandiseRepo {
	return &merchandiseRepo{dbConn}
}

type MerchandiseRepo interface {
	GetCategories() ([]model.Category, error)
	GetMerchandise(categoryId int) ([]model.Merchandise, error)
	GetDetailMerchandise(merchandiseId int, username string) (model.DetailMerchandise, error)
	AddToCart(merchandise model.MerchandiseAddCart, username string) error
	GetCart(username string) ([]model.MerchandiseInCart, error)
}

func (r *merchandiseRepo) GetCategories() ([]model.Category, error) {
	db := r.dbConn.DB
	categories := make([]model.Category, 0)

	query := `SELECT id, "name", "desc", create_at, update_at
	FROM public.master_category;`
	err := db.Raw(query).Scan(&categories).Error

	return categories, err
}

func (r *merchandiseRepo) GetMerchandise(categoryId int) ([]model.Merchandise, error) {
	db := r.dbConn.DB
	merchandise := make([]model.Merchandise, 0)

	query := `SELECT id, "name", "desc", category_id, create_at, update_at
	FROM public.master_merchandise
	WHERE category_id=?;`
	err := db.Raw(query, categoryId).Scan(&merchandise).Error

	return merchandise, err
}

func (r *merchandiseRepo) GetDetailMerchandise(merchandiseId int, username string) (model.DetailMerchandise, error) {
	db := r.dbConn.DB
	merchandise := model.DetailMerchandise{}
	params := make([]interface{}, 0)
	params = append(params, username)
	params = append(params, merchandiseId)

	query := `SELECT mm.id as id, mm."name" as "name", mm."desc" as "desc", mm.category_id as category_id, 
	mc."name" as category_name, c.qty as in_cart, mm.create_at, mm.update_at
	FROM public.master_merchandise mm
	left join public.master_category mc on mm.category_id = mc.id 
	left join public.cart c on mm.id = c.merchandise_id and c.username like ?
	WHERE mm.id=?;`
	err := db.Raw(query, params...).Scan(&merchandise).Error

	return merchandise, err
}

func (r *merchandiseRepo) AddToCart(merchandise model.MerchandiseAddCart, username string) error {
	db := r.dbConn.DB
	params := make([]interface{}, 0)
	params = append(params, merchandise.ID)
	params = append(params, username)
	params = append(params, merchandise.Qty)
	params = append(params, merchandise.Qty)

	query := `INSERT INTO public.cart
	(merchandise_id, username, qty)
	VALUES(?, ?, ?)
	on conflict ON CONSTRAINT cart_un
	do update set qty = ?`
	return db.Exec(query, params...).Error
}

func (r *merchandiseRepo) GetCart(username string) ([]model.MerchandiseInCart, error) {
	db := r.dbConn.DB
	merchandise := make([]model.MerchandiseInCart, 0)

	query := `select c.id as id, c.merchandise_id as merchandise_id, mm."name" as merchandise_name, c.qty as qty
	from public.cart c 
	left join public.master_merchandise mm on c.merchandise_id = mm.id 
	where c.username like ?;`
	err := db.Raw(query, username).Scan(&merchandise).Error

	return merchandise, err
}

package model

type Category struct {
	ID   int    `json:"id" gorm:"column:id"`
	Name string `json:"name" gorm:"column:name"`
	Desc string `json:"desc" gorm:"column:desc"`
}

type Merchandise struct {
	ID   int    `json:"id" gorm:"column:id"`
	Name string `json:"name" gorm:"column:name"`
	Desc string `json:"desc" gorm:"column:desc"`
}

type DetailMerchandise struct {
	ID      int    `json:"id" gorm:"column:id"`
	Name    string `json:"name" gorm:"column:name"`
	Desc    string `json:"desc" gorm:"column:desc"`
	CatID   string `json:"category_id" gorm:"column:category_id"`
	CatName string `json:"category_name" gorm:"column:category_name"`
	InCart  int    `json:"in_cart" gorm:"column:in_cart"`
}

type MerchandiseAddCart struct {
	ID  int `json:"id" gorm:"column:id"`
	Qty int `json:"qty" gorm:"column:qty"`
}

type MerchandiseInCart struct {
	ID              int    `json:"id" gorm:"column:id"`
	MerchandiseId   int    `json:"merchandise_id" gorm:"column:merchandise_id"`
	MerchandiseName string `json:"merchandise_name" gorm:"column:merchandise_name"`
	Qty             int    `json:"qty" gorm:"column:qty"`
}

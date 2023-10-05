package repository

import (
	"test-lion-superindo/config/database"
	"test-lion-superindo/lib/helper"
	"test-lion-superindo/package/auth/model"
)

type authRepo struct {
	dbConn *database.DbConnection
}

func NewAuthRepo(dbConn *database.DbConnection) AuthRepo {
	return &authRepo{dbConn}
}

type AuthRepo interface {
	AddUserAuth(user model.UserAuth) error
	CheckUser(user model.UserAuth) ([]model.UserAuth, error)
}

func (r *authRepo) AddUserAuth(user model.UserAuth) error {
	db := r.dbConn.DB
	params := make([]interface{}, 0)
	params = append(params, helper.EmptyStringToNull(user.Username))
	params = append(params, helper.EmptyStringToNull(user.Password))

	query := `INSERT INTO public.master_user
	(username, "password")
	VALUES(?, ?);`
	return db.Exec(query, params...).Error
}
func (r *authRepo) CheckUser(user model.UserAuth) ([]model.UserAuth, error) {
	db := r.dbConn.DB
	users := make([]model.UserAuth, 0)
	params := make([]interface{}, 0)
	params = append(params, helper.EmptyStringToNull(user.Username))
	params = append(params, helper.EmptyStringToNull(user.Password))

	query := `SELECT username, "password"
	FROM public.master_user mu
	WHERE mu.username = ? AND mu."password" = ?;`

	err := db.Raw(query, params...).Scan(&users).Error

	return users, err
}

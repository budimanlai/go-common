package gocommon

import (
	"github.com/go-playground/validator"
	"github.com/jmoiron/sqlx"
)

var (
	Db        *sqlx.DB
	Validator *validator.Validate
)

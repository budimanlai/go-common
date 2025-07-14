package models

import (
	"errors"
	"time"

	common "github.com/budimanlai/go-common"
	"github.com/jmoiron/sqlx"
)

type User struct {
	ID                 uint      `json:"id" db:"id"`
	Username           string    `json:"username" db:"username" validate:"required,min=3,max=16"`
	AuthKey            string    `json:"auth_key" db:"auth_key"`
	PasswordHash       string    `json:"password_hash" db:"password_hash" validate:"required,min=6"`
	PinHash            string    `json:"pin_hash" db:"pin_hash"`
	PasswordResetToken string    `json:"password_reset_token" db:"password_reset_token"`
	Fullname           string    `json:"fullname" db:"fullname" validate:"required,min=3"`
	Email              string    `json:"email" db:"email" validate:"required,min=3,email"`
	Handphone          string    `json:"handphone" db:"handphone" validate:"required,min=10,max=15"`
	Status             string    `json:"status" db:"status" validate:"required"`
	LoginDashboard     string    `json:"login_dashboard" db:"login_dashboard" validate:"required,oneof=Y N"`
	Dob                time.Time `json:"dob" db:"dob"`
	Gender             string    `json:"gender" db:"gender"`
	Address            *string   `json:"address" db:"address"`
	CountryID          string    `json:"country_id" db:"country_id"`
	ProvID             uint      `json:"prov_id" db:"prov_id"`
	CityID             uint      `json:"city_id" db:"city_id"`
	PostalCode         string    `json:"postal_code" db:"postal_code"`
	CreatedAt          time.Time `json:"created_at" db:"created_at"`
	UpdatedAt          time.Time `json:"updated_at" db:"updated_at"`
	VerificationToken  string    `json:"verification_token" db:"verification_token"`
	Avatar             string    `json:"avatar" db:"avatar"`
	AvatarSmall        string    `json:"avatar_small" db:"avatar_small"`
}

func (u *User) TableName() string {
	return "user"
}

// Validate melakukan validasi pada field User sesuai dengan tag 'validate'.
// Mengembalikan error jika ada field yang tidak valid.
func (u *User) Validate() error {
	return common.Validator.Struct(u)
}

// SetPassword meng-hash password dalam bentuk plain-text yang diberikan
// dan mengatur field PasswordHash pada User dengan hasil hash tersebut.
func (u *User) SetPassword(password string) {
	u.PasswordHash = common.HashPassword(password)
}

// GenerateAuthKey menghasilkan kunci otentikasi acak baru untuk pengguna
// dan menetapkannya ke field AuthKey. Kunci yang dihasilkan berupa string acak sepanjang 32 karakter.
func (u *User) GenerateAuthKey() {
	u.AuthKey = common.GenerateRandomString(32)
}

// FindUserByID mengambil data pengguna dari database berdasarkan ID uniknya.
// Fungsi ini mengembalikan pointer ke struct User jika ditemukan, atau error jika pengguna tidak ada
// atau terjadi kesalahan pada database.
//
// Parameter:
//
//	db - pointer ke koneksi database sqlx.DB
//	id - ID unik dari pengguna yang ingin diambil
//
// Return:
//
//	*User - pointer ke struct User jika ditemukan
//	error - error jika pengguna tidak ditemukan atau terjadi kesalahan database
func FindUserByID(db *sqlx.DB, id uint) (*User, error) {
	var user *User
	err := db.Get(&user, `SELECT id, username, auth_key, fullname, email, handphone
		status, address, country_id, prov_id, city_id FROM user where id = ?`, id)
	if err != nil {
		return nil, errors.New("user not found")
	}
	return user, nil
}

// FindUserByUsername mengambil data pengguna dari database berdasarkan username uniknya.
// Fungsi ini mengembalikan pointer ke struct User jika ditemukan, atau error jika pengguna tidak ada
// atau terjadi kesalahan pada database.
//
// Parameter:
//
//	db - pointer ke koneksi database sqlx.DB
//	username - username unik dari pengguna yang ingin diambil
//
// Return:
//
//	*User - pointer ke struct User jika ditemukan
//	error - error jika pengguna tidak ditemukan atau terjadi kesalahan database
func FindUserByUsername(db *sqlx.DB, username string) (*User, error) {
	var user User
	err := db.Get(&user, `SELECT id, username, auth_key, fullname, email, handphone
		status, address, country_id, prov_id, city_id FROM user WHERE username = ?`, username)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

// FindUserByEmail mengambil data pengguna dari database berdasarkan email uniknya.
// Fungsi ini mengembalikan pointer ke struct User jika ditemukan, atau error jika pengguna tidak ada
// atau terjadi kesalahan pada database.
//
// Parameter:
//
//	db - pointer ke koneksi database sqlx.DB
//	email - email unik dari pengguna yang ingin diambil
//
// Return:
//
//	*User - pointer ke struct User jika ditemukan
//	error - error jika pengguna tidak ditemukan atau terjadi kesalahan database
func FindUserByEmail(db *sqlx.DB, email string) (*User, error) {
	var user User
	err := db.Get(&user, `SELECT id, username, auth_key, fullname, email, handphone
		status, address, country_id, prov_id, city_id FROM user WHERE email = ?`, email)
	if err != nil {
		return nil, errors.New("user not found")
	}
	return &user, nil
}

// FindUserByHandphone mengambil data pengguna dari database berdasarkan nomor handphone uniknya.
// Fungsi ini mengembalikan pointer ke struct User jika ditemukan, atau error jika pengguna tidak ada
// atau terjadi kesalahan pada database.
//
// Parameter:
//
//	db - pointer ke koneksi database sqlx.DB
//	handphone - nomor handphone unik dari pengguna yang ingin diambil
//
// Return:
//
//	*User - pointer ke struct User jika ditemukan
//	error - error jika pengguna tidak ditemukan atau terjadi kesalahan database
func FindUserByHandphone(db *sqlx.DB, handphone string) (*User, error) {
	var user User
	err := db.Get(&user, `SELECT id, username, auth_key, fullname, email, handphone
		status, address, country_id, prov_id, city_id FROM user WHERE handphone = ?`, handphone)
	if err != nil {
		return nil, errors.New("user not found")
	}
	return &user, nil
}

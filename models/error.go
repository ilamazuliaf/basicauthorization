package models

import "errors"

var (
	ErrInternalServerError = errors.New("Ada Masalah Pada Server")
	ErrDataNotFound        = errors.New("Data tidak ditemukan")
	ErrConflict            = errors.New("Data sudah ada")
	ErrUserNotFound        = errors.New("Username atau password salah!")
)

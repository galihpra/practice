package coupons

import (
	"github.com/golang-jwt/jwt/v5"

	"github.com/labstack/echo/v4"
)

type Coupon struct {
	ID          uint
	NamaProgram string
	Kode        string
	Link        string
	Gambar      string
	UserID      uint
}

type Handler interface {
	Add() echo.HandlerFunc
	ReadByUser() echo.HandlerFunc
	Read() echo.HandlerFunc
}

type Service interface {
	TambahKupon(token *jwt.Token, newCoupon Coupon) (Coupon, error)
	GetKuponByUser(token *jwt.Token) ([]Coupon, error)
	GetKupon() ([]Coupon, error)
}

type Repository interface {
	InsertKupon(userID uint, newCoupon Coupon) (Coupon, error)
	ReadKuponByUser(userID uint) ([]Coupon, error)
	ReadKupon() ([]Coupon, error)
}

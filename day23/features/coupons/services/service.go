package services

import (
	"TugasDay23/features/coupons"
	"TugasDay23/helper/jwt"
	"errors"
	"strings"

	golangjwt "github.com/golang-jwt/jwt/v5"
)

type CouponService struct {
	m coupons.Repository
}

func New(model coupons.Repository) coupons.Service {
	return &CouponService{
		m: model,
	}
}

func (cs *CouponService) TambahKupon(token *golangjwt.Token, newCoupon coupons.Coupon) (coupons.Coupon, error) {
	userID, err := jwt.ExtractToken(token)
	if err != nil {
		return coupons.Coupon{}, err
	}

	result, err := cs.m.InsertKupon(userID, newCoupon)

	if err != nil {
		if strings.Contains(err.Error(), "duplicate") {
			return coupons.Coupon{}, errors.New("kupon sudah ada pada sistem")
		}
		return coupons.Coupon{}, errors.New("terjadi kesalahan server")
	}

	return result, nil
}

func (cs *CouponService) GetKuponByUser(token *golangjwt.Token) ([]coupons.Coupon, error) {
	userID, err := jwt.ExtractToken(token)
	if err != nil {
		return nil, err
	}

	result, err := cs.m.ReadKuponByUser(userID)

	if err != nil {
		return nil, errors.New("terjadi kesalahan server")
	}

	return result, nil
}

func (cs *CouponService) GetKupon() ([]coupons.Coupon, error) {
	result, err := cs.m.ReadKupon()

	if err != nil {
		return nil, errors.New("terjadi kesalahan server")
	}

	return result, nil
}

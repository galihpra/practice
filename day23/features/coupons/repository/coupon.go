package repository

import (
	"TugasDay23/features/coupons"

	"gorm.io/gorm"
)

type CouponModel struct {
	gorm.Model
	NamaProgram string
	Kode        string
	Link        string
	Gambar      string
	UserID      uint
}

type couponQuery struct {
	db *gorm.DB
}

func New(db *gorm.DB) coupons.Repository {
	return &couponQuery{
		db: db,
	}
}

func (cq *couponQuery) InsertKupon(userID uint, newCoupon coupons.Coupon) (coupons.Coupon, error) {
	var inputData = new(CouponModel)
	inputData.UserID = userID
	inputData.NamaProgram = newCoupon.NamaProgram
	inputData.Link = newCoupon.Link
	inputData.Kode = newCoupon.Kode
	inputData.Gambar = newCoupon.Gambar

	if err := cq.db.Create(&inputData).Error; err != nil {
		return coupons.Coupon{}, err
	}

	newCoupon.ID = inputData.ID

	return newCoupon, nil
}

func (cq *couponQuery) ReadKuponByUser(userID uint) ([]coupons.Coupon, error) {
	var couponModels []CouponModel

	if err := cq.db.Where("user_id = ?", userID).Find(&couponModels).Error; err != nil {
		return nil, err
	}

	var result []coupons.Coupon
	for _, model := range couponModels {
		result = append(result, coupons.Coupon{
			ID:          model.ID,
			NamaProgram: model.NamaProgram,
			Kode:        model.Kode,
			Link:        model.Link,
			Gambar:      model.Gambar,
			UserID:      model.UserID,
		})
	}

	return result, nil
}

func (cq *couponQuery) ReadKupon() ([]coupons.Coupon, error) {
	var couponModels []CouponModel

	if err := cq.db.Find(&couponModels).Error; err != nil {
		return nil, err
	}

	var result []coupons.Coupon
	for _, model := range couponModels {
		result = append(result, coupons.Coupon{
			ID:          model.ID,
			NamaProgram: model.NamaProgram,
			Kode:        model.Kode,
			Link:        model.Link,
			Gambar:      model.Gambar,
			UserID:      model.UserID,
		})
	}

	return result, nil
}

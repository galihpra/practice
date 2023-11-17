package detailpembelian

type CreateRequest struct {
	PembelianID string `json:"pembelian_id" form:"pembelian_id" validate:"required"`
	ProductID   string `json:"product_id" form:"product_id" validate:"required"`
	Qty         int    `json:"qty" form:"qty" validate:"required"`
}

type CreateResponse struct {
	PembelianID string `json:"pembelian_id"`
	ProductID   string `json:"product_id"`
	Qty         int    `json:"qty"`
}

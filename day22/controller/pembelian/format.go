package pembelian

type CreateRequest struct {
	No_invoice string `json:"no_invoice" form:"no_invoice" validate:"required"`
	CustomerID string `json:"customer_id" form:"customer_id" validate:"required"`
	UserID     string `json:"user_id" form:"user_id" validate:"required"`
}

type ListResponse struct {
	No_invoice string `json:"invoice"`
	UserID     string `json:"user_id"`
	CustomerID string `json:"customer_id"`
	Total      int    `json:"total"`
}

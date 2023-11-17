package pembelian

type CreateRequest struct {
	CustomerID string `json:"hp" form:"hp" validate:"required"`
}

type ListResponse struct {
	No_invoice string `json:"invoice"`
	UserID     string `json:"user_id"`
	CustomerID string `json:"customer_id"`
	Total      int    `json:"total"`
}

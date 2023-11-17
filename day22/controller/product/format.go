package product

type CreateRequest struct {
	Barcode string `json:"barcode" form:"barcode" validate:"required"`
	Nama    string `json:"nama" form:"nama" validate:"required"`
	Harga   int    `json:"harga" form:"harga" validate:"required,number"`
	Stok    int    `json:"stok" form:"stok" validate:"required,number"`
	UserID  string `json:"user_id" form:"user_id" validate:"required"`
}

type CreateResponse struct {
	Barcode string `json:"barcode"`
	Nama    string `json:"nama"`
	Harga   int    `json:"harga"`
	Stok    int    `json:"stok"`
	UserID  string `json:"user_id"`
}

type ListResponse struct {
	Barcode string `json:"barcode"`
	Nama    string `json:"nama"`
	Harga   int    `json:"harga"`
	Stok    int    `json:"stok"`
	UserID  string `json:"user_id"`
}

type UpdateRequest struct {
	Nama   string `json:"nama" form:"nama" validate:"required"`
	Harga  int    `json:"harga" form:"harga" validate:"required,number"`
	Stok   int    `json:"stok" form:"stok" validate:"required,number"`
	UserID string `json:"user_id" form:"user_id" validate:"required"`
}

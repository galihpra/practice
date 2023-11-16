package customer

type CreateRequest struct {
	Hp   string `json:"hp" form:"hp" validate:"required"`
	Nama string `json:"nama" form:"nama" validate:"required"`
}

type CreateResponse struct {
	Hp   string `json:"hp"`
	Nama string `json:"nama"`
}

type UpdateRequest struct {
	Nama string `json:"nama" form:"nama" validate:"required"`
}

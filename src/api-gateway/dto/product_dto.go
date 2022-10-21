package dto

type ProductResponse struct {
	ID          int      `json:"id"`
	MerchantID  int      `json:"merchant_id"`
	Name        string   `json:"name"`
	Description string   `json:"description"`
	Price       int      `json:"price"`
	CreatedAt   DateTime `json:"created_at"`
	UpdatedAt   DateTime `json:"updated_at"`
}

type ProductCreateReq struct {
	MerchantID  int    `json:"merchant_id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Price       int    `json:"price"`
}

type ProductUpdateReq struct {
	ID          int
	Name        string `json:"name"`
	Description string `json:"description"`
	Price       int    `json:"price"`
}

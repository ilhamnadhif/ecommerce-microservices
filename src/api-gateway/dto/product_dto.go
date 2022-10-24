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
type ProductResponseWithMerchant struct {
	ProductResponse
	Merchant MerchantResponse `json:"merchant"`
}

type ProductCreateReq struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Price       int    `json:"price"`
	QueryData   QueryData
}

type ProductUpdateReq struct {
	ID          int
	Name        string `json:"name"`
	Description string `json:"description"`
	Price       int    `json:"price"`
	QueryData   QueryData
}
type ProductDeleteReq struct {
	ID        int
	QueryData QueryData
}

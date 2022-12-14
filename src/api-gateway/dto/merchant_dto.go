package dto

type MerchantResponse struct {
	ID        int      `json:"id"`
	Name      string   `json:"name"`
	Email     string   `json:"email"`
	Password  string   `json:"password"`
	CreatedAt DateTime `json:"created_at"`
	UpdatedAt DateTime `json:"updated_at"`
}

type MerchantResponseWithProducts struct {
	MerchantResponse
	Products []ProductResponse `json:"products"`
}

type MerchantCreateReq struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type MerchantUpdateReq struct {
	ID        int
	Name      string `json:"name"`
	Email     string `json:"email"`
	Password  string `json:"password"`
	QueryData QueryData
}

type MerchantDeleteReq struct {
	ID        int
	QueryData QueryData
}

package dto

type CartResponse struct {
	ID         int      `json:"id"`
	CustomerID int      `json:"customer_id"`
	ProductID  int      `json:"product_id"`
	Quantity   int      `json:"quantity"`
	CreatedAt  DateTime `json:"created_at"`
	UpdatedAt  DateTime `json:"updated_at"`
}

type CartResponseWithProduct struct {
	ID         int             `json:"id"`
	CustomerID int             `json:"customer_id"`
	ProductID  int             `json:"product_id"`
	Quantity   int             `json:"quantity"`
	CreatedAt  DateTime        `json:"created_at"`
	UpdatedAt  DateTime        `json:"updated_at"`
	Product    ProductResponse `json:"product"`
}

type CartCreateReq struct {
	ProductID int `json:"product_id"`
	Quantity  int `json:"quantity"`
	QueryData QueryData
}

type CartUpdateReq struct {
	ID        int
	Quantity  int `json:"quantity"`
	QueryData QueryData
}
type CartDeleteReq struct {
	ID        int
	QueryData QueryData
}

package dto

type CustomerResponse struct {
	ID        int      `json:"id"`
	Name      string   `json:"name"`
	Email     string   `json:"email"`
	Password  string   `json:"password"`
	CreatedAt DateTime `json:"created_at"`
	UpdatedAt DateTime `json:"updated_at"`
}

type CustomerResponseWithCartProducts struct {
	ID        int                       `json:"id"`
	Name      string                    `json:"name"`
	Email     string                    `json:"email"`
	Password  string                    `json:"password"`
	CreatedAt DateTime                  `json:"created_at"`
	UpdatedAt DateTime                  `json:"updated_at"`
	Carts     []CartResponseWithProduct `json:"carts"`
}

type CustomerCreateReq struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type CustomerUpdateReq struct {
	ID        int
	Name      string `json:"name"`
	Email     string `json:"email"`
	Password  string `json:"password"`
	QueryData QueryData
}

type CustomerDeleteReq struct {
	ID        int
	QueryData QueryData
}

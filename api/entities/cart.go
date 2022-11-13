package entities

type CartProduct struct {
	ID int64 `json:"id"`
	Quantity int64 `json:"quantity"`
	Product  Coffe `json:"product"`
}

type CartSimple struct {
	ID int64 `json:"id"`
	Quantity int64 `json:"quantity"`
	ProductId  int64 `json:"productId"`
}

package entities

type CartProduct struct {
	Product  Coffe[[]string] `json:"product"`
	Quantity int64 `json:"quantity"`
	ID int64 `json:"id"`
}

package presenter

type ProductCart struct {
	ID       int64 `json:"id"`
	Quantity int64 `json:"quantity"`
	Product  Coffe `json:"product"`
}

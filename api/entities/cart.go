package entities

type ProductCart struct {
	ID       int64 `json:"id"`
	Quantity int64 `json:"quantity"`
	Product  Coffe `json:"product"`
}

type ProductCartSimple struct {
	ID        int64 `json:"id"`
	Quantity  int64 `json:"quantity"`
	ProductId int64 `json:"productId"`
}

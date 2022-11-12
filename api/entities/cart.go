package entities

type CartProduct struct {
	ID int64 `json:"id"`
	Quantity int64 `json:"quantity"`
	Product  Coffe[[]string] `json:"product"`
}

type CartProductSimple struct {
	ProductId int64
	Quantity  int64
	ID        int64
}
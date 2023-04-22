package entities

type Cart struct {
	ID       int64         `json:"id"`
	Products []CartProduct `json:"products"`
}

type CartProduct struct {
	Quantity int64   `json:"quantity"`
	Product  Product `json:"product"`
}

type CartProductSimple struct {
	Quantity  int64 `json:"quantity"`
	ProductId int64 `json:"productId"`
}

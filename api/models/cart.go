package models

// TODO
// ALTERAR PARA ---> 1 para muitos com os produtos ---> 1 carrinho para muitos produtos
type Cart struct {
	ID       int64 `json:"id"`
	Quantity int64 `json:"quantity"`
	Product  Coffe `json:"product"`
}

type CartSimple struct {
	ID        int64 `json:"id"`
	Quantity  int64 `json:"quantity"`
	ProductId int64 `json:"productId"`
}

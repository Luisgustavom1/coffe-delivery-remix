package entity

const (
	SPECIAL     = "especial"
	TRADITIONAL = "tradicional"
	ALCOHOLIC   = "alcoólico"
	WITH_MILK   = "com leite"
	ICE_COLD    = "gelado"
)

type Coffe struct {
	ID          int64  `json:"id"`
	Img         string `json:"img"`
	Price       int64  `json:"price"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Stok        int64  `json:"stok"`
	Categories  string `json:"categories"`
}

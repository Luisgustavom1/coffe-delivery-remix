export enum Categories {
  SPECIAL = 'especial',
  TRADITIONAL = 'tradicional',
  ALCOHOLIC = "alcoólico",
  WITH_MILK = "com leite",
  ICE_COLD = 'gelado'
}

export interface Coffe {
  id: number,
  img: string,
  price: number,
  title: string,
  description: string,
  categories: Array<Categories>,
  stok: number
}

export interface CartProduct {
  product: Coffe,
  quantity: number
}
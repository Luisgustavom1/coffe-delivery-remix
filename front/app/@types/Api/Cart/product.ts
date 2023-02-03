import type { Categories } from ".";

interface Coffe {
  id: number,
  img: string,
  price: number,
  title: string,
  description: string,
  categories: Array<Categories>,
  stok: number
}

interface Product {
  product: Coffe,
  quantity: number,
  id: number
}

export type { Coffe, Product }
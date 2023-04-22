import type { Categories } from ".";

interface Product {
  id: number,
  img: string,
  price: number,
  title: string,
  description: string,
  categories: Array<Categories>,
  type: string
  stok: number
}

interface Cart {
  id: number
  products: Array<CartProduct>
}

interface CartProduct {
  product: Product,
  quantity: number,
}

export type { Product, Cart, CartProduct }
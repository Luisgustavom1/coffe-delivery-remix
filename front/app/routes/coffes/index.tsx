import type { LoaderFunction } from "@remix-run/node";
import { json } from "@remix-run/node";
import { useLoaderData } from "@remix-run/react";
import { CoffeeCards } from "@/components/CoffeeCards";
import api from "@/services/api";
import type { Cart, Product } from "@/@types/Api/Cart";
import { useDispatch } from "react-redux";
import { CartActions } from "@/features/cart/slice";

let didInit = false;

type ApiProductsResponse = Array<Product>;
type ApiCartResponse = Cart;

interface LoaderResponse {
  products: ApiProductsResponse;
  cart: ApiCartResponse;
}

export function ErrorBoundary({ error }: { error: Error }) {
  console.error(error);

  return (
    <div className="mt-14 typography-title-s text-red-500">
      Opss algo de errado aconteceu...
    </div>
  );
}

export const loader: LoaderFunction = async () => {
  const { data: products } = await api.get<ApiProductsResponse>("/products");
  const { data: cart } = await api.get<ApiCartResponse>("/cart/101");

  return json<LoaderResponse>({
    products,
    cart,
  });
};

const ProductsIndexRoute = () => {
  const { cart, products } = useLoaderData<LoaderResponse>();
  const dispatch = useDispatch();

  if (!didInit) {
    didInit = true
    dispatch(CartActions.setCartProduct(cart));
  }
  return (
    <main className="mt-32">
      <h2 className="typography-title-l text-base-subtitle">Nossos cafés</h2>

      <div className="mt-14 flex flex-wrap justify-center gap-x-8 gap-y-10">
        {products.map((coffee) => (
          <CoffeeCards key={coffee.id} coffee={coffee} />
        ))}
      </div>
    </main>
  );
};

export default ProductsIndexRoute;

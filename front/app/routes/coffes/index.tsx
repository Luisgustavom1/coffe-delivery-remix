import type { LoaderFunction } from "@remix-run/node";
import { json } from "@remix-run/node";
import { useLoaderData } from "@remix-run/react";
import { CoffeCards } from "@/components/CoffeCards";
import api from "@/services/api";
import type { Product, Coffe } from "@/@types/Api/Cart";
import { useDispatch } from "react-redux";
import { CartActions } from "@/features/cart/slice";

let didInit = false;

type ApiCoffesResponse = Array<Coffe>;
type ApiCartResponse = Array<Product>;

interface LoaderResponse {
  coffes: ApiCoffesResponse;
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
  const { data: allCoffes } = await api.get<ApiCoffesResponse>("/coffes");
  const { data: cartProducts } = await api.get<ApiCartResponse>("/cart");
  
  return json<LoaderResponse>({
    coffes: allCoffes,
    cart: cartProducts,
  });
};

const CoffesIndexRoute = () => {
  const { cart, coffes } = useLoaderData<LoaderResponse>();
  const dispatch = useDispatch();

  if (!didInit) {
    didInit = true
    dispatch(CartActions.setCartProduct(cart));
  }
  return (
    <main className="mt-32">
      <h2 className="typography-title-l text-base-subtitle">Nossos cafés</h2>

      <div className="mt-14 flex flex-wrap justify-center gap-x-8 gap-y-10">
        {coffes.map((coffe) => (
          <CoffeCards key={coffe.id} coffe={coffe} />
        ))}
      </div>
    </main>
  );
};

export default CoffesIndexRoute;

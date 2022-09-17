import type { LoaderFunction } from "@remix-run/node";
import type { CartProduct } from "@/@types/Api";
import { json } from "@remix-run/node";
import { useLoaderData } from "@remix-run/react";
import api from "@/services/api";
import { InputNumber } from "@/components/UI/InputNumber";
import { formatPrice } from "@/utils/formats";
import { Chips } from "@/components/UI/Chips";
import { Trash } from "phosphor-react";
import { Button } from "@/components/UI/Button";

type LoaderResponse = Array<CartProduct>;

export const loader: LoaderFunction = async () => {
  const { data } = await api.get<LoaderResponse>("/cart");

  return json<LoaderResponse>(data);
};

const CheckoutIndexRoute = () => {
  const cart = useLoaderData<LoaderResponse>();

  return (
    <>
      {cart.map(({ product, quantity }) => (
        <section
          key={product.id}
          className="flex pb-8 border-b border-base-button mb-6"
        >
          <img
            className="w-16 h-16"
            src={product.img}
            alt={`Imagem do café ${product.title}`}
          />
          <div className="ml-5">
            <p className="typography-regular-m text-base-subtitle mb-2">
              {product.title}
            </p>
            <div className="flex gap-2">
              <InputNumber defaultValue={quantity} />
              <span className="p-2 flex gap-1 items-center bg-base-button rounded-md">
                <Trash size={16} color="#8047F8" />
                <p className="typography-button-s text-base-text">Remover</p>
              </span>
            </div>
          </div>
          <p className="typography-bold-m text-base-text ml-auto">
            R$ {formatPrice(product.price)}
          </p>
        </section>
      ))}
      <article className="mb-6">
        <span className="typography-regular-s text-base-text flex justify-between mb-3">
          <p>Total de itens</p>
          <p>R$ 29,70</p>
        </span>
        <span className="typography-regular-s text-base-text flex justify-between mb-3">
          <p>Entrega</p>
          <p>R$ 3,70</p>
        </span>
        <span className="typography-bold-l text-base-subtitle flex justify-between mb-3">
          <p>Total</p>
          <p>R$ 33,70</p>
        </span>
      </article>
      <Button variant="secondary">
        Confirmar pedido
      </Button>
    </>
  );
};

export default CheckoutIndexRoute;

import type { LoaderFunction } from "@remix-run/node";
import type { CartProduct } from "@/@types/Api";
import { json } from "@remix-run/node";
import { useLoaderData } from "@remix-run/react";
import { Trash } from "phosphor-react";
import api from "@/services/api";
import { InputNumber } from "@/components/UI/InputNumber";
import { formatPrice } from "@/utils/formats";
import { Button } from "@/components/UI/Button";
import { useDispatch, useSelector } from "react-redux";
import { cartTotalSelector } from "@/features/cart/selectors";
import { cartActions } from "@/features/cart/slice";
import { useEffect } from "react";

type LoaderResponse = Array<CartProduct>;

export const loader: LoaderFunction = async () => {
  const { data } = await api.get<LoaderResponse>("/cart");

  return json<LoaderResponse>(data);
};

const CheckoutIndexRoute = () => {
  const cart = useLoaderData<LoaderResponse>();
  const dispatch = useDispatch();
  const cartTotal = useSelector(cartTotalSelector);

  useEffect(() => {
    dispatch(cartActions.setCartProduct(cart));
  }, []);
  return (
    <>
      {cart.map(({ product, quantity, id }) => (
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
            <div className="flex gap-2 cursor-pointer">
              <InputNumber
                min={1}
                defaultValue={quantity}
                onChange={(quantityUpdated) => {
                  dispatch(
                    cartActions.updateCartProduct({
                      id,
                      product,
                      quantity: Number(quantityUpdated),
                    })
                  );
                }}
              />
              <button
                onClick={() => {
                  dispatch(cartActions.deleteCartProduct(id));
                }}
              >
                <span className="p-2 flex gap-1 items-center bg-base-button rounded-md">
                  <Trash size={16} color="#8047F8" />
                  <p className="typography-button-s text-base-text">Remover</p>
                </span>
              </button>
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
          <p>R$ {formatPrice(cartTotal.items)}</p>
        </span>
        <span className="typography-regular-s text-base-text flex justify-between mb-3">
          <p>Entrega</p>
          <p>R$ {formatPrice(cartTotal.freight)}</p>
        </span>
        <span className="typography-bold-l text-base-subtitle flex justify-between mb-3">
          <p>Total</p>
          <p>R$ {formatPrice(cartTotal.items + cartTotal.freight)}</p>
        </span>
      </article>
      <Button variant="secondary">Confirmar pedido</Button>
    </>
  );
};

export default CheckoutIndexRoute;

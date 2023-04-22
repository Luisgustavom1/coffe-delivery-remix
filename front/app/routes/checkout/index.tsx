import type { LoaderFunction } from "@remix-run/node";
import type { Cart, CartProduct } from "@/@types/Api/Cart";
import { json } from "@remix-run/node";
import { Link, useLoaderData } from "@remix-run/react";
import { Trash } from "phosphor-react";
import api from "@/services/api";
import { InputNumber } from "@/components/UI/InputNumber";
import { formatPrice } from "@/utils/formats";
import { Button } from "@/components/UI/Button";
import { useDispatch, useSelector } from "react-redux";
import {
  cartSelector,
  cartTotalSelector,
} from "@/features/cart/slice/selectors";
import { CartActions } from "@/features/cart/slice";

let didInit = false;

export const loader: LoaderFunction = async () => {
  const { data } = await api.get<Cart>("/cart/101");

  return json<Cart>(data);
};

const CheckoutIndex = () => {
  const cart = useLoaderData<Cart>();
  const dispatch = useDispatch();
  const cartState = useSelector(cartSelector);
  const cartTotal = useSelector(cartTotalSelector);

  if (!didInit) {
    didInit = true;
    if (cart) dispatch(CartActions.setCartProduct(cart));
  }

  const updateCartQuantity = ({
    product,
    quantity
  }: CartProduct) => {
    if (quantity === 0) {
      dispatch(CartActions.deleteCartProduct(product.id));
      return;
    }

    dispatch(
      CartActions.updateCartProduct({
        product,
        quantity,
      })
    );
  };
  return (
    <>
      {cartState.products.map(({ product, quantity }) => (
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
                defaultValue={quantity}
                onChange={(quantityUpdated) => {
                  updateCartQuantity({
                    product,
                    quantity: quantityUpdated,
                  });
                }}
              />
              <button
                onClick={() => {
                  dispatch(CartActions.deleteCartProduct(product.id));
                }}
                className="disabled:brightness-75 disabled:opacity-80"
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
      {cartState.products.length === 0 ? (
        <p className="typography-bold-m text-yellow-dark text-center mt-2">
          Ops! Parece que você não tem produtos no carrinho{" "}
          <Link
            to="/coffes"
            className="typography-bold-m text-purple hover:underline"
          >
            clique aqui
          </Link>{" "}
          para fazer seu pedido
        </p>
      ) : (
        <Button variant="secondary" type="submit">
          Confirmar pedido
        </Button>
      )}
    </>
  );
};

export default CheckoutIndex;

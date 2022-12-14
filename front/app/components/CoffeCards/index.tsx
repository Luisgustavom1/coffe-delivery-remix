import type { Coffe } from "@/@types/Api";
import * as CartActions from "@/features/cart/actions";
import { cartSelector } from "@/features/cart/selectors";
import { formatPrice } from "@/utils/formats";
import classNames from "classnames";
import { ShoppingCart } from "phosphor-react";
import React from "react";
import { useDispatch, useSelector } from "react-redux";
import { InputNumber } from "@/components/UI/InputNumber";

interface ICoffeCardsProps {
  coffe: Coffe;
}

export const CoffeCards = ({ coffe }: ICoffeCardsProps) => {
  const dispatch = useDispatch();
  const cart = useSelector(cartSelector);

  const productAlreadyExistsInCart = cart.find(({ product }) => product?.id === coffe.id);

  const handleSubmit = <T extends HTMLFormElement>(e: React.FormEvent<T>) => {
    e.preventDefault();

    const formData = new FormData(e.target as T);
    const quantity = Number(formData.get("quantity"));

    if (productAlreadyExistsInCart && !quantity) {
      dispatch(CartActions.deleteCartProduct(productAlreadyExistsInCart.id));
      return;
    }

    if (!productAlreadyExistsInCart) {
      dispatch(
        CartActions.addCartProduct({
          product: coffe,
          quantity: quantity,
        })
      );
    }

    if (productAlreadyExistsInCart) {
      dispatch(
        CartActions.updateCartProduct({
          quantity: quantity,
          id: productAlreadyExistsInCart.id,
          product: coffe,
        })
      );
    }
  };

  return (
    <div
      className={classNames(
        "flex flex-col items-center justify-center px-6 py-5 w-64 bg-base-card rounded-tl-md rounded-br-md rounded-tr-[36px] rounded-bl-[36px]",
        productAlreadyExistsInCart && "coffe-card-selected"
      )}
    >
      <header className="-mt-14 mb-4">
        <img
          src={coffe.img}
          alt={`Foto do ${coffe.title} - ${coffe.description}`}
          className="mx-auto"
        />
        <div className="flex justify-center gap-1 select-none">
          {coffe.categories.map((categorie, i) => (
            <span
              key={i}
              className="bg-yellow-light text-yellow-dark inline-flex justify-center rounded-full py-1 px-2 mt-3"
            >
              <p className="typography-tag-main" key={categorie}>
                {categorie}
              </p>
            </span>
          ))}
        </div>
      </header>

      <article className="text-center mb-8">
        <h3 className="typography-title-s text-base-subtitle mb-2">
          {coffe.title}
        </h3>
        <p className="typography-regular-s text-base-label">
          {coffe.description}
        </p>
      </article>

      <footer className="flex items-center gap-6">
        <div className="text-base-text flex gap-1">
          <p className="typography-regular-s">R$</p>
          <h6 className="typography-title-m">{formatPrice(coffe.price)}</h6>
        </div>

        <form className="flex gap-4" onSubmit={handleSubmit}>
          <InputNumber
            defaultValue={productAlreadyExistsInCart?.quantity || 0}
            name="quantity"
            min={0}
          />
          <button
            className={classNames(
              "p-2 rounded-md flex bg-purple-dark border-none hover:transition-all hover:brightness-90",
              productAlreadyExistsInCart && 'bg-opacity-70 hover:brightness-100'
            )}
            type="submit"
            disabled={!!productAlreadyExistsInCart}
          >
            <ShoppingCart size={22} weight="fill" color="#fff" />
          </button>
        </form>
      </footer>
    </div>
  );
};

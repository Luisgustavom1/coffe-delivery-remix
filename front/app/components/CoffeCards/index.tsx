import type { Coffe } from "@/@types/Api";
import * as CartActions from "@/features/cart/actions";
import { cartSelector } from "@/features/cart/selectors";
import { formatPrice } from "@/utils/formats";
import classNames from "classnames";
import { ShoppingCart } from "phosphor-react";
import React, { useState } from "react";
import { useDispatch, useSelector } from "react-redux";
import { InputNumber } from "@/components/UI/InputNumber";

interface ICoffeCardsProps {
  coffe: Coffe;
}

export const CoffeCards = ({ coffe }: ICoffeCardsProps) => {
  const dispatch = useDispatch();
  const cart = useSelector(cartSelector);
  
  const productAlreadyExistsInCart = cart.find(({ product }) => product?.id === coffe.id);

  const [quantityInputValue, setQuantityInputValue] = useState(productAlreadyExistsInCart?.quantity || 0)

  const handleSubmit = <T extends HTMLFormElement>(e: React.FormEvent<T>) => {
    e.preventDefault();

    if (productAlreadyExistsInCart && !quantityInputValue) {
      dispatch(CartActions.deleteCartProduct(productAlreadyExistsInCart.id));
      return;
    }

    if (!productAlreadyExistsInCart) {
      dispatch(
        CartActions.addCartProduct({
          product: coffe,
          quantity: quantityInputValue,
        })
      );
    }

    if (productAlreadyExistsInCart) {
      dispatch(
        CartActions.updateCartProduct({
          quantity: quantityInputValue,
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
            onChange={(newQuantity) => setQuantityInputValue(newQuantity)}
            value={quantityInputValue}
          />
          <button
            className={classNames(
              "p-2 rounded-md flex bg-purple-dark border-none hover:transition-all hover:brightness-90",
              productAlreadyExistsInCart && quantityInputValue === productAlreadyExistsInCart.quantity && 'bg-opacity-70 brightness-75 hover:brightness-75'
            )}
            type="submit"
            disabled={!!productAlreadyExistsInCart && quantityInputValue === productAlreadyExistsInCart.quantity}
          >
            <ShoppingCart size={22} weight="fill" color="#fff" />
          </button>
        </form>
      </footer>
    </div>
  );
};

import Logo from "public/assets/svg/Logo.svg";
import { cartSelector } from "@/features/cart/slice/selectors";
import { MapPin, ShoppingCart } from "phosphor-react";
import { useSelector } from "react-redux";
import { Link } from "react-router-dom";
import { Chips } from "../Chips";

export const Header = () => {
  const cart = useSelector(cartSelector);

  return (
    <header className="max-w-screen-xl flex items-center justify-between mx-auto py-8 px-4">
      <a href="/coffes">
        <img
          src={Logo}
          alt="Logo com um escrito Coffe delivery, 'coffe' em cima com fonte maior e em negrito e 'delivery' em baixo com fonte menor, além de um desenho de um copo de café ao lado na cor roxa clara"
        />
      </a>

      <span className="flex gap-3">
        <Chips
          variant="secondary"
          text="Porto Alegre, RS"
          icon={<MapPin size={22} weight="fill" />}
        />
        <div className="relative">
          <span className="typography-tag-s leading-none w-5 h-5 bg-yellow-dark py-1 text-white rounded-full flex items-center justify-center absolute -top-[10px] -right-[10px]">
            {cart.length}
          </span>
          <Link to="/checkout">
            <Chips
              icon={
                <ShoppingCart size={22} color="currentColor" weight="fill" />
              }
            />
          </Link>
        </div>
      </span>
    </header>
  );
};

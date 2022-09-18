import Logo from "@/assets/svg/Logo.svg";
import { quantityOfProductsSelector } from "@/features/cart/selectors";
import { MapPin, ShoppingCart } from "phosphor-react";
import { useSelector } from "react-redux";
import { Chips } from "../Chips";

export const Header = () => {
  const quantityOfProducts = useSelector(quantityOfProductsSelector);

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
            {quantityOfProducts}
          </span>
          <Chips
            icon={<ShoppingCart size={22} color="currentColor" weight="fill" />}
          />
        </div>
      </span>
    </header>
  );
};

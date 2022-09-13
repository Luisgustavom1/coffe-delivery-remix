import Logo from "@/assets/svg/Logo.svg";
import { MapPin, ShoppingCart } from "phosphor-react";
import { Chips } from "../Chips";

export const Header = () => {
  return (
    <header className="max-w-screen-xl flex items-center justify-between mx-auto py-8 px-4">
      <img
        src={Logo}
        alt="Logo com um escrito Coffe delivery, 'coffe' em cima com fonte maior e em negrito e 'delivery' em baixo com fonte menor, além de um desenho de um copo de café ao lado na cor roxa clara"
      />

      <span className="flex gap-3">
        <Chips
          variant="secondary"
          text="Porto Alegre, RS"
          icon={<MapPin size={22} weight="fill" />}
        />
        <Chips
          icon={<ShoppingCart size={22} color="currentColor" weight="fill" />}
        />
      </span>
    </header>
  );
};

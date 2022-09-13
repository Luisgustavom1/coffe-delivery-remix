import CoffeTeste from "@/assets/img/coffes/default.png";
import { ShoppingCart } from "phosphor-react";
import { InputNumber } from "../UI/InputNumber";

export const CoffeCards = () => {
  return (
    <div className="flex flex-col items-center justify-center px-6 py-5 w-64 bg-base-card rounded-tl-md rounded-br-md rounded-tr-[36px] rounded-bl-[36px]">
      <header className="text-center -mt-10 mb-4">
        <img src={CoffeTeste} alt="Foto do café tradicional" />
        <span className="bg-yellow-light text-yellow-dark inline-flex justify-center rounded-full py-1 px-2 mt-3">
          <p className="typography-tag-main">Tradicional</p>
        </span>
      </header>

      <article className="text-center mb-8">
        <h3 className="typography-title-s text-base-subtitle mb-2">
          Expresso Cremoso
        </h3>
        <p className="typography-regular-s text-base-label">
          Café expresso tradicional com espuma cremosa
        </p>
      </article>

      <footer className="flex items-center gap-6">
        <div className="text-base-text flex gap-1">
          <p className="typography-regular-s">R$</p>
          <h6 className="typography-title-m">9,90</h6>
        </div>

        <div className="flex gap-4">
          <InputNumber value={1} onChange={() => ({})} />
          <span className="p-2 rounded-md flex bg-purple-dark cursor-pointer hover:transition-all hover:brightness-90">
            <ShoppingCart size={22} weight="fill" color="#fff" />
          </span>
        </div>
      </footer>
    </div>
  );
};

import type { Categories } from "@/@types/api";
import { ShoppingCart } from "phosphor-react";
import { InputNumber } from "../UI/InputNumber";

interface ICoffeCardsProps {
  imgPath: string;
  title: string;
  description: string;
  categories: Array<Categories>;
  price: number;
}

export const CoffeCards = ({
  imgPath,
  categories,
  title,
  description,
  price,
}: ICoffeCardsProps) => {
  return (
    <div className="flex flex-col items-center justify-center px-6 py-5 w-64 bg-base-card rounded-tl-md rounded-br-md rounded-tr-[36px] rounded-bl-[36px]">
      <header className="text-center -mt-14 mb-4">
        <img src={imgPath} alt={`Foto do ${title} - ${description}`} />
        <span className="bg-yellow-light text-yellow-dark inline-flex justify-center rounded-full py-1 px-2 mt-3">
          {categories.map((categorie) => (
            <p className="typography-tag-main" key={categorie}>
              {categorie}
            </p>
          ))}
        </span>
      </header>

      <article className="text-center mb-8">
        <h3 className="typography-title-s text-base-subtitle mb-2">{title}</h3>
        <p className="typography-regular-s text-base-label">{description}</p>
      </article>

      <footer className="flex items-center gap-6">
        <div className="text-base-text flex gap-1">
          <p className="typography-regular-s">R$</p>
          <h6 className="typography-title-m">{price}</h6>
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

import { Coffee, Package, ShoppingCart, Timer } from "phosphor-react";
import CoffeIlustration from "@/assets/svg/Coffe.svg";
import { CircleIcon } from "@/components/UI/CircleIcon";
import { CoffeCards } from "@/components/CoffeCards";

export default function Home() {
  return (
    <div className="max-w-screen-xl mt-24 px-4 mx-auto">
      <section className="flex items-center justify-center gap-9 flex-wrap md:flex-nowrap">
        <article className="w-full md:w-3/5">
          <div>
            <h1 className="mb-4 typography-title-xl text-base-title">
              Encontre o café perfeito para qualquer hora do dia
            </h1>
            <p className="text-xl text-base-subtitle">
              Com o Coffee Delivery você recebe seu café onde estiver, a
              qualquer hora
            </p>
          </div>
          <div className="flex flex-wrap gap-4 mt-16 text-base-text">
            <span className="flex items-center gap-3">
              <CircleIcon color="yellow-dark">
                <ShoppingCart size={16} weight="fill" />
              </CircleIcon>
              <p className="typography-regular-m">Compra simples e segura</p>
            </span>
            <span className="flex items-center gap-3">
              <CircleIcon color="base-text">
                <Package size={22} weight="fill" />
              </CircleIcon>
              <p className="typography-regular-m">
                Embalagem mantém o café intacto
              </p>
            </span>
            <span className="flex items-center gap-3">
              <CircleIcon color="yellow">
                <Timer size={22} weight="fill" />
              </CircleIcon>
              <p className="typography-regular-m">Entrega rápida e rastreada</p>
            </span>
            <span className="flex items-center gap-3">
              <CircleIcon color="purple">
                <Coffee size={22} weight="fill" />
              </CircleIcon>
              <p className="typography-regular-m">
                O café chega fresquinho até você
              </p>
            </span>
          </div>
        </article>
        <aside className="min-w-[320px]">
          <img
            src={CoffeIlustration}
            alt="Ilustração de um copo de café branco e tampa preta, com um fundo amarelo escuro, grãos de café em baixo e ao lado do copo, além de pó de café no canto direito do copo"
          />
        </aside>
      </section>
      <main className="mt-32">
        <h2 className="typography-title-l text-base-subtitle">Nossos cafés</h2>

        <div className="mt-14 flex flex-wrap gap-x-8 gap-y-10">
          <CoffeCards />
          <CoffeCards />
          <CoffeCards />
          <CoffeCards />
        </div>
      </main>
    </div>
  );
}

import { Coffee, Package, ShoppingCart, Timer } from "phosphor-react";
import CoffeIlustration from "@/assets/svg/Coffe.svg";
import { IconChips } from "@/components/UI/IconChips";

export default function Home() {
  return (
    <div className="max-w-screen-xl mt-24 px-4 mx-auto">
      <section className="flex items-center justify-center gap-9 flex-wrap md:flex-nowrap">
        <article className="w-full md:w-3/5">
          <div>
            <h1 className="mb-4 text-5xl font-extrabold">
              Encontre o café perfeito para qualquer hora do dia
            </h1>
            <p className="text-xl">
              Com o Coffee Delivery você recebe seu café onde estiver, a
              qualquer hora
            </p>
          </div>
          <div className="flex flex-wrap gap-4 mt-16">
            <span className="flex items-center gap-3">
              <IconChips color="yellow-dark">
                <ShoppingCart size={16} weight="fill" />
              </IconChips>
              <p className="typography-regular-m">Compra simples e segura</p>
            </span>
            <span className="flex items-center gap-3">
              <IconChips color="base-text">
                <Package size={22} weight="fill" />
              </IconChips>
              <p className="typography-regular-m">
                Embalagem mantém o café intacto
              </p>
            </span>
            <span className="flex items-center gap-3">
              <IconChips color="yellow">
                <Timer size={22} weight="fill" />
              </IconChips>
              <p className="typography-regular-m">Entrega rápida e rastreada</p>
            </span>
            <span className="flex items-center gap-3">
              <IconChips color="purple">
                <Coffee size={22} weight="fill" />
              </IconChips>
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
        <h2 className="text-2xl">Nossos cafés</h2>

        <div className="mt-14">

        </div>
      </main>
    </div>
  );
}

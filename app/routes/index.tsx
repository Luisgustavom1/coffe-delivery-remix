import CoffeIlustration from "@/assets/svg/Coffe.svg";

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
        </article>
        <aside className="min-w-[320px]">
          <img src={CoffeIlustration} alt="Ilustração de um copo de café branco e tampa preta, com um fundo amarelo escuro, grãos de café em baixo e ao lado do copo, além de pó de café no canto direito do copo" />
        </aside>
      </section>
    </div>
  );
}

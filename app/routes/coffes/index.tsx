import type { LoaderFunction } from "@remix-run/node";
import { json } from "@remix-run/node";
import { CoffeCards } from "@/components/CoffeCards";
import api from "@/services/api";
import type { Coffe } from "@/@types/api";
import { useLoaderData } from "@remix-run/react";

type ApiCoffesResponse = Array<Coffe>;

export function ErrorBoundary({ error }: { error: Error }) {
  console.error(error);

  return <div className="mt-14 typography-title-s text-red-500">Opss algo de errado aconteceu...</div>;
}

export const loader: LoaderFunction = async () => {
  const { data } = await api.get<ApiCoffesResponse>("/coffes");

  return json<ApiCoffesResponse>(data);
};

const CoffesIndexRoute = () => {
  const coffes = useLoaderData<ApiCoffesResponse>();

  return (
    <main className="mt-32">
      <h2 className="typography-title-l text-base-subtitle">Nossos cafés</h2>

      <div className="mt-14 flex flex-wrap gap-x-8 gap-y-10">
        {coffes.map(({ id, img, ...rest }) => (
          <CoffeCards key={id} imgPath={img} {...rest} />
        ))}
      </div>
    </main>
  );
};

export default CoffesIndexRoute;

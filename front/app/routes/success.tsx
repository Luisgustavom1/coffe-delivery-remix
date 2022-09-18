import { CircleIcon } from "@/components/UI/CircleIcon";
import { Clock, CurrencyDollar, MapPin } from "phosphor-react";

interface ISuccessRow {
  title: string
  subtitle: string
  icon: JSX.Element
}

const SuccessRow = ({ title, subtitle, icon }: ISuccessRow) => {
  return (
    <span className="flex gap-4 items-center mt-8">
      {icon}
      <div className="typography-regular-m text-base-text">
        <p>{title}</p>
        <strong>{subtitle}</strong>
      </div>
    </span>
  );
};

const Success = () => {
  return (
    <div className="p-4">
      <article>
        <div className="mb-10">
          <h1 className="typography-title-l text-yellow-dark">
            Uhu! Pedido confirmado
          </h1>
          <p className="typography-regular-l text-base-subtitle">
            Agora é só aguardar que logo o café chegará até você
          </p>
        </div>
        <section className="bg-white p-10 border-success relative rounded-tl-md rounded-tr-[36px] rounded-br-md rounded-bl-[36px] after:absolute after:-left-[1px] after:-right-[1px] after:-bottom-[1px] after:-top-[1px] after:-z-10 after:rounded-tl-md after:rounded-tr-[36px] after:rounded-br-md after:rounded-bl-[36px]">
          <span className="flex gap-4">
            <CircleIcon color="purple">
              <MapPin size={16} weight="fill" color="#FFF" />
            </CircleIcon>
            <div className="typography-regular-m text-base-text">
              <p>
                Entrega em <strong> Rua João Daniel Martinelli, 102</strong>
              </p>
              <p>Farrapos - Porto Alegre, RS</p>
            </div>
          </span>
          <SuccessRow
            icon={
              <CircleIcon color="yellow">
                <Clock size={16} weight="fill" />
              </CircleIcon>
            }
            title="Previsão de entrega"
            subtitle="20 min - 30 min"
          />
          <SuccessRow
            icon={
              <CircleIcon color="yellow-dark">
                <CurrencyDollar size={16} weight="fill" />
              </CircleIcon>
            }
            title="Pagamento de entrega"
            subtitle="Cartão de Crédito"
          />
        </section>
      </article>
      <aside></aside>
    </div>
  );
};

export default Success;

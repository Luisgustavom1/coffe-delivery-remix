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
        <section className="bg-white p-[1px] border-success rounded-tl-md rounded-tr-[36px] rounded-br-md rounded-bl-[36px]">
          <div className="p-[39px]">
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
          </div>
        </section>
      </article>
      <aside></aside>
    </div>
  );
};

export default Success;

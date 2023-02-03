import { Clock, CurrencyDollar, MapPin } from "phosphor-react";
import { CircleIcon } from "@/components/UI/CircleIcon";
import Illustration from "public/assets/svg/Illustration.svg";
import { useSelector } from "react-redux";
import { checkoutDataSelector } from "@/features/checkout/slice/selectors";
import type { PaymentType } from "@/@types/Api/Cart";

interface ISuccessRow {
  title: string;
  subtitle: string;
  icon: JSX.Element;
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

const PaymentTypeLabelMap: Record<PaymentType, string>  = {
  cash: "Dinheiro",
  credit_card: "Crédito",
  debit_card: "Débito"
}

const Success = () => {
  const { address, paymentType } = useSelector(checkoutDataSelector);

  return (
    <div className="mt-20 px-4 flex flex-wrap items-end justify-center gap-4 max-w-screen-xl mx-auto">
      <article className="flex-1 min-w-[340px]">
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
                  Entrega em <strong> {address?.street}, {address?.number}</strong>
                </p>
                <p>{address?.neighborhood} - {address?.city}</p>
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
              subtitle={PaymentTypeLabelMap[paymentType as PaymentType]}
            />
          </div>
        </section>
      </article>
      <img className="h-72" src={Illustration} alt="" />
    </div>
  );
};

export default Success;

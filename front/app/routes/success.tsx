import { CircleIcon } from "@/components/UI/CircleIcon";
import { MapPin } from "phosphor-react";

interface ISuccessRow {
  
}

const SuccessRow = ({
  title,
  subtitle,
  icon
}) => {
  return (
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
        <section className="p-10">
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
        </section>
      </article>
      <aside></aside>
    </div>
  );
};

export default Success;

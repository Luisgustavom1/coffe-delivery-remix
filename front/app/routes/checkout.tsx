import { CartPaymentType } from "@/@types/Api";
import { Button } from "@/components/UI/Button";
import { Card } from "@/components/UI/Card";
import { InputText } from "@/components/UI/InputText";
import { Outlet } from "@remix-run/react";
import {
  Bank,
  CreditCard,
  CurrencyDollar,
  MapPinLine,
  Money,
} from "phosphor-react";
import React from "react";

interface ICheckoutPaymenyTypeButton extends React.ComponentPropsWithoutRef<'button'> {
  active: boolean
  Icon: React.ElementType
}

const CheckoutPaymenyTypeButton = ({
  active,
  Icon,
  children,
  ...rest
}: React.PropsWithChildren<ICheckoutPaymenyTypeButton>) => (
  <Button active={active} {...rest}>
    <Icon size={16} color="#8047F8" />
    {children}
  </Button>
)

const CheckoutRoute = () => {
  const [paymentType, setPaymentType] = React.useState<CartPaymentType>(CartPaymentType.Crédito)

  return (
    <div className="p-4 flex flex-wrap gap-8 max-w-screen-xl mx-auto">
      <section className="flex-1">
        <h1 className="text-base-subtitle typography-title-xs mb-4">
          Complete seu pedido
        </h1>
        <Card className="mb-3">
          <span className="flex items-start gap-2 mb-8">
            <MapPinLine size={22} color="#C47F17" />
            <div>
              <h3 className="typography-regular-m text-base-subtitle mb-[2px]">
                Endereço de Entrega
              </h3>
              <p className="typography-regular-s text-base-text">
                Informe o endereço onde deseja receber seu pedido
              </p>
            </div>
          </span>
          <div className="flex flex-col gap-4">
            <InputText name="cep" placeholder="CEP" />
            <InputText name="street" placeholder="Rua" />
            <span className="flex flex-wrap gap-3">
              <InputText name="number" placeholder="Número" />
              <InputText name="complement" placeholder="Complemento" />
            </span>
            <span className="flex flex-wrap gap-3">
              <InputText name="neighbordhood" placeholder="Bairro" />
              <InputText name="city" placeholder="Cidade" />
            </span>
          </div>
        </Card>
        <Card>
          <span className="flex items-start gap-2 mb-8">
            <CurrencyDollar size={22} color="#8047F8" />
            <div>
              <h3 className="typography-regular-m text-base-subtitle mb-[2px]">
                Pagamento
              </h3>
              <p className="typography-regular-s text-base-text">
                O pagamento é feito na entrega. Escolha a forma que deseja pagar
              </p>
            </div>
          </span>
          <section className="flex gap-3">
            <CheckoutPaymenyTypeButton
              active={paymentType === CartPaymentType.Crédito}
              onClick={() => setPaymentType(CartPaymentType.Crédito)}
              Icon={CreditCard}
            >
              Crédito
            </CheckoutPaymenyTypeButton>
            <CheckoutPaymenyTypeButton
              active={paymentType === CartPaymentType.Débito}
              onClick={() => setPaymentType(CartPaymentType.Débito)}
              Icon={Bank}
            >
              Débito
            </CheckoutPaymenyTypeButton>
            <CheckoutPaymenyTypeButton
              active={paymentType === CartPaymentType.Dinheiro}
              onClick={() => setPaymentType(CartPaymentType.Dinheiro)}
              Icon={Money}
            >
              Dinheiro
            </CheckoutPaymenyTypeButton>
          </section>
        </Card>
      </section>
      <aside className="flex-1">
        <h1 className="text-base-subtitle typography-title-xs mb-4">
          Complete seu pedido
        </h1>
        <Card className="min-w-[440px]">
          <Outlet />
        </Card>
      </aside>
    </div>
  );
};

export default CheckoutRoute;

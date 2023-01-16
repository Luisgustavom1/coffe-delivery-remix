import React from "react";
import { zodResolver } from '@hookform/resolvers/zod';
import { CartPaymentType } from "@/@types/Api";
import Formulary from "@/components/Form/Formulary";
import Input from "@/components/Form/Input";
import { Button } from "@/components/UI/Button";
import { Card } from "@/components/UI/Card";
import { Outlet, useNavigate,  } from "@remix-run/react";
import {
  Bank,
  CreditCard,
  CurrencyDollar,
  MapPinLine,
  Money,
} from "phosphor-react";
import { formCheckout } from "@/features/cart/schemas/address";
import api from "@/services/api";
import { render } from "@react-email/render";
import { CheckoutDone } from "@/features/checkout/emails/checkout-done";
import { useDispatch } from "react-redux";
import * as CartActions from "@/features/cart/slice/actions";

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
  const navigate = useNavigate()
  const dispatch = useDispatch()
  const [paymentType, setPaymentType] = React.useState<CartPaymentType>(CartPaymentType.Crédito)

  return (
    <Formulary
      onSubmit={async (formValues) => { 
        await api.post(`/checkout`, {
          to: formValues.email,
          subject: "Congratulations, thank you for shopping with us",
          message: render(<CheckoutDone />)
        })

        dispatch(CartActions.cleaCartProduct())
        navigate('/coffes')        
      }}
      defaultValues={{
        email: '',
        cep: '',
        street: '',
        number: '',
        complement: '',
        neighborhood: '',
        city: ''
      }}
      resolver={zodResolver(formCheckout)}
    >
    <div className="p-4 flex flex-wrap gap-8 max-w-screen-xl mx-auto">
      <section className="flex-1">
        <h1 className="text-base-subtitle typography-title-xs mb-4">
          Complete seu pedido
        </h1>
        <Card className="mb-3">
          <section className="mb-8">
            <Input required name="email" placeholder="Digite seu email" />
          </section>

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
            <Input required name="cep" placeholder="CEP" />
            <Input required name="street" placeholder="Rua" />
            <span className="flex flex-wrap gap-3">
              <Input required name="number" placeholder="Número" />
              <Input required name="complement" placeholder="Complemento" />
            </span>
            <span className="flex flex-wrap gap-3">
              <Input required name="neighborhood" placeholder="Bairro" />
              <Input required name="city" placeholder="Cidade" />
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
    </Formulary>
  );
};

export default CheckoutRoute;

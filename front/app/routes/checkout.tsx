import React from "react";
import { zodResolver } from '@hookform/resolvers/zod';
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
import { formCheckout } from "@/features/checkout/schemas";
import api from "@/services/api";
import { render } from "@react-email/render";
import { CheckoutDone } from "@/features/checkout/emails/checkout-done";
import { useDispatch } from "react-redux";
import { PAYMENT_TYPE } from "@/@types/Api/Cart";
import { CartActions } from "@/features/cart/slice";
import { CheckoutActions } from "@/features/checkout/slice";

interface CheckoutPaymentTypeButtonProps extends React.ComponentPropsWithoutRef<'input'> {
  Icon: React.ElementType
}

const CheckoutPaymentTypeButton = ({
  Icon,
  children,
  ...rest
}: React.PropsWithChildren<CheckoutPaymentTypeButtonProps>) => (
  <label className="flex-1" htmlFor={rest.id}>
    <input {...rest} className="sr-only peer/input" type='radio' name="paymentType" />
    <Button as='div' className="peer-checked/input:button-active" type='button'>
      <Icon size={16} color="#8047F8" />
      {children}
    </Button>
  </label>
)

const CheckoutRoute = () => {
  const navigate = useNavigate()
  const dispatch = useDispatch()

  return (
    <Formulary
      onSubmit={async (formValues) => { 
        await api.post(`/checkout`, {
          to: formValues.email,
          subject: "Congratulations, thank you for shopping with us",
          message: render(<CheckoutDone />)
        })

        dispatch(CheckoutActions.setCheckoutData({
          address: formValues.address,
          paymentType: formValues.paymentType
        }))
        dispatch(CartActions.clearCart())
        navigate('/success')        
      }}
      defaultValues={{
        email: '',
        paymentType: PAYMENT_TYPE.enum.CREDIT,
        address: {
          cep: '',
          street: '',
          number: '',
          complement: '',
          neighborhood: '',
          city: '',
        }
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
            <Input required name="email" type='email' placeholder="Digite seu email" />
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
            <Input required name="address.cep" placeholder="CEP" />
            <Input required name="address.street" placeholder="Rua" />
            <span className="flex flex-wrap gap-3">
              <Input required name="address.number" placeholder="Número" />
              <Input name="address.complement" placeholder="Complemento" />
            </span>
            <span className="flex flex-wrap gap-3">
              <Input required name="address.neighborhood" placeholder="Bairro" />
              <Input required name="address.city" placeholder="Cidade" />
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
            <CheckoutPaymentTypeButton 
              id={PAYMENT_TYPE.enum.CREDIT}
              Icon={CreditCard}
              defaultChecked
            >
              Crédito
            </CheckoutPaymentTypeButton>
            <CheckoutPaymentTypeButton 
              id={PAYMENT_TYPE.enum.DEBIT}
              Icon={Bank}
            >
              Débito
            </CheckoutPaymentTypeButton>
            <CheckoutPaymentTypeButton 
              id={PAYMENT_TYPE.enum.CASH}
              Icon={Money}
            >
              Dinheiro
            </CheckoutPaymentTypeButton>
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

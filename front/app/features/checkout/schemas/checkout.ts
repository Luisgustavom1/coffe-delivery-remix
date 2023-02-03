import { PAYMENT_TYPE } from "@/@types/Api/Cart";
import { z } from "zod";
import { address } from ".";

export const formCheckout = z.object({
  email: z.string().min(1, { message: 'Email é obrigatório' }),
  paymentType: PAYMENT_TYPE,
  address
})

type FormCheckout = z.infer<typeof formCheckout>;

export type { FormCheckout }

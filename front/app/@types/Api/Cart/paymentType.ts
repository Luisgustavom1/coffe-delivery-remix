import { z } from "zod";

const paymentType = {
  "CREDIT": "credit_card",
  "DEBIT": "debit_card",
  "CASH": "cash"
} as const

const PAYMENT_TYPE = z.nativeEnum(paymentType)

type PaymentType = z.infer<typeof PAYMENT_TYPE>

export { PAYMENT_TYPE } 
export type { PaymentType }
import { z } from "zod";

export const formCheckout = z.object({
  email: z.string().min(1, { message: 'Email é obrigatório' }),
  cep: z.string().min(1, { message: 'CEP é obrigatório' }),
  street: z.string().min(1, { message: 'Rua é obrigatório' }),
  number: z.string().min(1, { message: 'Número é obrigatório' }),
  complement: z.string().nullable(),
  neighborhood: z.string().min(1, { message: 'Bairro é obrigatório' }),
  city: z.string().min(1, { message: 'Cidade é obrigatório' })
})

type FormCheckout = z.infer<typeof formCheckout>;

export type { FormCheckout }

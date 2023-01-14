import { z } from "zod";

export const address = z.object({
  cep: z.string().min(1, { message: 'CEP é obrigatório' }),
  street: z.string().min(1, { message: 'Rua é obrigatório' }),
  number: z.string().min(1, { message: 'Número é obrigatório' }),
  complement: z.string().nullable(),
  neighborhood: z.string().min(1, { message: 'Bairro é obrigatório' }),
  city: z.string().min(1, { message: 'Cidade é obrigatório' })
})

type Address = z.infer<typeof address>;

export type { Address }

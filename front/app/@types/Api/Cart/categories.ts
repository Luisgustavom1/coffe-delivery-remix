import { z } from "zod";

const categories = {
  SPECIAL: "especial",
  TRADITIONAL: "tradicional",
  ALCOHOLIC: "alcoólico",
  WITH_MILK: "com leite",
  ICE_COLD: "gelado"
} as const

const CATEGORIES = z.nativeEnum(categories)

type Categories = z.infer<typeof CATEGORIES>

export { CATEGORIES }
export type { Categories }
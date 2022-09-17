import type { AxiosError } from "axios";

export function* handleError(error: AxiosError<string>): Generator<string> {
  const errorMessage = 
    yield error.response?.data 
      || error.request 
      || error.message
      || 'Algo de errado aconteceu, tente novamente mais tarde'

  return errorMessage
}
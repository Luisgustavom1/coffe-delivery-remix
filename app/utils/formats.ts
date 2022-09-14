export const formatDecimals = (value: number) => String(value)
  .replace(/\D/i, '')
  .replace(/(\d)(\d{2})$/, '$1,$2')

import type { ComponentPropsWithoutRef, ReactNode } from "react";
import type { FieldValues} from "react-hook-form";
import { FormProvider, useForm } from "react-hook-form";

interface IFormularyProps extends ComponentPropsWithoutRef<'form'> {
  children: ReactNode;
  onSubmit: (data: FieldValues) => void;
}

export default function Formulary({
  children,
  onSubmit,
  ...rest
}: IFormularyProps) {
  const methods = useForm({
    reValidateMode: "onSubmit",
    mode: "onSubmit",
  });

  return (
    <FormProvider {...methods}>
      <form onSubmit={methods.handleSubmit(onSubmit)} {...rest}>
        {children}
      </form>
    </FormProvider>
  );
}
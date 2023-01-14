import type { ReactNode } from "react";
import type { FieldValues, UseFormProps} from "react-hook-form";
import { FormProvider, useForm } from "react-hook-form";

interface IFormularyProps<TFieldValues extends FieldValues> extends UseFormProps<TFieldValues> {
  children: ReactNode;
  onSubmit: (data: TFieldValues) => void;
}

export default function Formulary<TFieldValues extends FieldValues>({
  children,
  onSubmit,
  defaultValues,
  ...rest
}: IFormularyProps<TFieldValues>) {
  const methods = useForm({
    reValidateMode: "onSubmit",
    mode: "onSubmit",
    defaultValues,
    ...rest
  });

  return (
    <FormProvider {...methods}>
      <form onSubmit={methods.handleSubmit(onSubmit)}>
        {children}
      </form>
    </FormProvider>
  );
}
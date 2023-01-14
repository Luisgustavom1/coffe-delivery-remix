import type { ComponentPropsWithoutRef } from "react";
import { Controller } from "react-hook-form";
import { InputText } from "@/components/UI/InputText";

interface ruleContent {
  value: number;
  message: string;
}

interface IInput extends ComponentPropsWithoutRef<"input"> {
  name: string;
  rules?: Record<string, ruleContent>;
}

export default function Input({
  name,
  defaultValue,
  onChange,
  onBlur,
  required = false,
  rules,
  ...rest
}: IInput) {
  return (
    <Controller
      name={name}
      defaultValue={defaultValue || ""}
      rules={{
        required: required ? "Este campo é obrigatório" : false,
        ...rules,
      }}
      render={({ field, fieldState }) => (
        <InputText
          name={name}
          onBlur={(e) => {
            field.onBlur();

            if (onBlur) {
              onBlur(e);
            }
          }}
          onChange={(e) => {
            field.onChange(e);

            if (onChange) {
              onChange(e);
            }
          }}
          value={field.value}
          error={fieldState.error?.message}
          {...rest}
        />
      )}
    />
  );
}
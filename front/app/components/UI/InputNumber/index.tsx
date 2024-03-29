import React from "react";
import { Plus, Minus } from "phosphor-react";

type InputNumberProps = {
  defaultValue?: number;
  onChange?: (value: number) => void;
} & Omit<React.ComponentPropsWithoutRef<"input">, 'onChange'>;

export const InputNumber = ({
  defaultValue: value,
  onChange,
  ...rest
}: InputNumberProps) => {
  const [currentValue, setCurrentValue] = React.useState(value || 0);

  React.useEffect(() => {
    if (onChange && currentValue !== value) onChange(currentValue);
  }, [currentValue, onChange, value]);
  return (
    <span className="rounded-md bg-base-button p-2 flex gap-2 w-min">
      <Minus
        onClick={() => {
          if (rest.min !== undefined && rest.min === currentValue) return;
          setCurrentValue((prevState) => prevState - 1);
        }}
        size={14}
        weight="fill"
        className="text-purple cursor-pointer hover:text-purple-dark transition-colors"
      />
      <p className="text-base-title typography-regular-m leading-none select-none">
        {currentValue}
      </p>
      <input
        className="sr-only"
        type="number"
        value={currentValue}
        readOnly
        {...rest}
      />
      <Plus
        onClick={() => setCurrentValue((prevState) => prevState + 1)}
        size={14}
        weight="fill"
        className="text-purple cursor-pointer hover:text-purple-dark transition-colors"
      />
    </span>
  );
};

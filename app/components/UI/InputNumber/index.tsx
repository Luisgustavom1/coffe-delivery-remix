import { Plus, Minus } from "phosphor-react";
import React from "react";

interface IInputNumberProps {
  value: number;
  onChange: (value: number) => void;
}

export const InputNumber = ({
  value,
  onChange,
}: IInputNumberProps) => {
  const [currentValue, setCurrentValue] = React.useState(value || 0);

  React.useEffect(() => {
    onChange(currentValue);
  }, [currentValue, onChange]);
  return (
    <span className="rounded-md bg-base-button p-2 flex gap-2 items-center">
      <Minus
        onClick={() => setCurrentValue((prevState) => prevState - 1)}
        size={14}
        weight="fill"
        className="text-purple cursor-pointer"
      />
      <p className="text-base-title typography-regular-m leading-none">
        {currentValue}
      </p>
      <Plus
        onClick={() => setCurrentValue((prevState) => prevState + 1)}
        size={14}
        weight="fill"
        className="text-purple cursor-pointer"
      />
    </span>
  );
};

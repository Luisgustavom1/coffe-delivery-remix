import classNames from "classnames";
import type { ComponentPropsWithoutRef } from "react";

interface IInputTextProps extends ComponentPropsWithoutRef<"input"> {
  error?: string
}

export const InputText = ({ className, error, ...rest }: IInputTextProps) => {
  return (
    <label className="w-auto flex-1 h-min grid">
      <input
        className={classNames(
          className,
          "flex-1 bg-base-button p-3 rounded placeholder:typography-regular-s placeholder:text-base-label focus:outline-purple"
        )}
        {...rest}
      />
      {error && <p className="mt-1 pl-1 typography-button-s text-red-600">{error}</p>}
    </label>
  );
};

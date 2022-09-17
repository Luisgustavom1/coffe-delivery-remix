import classNames from "classnames";
import type { ComponentPropsWithoutRef } from "react";

interface IInputTextProps extends ComponentPropsWithoutRef<"input"> {}

export const InputText = ({ className, ...rest }: IInputTextProps) => {
  return (
    <input
      className={classNames(
        className,
        "flex-1 bg-base-button p-3 rounded placeholder:typography-regular-s placeholder:text-base-label focus:outline-purple"
      )}
      {...rest}
    />
  );
};

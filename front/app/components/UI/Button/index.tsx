import classNames from "classnames";
import type { ComponentPropsWithoutRef } from "react";

type ButtonProps = ComponentPropsWithoutRef<'button'> & {
  variant?: "primary" | "secondary";
  active?: boolean;
}

export const Button = ({
  className,
  active,
  variant = "primary",
  as: Element = 'button',
  ...rest
}: ComponentPropsWithAs<ButtonProps>) => {
  return (
    <Element
      className={classNames(
        className,
        variant === "primary" &&
          "bg-base-button text-base-text typography-button-s",
        variant === "secondary" && "bg-yellow text-white text-sm font-bold",
        active && "button-active",
        "w-full flex items-center justify-center gap-3 flex-1 p-4 rounded-md border-none uppercase cursor-pointer hover:transition-all hover:brightness-95"
      )}
      {...rest}
    />
  );
};

import classNames from "classnames";
import React from "react";

export const Card = ({
  children,
  className,
  ...rest
}: React.PropsWithChildren<React.ComponentPropsWithoutRef<"div">>) => {
  return (
    <div
      className={classNames(className, "rounded-md bg-base-card p-10")}
      {...rest}
    >
      {children}
    </div>
  );
};

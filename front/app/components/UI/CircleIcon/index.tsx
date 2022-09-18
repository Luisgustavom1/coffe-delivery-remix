import React from "react";

interface CircleIconProps {
  color: "yellow-dark" | "yellow" | "purple" | "base-text";
}

export const CircleIcon = ({
  children,
  color,
}: React.PropsWithChildren<CircleIconProps>) => {
  const currentBgChip = {
    "yellow-dark": "bg-yellow-dark",
    yellow: "bg-yellow",
    purple: "bg-purple",
    "base-text": "bg-base-text",
  }[color];

  return (
    <span
      className={`w-8 h-8 flex items-center justify-center flex-shrink-0 rounded-full ${currentBgChip} p-2 text-white`}
    >
      {children}
    </span>
  );
};

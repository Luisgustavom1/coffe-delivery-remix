import classNames from "classnames";

interface IChipsProps {
  text?: string;
  icon?: JSX.Element;
  variant?: "primary" | "secondary";
}

export const Chips = ({ text, icon, variant = "primary" }: IChipsProps) => {
  return (
    <span
      className={classNames(
        variant === "primary" && `bg-yellow-light text-yellow-dark`,
        variant === "secondary" && `bg-purple-light text-purple-dark`,
        icon && text && "gap-1 py-[10px]",
        !icon && text && "py-1",
        "inline-flex justify-center items-center rounded-md px-2 py-2 typography-regular-s"
      )}
    >
      <>
        {icon}
        <p>{text}</p>
      </>
    </span>
  );
};

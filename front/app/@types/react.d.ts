type As<P = any> = React.ElementType<P>;

type ComponentPropsWithAs<P, T extends As = As> = P & {
  as?: T
}
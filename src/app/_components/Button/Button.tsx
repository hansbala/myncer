import { type PropsWithChildren } from "react";

interface ButtonProps extends React.ButtonHTMLAttributes<HTMLButtonElement> {
  // more props here
}

export default function Button({ children, ...props }: PropsWithChildren<ButtonProps>) {
  return (
    <button type="button" className="rounded-md flex-grow bg-white/10 px-10 py-3 font-semibold transition hover:bg-white/20" {...props}>{children}</button>
  )
}

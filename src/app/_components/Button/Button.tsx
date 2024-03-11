import { type PropsWithChildren } from 'react'

interface ButtonProps extends React.ButtonHTMLAttributes<HTMLButtonElement> {
  // more props here
  className?: string
}

export default function Button({
  children,
  className,
  ...props
}: PropsWithChildren<ButtonProps>) {
  return (
    <button
      type="button"
      className={`flex-grow rounded-md bg-white/10 px-10 py-3 font-semibold transition hover:bg-white/20 ${className}`}
      {...props}
    >
      {children}
    </button>
  )
}

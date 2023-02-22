import "./button.scss"

type Props = {
    title: string
    size?: string
}

export default function Button({title, size="medium"}: Props) {
  return (
    <button className={`btn btn-${size}`}>{title}</button>
  )
}

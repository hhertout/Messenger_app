import "./button.scss"

type Props = {
    title: string
    size?: string
    color?: string
}

export default function Button({title, size="medium", color="primary"}: Props) {
  return (
    <button className={`btn btn-${size} btn-${color}`}>{title}</button>
  )
}

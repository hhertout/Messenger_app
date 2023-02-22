import "./formraw.scss"

type Props = {
  name: string
  title: string
  type: string
  onChange?: React.ChangeEventHandler<HTMLInputElement>
  value?: string
  inputRef?: React.RefObject<HTMLInputElement>
  isRequired?: boolean
  isValid?: boolean
}

export const FormRaw = ({ name, title, type, onChange, value, inputRef, isRequired = false, isValid= true }: Props) => {
  return (
    <div className="form-raw">
      <label htmlFor={name}>{title}</label>
      <input id={name} name={name} type={type} onChange={onChange} value={value} ref={inputRef} required={isRequired} className={isValid ? "outline-invalid": ""}/>
    </div>
  )
}

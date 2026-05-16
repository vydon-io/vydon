import { ReactElement } from 'react';
import { Control, FieldValues, UseFormReturn } from 'react-hook-form';
import useFormPersist from './useFormPersist';

interface FormPersistProps<T extends FieldValues> {
  form: UseFormReturn<T>;
  formKey: string;
}
const isBrowser = () => typeof window !== 'undefined';

export default function FormPersist<T extends FieldValues>(
  props: FormPersistProps<T>
): ReactElement {
  const { form, formKey } = props;
  useFormPersist(formKey, {
    // react-hook-form 7.56 introduced a 3-parameter `Control<T, TContext, TTransformedValues>`;
    // useFormPersist consumes the legacy single-parameter form, so widen here.
    control: form.control as Control<FieldValues>,
    setValue: form.setValue,
    storage: isBrowser() ? window.sessionStorage : undefined,
  });
  return <></>;
}

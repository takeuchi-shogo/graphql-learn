import { Button, Form, TextField } from '@/components/ui'
import { useForm } from 'react-hook-form'
import type { FieldValues } from 'react-hook-form'

export function LoginForm() {
  const { handleSubmit } = useForm({
    defaultValues: {
      email: '',
      password: '',
    },
  })

  const onSubmit = (data: FieldValues) => {
    console.log('data', data)
  }

  return (
    <>
      <div className='flex flex-col gap-4'>Form</div>
      <Form onSubmit={handleSubmit(onSubmit)} className='flex flex-col gap-4'>
        <TextField name='email' label='Email' isRequired errorMessage={'email is required'} />
        <TextField
          name='password'
          label='Password'
          isRequired
          errorMessage={'password is required'}
        />
        <Button type='submit'>Login</Button>
      </Form>
    </>
  )
}

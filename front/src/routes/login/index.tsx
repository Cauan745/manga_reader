import { Button } from '#/components/ui/button'
import { Field, FieldGroup, FieldLabel } from '#/components/ui/field'
import { Input } from '#/components/ui/input'
import { createFileRoute } from '@tanstack/react-router'

export const Route = createFileRoute('/login/')({
  component: RouteComponent,
})

function RouteComponent() {
  return (
    <div className='bg-gray-600 flex justify-center items-center flex-col gap-5 h-screen'>
      <div>
        <FieldGroup>
          <Field>
            <FieldLabel>Email</FieldLabel>
            <Input type='email' placeholder='name@example.com' className='bg-black'></Input>
          </Field>
          <Field>
            <FieldLabel>Password</FieldLabel>
            <Input type='password' placeholder="********"></Input>
          </Field>
          <Field orientation="horizontal">
            <Button variant="outline" type='reset'>Reset</Button>
            <Button type='submit'>Submit</Button>
          </Field>
        </FieldGroup>
      </div>
    </div>
  )
}

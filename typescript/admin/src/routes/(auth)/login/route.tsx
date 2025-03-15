import { createFileRoute } from "@tanstack/react-router";

export const Route = createFileRoute('/(auth)/login')({
  component: LoginRoute,
})

export function LoginRoute() {
  return <div>Login</div>
}

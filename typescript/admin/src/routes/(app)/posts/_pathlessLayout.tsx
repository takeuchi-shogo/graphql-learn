import { createFileRoute, Outlet } from '@tanstack/react-router'

export const Route = createFileRoute('/(app)/posts/_pathlessLayout')({
  component: RouteComponent,
})

function RouteComponent() {
  return (
    <div className="pathless-layout">
      <Outlet />
    </div>
  )
}

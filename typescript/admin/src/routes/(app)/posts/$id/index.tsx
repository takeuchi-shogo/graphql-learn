import { createFileRoute } from '@tanstack/react-router'

export const Route = createFileRoute('/(app)/posts/$id/')({
  // In a loader
  loader: ({ params }) => params.id,
  // Or in a component
  component: PostIdComponent,
})

function PostIdComponent() {
  // In a component!
  const { id } = Route.useParams()
  return <div>Post ID: {id}</div>
}

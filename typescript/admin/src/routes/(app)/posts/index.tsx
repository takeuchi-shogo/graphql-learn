import { createFileRoute } from "@tanstack/react-router";
import PostsPage from "@/pages/posts/page";
export const Route = createFileRoute('/(app)/posts/')({
  component: PostsRoute,
})

export function PostsRoute() {
  console.log('PostsRoute')
  return <PostsPage />
}

import { createFileRoute } from "@tanstack/react-router";
import PostsPage from "../../pages/posts/page";
export const Route = createFileRoute('/posts/')({
  component: PostsRoute,
})

export function PostsRoute() {
  console.log('PostsRoute')
  return <PostsPage />
}

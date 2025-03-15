import type { PostFragment } from '@/__generated__/graphql'
import { PostListItem } from '../post-list-item'

type Props = {
  posts: PostFragment[]
}

export function PostList({ posts }: Props) {
  return (
    <div>
      {posts.map((post) => (
        <PostListItem key={post.id} post={post} />
      ))}
    </div>
  )
}

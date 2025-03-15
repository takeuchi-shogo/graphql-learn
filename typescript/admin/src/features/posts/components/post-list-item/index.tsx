import type { PostFragment } from '@/__generated__/graphql'

type Props = {
  post: PostFragment
}

export function PostListItem({ post }: Props) {
  return (
    <div>
      <p>{post.content}</p>
      <p>{post.author.displayName}</p>
    </div>
  )
}

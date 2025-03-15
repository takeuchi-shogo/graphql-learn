import { PostFragment } from '@/__generated__/graphql'

type Props = {
  posts: PostFragment[]
}

export function PostList({ posts }: Props) {
  return (
    <div>
      {posts.map((post) => (
        <div key={post.id}>
          <p>{post.content}</p>
          <p>{post.author.displayName}</p>
        </div>
      ))}
    </div>
  )
}

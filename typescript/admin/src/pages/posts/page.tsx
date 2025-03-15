import { useGetPostsQuery } from '@/__generated__/graphql'
import { PostList } from '@/features/posts/components/post-list'
export default function PostsPage() {
  const { data, loading, error } = useGetPostsQuery({
    variables: {
      first: 10,
      after: "1",
    },
  })
  if (loading) return <div>Loading...</div>
  if (error) return <div>Error: {error.message}</div>
  return (
    <>
      {data && <PostList posts={data.posts.edges.map((edge) => edge.node)} />}
    </>
  )
}

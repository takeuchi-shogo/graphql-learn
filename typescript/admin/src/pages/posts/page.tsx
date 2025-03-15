import { useGetPostsQuery } from '../../__generated__/graphql'

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
      {data && data?.posts.edges.map((edge) => (
        <div key={edge.node.id}>
          <p>{edge.node.content}</p>
          <p>{edge.node.author.displayName}</p>
        </div>
      ))}
    </>
  )
}

import { Providers } from '@/components/providers'
import { Toast } from '@/components/ui'
import { Link, Outlet, createRootRoute } from '@tanstack/react-router'
import { TanStackRouterDevtools } from '@tanstack/react-router-devtools'

export const Route = createRootRoute({
  component: () => (
    <Providers>
      <div className='flex gap-2 p-2'>
        <Link to='/' className='[&.active]:font-bold'>
          Home
        </Link>{' '}
        <Link to='/posts' className='text-blue-500 [&.active]:font-bold'>
          Posts
        </Link>
        <Link to={'/posts/$id'} params={{ id: '1' }} className='[&.active]:font-bold'>
          Post 1
        </Link>
        <Link to='/login' className='[&.active]:font-bold'>
          Login
        </Link>
      </div>
      <hr />
      <Toast />
      <Outlet />
      <TanStackRouterDevtools />
    </Providers>
  ),
})

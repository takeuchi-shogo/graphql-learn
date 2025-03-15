import { router } from '@inertiajs/react'
import { ThemeProvider } from './theme-provider'
import { RouterProvider } from 'react-aria-components'
import type { ReactNode } from 'react'

export function Providers({ children }: { children: ReactNode }) {
    return (
        <RouterProvider navigate={(to, options) => router.visit(to, options)}>
            <ThemeProvider defaultTheme="system" storageKey="ui-theme">
                {children}
            </ThemeProvider>
        </RouterProvider>
    )
}

'use client';

import { useRouter } from 'next/navigation';

interface LoginButtonProps {
    children: React.ReactNode;
    mode?: 'modal' | 'redirect';
    asChild?: boolean;
}

export const LoginButton = ({ children, asChild, mode = 'redirect' }: LoginButtonProps) => {
    const router = useRouter();

    const onClick = () => router.push('/sign-in');

    if (mode === 'modal') {
        return <span>TODO: Implementar modal</span>;
    }

    return <span onClick={onClick}>{children}</span>;
};

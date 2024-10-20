import { cn } from '@/lib/utils';

import { Roboto } from 'next/font/google';

const font = Roboto({ subsets: ['latin'], weight: '500' });

interface HeaderProps {
    label: string;
}

export const Header = ({ label }: HeaderProps) => {
    return (
        <div className='w-full flex flex-col gap-y-4 items-center justify-center tracking-tight'>
            <h1 className={cn('text-3xl font-semibold', font.className)}>{label}</h1>
        </div>
    );
};

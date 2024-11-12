'use client';
import { usePathname } from 'next/navigation';

export default function Header() {
    const pathname = usePathname().split('/r/')[1];

    const translations: Record<string, string> = {
        study: 'Estudos',
        settings: 'Configurações'
    };

    return (
        <div className='flex flex-col gap-2'>
            <span className='text-xs text-muted-foreground'>{'publico'.toLocaleUpperCase()}</span>
            <h1 className='text-4xl font-medium'>
                {pathname
                    ? translations[pathname.toLowerCase()] || pathname.charAt(0).toUpperCase() + pathname.slice(1).toLowerCase()
                    : 'Meus Estudos'}
            </h1>
        </div>
    );
}

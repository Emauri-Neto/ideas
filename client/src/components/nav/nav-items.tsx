'use client';

import { HomeIcon, BookCopy } from 'lucide-react';
import { usePathname } from 'next/navigation';
import { Tooltip, TooltipContent, TooltipProvider, TooltipTrigger } from '@/components/ui/tooltip';
import Link from 'next/link';
import { buttonVariants } from '@/components/ui/button';

const NAV = [
    {
        description: 'Inicio',
        href: '/',
        icon: HomeIcon
    },
    {
        description: 'Estudos',
        href: '/r/study',
        icon: BookCopy
    }
];

const NavItems = () => {
    const path = usePathname();

    return (
        <div className='w-full'>
            <TooltipProvider>
                {NAV.map(item => (
                    <Tooltip key={item.href}>
                        <TooltipTrigger asChild>
                            <Link
                                href={item.href}
                                className={buttonVariants({
                                    variant: path === item.href ? 'default' : 'ghost',
                                    className: `mx-1 ${path !== item.href ? 'hover:bg-gray-300 hover:dark:bg-gray-900' : ''}`
                                })}
                            >
                                <item.icon />
                            </Link>
                        </TooltipTrigger>
                        <TooltipContent>
                            <p>{item.description}</p>
                        </TooltipContent>
                    </Tooltip>
                ))}
            </TooltipProvider>
        </div>
    );
};

export default NavItems;

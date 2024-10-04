'use client';

import React from 'react';
import { buttonVariants } from '@/components/ui/button';
import { BookAIcon, HomeIcon } from 'lucide-react';
import { Tooltip, TooltipContent, TooltipProvider, TooltipTrigger } from '@/components/ui/tooltip';
import Link from 'next/link';
import { usePathname } from 'next/navigation';

export const NAV = [
    {
        description: 'Inicio',
        href: '/',
        icon: HomeIcon
    },
    {
        description: 'Estudos',
        href: '/r/study',
        icon: BookAIcon
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
                            <Link href={item.href} className={buttonVariants({ variant: path === item.href ? 'secondary' : 'ghost' })}>
                                <item.icon className='w-5 h-5' />
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

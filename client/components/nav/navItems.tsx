import React from 'react';
import { buttonVariants } from '@/components/ui/button';
import { HomeIcon } from 'lucide-react';
import { Tooltip, TooltipContent, TooltipProvider, TooltipTrigger } from '@/components/ui/tooltip';
import Link from 'next/link';

export const NAV = [
    {
        description: 'Inicio',
        href: '/',
        icon: HomeIcon
    }
];

const NavItems = () => {
    return (
        <div className='w-full'>
            <TooltipProvider>
                {NAV.map(item => (
                    <Tooltip key={item.href}>
                        <TooltipTrigger asChild>
                            <Link href={item.href} className={buttonVariants({ variant: 'ghost' })}>
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

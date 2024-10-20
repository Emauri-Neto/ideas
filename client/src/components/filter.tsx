'use client';

import { useState } from 'react';
import { cn } from '@/lib/utils';
import { Network, Popcorn, BookDashed } from 'lucide-react';

const opts = [
    {
        title: 'Recentes',
        description: 'Ache os estudos mais recentes',
        value: 'recent',
        color: 'text-blue-500',
        icon: Network
    },
    {
        title: 'Populares',
        description: 'Mais famosos entre o público',
        value: 'popular',
        color: 'text-emerald-500',
        icon: Popcorn
    },
    {
        title: 'Meus Estudos',
        description: 'Ver estudos privados',
        value: 'owner',
        color: 'text-purple-600',
        icon: BookDashed
    }
];

const FilterMenu = () => {
    const [filter, setFilter] = useState<string>('recent');

    return (
        <div className='hidden md:flex md:flex-col mr-10'>
            <div className='rounded-md bg-secondary'>
                {opts.map(item => (
                    <div
                        className={cn(
                            'flex flex-row gap-4 items-center m-3 p-3 cursor-pointer',
                            filter === item.value ? 'bg-gray-300 dark:bg-gray-900 rounded-lg' : null
                        )}
                        key={item.value}
                        onClick={() => setFilter(item.value)}
                    >
                        <item.icon className={`h-8 w-8 ${item.color}`} />
                        <div>
                            <h6>{item.title}</h6>
                            <p className='text-sm text-muted-foreground hidden lg:block'>{item.description}</p>
                        </div>
                    </div>
                ))}
            </div>
        </div>
    );
};

export default FilterMenu;

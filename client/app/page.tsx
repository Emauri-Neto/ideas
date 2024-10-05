'use client';

import Footer from '@/components/footer';
import Navbar from '@/components/nav/navbar';
import StudiesArea from '@/components/study/studiesArea';
import MaxWidthWrapper from '@/components/wrapper';
import { cn } from '@/lib/utils';
import { Network, Popcorn, BookDashed } from 'lucide-react';
import { useState } from 'react';

const OPTS = [
    {
        title: 'Recentes',
        description: 'Ache os estudos mais recentes',
        value: 'recent',
        icon: Network
    },
    {
        title: 'Populares',
        description: 'Mais famosos entre o público',
        value: 'popular',
        icon: Popcorn
    },
    {
        title: 'Meus Estudos',
        description: 'Ver estudos privados',
        value: 'owner',
        icon: BookDashed
    }
];

export default function Home() {
    const [filter, setFilter] = useState<string>('recent');

    return (
        <MaxWidthWrapper>
            <Navbar />
            <div className='flex flex-col grid-cols-4 md:grid'>
                <div className='w-11/12'>
                    <div className='hidden md:flex md:flex-col'>
                        <div className='bg-secondary rounded-md'>
                            {OPTS.map(item => (
                                <div
                                    className={cn(
                                        'flex flex-row gap-4 items-center m-4 p-4 cursor-pointer',
                                        filter === item.value ? 'bg-slate-200 dark:bg-gray-900 rounded-lg' : null
                                    )}
                                    key={item.value}
                                    onClick={() => setFilter(item.value)}
                                >
                                    <item.icon className='h-8 w-8 text-blue-500' />
                                    <div>
                                        <h6>{item.title}</h6>
                                        <p className='text-sm text-muted-foreground'>{item.description}</p>
                                    </div>
                                </div>
                            ))}
                        </div>
                    </div>
                </div>
                <div className='col-span-3'>
                    <StudiesArea />
                </div>
            </div>
            <Footer />
        </MaxWidthWrapper>
    );
}

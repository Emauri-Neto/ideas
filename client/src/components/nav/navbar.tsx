'use client';

import { Bell, Search, Moon, Sun } from 'lucide-react';
import { Input } from '@/components/ui/input';
import { Button } from '@/components/ui/button';
import { useTheme } from 'next-themes';

export default function Navbar() {
    const { theme, setTheme } = useTheme();
    return (
        <div className='pl-20 md:pl-64 fixed p-7 w-full flex items-center justify-between gap-4'>
            <div className='flex gap-4 justify-center items-center px-4 w-1/4 max-w-lg'>
                <Input className='bg-transparent !placeholder-muted-foreground' placeholder='Pesquise por pessoas, estudos e linhas de discussão' />
            </div>

            <div className='flex items-center gap-4'>
                <Button className='flex items-center gap-2'>
                    <Bell />
                </Button>
                <Button onClick={() => setTheme(theme === 'light' ? 'dark' : 'light')}>
                    <Sun className='h-[1.2rem] w-[1.2rem] rotate-0 scale-100 transition-all dark:-rotate-90 dark:scale-0' />
                    <Moon className='absolute h-[1.2rem] w-[1.2rem] rotate-90 scale-0 transition-all dark:rotate-0 dark:scale-100' />
                </Button>
            </div>
        </div>
    );
}

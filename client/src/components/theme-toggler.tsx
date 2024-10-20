'use client';

import { Moon, Sun } from 'lucide-react';
import { useTheme } from 'next-themes';
import { Button } from '@/components/ui/button';

const ThemeToggler = () => {
    const { theme, setTheme } = useTheme();

    return (
        <Button onClick={() => setTheme(theme === 'light' ? 'dark' : 'light')} variant='ghost'>
            <Sun className='h-[1.2rem] w-[1.2rem] rotate-0 scale-100 transition-all dark:-rotate-90 dark:scale-0' />
            <Moon className='absolute h-[1.2rem] w-[1.2rem] rotate-90 scale-0 transition-all dark:rotate-0 dark:scale-100' />
        </Button>
    );
};

export default ThemeToggler;

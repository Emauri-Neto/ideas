'use client';

import { Bell, Moon, Sun, ChevronDown, User2Icon, LogOutIcon, Loader2Icon } from 'lucide-react';
import { Input } from '@/components/ui/input';
import { Button } from '@/components/ui/button';
import { useTheme } from 'next-themes';
import { GetUser } from '@/actions/user';
import { User } from '@/types';
import { useQuery } from '@tanstack/react-query';
import {
    DropdownMenu,
    DropdownMenuContent,
    DropdownMenuGroup,
    DropdownMenuLabel,
    DropdownMenuTrigger,
    DropdownMenuSeparator,
    DropdownMenuItem
} from '@/components/ui/dropdown-menu';
import { Avatar, AvatarFallback } from '@/components/ui/avatar';
import { Logout } from '@/actions/auth';
import { useRouter } from 'next/navigation';

export default function Navbar() {
    const { theme, setTheme } = useTheme();
    const router = useRouter();

    const { data, isPending, isFetched } = useQuery({
        queryKey: ['user-data'],
        queryFn: GetUser
    });

    const user = data as unknown as User;

    async function SignOut() {
        await Logout();
        router.refresh();
    }

    return (
        <div className='pl-20 md:pl-64 fixed p-7 w-full flex items-center justify-between gap-4 pr-10'>
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

                <DropdownMenu>
                    <DropdownMenuTrigger className='flex flex-row items-center justify-center'>
                        <Avatar>
                            <AvatarFallback>
                                <Loader2Icon className='animate-spin w-4 h-4' />
                            </AvatarFallback>
                        </Avatar>
                    </DropdownMenuTrigger>
                    <DropdownMenuContent className='mr-4'>
                        <DropdownMenuLabel>{user.email}</DropdownMenuLabel>
                        <DropdownMenuSeparator />
                        <DropdownMenuGroup>
                            <DropdownMenuItem className='text-muted-foreground font-semibold'>
                                <User2Icon />
                                Perfil
                            </DropdownMenuItem>
                        </DropdownMenuGroup>
                        <DropdownMenuSeparator />
                        <DropdownMenuItem className='text-muted-foreground font-semibold' onClick={SignOut}>
                            <LogOutIcon />
                            Sair
                        </DropdownMenuItem>
                    </DropdownMenuContent>
                </DropdownMenu>
            </div>
        </div>
    );
}

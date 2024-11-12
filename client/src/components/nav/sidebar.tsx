'use client';

import { Icons } from '@/components/global/icons';
import { Select, SelectContent, SelectGroup, SelectItem, SelectLabel, SelectTrigger, SelectValue } from '@/components/ui/select';
import { usePathname, useRouter } from 'next/navigation';
import { Separator } from '@/components/ui/separator';
import { MenuItens } from '@/config/constants';
import SidebarItens from '@/components/nav/sidebar-itens';
import { BookA, MenuIcon } from 'lucide-react';
import { Sheet, SheetContent, SheetTrigger } from '@/components/ui/sheet';
import { Button } from '@/components/ui/button';
import Navbar from '@/components/nav/navbar';

interface SidebarProps {
    activeStudyId: string;
}

const studies = [
    {
        id: 'aaaaaa',
        name: 'Estudo legal'
    },
    {
        id: 'bbbbb',
        name: 'Estudo legal3'
    },
    {
        id: 'cccccc',
        name: 'Estudo legal2'
    }
];

export default function Sidebar({ activeStudyId }: SidebarProps) {
    const router = useRouter();
    const pathname = usePathname();

    function OnChange(id: string) {
        router.push(`r/study/${id}`);
    }

    const menu = MenuItens(activeStudyId);

    const SidebarSection = (
        <div className='bg-secondary flex-none relative p-4 h-full w-[250px] flex flex-col gap-4 overflow-hidden'>
            <div className='bg-secondary p-4 justify-center items-center mb-4 absolute top-0 left-0 right-0'>
                <Icons.logo className='w-44 h-16 fill-black dark:fill-white' />
            </div>

            <Select defaultValue={activeStudyId} onValueChange={OnChange}>
                <SelectTrigger className='mt-24 text-muted-foreground bg-transparent'>
                    <SelectValue placeholder='Escolha um estudo'></SelectValue>
                </SelectTrigger>
                <SelectContent className=' backdrop-blur-xl'>
                    <SelectGroup>
                        <SelectLabel>Estudos</SelectLabel>
                        <Separator />
                        {studies.map(ws => (
                            <SelectItem key={ws.id} value={ws.id}>
                                {ws.name}
                            </SelectItem>
                        ))}
                    </SelectGroup>
                </SelectContent>
            </Select>

            <h2 className='mt-4 w-full font-semibold'>Menu</h2>
            <nav className='w-full'>
                <ul>
                    {menu.map(item => (
                        <SidebarItens
                            href={item.href}
                            icon={item.icon}
                            selected={pathname === item.href}
                            title={item.title}
                            key={item.href}
                            notifications={1}
                        />
                    ))}
                </ul>
            </nav>

            <Separator className='w-full bg-gray-400 dark:bg-gray-600' />
            <h2 className='mt-4 w-full font-semibold'>Estudos</h2>
            <nav className='w-full'>
                <ul>
                    {studies.map(item => (
                        <SidebarItens
                            href={`/r/study/${item.id}`}
                            selected={pathname === `/r/study/${item.id}`}
                            title={item.name}
                            notifications={0}
                            key={item.id}
                            icon={BookA}
                        />
                    ))}
                </ul>
            </nav>
        </div>
    );

    return (
        <div>
            <Navbar />
            <div className='md:hidden fixed my-4'>
                <Sheet>
                    <SheetTrigger asChild className='ml-2'>
                        <Button className='mt-1' variant='ghost'>
                            <MenuIcon />
                        </Button>
                    </SheetTrigger>
                    <SheetContent side='left' className='p-0 w-fit h-full'>
                        {SidebarSection}
                    </SheetContent>
                </Sheet>
            </div>

            <div className='md:block hidden h-full'>{SidebarSection}</div>
        </div>
    );
}

'use client';

import { Icons } from '@/components/global/icons';
import { Select, SelectContent, SelectGroup, SelectItem, SelectLabel, SelectTrigger, SelectValue } from '@/components/ui/select';
import { useParams, usePathname, useRouter } from 'next/navigation';
import { Separator } from '@/components/ui/separator';
import { MenuItens } from '@/config/constants';
import SidebarItens from '@/components/nav/sidebar-itens';
import { BookA, MenuIcon, TelescopeIcon } from 'lucide-react';
import { Sheet, SheetContent, SheetTrigger } from '@/components/ui/sheet';
import { Button, buttonVariants } from '@/components/ui/button';
import Navbar from '@/components/nav/navbar';
import { useEffect, useState } from 'react';
import { ListStudies } from '@/actions/study';
import { Study } from '@/types';
import Link from 'next/link';
import { cn } from '@/lib/utils';
import { useQuery } from '@tanstack/react-query';

export default function Sidebar() {
    const router = useRouter();
    const pathname = usePathname();
    const params = useParams();

    const [selectedId, setSelectedId] = useState(params.studyID || '');

    function OnChange(id: string) {
        router.push(`/r/study/${id}`);
    }

    useEffect(() => {
        setSelectedId(params.studyID || '');
    }, [pathname]);

    const { data, isPending, isFetched } = useQuery({
        queryKey: ['study-data'],
        queryFn: ListStudies
    });

    const studies = (data as unknown as Study[]) || [];

    const menu = MenuItens(selectedId as string);

    const SidebarSection = (
        <div className='bg-secondary flex-none relative p-4 h-full w-[250px] flex flex-col gap-4 overflow-hidden'>
            <div className='bg-secondary p-4 justify-center items-center mb-4 absolute top-0 left-0 right-0'>
                <Icons.logo className='w-44 h-16 fill-black dark:fill-white' />
            </div>

            <Select value={selectedId as string} onValueChange={OnChange}>
                <SelectTrigger className='mt-24 text-muted-foreground bg-transparent border-gray-300 dark:border-gray-900'>
                    <SelectValue placeholder='Escolha um estudo'></SelectValue>
                </SelectTrigger>
                <SelectContent className=' backdrop-blur-xl'>
                    <SelectGroup>
                        <SelectLabel>Estudos</SelectLabel>
                        <Separator className='mb-2' />
                        {studies.map(ws => (
                            <SelectItem key={ws.id} value={ws.id}>
                                {ws.title}
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

            <Separator className='w-full bg-secondary' />
            <h2 className='mt-4 w-full font-semibold'>Estudos</h2>
            <nav className='w-full'>
                {studies.length > 0 ? (
                    <ul>
                        {studies.map(item => (
                            <SidebarItens
                                href={`/r/study/${item.id}`}
                                selected={pathname === `/r/study/${item.id}`}
                                title={item.title}
                                notifications={0}
                                key={item.id}
                                icon={BookA}
                            />
                        ))}
                    </ul>
                ) : (
                    <span className='text-muted-foreground text-sm'>
                        <p>
                            Nenhum estudo foi adicionado ainda.{' '}
                            <Link href='/r/study' className={cn(buttonVariants({ variant: 'link', className: 'p-0 text-sm' }))}>
                                {' '}
                                Que tal começar agora?
                            </Link>
                        </p>
                    </span>
                )}
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

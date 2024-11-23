'use client';
import { cn } from '@/lib/utils';
import { HoverCard, HoverCardContent, HoverCardTrigger } from '../ui/hover-card';
import { GetUser } from '@/actions/user';
import { User } from '@/types';
import { useQuery } from '@tanstack/react-query';
import Link from 'next/link';
import { Button, buttonVariants } from '../ui/button';
import { Avatar, AvatarFallback, AvatarImage } from '../ui/avatar';
import { CalendarIcon } from 'lucide-react';
import { useRouter } from 'next/navigation';

const HoverButton = ({ id }: { id: string }) => {
    const { data, isPending, isFetched } = useQuery({
        queryKey: ['user-data'],
        queryFn: GetUser
    });

    const user = data as unknown as User;

    const router = useRouter();

    function onClick() {
        router.push(`/r/user/${id}`);
    }
    return (
        <HoverCard>
            <HoverCardTrigger>
                <Button variant='link' onClick={onClick} className='p-0'>
                    {id === user.id ? 'Você' : user.email}
                </Button>
            </HoverCardTrigger>
            <HoverCardContent className='w-70'>
                <div className='flex justify-between space-x-4'>
                    <Avatar>
                        <AvatarImage src='' />
                        <AvatarFallback>YOU</AvatarFallback>
                    </Avatar>
                    <div className='space-y-1'>
                        <h4 className='text-sm font-semibold'>{user.email}</h4>
                        <p className='text-sm'>Descricao que eu ainda nao fiz</p>
                        <div className='flex items-center pt-2'>
                            <CalendarIcon className='mr-2 h-4 w-4 opacity-70' />{' '}
                            <span className='text-xs text-muted-foreground'>
                                Se juntou em{' '}
                                {new Date(user.created_at).toLocaleDateString('pt-BR', {
                                    year: 'numeric',
                                    month: 'numeric',
                                    day: 'numeric'
                                })}
                            </span>
                        </div>
                    </div>
                </div>
            </HoverCardContent>
        </HoverCard>
    );
};

export default HoverButton;

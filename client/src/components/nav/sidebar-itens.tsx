import { MenuProps } from '@/config/constants';
import { cn } from '@/lib/utils';
import { LucideProps } from 'lucide-react';
import Link from 'next/link';
import { buttonVariants } from '../ui/button';

interface SidebarItensProps extends MenuProps {
    selected: boolean;
    notifications: number;
}

export default function SidebarItens({ href, icon: Icon, notifications, selected, title }: SidebarItensProps) {
    return (
        <div className='cursor-pointer my-2'>
            <Link
                className={cn(
                    buttonVariants({
                        variant: selected ? 'default' : 'ghost',
                        className: selected ? '' : 'hover:bg-gray-300 hover:dark:bg-gray-900'
                    })
                )}
                href={href}
            >
                <div className='flex items-center gap-2 transition-all p-2'>
                    <Icon />
                    <span className={cn('font-medium transition-all truncate w-32')}>{title}</span>
                </div>
            </Link>
        </div>
    );
}

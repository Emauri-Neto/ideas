import { listThreads } from '@/actions/list';
import { ArrowBigRightDashIcon } from 'lucide-react';
import Link from 'next/link';
import { buttonVariants } from '../ui/button';
import { cn } from '@/lib/utils';
import { Badge } from '@/components/ui/badge';

type Props = {
    id: string;
};

const ThreadArea = async ({ id }: Props) => {
    const thread = await listThreads(id);

    function capitalize(str: string) {
        return str.charAt(0).toUpperCase() + str.slice(1);
    }

    function NormalizeDate(date: Date): string {
        const dt = new Date(date);
        const day = String(dt.getDate()).padStart(2, '0');
        const month = String(dt.getMonth() + 1).padStart(2, '0');
        const year = dt.getFullYear();

        return `${day}/${month}/${year}`;
    }
    return (
        <div className='flex flex-col'>
            {thread.map(t => (
                <div className='flex flex-row items-center justify-between' key={t.id}>
                    <Link href='' className={cn(buttonVariants({ variant: 'link', className: 'p-1' }))}>
                        <ArrowBigRightDashIcon />
                        <Badge className='rounded-sm'>{capitalize(t.status)}</Badge>
                        {t.name}
                    </Link>
                    <p className='text-sm text-muted-foreground'>Criado em {NormalizeDate(t.created_at)}</p>
                </div>
            ))}
        </div>
    );
};

export default ThreadArea;

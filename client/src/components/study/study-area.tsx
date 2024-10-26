import { Button, buttonVariants } from '@/components/ui/button';
import { Ghost, User2Icon } from 'lucide-react';
import { Badge } from '@/components/ui/badge';
import Link from 'next/link';
import { cn } from '@/lib/utils';
import FormSugestion from "@/components/thread/form/formSugestion"

const study = [
    {
        id: 'cbf5e25d-1e4d-4cc5-83bc-9266a21aee02',
        name: 'Meu estudo',
        objective: 'Descobrir o mundo',
        methodology: 'ChatGPT',
        created_at: '2024-10-17T21:50:04.882913Z',
        responsible_id: 'dc65b81f-4c35-4659-bb60-fe5b73a0d98e'
    },
    {
        id: 'cbf5e25d-1e4d-4cc5-83bc-9266a21aee02',
        name: 'Seu estudo',
        objective: 'Descobrir o mundo',
        methodology: 'ChatGPT',
        created_at: '2024-10-17T21:50:04.882913Z',
        responsible_id: 'dc65b81f-4c35-4659-bb60-fe5b73a0d98e'
    }
];

const StudyArea = () => {
    function truncateString(str: string, maxLength: number) {
        if (str.length > maxLength) {
            return str.slice(0, maxLength) + '...';
        }
        return str;
    }

    if (study.length < 1) {
        return (
            <div className='flex items-center rounded-md flex-col justify-center h-full gap-y-2'>
                <span>
                    <Ghost className='h-7 w-7' />
                </span>
                <h2 className='text-xl'>Sem estudos disponíveis</h2>
                <p className='text-muted-foreground'>Seja o primeiro a começar um estudo</p>
                <Button>Começar</Button>
            </div>
        );
    }

    return study.map(s => (
        <div className='mb-4'>
            <div className='bg-secondary p-4 rounded-md flex gap-12'>
                <div className='flex flex-col w-full'>
                    <div className='flex flex-row items-center w-full bg-gray-300 dark:bg-gray-900 px-3 py-3 rounded-sm'>
                        <div className='flex flex-row gap-2 items-center'>
                            <Link href={`/r/${s.id}`} className={cn(buttonVariants({ variant: 'link', className: 'p-0 font-semibold text-lg' }))}>
                                {truncateString(s.name, 20)}
                            </Link>
                            <span className='h-6 w-px bg-gray-300' aria-hidden='true' />
                            <p className='text-muted-foreground text-sm'>{truncateString(s.objective, 30)}</p>
                            <span className='text-muted-foreground'>-</span>
                            <Badge className='rounded-sm'>{s.methodology}</Badge>
                        </div>
                        <div className='ml-auto flex flex-row items-center gap-4'>
                            <div className='flex flex-row gap-1 items-center'>
                                <p>
                                    <span className='text-muted-foreground'>1</span>
                                    <span className='px-1'>/</span>
                                    10
                                </p>
                                <User2Icon className='w-4 h-4' />
                            </div>
                            <FormSugestion title='Sugerir Nova Linha de Discussão'/>
                        </div>
                    </div>
                    <div>ss</div>
                </div>
            </div>
        </div>
    ));
};

export default StudyArea;

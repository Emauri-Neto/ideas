import { Study } from '@/types';
import { Card, CardContent, CardHeader, CardTitle } from '@/components/ui/card';
import Link from 'next/link';
import { cn } from '@/lib/utils';
import { buttonVariants } from '@/components/ui/button';
import CopyToClipBoard from '../clipboard';
import { Separator } from '@/components/ui/separator';
import { Badge } from '@/components/ui/badge';
import { AlarmClock } from 'lucide-react';
import HoverButton from '../hover-btn';

interface Props {
    st: Study;
}

const StudyCard = ({ st }: Props) => {
    console.log(st);

    return (
        <Card className='w-[450px] shadow-md'>
            <CardHeader className='flex flex-row items-center justify-between'>
                <CardTitle className='flex items-center'>
                    <Link href={`/r/study/${st.id}`} className={cn(buttonVariants({ variant: 'link', className: 'p-0 text-xl' }))}>
                        {st.title}
                    </Link>
                    <Badge className='mx-2'>{st.methodology}</Badge>
                </CardTitle>
                <CopyToClipBoard studyID={st.id} />
            </CardHeader>
            <Separator />
            <CardContent className='my-4'>
                <div className='flex flex-col'>
                    {st.threads.map(th => (
                        <div className='flex flex-col my-1 text-sm gap-1 items-start' key={th.id}>
                            <div className='flex gap-2 items-center'>
                                <Link href={`r/study/${st.id}/${th.id}`} className={cn(buttonVariants({ variant: 'link', className: 'p-0' }))}>
                                    {th.name}
                                </Link>
                                <span className='text-muted-foreground'>-</span>
                                <HoverButton id={th.responsible_user} />
                            </div>
                            <div className='text-sm flex items-center gap-1 text-muted-foreground'>
                                <span>Prazo:</span>
                                {new Date(th.deadline).toLocaleDateString('pt-BR', {
                                    year: 'numeric',
                                    month: 'numeric',
                                    day: 'numeric'
                                })}
                            </div>
                        </div>
                    ))}
                </div>
            </CardContent>
        </Card>
    );
};

export default StudyCard;

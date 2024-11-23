'use client';

import HoverButton from '@/components/study/hover-btn';
import { Badge } from '@/components/ui/badge';
import { Separator } from '@/components/ui/separator';
import { Calendar1Icon } from 'lucide-react';
import { useRouter } from 'next/navigation';

const study = {
    id: 'dfa9f9ad-971e-4089-afe7-55cd44955a12',
    title: 'Meu Estudo',
    objective: 'Descobrir o mundo',
    methodology: 'ChatGPT',
    max_participants: null,
    num_participants: null,
    participation_type: null,
    _private: false,
    user_id: '2a2f24f1-96ef-40e2-880b-6165ab83e2fb',
    threads: [
        {
            id: '4b304b46-2217-4b33-8164-dc69f0f01fc8',
            name: 'Thread legal',
            deadline: '2024-11-21T12:47:12.797888Z',
            responsible_user: '2a2f24f1-96ef-40e2-880b-6165ab83e2fb',
            study_id: 'dfa9f9ad-971e-4089-afe7-55cd44955a12'
        },
        {
            id: 'ec043cd9-4c1c-4049-864e-bb6011753c52',
            name: 'Thread show',
            deadline: '2024-11-21T12:47:12.797717Z',
            responsible_user: '2a2f24f1-96ef-40e2-880b-6165ab83e2fb',
            study_id: 'dfa9f9ad-971e-4089-afe7-55cd44955a12'
        }
    ],
    created_at: '2024-11-21T12:47:12.797717Z',
    updated_at: '2024-11-21T12:47:12.797717Z'
};

const Study = () => {
    const router = useRouter();
    function onClick(id: string) {
        router.push(`/r/study/${study.id}/${id}`);
    }
    return (
        <div className='py-2'>
            <div className='text-4xl font-bold'>{study.title}</div>
            <span className='text-muted-foreground'>Criado em {new Date(study.created_at).toLocaleDateString()} por </span>
            <HoverButton id={study.user_id} />
            <Separator className='p-1 mt-4 mb-4' />
            <div className='flex flex-row flex-wrap gap-10 mb-10'>
                {study.objective && (
                    <div className='flex flex-col'>
                        <h3 className='font-semibold text-xl'>Objetivo</h3>
                        <p className='text-muted-foreground'>{study.objective}</p>
                    </div>
                )}

                <div>
                    <h3 className='font-semibold text-xl'>Metodologia</h3>
                    <p className='text-muted-foreground'>{study.methodology}</p>
                </div>
                <div>
                    <h3 className='font-semibold text-xl'>Participantes</h3>
                    <p className='text-muted-foreground'>
                        {study.num_participants || 1} / {study.max_participants || 'Unlimited'}
                    </p>
                </div>
            </div>
            <div>
                <h3 className='font-semibold text-xl'>Threads</h3>
                {study.threads.length > 0 ? (
                    <div className='mt-4 grid grid-cols-4 gap-4'>
                        {study.threads.map(thread => (
                            <div key={thread.id}>
                                <div className='bg-secondary p-4 rounded-md cursor-pointer' onClick={() => onClick(thread.id)}>
                                    <div className='flex gap-2 items-center'>
                                        <h3 className='font-semibold text-base'>{thread.name}</h3>
                                        <span>-</span>
                                        <HoverButton id={thread.responsible_user} />
                                    </div>
                                    <div className='flex flex-row items-center text-muted-foreground text-sm gap-2'>
                                        <Calendar1Icon size={16} />
                                        Finaliza em:{' '}
                                        {new Date(thread.deadline).toLocaleDateString('pt-BR', {
                                            weekday: 'long',
                                            year: 'numeric',
                                            month: 'long',
                                            day: 'numeric'
                                        })}
                                    </div>
                                </div>
                            </div>
                        ))}
                    </div>
                ) : (
                    <p className='text-muted-foreground'>Nenhuma linha de discussão criada.</p>
                )}
            </div>
        </div>
    );
};

export default Study;

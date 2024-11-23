'use client';

import { ListStudies } from '@/actions/study';
import StudyCard from '@/components/study/card/card';
import CreateStudyForm from '@/components/study/form/create';
import { Button } from '@/components/ui/button';
import Modal from '@/components/ui/modal';
import { Tabs, TabsContent, TabsList, TabsTrigger } from '@/components/ui/tabs';
import { Study as St } from '@/types';
import { useQuery } from '@tanstack/react-query';
import { BookPlusIcon } from 'lucide-react';
import React from 'react';

type Props = {};

const Study = ({}: Props) => {
    const { data, isFetched, isLoading } = useQuery({
        queryKey: ['study-data'],
        queryFn: ListStudies
    });

    const study = data as unknown as St[] | undefined;

    return (
        <div>
            <Tabs defaultValue='study' className='mt-6'>
                <div className='flex w-full items-center gap-6'>
                    <TabsList>
                        <TabsTrigger value='study'>Estudos</TabsTrigger>
                        <TabsTrigger value='thread'>Linhas de Discussão</TabsTrigger>
                    </TabsList>

                    <Modal
                        title='Crie seu Estudo'
                        description=''
                        trigger={
                            <Button variant='ghost'>
                                <BookPlusIcon />
                                Criar estudo
                            </Button>
                        }
                    >
                        <CreateStudyForm />
                    </Modal>
                </div>

                <TabsContent value='study'>
                    <div className='grid grid-cols-5 my-5'>
                        {study?.map(st => (
                            <StudyCard key={st.id} st={st} />
                        ))}
                    </div>
                </TabsContent>
            </Tabs>
        </div>
    );
};

export default Study;

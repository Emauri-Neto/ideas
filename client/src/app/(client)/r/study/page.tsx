import CreateStudyForm from '@/components/study/form/create';
import { Button } from '@/components/ui/button';
import Modal from '@/components/ui/modal';
import { Tabs, TabsList, TabsTrigger } from '@/components/ui/tabs';
import { BookPlusIcon } from 'lucide-react';
import React from 'react';

type Props = {};

const Study = (props: Props) => {
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
            </Tabs>
        </div>
    );
};

export default Study;

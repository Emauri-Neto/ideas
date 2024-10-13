import React from 'react';
import { Button } from '../ui/button';
import { VscEmptyWindow } from 'react-icons/vsc';
import { Badge } from '../ui/badge';
import { BsEyeFill } from 'react-icons/bs';
import { User } from 'lucide-react';

const study = [
    {
        name: 'Estudo teste',
        objective: 'Estudar algo do mundo',
        methodology: 'Metodologia agil',
        owner: 'Me',
        created_at: 'Dois dias atrás',
        participantes: 2
    },
    {
        name: 'Estudo sobre Coisas',
        objective: 'Estudar Coisas',
        methodology: 'Metodologia Legal',
        owner: 'Cledosvaldo',
        created_at: 'Dez dias atrás',
        participantes: 17
    }
];

interface StudiesAreaProps {}

const StudiesArea: React.FC<StudiesAreaProps> = () => {
    if (study.length < 1) {
        return (
            <div className='flex items-center rounded-md flex-col justify-center h-full bg-secondary gap-y-2'>
                <span>
                    <VscEmptyWindow className='h-7 w-7' />
                </span>
                <h2 className='text-xl'>Sem estudos disponíveis</h2>
                <p className='text-muted-foreground'>Seja o primeiro a começar um estudo</p>
                <Button>Começar</Button>
            </div>
        );
    }

    return study.map(study => (
        <div className='mb-4 cursor-pointer'>
            <div className='bg-secondary p-4 rounded-md flex gap-12'>
                <div>Image</div>
                <div className='flex flex-col gap-y-1 w-full mx-12'>
                    <h3 className='font-semibold mx-1'>{study.name}</h3>
                    <div>
                        <Badge className='bg-slate-200 dark:bg-gray-900 text-primary '>{study.methodology}</Badge>
                    </div>
                    <div className='flex flex-row mt-6 justify-between'>
                        <div className='flex flex-col'>
                            <p>{study.owner}</p>
                            <span className='text-muted-foreground text-sm'>{study.created_at}</span>
                        </div>
                        <span className='flex gap-2 items-center justify-center'>
                            <User className='h-4 w-4' />
                            <p>{study.participantes || 1}</p>
                        </span>
                    </div>
                </div>
            </div>
        </div>
    ));
};

export default StudiesArea;

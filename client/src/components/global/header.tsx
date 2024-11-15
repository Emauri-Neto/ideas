'use client';
import { GetStudy, ListStudies } from '@/actions/study';
import { Study } from '@/types';
import { useQuery } from '@tanstack/react-query';
import { usePathname } from 'next/navigation';
import { useEffect, useState } from 'react';

export default function Header() {
    const pathname = usePathname();
    const [basePath, subPath] = pathname.split('/r/')[1]?.split('/') || [];

    const { data, isFetched } = useQuery({
        queryKey: ['get-study-data', subPath],
        queryFn: () => (subPath ? GetStudy(subPath) : null),
        enabled: !!subPath
    });

    const study = data as unknown as Study | undefined;

    const translations: Record<string, string> = {
        study: 'Estudos',
        settings: 'Configurações'
    };

    const text = basePath === 'study' && subPath ? (study?._private ? 'Privado' : 'Público') : 'Privado';

    return (
        <div className='flex flex-col gap-2'>
            <span className='text-xs text-muted-foreground'>{text.toLocaleUpperCase()}</span>
            <h1 className='text-4xl font-medium'>
                {basePath
                    ? translations[basePath.toLowerCase()] || basePath.charAt(0).toUpperCase() + basePath.slice(1).toLowerCase()
                    : 'Meus Estudos'}
            </h1>
        </div>
    );
}

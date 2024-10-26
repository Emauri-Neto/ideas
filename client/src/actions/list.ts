'use server';

import * as z from 'zod';
import { studySchema, threadSchema } from '@/config/schemas';
import api from '@/lib/api';

export const listStudies = async (): Promise<z.infer<typeof studySchema>[]> => {
    const { data, status } = await api.get<z.infer<typeof studySchema>[]>('/r/study');

    if (status !== 200) {
        return [];
    }

    return data;
};

export const listThreads = async (id: string): Promise<z.infer<typeof threadSchema>[]> => {
    const { data, status } = await api.get<z.infer<typeof threadSchema>[]>(`/r/study/${id}/thread`);

    if (status !== 200) {
        return [];
    }

    return data;
};

'use server';

import api from '@/lib/api';
import { registerStudySchema } from '@/lib/schemas';
import { cookies } from 'next/headers';
import * as z from 'zod';

export const ListStudies = async () => api.get('/api/study', { headers: { Authorization: `Bearer ${cookies().get('refresh')?.value}` } });

export const GetStudy = async (id: string) =>
    api.get(`/api/study/${id}`, { headers: { Authorization: `Bearer ${cookies().get('refresh')?.value}` } });

export const CreateStudy = async (data: z.infer<typeof registerStudySchema>) =>
    api.post('/api/study', data, { headers: { Authorization: `Bearer ${cookies().get('refresh')?.value}` } });

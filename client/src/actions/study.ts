'use server';

import api from '@/lib/api';
import { cookies } from 'next/headers';

export const ListStudies = async () => api.get('/api/study', { headers: { Authorization: `Bearer ${cookies().get('refresh')?.value}` } });

export const GetStudy = async (id: string) =>
    api.get(`/api/study/${id}`, { headers: { Authorization: `Bearer ${cookies().get('refresh')?.value}` } });

'use server';

import api from '@/lib/api';
import { cookies } from 'next/headers';

export const GetUser = async () => api.get('/api/user', { headers: { Authorization: `Bearer ${cookies().get('refresh')?.value}` } });

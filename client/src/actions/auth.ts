'use server';

import api from '@/lib/api';
import { loginSchema, registerSchema } from '@/lib/schemas';
import { cookies } from 'next/headers';
import { z } from 'zod';

export const Login = async (values: z.infer<typeof loginSchema>) => {
    const data = await api.post<{ Access: string; Refresh: string }>('/auth/login', values);

    cookies().set({
        name: 'access',
        //@ts-expect-error
        value: data.Access,
        path: '/',
        httpOnly: true,
        secure: false,
        maxAge: 15 * 60
    });

    cookies().set({
        name: 'refresh',
        //@ts-expect-error
        value: data.Refresh,
        path: '/',
        httpOnly: true,
        secure: false,
        maxAge: 7 * 24 * 60 * 60
    });

    return data;
};

export const Register = async (values: z.infer<typeof registerSchema>) => api.post('/auth/register', values);

export const Logout = async () => {
    cookies().delete('refresh');
    cookies().delete('access');
    return api.get('/auth/logout');
};

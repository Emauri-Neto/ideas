'use server';

import * as z from 'zod';
import { loginSchema } from '@/lib/schemas';
import { cookies } from 'next/headers';
import { redirect } from 'next/navigation';
import Auth from './auth';

export const login = async (values: z.infer<typeof loginSchema>) => {
    const fields = loginSchema.safeParse(values);

    if (!fields.success) {
        return { error: 'Campos inválidos.' };
    }

    const res = await fetch('http://localhost:4367/auth/login', {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json'
        },
        body: JSON.stringify(fields.data)
    });

    if (!res.ok) {
        const errorData = await res.json();
        return { error: errorData.message || 'Erro ao Entrar.' };
    }

    const data = await res.json();

    cookies().set({
        name: 'IAUTH',
        value: data.token!,
        httpOnly: true,
        sameSite: 'strict',
        secure: false,
        maxAge: 3600
    });

    return redirect('/');
};

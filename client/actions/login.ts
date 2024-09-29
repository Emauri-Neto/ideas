'use server';

import * as z from 'zod';
import { loginSchema } from '@/config/schemas';
import { cookies } from 'next/headers';
import { redirect } from 'next/navigation';

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
        return { error: errorData.message || 'Erro ao registrar.' };
    }

    const data = await res.json();

    cookies().set({
        name: 'IAUTH',
        value: data.token!,
        httpOnly: true,
        sameSite: 'strict',
        secure: process.env.NODE_ENV !== 'development',
        maxAge: 3600
    });

    return redirect('/');
};

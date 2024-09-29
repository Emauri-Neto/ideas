'use server';

import * as z from 'zod';
import { registerSchema } from '@/config/schemas';

export const register = async (values: z.infer<typeof registerSchema>) => {
    const fields = registerSchema.safeParse(values);

    if (!fields.success) {
        return { error: 'Campos inválidos.' };
    }

    const { email, confirmPassword, name, password } = fields.data;

    const res = await fetch('http://localhost:4367/auth/register', {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json'
        },
        body: JSON.stringify({
            email,
            name,
            password,
            confirmPassword
        })
    });

    if (!res.ok) {
        const errorData = await res.json();
        return { error: errorData.message || 'Erro ao registrar.' };
    }
};

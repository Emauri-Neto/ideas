'use server';

import * as z from 'zod';
import { loginSchema } from '@/config/schemas';

export const login = async (values: z.infer<typeof loginSchema>) => {
    const fields = loginSchema.safeParse(values);

    if (!fields.success) {
        return { error: 'Campos inválidos.' };
    }
};

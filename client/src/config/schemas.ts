import * as z from 'zod';

export const emailSchema = z.string().email({ message: 'Email inválido.' }).min(1).max(255);
export const passwordSchema = z.string().min(6, { message: 'Senha é obrigatório.' });

export const loginSchema = z.object({
    email: emailSchema,
    password: passwordSchema
});

export const registerSchema = loginSchema
    .extend({
        confirmPassword: z
            .string()
            .min(6)
            .regex(/^(?=.*[a-z])(?=.*[A-Z])(?=.*\d)(?=.*[\W_]).+$/, {
                message: 'A senha deve ter no mínimo 6 caracteres, incluindo 1 letra maiúscula, 1 letra minúscula, 1 número e 1 caractere especial.'
            }),
        name: z.string().min(1, { message: 'Nome é obrigatório.' })
    })
    .refine((data) => data.password === data.confirmPassword, {
        message: 'As senhas não coincidem.',
        path: ['confirmPassword']
    });

export const studySchema = z.object({
    id: z.string().uuid(),
    name: z.string(),
    objective: z.string(),
    methodology: z.string(),
    num_participants: z.number(),
    max_participants: z.number().optional(),
    participation_type: z.string().optional(),
    created_at: z.date(),
    updated_at: z.date(),
    responsible_id: z.string()
});

export const threadSchema = z.object({
    id: z.string().uuid(),
    name: z.string(),
    max_participants: z.number().optional(),
    deadline: z.date().optional(),
    status: z.string(),
    created_at: z.date(),
    responsible_id: z.string(),
    study_id: z.string()
});

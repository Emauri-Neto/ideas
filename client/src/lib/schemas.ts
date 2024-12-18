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
            })
    })
    .refine((data) => data.password === data.confirmPassword, {
        message: 'As senhas não coincidem.',
        path: ['confirmPassword']
    });

export const registerStudySchema = z.object({
    title: z.string().min(3),
    objective: z.string().min(6).max(120).optional(),
    methodology: z.enum(['prospective', 'intuitive', 'probabilistic', 'custom']),
    _private: z.boolean().default(true)
});

'use client';

import * as z from 'zod';
import { registerSchema } from '@/config/schemas';

import { useForm } from 'react-hook-form';
import { zodResolver } from '@hookform/resolvers/zod';
import { Form, FormControl, FormField, FormItem, FormLabel, FormMessage } from '@/components/ui/form';
import { Input } from '@/components/ui/input';
import { Button } from '@/components/ui/button';
import { AuthWrapper } from '@/components/auth/auth-wrapper';
import FormError from '../form-error';
import { useState, useTransition } from 'react';
import { EnterIcon } from '@radix-ui/react-icons';
import { register } from '@/actions/register';

const RegisterForm = () => {
    const [isPending, startTransition] = useTransition();

    const [err, setErr] = useState<string | undefined>('');

    const form = useForm<z.infer<typeof registerSchema>>({
        resolver: zodResolver(registerSchema),
        defaultValues: {
            email: '',
            password: '',
            confirmPassword: '',
            name: ''
        }
    });

    const onSubmit = (values: z.infer<typeof registerSchema>) => {
        startTransition(() => {
            register(values).then(data => setErr(data?.error));
        });
    };

    return (
        <AuthWrapper headerLabel='Crie sua Conta!' buttonHref='/sign-in' buttonLabel='Já tem uma conta? Faça login!' showSocials>
            <Form {...form}>
                <form onSubmit={form.handleSubmit(onSubmit)} className='space-y-6'>
                    <div className='space-y-4'>
                        <FormField
                            control={form.control}
                            name='name'
                            disabled={isPending}
                            render={({ field }) => (
                                <FormItem>
                                    <FormLabel>Nome</FormLabel>
                                    <FormControl>
                                        <Input {...field} placeholder='João Exemplo' />
                                    </FormControl>
                                    <FormMessage />
                                </FormItem>
                            )}
                        />
                        <FormField
                            control={form.control}
                            name='email'
                            disabled={isPending}
                            render={({ field }) => (
                                <FormItem>
                                    <FormLabel>Email</FormLabel>
                                    <FormControl>
                                        <Input {...field} placeholder='pessoa@exemplo.com' />
                                    </FormControl>
                                    <FormMessage />
                                </FormItem>
                            )}
                        />
                        <FormField
                            control={form.control}
                            name='password'
                            disabled={isPending}
                            render={({ field }) => (
                                <FormItem>
                                    <FormLabel>Senha</FormLabel>
                                    <FormControl>
                                        <Input {...field} placeholder='******' type='password' />
                                    </FormControl>
                                    <FormMessage />
                                </FormItem>
                            )}
                        />
                        <FormField
                            control={form.control}
                            name='confirmPassword'
                            disabled={isPending}
                            render={({ field }) => (
                                <FormItem>
                                    <FormLabel>Confirmar Senha</FormLabel>
                                    <FormControl>
                                        <Input {...field} placeholder='******' type='password' />
                                    </FormControl>
                                    <FormMessage />
                                </FormItem>
                            )}
                        />
                    </div>
                    <FormError message={err} />
                    <Button type='submit' className='w-full gap-1 font-semibold' disabled={!form.formState.isValid || isPending}>
                        Criar conta
                        <EnterIcon className='w-4 h-4' />
                    </Button>
                </form>
            </Form>
        </AuthWrapper>
    );
};

export default RegisterForm;

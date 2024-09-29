'use client';

import * as z from 'zod';
import { loginSchema } from '@/config/schemas';

import { useForm } from 'react-hook-form';
import { zodResolver } from '@hookform/resolvers/zod';
import { Form, FormControl, FormField, FormItem, FormLabel, FormMessage } from '@/components/ui/form';
import { Input } from '@/components/ui/input';
import { Button } from '@/components/ui/button';
import { AuthWrapper } from '@/components/auth/auth-wrapper';
import FormError from '../form-error';
import { login } from '@/actions/login';
import { useState, useTransition } from 'react';

import { EnterIcon } from '@radix-ui/react-icons';

const LoginForm = () => {
    const [isPending, startTransition] = useTransition();

    const [err, setErr] = useState<string | undefined>('');

    const form = useForm<z.infer<typeof loginSchema>>({
        resolver: zodResolver(loginSchema),
        defaultValues: {
            email: '',
            password: ''
        }
    });

    const onSubmit = (values: z.infer<typeof loginSchema>) => {
        startTransition(() => {
            login(values).then(data => setErr(data?.error));
        });
    };

    return (
        <AuthWrapper headerLabel='Bem vindo de volta!' buttonHref='/sign-up' buttonLabel='Novo por aqui? Registre-se!' showSocials>
            <Form {...form}>
                <form onSubmit={form.handleSubmit(onSubmit)} className='space-y-6'>
                    <div className='space-y-4'>
                        <FormField
                            control={form.control}
                            name='email'
                            disabled={isPending}
                            render={({ field }) => (
                                <FormItem>
                                    <FormLabel>Email</FormLabel>
                                    <FormControl>
                                        <Input {...field} placeholder='pessoa@exemplo.com' type='text' />
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
                    </div>
                    <FormError message={err} />
                    <Button type='submit' className='w-full gap-1 font-semibold' disabled={!form.formState.isValid || isPending}>
                        Entrar
                        <EnterIcon className='w-4 h-4' />
                    </Button>
                </form>
            </Form>
        </AuthWrapper>
    );
};

export default LoginForm;

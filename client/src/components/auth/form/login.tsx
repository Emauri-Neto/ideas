'use client';

import * as z from 'zod';
import { loginSchema } from '@/lib/schemas';

import { useForm } from 'react-hook-form';
import { zodResolver } from '@hookform/resolvers/zod';
import { Form, FormControl, FormField, FormItem, FormLabel, FormMessage } from '@/components/ui/form';
import { Input } from '@/components/ui/input';
import { Button } from '@/components/ui/button';
import { AuthWrapper } from '@/components/auth/auth-wrapper';
import { Login } from '@/actions/auth';
import { Loader2Icon, LogIn } from 'lucide-react';
import FormError from '@/components/auth/form/error';
import { useMutation } from '@tanstack/react-query';
import { useRouter } from 'next/navigation';

const LoginForm = () => {
    const router = useRouter();
    const {
        mutate: SignIn,
        isPending,
        isError,
        error
    } = useMutation({
        mutationFn: Login,
        onSuccess: () => {
            router.replace('/');
        }
    });

    const form = useForm<z.infer<typeof loginSchema>>({
        resolver: zodResolver(loginSchema),
        defaultValues: {
            email: '',
            password: ''
        }
    });

    const onSubmit = (values: z.infer<typeof loginSchema>) => {
        SignIn(values);
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
                    <FormError message={error?.message} isError={isError} />
                    <Button type='submit' className='w-full gap-1 font-semibold' disabled={!form.formState.isValid || isPending}>
                        {isPending ? (
                            <Loader2Icon className='animate-spin w-4 h-4' />
                        ) : (
                            <>
                                <span>Entrar</span>
                                <LogIn className='w-4 h-4' />
                            </>
                        )}
                    </Button>
                </form>
            </Form>
        </AuthWrapper>
    );
};

export default LoginForm;

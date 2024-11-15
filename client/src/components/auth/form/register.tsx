'use client';

import * as z from 'zod';
import { registerSchema } from '@/lib/schemas';

import { useForm } from 'react-hook-form';
import { zodResolver } from '@hookform/resolvers/zod';
import { Form, FormControl, FormField, FormItem, FormLabel, FormMessage } from '@/components/ui/form';
import { Input } from '@/components/ui/input';
import { Button } from '@/components/ui/button';
import { AuthWrapper } from '@/components/auth/auth-wrapper';
import FormError from '@/components/auth/form/error';
import { Loader2Icon, LogIn } from 'lucide-react';
import { Register } from '@/actions/auth';
import { useMutation } from '@tanstack/react-query';
import { useRouter } from 'next/navigation';

const RegisterForm = () => {
    const router = useRouter();

    const {
        mutate: SignUp,
        isPending,
        isError
    } = useMutation({
        mutationFn: Register,
        onSuccess: () => {
            router.replace('/');
        }
    });

    const form = useForm<z.infer<typeof registerSchema>>({
        resolver: zodResolver(registerSchema),
        defaultValues: {
            email: '',
            password: '',
            confirmPassword: ''
        }
    });

    const onSubmit = (values: z.infer<typeof registerSchema>) => {
        SignUp(values);
    };

    return (
        <AuthWrapper headerLabel='Crie sua Conta!' buttonHref='/sign-in' buttonLabel='Já tem uma conta? Faça login!' showSocials>
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
                    <FormError message={'Email ou senha inválidos'} isError={isError} />
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

export default RegisterForm;

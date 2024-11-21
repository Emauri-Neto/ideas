'use client';

import { Form, FormControl, FormField, FormItem, FormLabel, FormMessage } from '@/components/ui/form';
import * as z from 'zod';
import { zodResolver } from '@hookform/resolvers/zod';
import { useForm } from 'react-hook-form';
import { registerStudySchema } from '@/lib/schemas';
import { useMutation } from '@tanstack/react-query';
import { Input } from '@/components/ui/input';
import { Textarea } from '@/components/ui/textarea';
import { Select, SelectContent, SelectGroup, SelectItem, SelectLabel, SelectTrigger, SelectValue } from '@/components/ui/select';
import { CreateStudy } from '@/actions/study';
import FormError from '@/components/auth/form/error';
import { Button } from '@/components/ui/button';
import { BookPlusIcon, Loader2Icon } from 'lucide-react';
import { useRouter } from 'next/navigation';
import { useToast } from '@/hooks/use-toast';

const CreateStudyForm = () => {
    const router = useRouter();
    const { toast } = useToast();

    const form = useForm<z.infer<typeof registerStudySchema>>({
        resolver: zodResolver(registerStudySchema)
    });

    const {
        mutate: RegisterStudy,
        isPending,
        isError,
        error
    } = useMutation({
        mutationFn: CreateStudy,
        onSuccess(data, variables, context) {
            toast({
                title: 'Estudo criado com sucesso!'
            });
            router.push(`/`);
        }
    });

    const onSubmit = (values: z.infer<typeof registerStudySchema>) => {
        RegisterStudy({ ...values });
    };

    return (
        <Form {...form}>
            <form onSubmit={form.handleSubmit(onSubmit)} className='space-y-3'>
                <div className='space-y-3'>
                    <FormField
                        control={form.control}
                        name='title'
                        disabled={isPending}
                        render={({ field }) => (
                            <FormItem>
                                <FormLabel>Titulo</FormLabel>
                                <FormControl>
                                    <Input {...field} placeholder='Meu estudo' />
                                </FormControl>
                                <FormMessage />
                            </FormItem>
                        )}
                    />

                    <FormField
                        control={form.control}
                        name='objective'
                        disabled={isPending}
                        render={({ field }) => (
                            <FormItem>
                                <FormLabel>Objetivo</FormLabel>
                                <FormControl>
                                    <Textarea {...field} placeholder='Descreva o objetivo do estudo numa descrição objetiva e curta.' />
                                </FormControl>
                                <FormMessage />
                            </FormItem>
                        )}
                    />

                    <FormField
                        control={form.control}
                        name='methodology'
                        disabled={isPending}
                        render={({ field }) => (
                            <FormItem>
                                <FormLabel>Metodologia</FormLabel>
                                <FormControl>
                                    <Select onValueChange={field.onChange} value={field.value}>
                                        <SelectTrigger>
                                            <SelectValue placeholder='Metodologia utilizada no estudo' />
                                        </SelectTrigger>

                                        <SelectContent>
                                            <SelectGroup>
                                                <SelectLabel>Metodologia disponíveis:</SelectLabel>
                                                <SelectItem value='prospective'>Prospectiva</SelectItem>
                                                <SelectItem value='intuitive'>Intuitivo</SelectItem>
                                                <SelectItem value='probabilistic'>Probabilístico</SelectItem>
                                                <SelectItem value='custom'>Customizado</SelectItem>
                                            </SelectGroup>
                                        </SelectContent>
                                    </Select>
                                </FormControl>
                                <FormMessage />
                            </FormItem>
                        )}
                    />
                </div>
                <FormError message={error?.message} isError={isError} />
                <Button type='submit' className='w-full gap-1 font-semibold' disabled={isPending}>
                    {isPending ? (
                        <Loader2Icon className='animate-spin w-4 h-4' />
                    ) : (
                        <>
                            <span>Criar</span>
                            <BookPlusIcon />
                        </>
                    )}
                </Button>
            </form>
        </Form>
    );
};

export default CreateStudyForm;

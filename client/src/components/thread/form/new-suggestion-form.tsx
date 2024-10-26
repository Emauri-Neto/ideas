'use client';
import { Button } from '@/components/ui/button';
import { Dialog, DialogContent, DialogHeader, DialogTitle, DialogTrigger } from '@/components/ui/dialog';
import { Input } from '@/components/ui/input';
import { Form, FormField, FormControl, FormItem, FormLabel, FormMessage } from '@/components/ui/form';
import { useForm } from 'react-hook-form';
import { Textarea } from '@/components/ui/textarea';
import { threadSugestionSchema } from '@/config/schemas';
import { z } from 'zod';
import { zodResolver } from '@hookform/resolvers/zod';

interface Props {
    variant?: any;
    className?: any;
    title?: string;
}

export default function AddThreadSuggestion(props: Props) {
    const form = useForm<z.infer<typeof threadSugestionSchema>>({
        resolver: zodResolver(threadSugestionSchema),
        defaultValues: {
            study: '',
            title: '',
            description: ''
        }
    });
    const onSubmit = (data: any) => {
        console.log(data);
    };
    return (
        <Dialog>
            <DialogTrigger asChild>
                <Button variant='default' className='p-3'>
                    {props.title}
                </Button>
            </DialogTrigger>
            <DialogContent className='sm:max-w-md flex flex-col'>
                <DialogHeader className='my-3 flex items-center'>
                    <DialogTitle>{props.title}</DialogTitle>
                </DialogHeader>
                <Form {...form}>
                    <form onSubmit={form.handleSubmit(onSubmit)} className='space-y-6'>
                        <div className='space-y-2'>
                            <FormField
                                control={form.control}
                                name='study'
                                render={({ field }) => (
                                    <FormItem>
                                        <FormLabel>Estudo</FormLabel>
                                        <FormControl>
                                            <Input {...field} placeholder='' type='text' />
                                        </FormControl>
                                        <FormMessage />
                                    </FormItem>
                                )}
                            />
                            <FormField
                                control={form.control}
                                name='title'
                                render={({ field }) => (
                                    <FormItem>
                                        <FormLabel>Titulo</FormLabel>
                                        <FormControl>
                                            <Input {...field} placeholder='' type='text' />
                                        </FormControl>
                                        <FormMessage />
                                    </FormItem>
                                )}
                            />
                            <FormField
                                control={form.control}
                                name='description'
                                render={({ field }) => (
                                    <FormItem>
                                        <FormLabel>Descrição</FormLabel>
                                        <FormControl>
                                            <Textarea {...field} className='resize-none' />
                                        </FormControl>
                                        <FormMessage />
                                    </FormItem>
                                )}
                            />
                        </div>
                        <div className='flex justify-center'>
                            <Button type='submit' className='w-full gap-1 font-semibold' disabled={!form.formState.isValid}>
                                Sugerir
                            </Button>
                        </div>
                    </form>
                </Form>
            </DialogContent>
        </Dialog>
    );
}

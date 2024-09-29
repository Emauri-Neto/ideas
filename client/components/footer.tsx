import MaxWidthWrapper from '@/components/wrapper';
import Link from 'next/link';
import { cn } from '@/lib/utils';
import { buttonVariants } from '@/components/ui/button';

const Footer = () => {
    return (
        <footer className='bg-background fixed bottom-0 inset-x-0 flex-grow-0 mt-auto'>
            <MaxWidthWrapper>
                <div className='border-t border-muted-foreground'></div>

                <div className='py-5 md:flex md:items-center md:justify-between'>
                    <div className='text-center md:text-left'>
                        <p className='text-sm text-muted-foreground'>&copy; {new Date().getFullYear()} IDEAS - Todos os direitos reservados</p>
                    </div>

                    <div className='mt-4 flex items-center justify-center md:mt-0'>
                        <div className='flex flex-col md:flex-row'>
                            <Link href='/' className={cn(buttonVariants({ variant: 'link', className: 'text-sm text-muted-foreground' }))}>
                                Termos de uso
                            </Link>
                            <Link href='/' className={cn(buttonVariants({ variant: 'link', className: 'text-sm text-muted-foreground' }))}>
                                Politica de privacidade
                            </Link>
                            <Link
                                href='https://github.com/Emauri-Neto/ideas/wiki'
                                className={cn(buttonVariants({ variant: 'link', className: 'text-sm text-muted-foreground' }))}
                            >
                                Entre em contato
                            </Link>
                        </div>
                    </div>
                </div>
            </MaxWidthWrapper>
        </footer>
    );
};

export default Footer;

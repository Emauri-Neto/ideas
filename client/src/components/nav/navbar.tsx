import { Icons } from '@/config/icons';
import Link from 'next/link';
import NavItems from '@/components/nav/nav-items';
import ThemeToggler from '@/components/theme-toggler';
import { Bell } from 'lucide-react';
import { Button, buttonVariants } from '@/components/ui/button';

const Navbar = () => {
    const user = undefined;
    return (
        <div className='sticky z-100 h-20 inset-x-0 top-0 w-full backdrop-blur-lg transition-all bg-secondary'>
            <header className='relative'>
                <nav className='border-b px-2 md:px-10'>
                    <div className='flex h-20 items-center'>
                        <Link href='/' className='text-3xl font-medium italic mr-5'>
                            <Icons.logo className='w-44 h-16 fill-black dark:fill-white' />
                        </Link>

                        <div className='hidden z-50 lg:ml-8 lg:block lg:self-center'>
                            <NavItems />
                        </div>

                        <div className='ml-auto flex items-center'>
                            <div className='hidden lg:flex lg:flex-1 lg:items-center lg:justify-end lg:space-x-3'>
                                <div className='flow-root lg:ml-6'>
                                    <ThemeToggler />
                                </div>

                                <div className='flow-root lg:ml-6'>
                                    <Button variant='ghost' className='hover:bg-gray-400/60 hover:dark:bg-gray-900'>
                                        <Bell />
                                    </Button>
                                </div>

                                {user ? <span className='h-6 w-px bg-gray-300 dark:bg-gray-300' aria-hidden='true' /> : null}

                                {user ? null : (
                                    <div className='flex lg:ml-6 '>
                                        <span className='h-6 w-px bg-gray-300 dark:bg-gray-300' aria-hidden='true' />
                                    </div>
                                )}

                                {user ? null : (
                                    <Link
                                        href='/sign-in'
                                        className={buttonVariants({ variant: 'ghost', className: 'hover:bg-gray-300 hover:dark:bg-gray-900' })}
                                    >
                                        Entrar
                                    </Link>
                                )}

                                {user ? null : <span className='h-6 w-px bg-gray-300' aria-hidden='true' />}

                                {user ? (
                                    <>
                                        {/* <NavAccount /> */}
                                        aaa
                                    </>
                                ) : (
                                    <Link
                                        href='/sign-up'
                                        className={buttonVariants({ variant: 'ghost', className: 'hover:bg-gray-300 hover:dark:bg-gray-900' })}
                                    >
                                        Criar Conta
                                    </Link>
                                )}
                            </div>
                        </div>
                    </div>
                </nav>
            </header>
        </div>
    );
};

export default Navbar;

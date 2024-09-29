import Link from 'next/link';
import { Button, buttonVariants } from '@/components/ui/button';
import ThemeToggler from '@/components/themeToggler';
import NavItems from '@/components/nav/navItems';
import { LoginButton } from '../auth/authBtn';

const Navbar = () => {
    return (
        <div className='sticky z-50 h-14 inset-x-0 top-0 w-full'>
            <header className='relative'>
                <div className='border-b border-gray-200 dark:border-gray-300'>
                    <div className='flex h-16 items-center'>
                        <div className='ml-4 flex lg:ml-0'>
                            <Link href='/' className='text-2xl font-medium italic'>
                                🚀 Ideas
                            </Link>
                        </div>

                        <div className='hidden z-50 lg:ml-8 lg:block lg:self-center bg-background'>
                            <NavItems />
                        </div>

                        <div className='ml-auto flex items-center'>
                            <div className='hidden lg:flex lg:flex-1 lg:items-center lg:justify-end lg:space-x-6'>
                                <div className='ml-4 flow-root lg:ml-6'>
                                    <ThemeToggler />
                                </div>

                                {true ? <span className='h-6 w-px bg-gray-200 dark:bg-gray-300' aria-hidden='true' /> : null}

                                {true ? null : (
                                    <div className='flex lg:ml-6 '>
                                        <span className='h-6 w-px bg-gray-200 dark:bg-gray-300' aria-hidden='true' />
                                    </div>
                                )}

                                {false ? null : (
                                    <LoginButton>
                                        <Button variant='ghost'>Entrar</Button>
                                    </LoginButton>
                                )}

                                {false ? null : <span className='h-6 w-px bg-gray-200' aria-hidden='true' />}

                                {false ? (
                                    <>
                                        {/* <NavAccount /> */}
                                        aaa
                                    </>
                                ) : (
                                    <Link href='/sign-up' className={buttonVariants({ variant: 'ghost' })}>
                                        Criar Conta
                                    </Link>
                                )}
                            </div>
                        </div>
                    </div>
                </div>
            </header>
        </div>
    );
};

export default Navbar;

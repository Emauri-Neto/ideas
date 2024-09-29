'use client';

import { Card, CardContent, CardFooter, CardHeader } from '@/components/ui/card';
import { Header } from '@/components/auth/card/header';
import Socials from '@/components/auth/card/socials';
import LinkButton from '@/components/auth/card/link-button';

interface CardProps {
    children: React.ReactNode;
    headerLabel: string;
    buttonHref: string;
    buttonLabel: string;
    showSocials?: boolean;
}

export const AuthWrapper = ({ children, headerLabel, buttonLabel, buttonHref, showSocials }: CardProps) => {
    return (
        <Card className='w-[400px] shadow-md'>
            <CardHeader className='flex flex-col items-center justify-center'>
                <Header label={headerLabel} />
                <LinkButton label={buttonLabel} href={buttonHref} />
            </CardHeader>
            <CardContent>{children}</CardContent>
            {showSocials && (
                <div className='flex flex-col gap-y-5 items-center'>
                    <div className='relative w-10/12'>
                        <div className='absolute inset-0 flex items-center' aria-hidden='true'>
                            <span className='w-full border-t' />
                        </div>
                        <div className='relative flex justify-center text-xs uppercase'>
                            <span className='bg-background mx-auto px-2 text-muted-foreground'>ou</span>
                        </div>
                    </div>
                    <CardFooter>
                        <Socials />
                    </CardFooter>
                </div>
            )}
        </Card>
    );
};

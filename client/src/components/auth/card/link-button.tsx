'use client';

import { Button } from '@/components/ui/button';
import Link from 'next/link';

type LinkButtonProps = {
    label: string;
    href: string;
};

const LinkButton = ({ label, href }: LinkButtonProps) => {
    return (
        <Button variant='link' className='font-normal w-fit text-muted-foreground' size='sm' asChild>
            <Link href={href}>{label}</Link>
        </Button>
    );
};

export default LinkButton;

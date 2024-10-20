'use client';

import { Button } from '@/components/ui/button';
import { FcGoogle } from 'react-icons/fc';

const Socials = () => {
    return (
        <div className='flex items-center w-full gap-x-2'>
            <Button className='w-full' variant='outline'>
                <FcGoogle className='h-5 w-5' />
            </Button>
        </div>
    );
};

export default Socials;

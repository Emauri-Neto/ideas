'use client';

import React from 'react';
import { Button } from '@/components/ui/button';
import { useToast } from '@/hooks/use-toast';
import { Share2Icon } from 'lucide-react';

interface Props {
    studyID: string;
    className?: string;
}

const CopyToClipBoard = ({ studyID, className }: Props) => {
    const { toast } = useToast();
    function copyToClipboard() {
        navigator.clipboard.writeText(`http://localhost:3000/r/study/${studyID}`);

        return toast({
            title: 'Copiado para a área de transferência.'
        });
    }
    return (
        <Button variant='ghost' onClick={copyToClipboard}>
            <Share2Icon />
        </Button>
    );
};

export default CopyToClipBoard;

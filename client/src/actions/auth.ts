'use server';

import { User } from '@/types';
import { cookies } from 'next/headers';

export default async function Auth() {
    const cookie = cookies().get('IAUTH')?.value;

    try {
        const response = await fetch('http://localhost:4367/api/user', {
            method: 'GET',
            headers: {
                Authorization: `Bearer ${cookie}`
            }
        });

        if (!response.ok) {
            console.log('Erro na req ->', response.statusText);
        }

        const res = await response.json();

        const user: User = res.data;

        console.log('user', user);

        return user;
    } catch (error) {
        console.log('Erro durante a requisição 🤓🤓🤓 ->', error);
    }
}

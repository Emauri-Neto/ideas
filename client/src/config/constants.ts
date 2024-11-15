import { Icons } from '@/components/global/icons';
import { HomeIcon, Library, LucideProps, Settings } from 'lucide-react';

export interface MenuProps {
    title: string;
    href: string;
    icon: React.FC<LucideProps>;
}

type Menu = (id?: string | number) => MenuProps[];

export const MenuItens: Menu = (id) => [
    {
        title: 'Home',
        href: '/',
        icon: HomeIcon
    },
    {
        title: 'Estudos',
        href: `/r/study`,
        icon: Library
    },
    {
        title: 'Configurações',
        href: `/r/user/settings`,
        icon: Settings
    }
];

export const API_BASE_URL = 'http://localhost:4367';

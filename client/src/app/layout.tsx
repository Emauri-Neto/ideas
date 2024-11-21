import type { Metadata } from 'next';
import { Manrope } from 'next/font/google';
import './globals.css';
import { ThemeProvider } from '@/components/theme/theme-provider';
import { QueryClientProvider } from '@tanstack/react-query';
import client from '@/lib/query-client';
import { Toaster } from '@/components/ui/toaster';

export const metadata: Metadata = {
    title: 'Ideas',
    description:
        'Uma plataforma integrada e aberta, projetada para facilitar o relacionamento e o compartilhamento de ideias entre especialistas, formuladores de estratégia, políticas públicas e cenaristas. Permita que discussões sobre o futuro e decisões estratégicas aconteçam em qualquer lugar, a qualquer momento.'
};

const font = Manrope({ subsets: ['latin'] });

export default function RootLayout({
    children
}: Readonly<{
    children: React.ReactNode;
}>) {
    return (
        <html lang='en'>
            <body className={`${font.className} antialiased min-h-screen h-full`}>
                <QueryClientProvider client={client}>
                    <ThemeProvider attribute='class' defaultTheme='light' enableSystem disableTransitionOnChange>
                        {children}
                    </ThemeProvider>
                    <Toaster />
                </QueryClientProvider>
            </body>
        </html>
    );
}

import type { Metadata } from 'next';
import { Manrope } from 'next/font/google';
import './globals.css';
import { ThemeProvider } from '@/components/theme/theme-provider';

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
                <ThemeProvider attribute='class' defaultTheme='light' enableSystem disableTransitionOnChange>
                    {children}
                </ThemeProvider>
            </body>
        </html>
    );
}

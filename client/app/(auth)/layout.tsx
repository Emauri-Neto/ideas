export default function AuthLayout({
    children
}: Readonly<{
    children: React.ReactNode;
}>) {
    return <div className='h-full min-h-screen flex items-center justify-center'>{children}</div>;
}

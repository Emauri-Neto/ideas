import Auth from '@/actions/auth';
import Header from '@/components/global/header';
import Sidebar from '@/components/nav/sidebar';

export default async function Layout({
    children
}: Readonly<{
    children: React.ReactNode;
}>) {
    const user = await Auth();
    const id: string = 'aaaaaa';

    return (
        <div className='flex h-screen w-screen'>
            <Sidebar activeStudyId={id} />
            <div className='w-full pt-28 p-6 overflow-x-hidden overflow-y-hidden'>
                <Header />
                <div className='mt-4'>{children}</div>
            </div>
        </div>
    );
}

import { ListStudies } from '@/actions/study';
import { GetUser } from '@/actions/user';
import Header from '@/components/global/header';
import Sidebar from '@/components/nav/sidebar';
import { dehydrate, HydrationBoundary, QueryClient } from '@tanstack/react-query';

export default async function Layout({
    children
}: Readonly<{
    children: React.ReactNode;
}>) {
    const query = new QueryClient();

    await query.prefetchQuery({
        queryKey: ['user-data'],
        queryFn: GetUser
    });

    await query.prefetchQuery({
        queryKey: ['study-data'],
        queryFn: ListStudies
    });

    return (
        <HydrationBoundary state={dehydrate(query)}>
            <div className='flex h-screen w-screen'>
                <Sidebar />
                <div className='w-full pt-28 p-6 overflow-x-hidden overflow-y-hidden'>
                    <Header />
                    <div className='mt-4'>{children}</div>
                </div>
            </div>
        </HydrationBoundary>
    );
}

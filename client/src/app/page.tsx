import FilterMenu from '@/components/filter';
import Navbar from '@/components/nav/navbar';
import StudyArea from '@/components/study/study-area';

export default function Home() {
    return (
        <div className='flex flex-col'>
            <Navbar />
            <div className='container flex items-center justify-center mt-5'>
                <div className='md:grid lg:grid-cols-4 md:grid-cols-3'>
                    <FilterMenu />
                    <div className='lg:col-span-3 md:col-span-2'>
                        <StudyArea />
                    </div>
                </div>
            </div>
        </div>
    );
}

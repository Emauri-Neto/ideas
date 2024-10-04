import Footer from '@/components/footer';
import Navbar from '@/components/nav/navbar';
import MaxWidthWrapper from '@/components/wrapper';

export default function Home() {
    return (
        <MaxWidthWrapper>
            <Navbar />
            <div className='mt-5 relative mx-auto'>
                <p>aaaaa</p>
            </div>
            <Footer />
        </MaxWidthWrapper>
    );
}

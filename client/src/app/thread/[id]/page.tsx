import Navbar from '@/components/nav/navbar';
import Message from '@/components/message/message';
import { Input } from '@/components/ui/input';

export default function Thread() {
    const thread = {
        name: 'Thread Name',
        description: `Lorem ipsum dolor sit amet consectetur adipisicing elit. Dolorem maxime, 
        error corporis repellendus in architecto aspernatur totam neque accusamus facere! Doloribus quo fugit vero quod debitis cum libero at magni?
        error corporis repellendus in architecto aspernatur totam neque accusamus facere! Doloribus quo fugit vero quod debitis cum libero at magni?
        error corporis repellendus in architecto aspernatur totam neque accusamus facere! Doloribus quo fugit vero quod debitis cum libero at magni?
        error corporis repellendus in architecto aspernatur totam neque accusamus facere! Doloribus quo fugit vero quod debitis cum libero at magni?
        error corporis repellendus in architecto aspernatur totam neque accusamus facere! Doloribus quo fugit vero quod debitis cum libero at magni?
        error corporis repellendus in architecto aspernatur totam neque accusamus facere! Doloribus quo fugit vero quod debitis cum libero at magni?
        `
    }
    
    const msg =[{
        user: "user 1",
        msg: `Lorem ipsum dolor sit amet consectetur adipisicing elit. Commodi ab eum fuga. Ut eius voluptatem sit, 
        id quod natus totam a quo provident modi, eaque molestiae expedita, beatae aspernatur asperiores?`
    }, 
    {
        user: "user 1",
        msg: `Lorem ipsum dolor sit amet consectetur adipisicing elit. Commodi ab eum fuga. Ut eius voluptatem sit, 
        id quod natus totam a quo provident modi, eaque molestiae expedita, beatae aspernatur asperiores?`
    }, 
    {
        user: "user 1",
        msg: `Lorem ipsum dolor sit amet consectetur adipisicing elit. Commodi ab eum fuga. Ut eius voluptatem sit, 
        id quod natus totam a quo provident modi, eaque molestiae expedita, beatae aspernatur asperiores?`
    },
    {
        user: "user 1",
        msg: `Lorem ipsum dolor sit amet consectetur adipisicing elit. Commodi ab eum fuga. Ut eius voluptatem sit, 
        id quod natus totam a quo provident modi, eaque molestiae expedita, beatae aspernatur asperiores?`
    }]

    return (
        <div className='flex flex-col h-screen'>
            <Navbar />
            <div className="container h-full flex flex-col items-center justify-left mt-5 mx-auto px-10">
                <div className='md:grid h-full 2xl:grid-cols-6 xl:grid-cols-5 lg:grid-cols-4 md:grid-cols-3'>
                    <div className='2xl:col-span-6 xl:col-span-5 lg:col-span-4 md:col-span-3 flex flex-col justify-center items-center'>
                            <div className='text-lg font-bold'>{thread.name}</div>
                            <div className='text-left'>{thread.description}</div>
                    </div>
                    <div className='2xl:col-span-6 xl:col-span-5 lg:col-span-4 md:col-span-3 flex flex-col justify-center mx-10'>
                        {msg.map((msg, i) => (
                            <div key={i} className='mt-5'>
                                <Message msg={msg}/>
                            </div>))}
                    </div>
                </div>
                
                <div className='sticky bottom-0 left-0 right-0 w-full rounded-lg p-2 bg-white dark:bg-gray-900 shadow-lg mt-10 mx-10'>
                    <div className='dark:bg-gray-900'>
                        <Input placeholder='Digite...' className='bg-gray-300 dark:bg-gray-900 w-full' />
                    </div>
                </div>
            </div>
            
        </div>
    );
}

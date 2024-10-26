'use client'

interface Props {
    msg: any
}

export default function message(props: Props){

    return(
            <div>
                <div className="font-bold mb-2">
                    {props.msg.user}
                </div>
                <div className="px-4 py-2 bg-gray-300 dark:bg-gray-900 rounded-lg">
                    {props.msg.msg}
                </div>
            </div>
    )
}
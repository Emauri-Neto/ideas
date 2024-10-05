'use server';

import { cookies } from 'next/headers';

interface CookieOpts {
    name: string;
    value: any;
    httpOnly?: boolean;
    sameSite?: any;
    secure?: boolean;
    maxAge?: number;
}

export const GetCookies = (cookie: string): any => {
    return cookies().get(cookie);
};

export const SetCookie = ({ name, value, httpOnly = true, maxAge, sameSite, secure }: CookieOpts): any => {
    cookies().set({
        name,
        value,
        httpOnly,
        sameSite,
        secure,
        maxAge
    });
};

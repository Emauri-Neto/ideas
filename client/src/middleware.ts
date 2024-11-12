import { NextRequest } from 'next/server';

const authRoutes: string[] = ['/sign-in', '/sign-up'];
const publicRoutes: string[] = [];

export default function middleware(req: NextRequest) {
    // const { nextUrl } = req;

    // const isLogged = !!req.cookies.get('refreshToken');

    // const isPublicRoute = publicRoutes.includes(nextUrl.pathname);
    // const isAuthRoute = authRoutes.includes(nextUrl.pathname);

    // if (isAuthRoute) {
    //     if (isLogged) {
    //         return Response.redirect(new URL('/', nextUrl));
    //     }
    //     return null;
    // }

    // if (!isLogged && !isPublicRoute) {
    //     return Response.redirect(new URL('/sign-in', nextUrl));
    // }

    return null;
}

export const config = {
    matcher: [
        '/((?!_next|[^?]*\\.(?:html?|css|js(?!on)|jpe?g|webp|png|gif|svg|ttf|woff2?|ico|csv|docx?|xlsx?|zip|webmanifest)).*)',
        '/(api|trpc)(.*)'
    ]
};

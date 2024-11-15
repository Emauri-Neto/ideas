'use client';

import { QueryClient } from '@tanstack/react-query';

const client = new QueryClient({
    defaultOptions: {
        queries: {
            retry: false
        }
    }
});

export default client;

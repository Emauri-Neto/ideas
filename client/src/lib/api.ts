'use server';

import { API_BASE_URL } from '@/config/constants';
import axios from 'axios';
import { cookies } from 'next/headers';

type AxiosCfg = {
    baseURL: string;
    withCredentials: boolean;
};

const opts: AxiosCfg = {
    baseURL: API_BASE_URL,
    withCredentials: true
};

const api = axios.create({
    ...opts,
    timeout: 1000,
    headers: {
        'Content-Type': 'application/json'
    }
});

api.interceptors.request.use(
    async (config) => {
        return config;
    },
    async (error) => {
        return Promise.reject(error);
    }
);

api.interceptors.response.use(
    (response) => response.data,
    async (error) => {
        return Promise.reject(error);
    }
);

export default api;

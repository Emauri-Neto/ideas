'use server';

import * as z from 'zod';

export const GetStudies = async () => {
    const res = await fetch('http://localhost:4367/study/get');

    const data = await res.json();

    return data;
};

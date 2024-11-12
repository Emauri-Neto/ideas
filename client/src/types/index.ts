export interface User {
    readonly id: string;
    name: string;
    role: 'admin' | 'user';
    email: string;
    created_at: Date;
    updated_at: Date;
}

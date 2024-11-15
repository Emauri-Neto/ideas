export interface User {
    readonly id: string;
    // name: string;
    // role: 'admin' | 'user';
    _verified: boolean;
    email: string;
    created_at: Date;
    updated_at: Date;
}

export interface Study {
    readonly id: string;
    title: string;
    objective?: string;
    methodology?: 'prospective' | 'intuitive' | 'probabilistic' | 'custom';
    max_participants?: number;
    num_participants?: number;
    _private: boolean;
    user_id: string;
    created_at: Date;
    updated_at: Date;
}

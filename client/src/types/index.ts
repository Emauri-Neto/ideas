export interface User {
    readonly id: string;
    // name: string;
    // role: 'admin' | 'user';
    _verified: boolean;
    email: string;
    created_at: Date;
    updated_at: Date;
}

export interface Thread {
    readonly id: string;
    name: string;
    deadline: Date;
    responsible_user: string;
    study_id: string;
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
    threads: Thread[];
    created_at: Date;
    updated_at: Date;
}

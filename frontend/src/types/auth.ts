export type LoginRequest = {
    login: string;
    password: string;
};

export type RegisterRequest = {
    username: string;
    email: string;
    password: string;
};

export type AuthResponse = {
    token: string;
};

export type JwtPayload = {
    exp: number;
    iat: number;
    user_id: number;
};

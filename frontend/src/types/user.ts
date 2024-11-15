export type User = {
    id: number;
    email: string;
    username: string;
};

export type UserStoreState = {
    user: User | undefined;
    isAuth: boolean;
    isUserFetching: boolean;
    setCurrentUser: (user: User) => void;
    deleteCurrentUser: () => void;
    checkAuth: () => Promise<void>;
};

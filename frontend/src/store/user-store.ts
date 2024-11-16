import { create, StateCreator } from "zustand";
import { UserStoreState } from "@/types/user";
import AuthService from "@/services/auth-service";
import { toast } from "sonner";

const userState: StateCreator<UserStoreState> = (set) => ({
    user: undefined,
    isAuth: false,
    isUserFetching: false,
    setCurrentUser: (user) => set({ user, isAuth: true }),
    deleteCurrentUser: () => set({ user: undefined, isAuth: false }),
    checkAuth: async () => {
        if (!localStorage.getItem("auth_token")) {
            set({ user: undefined, isAuth: false });
            return;
        }
        set({ isUserFetching: true });
        try {
            const response = await AuthService.identity();

            set({ user: response.data, isAuth: true });
        } catch (error) {
            console.error("Ошибка в проверке токена:", error);
            toast.error("Ошибка при проверке ауфа.");
        } finally {
            set({ isUserFetching: false });
        }
    },
});

const useUserStore = create(userState);

export default useUserStore;

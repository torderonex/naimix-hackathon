import useUserStore from "@/store/user-store";

export default function MainPage() {
    const { isAuth, user } = useUserStore();
    return (
        <div className="flex flex-col items-center min-h-screen px-16 py-8">
            <img src="/logo.svg" alt="logo" className="w-52 mb-3" />
            {isAuth ? `Авторизован - ${user?.username}` : "Не авторизован"}
        </div>
    );
}

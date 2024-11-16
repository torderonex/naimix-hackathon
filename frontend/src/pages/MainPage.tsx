import useUserStore from "@/store/user-store";
import { Link } from "react-router-dom";

export default function MainPage() {
    const { isAuth, user } = useUserStore();

    if (isAuth) {
        return (
            <div className="flex flex-col items-center min-h-screen px-16 py-8">
                <h1 className="text-2xl text-bald text-center mt-10">
                    Добро пожаловать, {user?.username}!
                </h1>
                <div className="mt-10 flex gap-5">
                    <Link to="/add-team" className="text-primary underline">
                        Добавить команду
                    </Link>
                    <Link
                        to="/add-participant"
                        className="text-primary underline"
                    >
                        Добавить участника
                    </Link>
                </div>
            </div>
        );
    }

    return (
        <h1 className="text-lg text-bald text-center mt-20">
            Вы не авторизованы!{" "}
            <Link to="/login" className="underline">
                Войди
            </Link>{" "}
            в аккаунт, чтобы пользоваться инструментом.
        </h1>
    );
}

import { Link, Outlet, useNavigate } from "react-router-dom";
import {
    NavigationMenu,
    NavigationMenuItem,
    NavigationMenuLink,
    NavigationMenuList,
    navigationMenuTriggerStyle,
} from "../ui/navigation-menu";
import {
    DropdownMenu,
    DropdownMenuContent,
    DropdownMenuItem,
    DropdownMenuTrigger,
} from "../ui/dropdown-menu";
import { Avatar, AvatarImage } from "../ui/avatar";
import { Button } from "../ui/button";
import { Skeleton } from "../ui/skeleton";
import useUserStore from "@/store/user-store";
import { useMediaQuery } from "react-responsive";

export default function HeaderLayout() {
    const responsive = useMediaQuery({
        query: "(min-width: 550px)",
    });

    return (
        <>
            <header className="flex justify-between py-3 px-16 items-center top-0 sticky w-full bg-white z-50 shadow h-[72px]">
                <div className="flex items-center gap-10">
                    {responsive && (
                        <Link to="/">
                            <img
                                src="/logo.svg"
                                alt="Naimix"
                                className="logo"
                                height={0}
                            />
                        </Link>
                    )}
                    <HeaderLayoutNav />
                </div>
                <HeaderLayoutRight />
            </header>
            <main>
                <Outlet />
            </main>
        </>
    );
}

export function HeaderLayoutRight() {
    const { isUserFetching, isAuth, deleteCurrentUser } = useUserStore();
    const navigate = useNavigate();

    const handleLogout = () => {
        deleteCurrentUser();
        localStorage.removeItem("auth_token");
        navigate("/");
    };

    if (isUserFetching) {
        return <Skeleton className="h-10 w-10 rounded-full" />;
    }

    if (isAuth) {
        return (
            <DropdownMenu>
                <DropdownMenuTrigger asChild>
                    <Avatar className="w-10 h-10 cursor-pointer">
                        <AvatarImage
                            src="/pfp.png"
                            alt="User Avatar"
                            className="rounded-full"
                        />
                    </Avatar>
                </DropdownMenuTrigger>
                <DropdownMenuContent>
                    <DropdownMenuItem
                        variant="destructive"
                        onClick={handleLogout}
                    >
                        Выйти
                    </DropdownMenuItem>
                </DropdownMenuContent>
            </DropdownMenu>
        );
    }

    return (
        <div className="flex gap-5">
            <Link to="/login">
                <Button variant={"outline"}>Войти</Button>
            </Link>
            <Link to="/register">
                <Button>Регистрация</Button>
            </Link>
        </div>
    );
}

export function HeaderLayoutNav() {
    return (
        <NavigationMenu>
            <NavigationMenuList>
                <NavigationMenuItem>
                    <NavigationMenuLink
                        href="/add-participant"
                        className={navigationMenuTriggerStyle()}
                    >
                        Участники
                    </NavigationMenuLink>
                    <NavigationMenuLink
                        href="/add-team"
                        className={navigationMenuTriggerStyle()}
                    >
                        Команды
                    </NavigationMenuLink>
                </NavigationMenuItem>
            </NavigationMenuList>
        </NavigationMenu>
    );
}

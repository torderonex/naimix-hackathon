import useUserStore from "@/store/user-store";
import { PropsWithChildren, useEffect } from "react";
import { useNavigate } from "react-router-dom";
import Spinner from "../ui/spinner";

export function OnlyForNotAuthGuard({ children }: PropsWithChildren) {
    const { isUserFetching, isAuth } = useUserStore();
    const navigate = useNavigate();

    useEffect(() => {
        if (isAuth) {
            navigate("/");
        }
    }, [isAuth, navigate]);

    if (isUserFetching) {
        return <Spinner />;
    }

    return !isAuth ? children : null;
}

export function OnlyForAuthGuard({ children }: PropsWithChildren) {
    const { isUserFetching, isAuth } = useUserStore();
    const navigate = useNavigate();

    useEffect(() => {
        if (isAuth) {
            navigate("/");
        }
    }, [isAuth, navigate]);

    if (isUserFetching) {
        return <Spinner />;
    }

    return isAuth ? children : null;
}

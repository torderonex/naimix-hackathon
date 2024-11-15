import { RouteProps } from "react-router-dom";
import MainPage from "@/pages/MainPage";
import LoginPage from "@/pages/LoginPage";
import RegisterPage from "@/pages/RegisterPage";
import { OnlyForNotAuthGuard } from "@/components/auth/guards";
import AddTeamPage from "@/pages/AddTeamPage";
import AddParticipantPage from "@/pages/AddParticipantPage";

export enum routesPath {
    MAIN = "/",
    LOGIN = "/login",
    REGISTER = "/register",

    ADD_PARTICIPANT = "/add-participant",
    ADD_TEAM = "/add-team",
}

export const routesConfig: RouteProps[] = [
    {
        path: routesPath.MAIN,
        element: <MainPage />,
    },
    {
        path: routesPath.LOGIN,
        element: (
            <OnlyForNotAuthGuard>
                <LoginPage />
            </OnlyForNotAuthGuard>
        ),
    },
    {
        path: routesPath.REGISTER,
        element: (
            <OnlyForNotAuthGuard>
                <RegisterPage />
            </OnlyForNotAuthGuard>
        ),
    },

    {
        path: routesPath.ADD_TEAM,
        element: <AddTeamPage />,
    },
    {
        path: routesPath.ADD_PARTICIPANT,
        element: <AddParticipantPage />,
    },
];

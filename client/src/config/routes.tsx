import { RouteProps } from "react-router-dom";
import MainPage from "@/pages/MainPage";
import ResultPage from "@/pages/ResultPage";

export enum routesPath {
    MAIN = "/",
    RESULT = "/result",
}

export const routesConfig: RouteProps[] = [
    {
        path: routesPath.MAIN,
        element: <MainPage />,
    },
    {
        path: routesPath.RESULT,
        element: <ResultPage />,
    },
];

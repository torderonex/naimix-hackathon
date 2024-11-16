import HeaderLayout from "@/components/layout/header-layout";
import { routesConfig } from "@/config/routes";
import { Suspense } from "react";
import { Navigate, Route, Routes } from "react-router-dom";

export default function AppRouter() {
    return (
        <Suspense fallback={<div>Загрузка...</div>}>
            <Routes>
                <Route element={<HeaderLayout />}>
                    {routesConfig.map(({ path, element }) => (
                        <Route path={path} element={element} key={path} />
                    ))}
                </Route>
                <Route path="*" element={<Navigate to="/" />} />
            </Routes>
        </Suspense>
    );
}

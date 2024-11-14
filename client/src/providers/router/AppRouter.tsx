import { routesConfig } from "@/config/routes";
import { Suspense } from "react";
import { Route, Routes } from "react-router-dom";

export default function AppRouter() {
    return (
        <Suspense fallback={<div>Загрузка...</div>}>
            <Routes>
                {routesConfig.map(({ path, element }) => (
                    <Route path={path} element={element} key={path} />
                ))}
            </Routes>
        </Suspense>
    );
}

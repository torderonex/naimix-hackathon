import { BrowserRouter } from "react-router-dom";
import AppRouter from "./providers/router/AppRouter";
import { Toaster } from "sonner";
import useUserStore from "./store/user-store";
import { useEffect } from "react";

function App() {
    const { checkAuth } = useUserStore();

    useEffect(() => {
        if (localStorage.getItem("auth_token")) {
            checkAuth();
        }
    }, [checkAuth]);

    return (
        <BrowserRouter>
            <AppRouter />
            <Toaster />
        </BrowserRouter>
    );
}

export default App;

import { BrowserRouter } from "react-router-dom";
import AppRouter from "./providers/router/AppRouter";
import { Toaster } from "sonner";

function App() {
    return (
        <BrowserRouter>
            <AppRouter />
            <Toaster />
        </BrowserRouter>
    );
}

export default App;

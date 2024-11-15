import { Loader } from "lucide-react";

export default function Spinner() {
    return (
        <div className="max-w-full h-screen flex justify-center items-center">
            <Loader className="animate-spin"/>
        </div>
    )
}
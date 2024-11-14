import CompatibilityForm from "@/components/CompatibilityForm";
import CompatibilityTable from "@/components/CompatibilityTable";

export default function MainPage() {
    return (
        <div className="flex flex-col items-center min-h-screen px-16 py-8">
            <img src="/logo.svg" alt="logo" className="w-52 mb-3" />
            <CompatibilityForm />
            <CompatibilityTable title="Последние проверки совместимости" />
        </div>
    );
}

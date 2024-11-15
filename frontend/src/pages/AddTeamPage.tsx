import AddTeamForm from "@/components/team/AddTeamForm";
import TeamsTable from "@/components/team/TeamsTable";

export default function AddTeamPage() {
    return (
        <div className="flex flex-col items-center px-16 py-8">
            <AddTeamForm />
            <TeamsTable />
        </div>
    );
}

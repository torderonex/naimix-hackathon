import AddTeamForm from "@/components/team/AddTeamForm";
import TeamsTable from "@/components/team/TeamsTable";
import { useTeams } from "@/hooks/useTeams";

export default function AddTeamPage() {
    const { teams, loading, getAllTeams } = useTeams();

    return (
        <div className="flex flex-col items-center px-16 py-8">
            <AddTeamForm refetchTeams={getAllTeams} />
            <TeamsTable
                teams={teams}
                loading={loading}
                refetchTeams={getAllTeams}
            />
        </div>
    );
}

import {
    Table,
    TableBody,
    TableCell,
    TableHead,
    TableHeader,
    TableRow,
} from "@/components/ui/table";
import TeamActions from "./TeamActions";
import { Team } from "@/types/team";

interface TeamsTableProps {
    teams: Team[] | null;
    loading: boolean;
    refetchTeams: () => void;
}

export default function TeamsTable({
    teams,
    loading,
    refetchTeams,
}: TeamsTableProps) {
    if (loading) {
        return <div className="mt-10">Загрузка...</div>;
    }

    return (
        <div className="mt-10 w-[1200px] max-w-full">
            <h3 className="text-center my-8 text-2xl">Добавленные команды</h3>
            <Table>
                <TableHeader>
                    <TableRow>
                        <TableHead>Название команды</TableHead>
                        <TableHead>Описание</TableHead>
                        <TableHead className="text-end">Действия</TableHead>
                    </TableRow>
                </TableHeader>
                <TableBody>
                    {teams === null || teams.length === 0 ? (
                        <TableRow>
                            <TableCell>Команд пока что нет.</TableCell>
                        </TableRow>
                    ) : (
                        teams.map((team) => (
                            <TableRow key={team.id}>
                                <TableCell>{team.name}</TableCell>
                                <TableCell>{team.description}</TableCell>
                                <TableCell align="right">
                                    <TeamActions
                                        team={team}
                                        refetchTeams={refetchTeams}
                                    />
                                </TableCell>
                            </TableRow>
                        ))
                    )}
                </TableBody>
            </Table>
        </div>
    );
}

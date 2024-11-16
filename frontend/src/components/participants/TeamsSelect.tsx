import { useTeams } from "@/hooks/useTeams";
import {
    Select,
    SelectContent,
    SelectItem,
    SelectTrigger,
    SelectValue,
} from "../ui/select";

interface TeamsSelectProps {
    selectedTeamId: string;
    onChange: (value: string) => void;
}

export default function TeamsSelect({
    selectedTeamId,
    onChange,
}: TeamsSelectProps) {
    const { teams, loading } = useTeams();

    if (loading) {
        return <span>Загрузка...</span>;
    }

    if (teams === null || teams.length === 0) {
        return <span>Нет команд для выбора.</span>;
    }

    return (
        <Select value={selectedTeamId} onValueChange={onChange}>
            <SelectTrigger className="w-[150px]">
                <SelectValue placeholder="Команда" />
            </SelectTrigger>
            <SelectContent>
                {teams.map((team) => (
                    <SelectItem key={team.id} value={team.id.toString()}>
                        {team.name}
                    </SelectItem>
                ))}
            </SelectContent>
        </Select>
    );
}

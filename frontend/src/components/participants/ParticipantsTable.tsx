import {
    Table,
    TableBody,
    TableCell,
    TableHead,
    TableHeader,
    TableRow,
} from "@/components/ui/table";
import ParticipantActions from "./ParticipantActions";
import { Participant } from "@/types/participant";

interface ParticipantsTableProps {
    participants: Participant[] | null;
    loading: boolean;
    refetchParticipants: () => void;
}

export default function ParticipantsTable({
    participants,
    loading,
    refetchParticipants,
}: ParticipantsTableProps) {
    if (loading) {
        return <div className="mt-10">Загрузка...</div>;
    }

    return (
        <div className="mt-10 w-[1200px] max-w-full">
            <h3 className="text-center my-8 text-2xl">Добавленные участники</h3>
            <Table>
                <TableHeader>
                    <TableRow>
                        <TableHead>Имя</TableHead>
                        <TableHead>Команда</TableHead>
                        <TableHead>Роль в команде</TableHead>
                        <TableHead>Время рождения</TableHead>
                        <TableHead>Место рождения</TableHead>
                        <TableHead className="text-end">Действия</TableHead>
                    </TableRow>
                </TableHeader>
                <TableBody>
                    {participants === null || participants.length === 0 ? (
                        <TableRow>
                            <TableCell>Участников пока что нет.</TableCell>
                        </TableRow>
                    ) : (
                        participants.map((participant) => (
                            <TableRow key={participant.id}>
                                <TableCell>{participant.name}</TableCell>
                                <TableCell>{participant.team_name}</TableCell>
                                <TableCell>{participant.role}</TableCell>
                                <TableCell>
                                    {participant.birthdate
                                        .replace("T", " ")
                                        .replace("Z", "")
                                        .slice(0, -4)}
                                </TableCell>
                                <TableCell>{participant.birthplace}</TableCell>
                                <TableCell align="right">
                                    <ParticipantActions
                                        participant={participant}
                                        refetchParticipants={
                                            refetchParticipants
                                        }
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

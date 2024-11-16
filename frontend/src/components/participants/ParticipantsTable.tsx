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
import ComparisonService from "@/services/comparison-service";
import { toast } from "sonner";
import { useLocation, useNavigate } from "react-router-dom";

interface ParticipantsTableProps {
    participants: Participant[] | null;
    loading: boolean;
    refetchParticipants: () => void;
    comparisonPercents: number[] | null;
    birthDate: string;
    setText: (s: string) => void;
}

export default function ParticipantsTable({
    participants,
    loading,
    refetchParticipants,
    comparisonPercents,
    birthDate,
    setText,
}: ParticipantsTableProps) {
    const { pathname } = useLocation();
    const navigate = useNavigate();

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
                        <TableHead>Дата рождения</TableHead>
                        <TableHead>Совместимость по натальной карте</TableHead>
                        <TableHead>Совместимость по картам таро</TableHead>
                        <TableHead className="text-end">Действия</TableHead>
                    </TableRow>
                </TableHeader>
                <TableBody>
                    {participants === null || participants.length === 0 ? (
                        <TableRow>
                            <TableCell>Участников пока что нет.</TableCell>
                        </TableRow>
                    ) : (
                        participants.map((participant, index) => (
                            <TableRow key={participant.id}>
                                <TableCell>{participant.name}</TableCell>
                                <TableCell
                                    onClick={() =>
                                        navigate(`/team/${participant.team_id}`)
                                    }
                                    className="cursor-pointer text-primary underline"
                                >
                                    {participant.team_name}
                                </TableCell>
                                <TableCell>{participant.role}</TableCell>
                                <TableCell>
                                    {participant.birthdate.split("T")[0]}
                                </TableCell>
                                <TableCell>{participant.birthplace}</TableCell>
                                <TableCell align="center">
                                    {comparisonPercents &&
                                    comparisonPercents.length > index
                                        ? `${comparisonPercents[index]}%`
                                        : "Не вычеслено"}
                                </TableCell>
                                <TableCell align="center">
                                    <span
                                        className="cursor-pointer text-amber-500 text-center"
                                        onClick={() => {
                                            if (
                                                pathname === "/add-participant"
                                            ) {
                                                toast.error(
                                                    "Перейди на страницу команды для этой функции."
                                                );
                                                return;
                                            }
                                            if (!birthDate) {
                                                toast.error(
                                                    "Установи дату рождения проверяемого участника."
                                                );
                                                return;
                                            }
                                            toast("Идет рассчет...", {
                                                duration: 1000,
                                            });
                                            const format = new Date(
                                                birthDate
                                            ).toISOString();
                                            ComparisonService.compare2(
                                                format,
                                                participant.birthdate
                                            ).then((resp) => {
                                                setText(resp.data.description);
                                                toast("Рассчет получен.", {
                                                    duration: 1000,
                                                });
                                            });
                                        }}
                                    >
                                        Рассчитать
                                    </span>
                                </TableCell>
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

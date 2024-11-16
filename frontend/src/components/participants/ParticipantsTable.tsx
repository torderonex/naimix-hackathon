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

interface ParticipantsTableProps {
    participants: Participant[] | null;
    loading: boolean;
    refetchParticipants: () => void;
    comparisonPercents : number[] | null;
    birthDate : string;
    setText : (s: string) => void;
}

export default function ParticipantsTable({
    participants,
    loading,
    refetchParticipants,
    comparisonPercents,
    birthDate,
    setText
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
                        <TableHead>Совместимость</TableHead>
                        <TableHead className="text-end">Действия</TableHead>
                    </TableRow>
                </TableHeader>
                <TableBody>
                    {participants === null || participants.length === 0 ? (
                        <TableRow>
                            <TableCell>Участников пока что нет.</TableCell>
                        </TableRow>
                    ) : (
                        participants.map((participant,index) => (
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
                                <TableCell>
                                    {comparisonPercents && comparisonPercents.length > index
                                        ? `${comparisonPercents[index]}%`
                                        : '?' 
                                    }
                                </TableCell>
                                <TableCell align="right">
                                    <ParticipantActions
                                        participant={participant}
                                        refetchParticipants={
                                            refetchParticipants
                                        }
                                    />
                                    <span className="cursor-pointer text-amber-500"  onClick={() => {
                                        if (!birthDate){
                                            setText('Установите дату рождения')
                                            return
                                        }
                                        const format = new Date(birthDate).toISOString()
                                        ComparisonService.compare2(format,participant.birthdate)
                                        .then(resp => {
                                            setText(resp.data.description )
                                        })
                                    }}>
                                        Совместимость
                                    </span>
                                </TableCell>
                            </TableRow>
                        ))
                    )}
                </TableBody>
            </Table>
        </div>
    );
}

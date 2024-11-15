import {
    Table,
    TableBody,
    TableCell,
    TableHead,
    TableHeader,
    TableRow,
} from "@/components/ui/table";
import ParticipantActions from "./ParticipantActions";

export default function ParticipantsTable() {
    return (
        <div className="mt-10 w-[1200px] max-w-full">
            <h3 className="text-center my-8 text-2xl">Добавленные команды</h3>
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
                    <TableRow>
                        <TableCell>Иван</TableCell>
                        <TableCell>Фронтенд разработка</TableCell>
                        <TableCell>Фронтенд разработчик</TableCell>
                        <TableCell>20.12.2004</TableCell>
                        <TableCell>Россия, Москва</TableCell>
                        <TableCell align="right">
                            <ParticipantActions />
                        </TableCell>
                    </TableRow>
                </TableBody>
            </Table>
        </div>
    );
}

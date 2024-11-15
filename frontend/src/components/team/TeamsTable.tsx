import {
    Table,
    TableBody,
    TableCell,
    TableHead,
    TableHeader,
    TableRow,
} from "@/components/ui/table";
import TeamActions from "./TeamActions";

export default function TeamsTable() {
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
                    <TableRow>
                        <TableCell>Фронтенд разработка</TableCell>
                        <TableCell>
                            Команда из Яндекса, 10 участников, ищут мидла
                        </TableCell>
                        <TableCell align="right">
                            <TeamActions />
                        </TableCell>
                    </TableRow>
                </TableBody>
            </Table>
        </div>
    );
}

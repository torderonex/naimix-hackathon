import {
    Table,
    TableBody,
    TableCell,
    TableHead,
    TableHeader,
    TableRow,
} from "@/components/ui/table";
import { Link } from "react-router-dom";

interface CompatibilityTableProps {
    title: string;
}

export default function CompatibilityTable({ title }: CompatibilityTableProps) {
    return (
        <>
            <h3 className="text-center my-8">{title}</h3>
            <Table>
                <TableHeader>
                    <TableRow>
                        <TableHead>Дата проверки</TableHead>
                        <TableHead className="w-[300px]">Участники</TableHead>
                        <TableHead>Коммуникация</TableHead>
                        <TableHead>Производительность</TableHead>
                        <TableHead>Общий результат</TableHead>
                        <TableHead>Комментарий</TableHead>
                        <TableHead>Ссылка</TableHead>
                    </TableRow>
                </TableHeader>
                <TableBody>
                    <TableRow>
                        <TableCell className="font-medium">
                            13.11.2024
                        </TableCell>
                        <TableCell>Александр и Евгений</TableCell>
                        <TableCell>80%</TableCell>
                        <TableCell>70%</TableCell>
                        <TableCell>75%</TableCell>
                        <TableCell>Хорошая совместимость</TableCell>
                        <TableCell>
                            <Link to="/result" className="underline">
                                Посмотреть
                            </Link>
                        </TableCell>
                    </TableRow>
                </TableBody>
            </Table>
        </>
    );
}

import { Edit2Icon, Trash2Icon } from "lucide-react";
import {
    Dialog,
    DialogContent,
    DialogHeader,
    DialogTitle,
    DialogTrigger,
} from "../ui/dialog";
import EditTeamForm from "./EditTeamForm";
import { Team } from "@/types/team";
import { toast } from "sonner";
import TeamService from "@/services/team-service";

export default function TeamActions({
    team,
    refetchTeams,
}: {
    team: Team;
    refetchTeams: () => void;
}) {
    return (
        <div className="flex justify-end">
            <EditWithDialogBtn team={team} refetchTeams={refetchTeams} />
            <DeleteBtn id={team.id} refetchTeams={refetchTeams} />
        </div>
    );
}

function EditWithDialogBtn({
    team,
    refetchTeams,
}: {
    team: Team;
    refetchTeams: () => void;
}) {
    return (
        <Dialog>
            <DialogTrigger asChild>
                <Edit2Icon className="w-4 cursor-pointer hover:fill-orange-300" />
            </DialogTrigger>
            <DialogContent className="sm:max-w-[425px]">
                <DialogHeader>
                    <DialogTitle>Редактировать команду</DialogTitle>
                </DialogHeader>
                <EditTeamForm refetchTeams={refetchTeams} team={team} />
            </DialogContent>
        </Dialog>
    );
}

function DeleteBtn({
    id,
    refetchTeams,
}: {
    id: number;
    refetchTeams: () => void;
}) {
    async function handleDelete() {
        try {
            await TeamService.delete(id);

            toast.success("Команда успешно удалена");
            refetchTeams();
        } catch (error) {
            console.error("EditTeamForm submission error", error);
            toast.error(`Ошибка при удалении команды ${id}.`);
        }
    }

    return (
        <Trash2Icon
            className="ml-2 w-5 hover:fill-red-500"
            onClick={handleDelete}
        />
    );
}

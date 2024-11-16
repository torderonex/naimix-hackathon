import { Participant } from "@/types/participant";
import {
    Dialog,
    DialogContent,
    DialogDescription,
    DialogHeader,
    DialogTitle,
    DialogTrigger,
} from "../ui/dialog";
import { Edit2Icon, Trash2Icon } from "lucide-react";
import { toast } from "sonner";
import ParticipantService from "@/services/participant-service";
import EditParticipantForm from "./EditParticipantForm";

export default function ParticipantActions({
    participant,
    refetchParticipants,
}: {
    participant: Participant;
    refetchParticipants: () => void;
}) {
    return (
        <div className="flex justify-end">
            <EditWithDialogBtn
                participant={participant}
                refetchParticipants={refetchParticipants}
            />
            <DeleteBtn
                id={participant.id}
                refetchParticipants={refetchParticipants}
            />
        </div>
    );
}

function EditWithDialogBtn({
    participant,
    refetchParticipants,
}: {
    participant: Participant;
    refetchParticipants: () => void;
}) {
    return (
        <Dialog>
            <DialogTrigger asChild>
                <Edit2Icon className="w-4 cursor-pointer hover:fill-primary" />
            </DialogTrigger>
            <DialogContent className="sm:max-w-[425px]">
                <DialogHeader>
                    <DialogTitle>Редактировать участника</DialogTitle>
                    <DialogDescription>
                        Редактирование добавленного участника.
                    </DialogDescription>
                </DialogHeader>
                <EditParticipantForm
                    participant={participant}
                    refetchParticipants={refetchParticipants}
                />
            </DialogContent>
        </Dialog>
    );
}

function DeleteBtn({
    id,
    refetchParticipants,
}: {
    id: number;
    refetchParticipants: () => void;
}) {
    async function handleDelete() {
        try {
            await ParticipantService.delete(id);

            toast.success("Команда успешно удалена");
            refetchParticipants();
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

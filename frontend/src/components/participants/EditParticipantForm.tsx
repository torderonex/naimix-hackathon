import { zodResolver } from "@hookform/resolvers/zod";
import { useForm } from "react-hook-form";
import { toast } from "sonner";
import { z } from "zod";
import {
    Form,
    FormControl,
    FormDescription,
    FormField,
    FormItem,
    FormLabel,
    FormMessage,
} from "../ui/form";
import { Input } from "../ui/input";
import { Button } from "../ui/button";
import { DatetimePicker } from "../ui/datetime-picker";
import TeamsSelect from "./TeamsSelect";
import ParticipantService from "@/services/participant-service";
import useUserStore from "@/store/user-store";
import { Participant } from "@/types/participant";
import { DialogClose } from "../ui/dialog";

const formSchema = z.object({
    name: z.string().min(2, {
        message: "Имя должно содержать больше двух символов.",
    }),
    role: z.string().min(2, {
        message: "Роль должно содержать больше двух символов.",
    }),
    birthdate: z.coerce.date(),
    birthplace: z.string().min(2, {
        message: "Место рождения должно содержать больше двух символов.",
    }),
    team_id: z.string().min(1, {
        message: "Необходимо выбрать команду.",
    }),
});

export default function EditParticipantForm({
    refetchParticipants,
    participant,
}: {
    refetchParticipants: () => void;
    participant: Participant;
}) {
    const { user } = useUserStore();
    const form = useForm<z.infer<typeof formSchema>>({
        resolver: zodResolver(formSchema),
        defaultValues: {
            name: participant.name,
            role: participant.role,
            birthdate:
                typeof participant.birthdate === "string"
                    ? new Date(participant.birthdate)
                    : participant.birthdate,
            birthplace: participant.birthplace,
            team_id: participant.team_id.toString(),
        },
    });

    const handleTeamChange = (teamId: string) => {
        form.setValue("team_id", teamId);
    };

    async function onSubmit(values: z.infer<typeof formSchema>) {
        try {
            if (user?.id === undefined) return;
            const req = {
                ...values,
                team_id: parseInt(values.team_id),
                user_id: user.id,
                id: participant.id,
                birthdate: new Date(
                    values.birthdate.getTime() + 180 * 60 * 1000
                ),
            };
            console.log(req);

            await ParticipantService.edit(req, participant.id);

            form.reset();
            refetchParticipants();
            toast.success("Участник успешно обновлен");
        } catch (error) {
            console.error("EditParticipantForm submission error", error);
            toast.error(
                "Ошибка при обновлении участника. Проверьте данные и попробуйте еще раз."
            );
        }
    }

    return (
        <div className="max-w-full">
            <Form {...form}>
                <form
                    onSubmit={form.handleSubmit(onSubmit)}
                    className="space-y-4"
                >
                    <FormField
                        control={form.control}
                        name="name"
                        render={({ field }) => (
                            <FormItem>
                                <FormLabel>Имя:</FormLabel>
                                <FormControl>
                                    <Input
                                        placeholder="Введите Имя:"
                                        {...field}
                                    />
                                </FormControl>
                                <FormDescription>Имя участника</FormDescription>
                                <FormMessage />
                            </FormItem>
                        )}
                    />

                    <TeamsSelect
                        selectedTeamId={form.watch("team_id")}
                        onChange={handleTeamChange}
                    />

                    <FormField
                        control={form.control}
                        name="role"
                        render={({ field }) => (
                            <FormItem>
                                <FormLabel>Роль:</FormLabel>
                                <FormControl>
                                    <Input
                                        placeholder="Введите Роль:"
                                        {...field}
                                    />
                                </FormControl>
                                <FormDescription>
                                    Роль в команде
                                </FormDescription>
                                <FormMessage />
                            </FormItem>
                        )}
                    />
                    <FormField
                        control={form.control}
                        name="birthdate"
                        render={({ field }) => (
                            <FormItem className="flex flex-col">
                                <FormLabel>Дата и время рождения:</FormLabel>
                                <DatetimePicker
                                    {...field}
                                    format={[
                                        ["days", "months", "years"],
                                        ["hours", "minutes", "seconds"],
                                    ]}
                                />
                                <FormDescription>
                                    Если время неизвестно - поставьте 00:00:00
                                </FormDescription>
                                <FormMessage />
                            </FormItem>
                        )}
                    />
                    <FormField
                        control={form.control}
                        name="birthplace"
                        render={({ field }) => (
                            <FormItem>
                                <FormLabel>Место рождения:</FormLabel>
                                <FormControl>
                                    <Input
                                        placeholder="Введите Место:"
                                        {...field}
                                    />
                                </FormControl>
                                <FormDescription>
                                    Введите Страну и Город
                                </FormDescription>
                                <FormMessage />
                            </FormItem>
                        )}
                    />

                    <DialogClose asChild>
                        <Button
                            type="submit"
                            className="w-full"
                            disabled={form.formState.isSubmitting}
                        >
                            {form.formState.isSubmitting
                                ? "Обновление..."
                                : "Обновить"}
                        </Button>
                    </DialogClose>
                </form>
            </Form>
        </div>
    );
}

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
    team_id: z.string(),
});

export default function AddParticipantForm({
    refetchParticipants,
}: {
    refetchParticipants: () => void;
}) {
    const { user } = useUserStore();
    const form = useForm<z.infer<typeof formSchema>>({
        resolver: zodResolver(formSchema),
        defaultValues: {
            name: "",
            role: "",
            birthdate: new Date(),
            birthplace: "",
            team_id: "",
        },
    });

    const handleTeamChange = (teamId: string) => {
        form.setValue("team_id", teamId);
    };

    async function onSubmit(values: z.infer<typeof formSchema>) {
        try {
            console.log("onSubmit");
            if (user?.id === undefined) return;
            console.log("values.team_id", values.team_id);
            if (values.team_id === "") {
                toast.error(
                    "Необходимо выбрать команду. Если ее нет - сперва создайте команду."
                );
                return;
            }
            const req = {
                ...values,
                team_id: parseInt(values.team_id),
                user_id: user.id,
                birthdate: new Date(
                    values.birthdate.getTime() + 180 * 60 * 1000
                ),
            };
            console.log(req);

            await ParticipantService.add(req);

            form.reset();
            refetchParticipants();
            toast.success("Участник успешно добавлен");
        } catch (error) {
            console.error("AddParticipantForm submission error", error);
            toast.error(
                "Ошибка при добавлении участника. Проверьте данные и попробуйте еще раз."
            );
        }
    }

    return (
        <div className="w-[500px] max-w-full">
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

                    <Button
                        type="submit"
                        className="w-full"
                        disabled={form.formState.isSubmitting}
                    >
                        {form.formState.isSubmitting
                            ? "Добавление..."
                            : "Добавить"}
                    </Button>
                </form>
            </Form>
        </div>
    );
}

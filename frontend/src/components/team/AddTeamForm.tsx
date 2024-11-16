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
import useUserStore from "@/store/user-store";
import TeamService from "@/services/team-service";

const formSchema = z.object({
    name: z.string().min(2, {
        message: "Название команды должно содержать больше двух символов.",
    }),
    description: z.string().min(2, {
        message: "Описание команды должно содержать больше двух символов.",
    }),
});

export default function AddTeamForm({
    refetchTeams,
}: {
    refetchTeams: () => void;
}) {
    const { user } = useUserStore();
    const form = useForm<z.infer<typeof formSchema>>({
        resolver: zodResolver(formSchema),
        defaultValues: {
            name: "",
            description: "",
        },
    });

    async function onSubmit(values: z.infer<typeof formSchema>) {
        try {
            if (user?.id === undefined) return;
            const req = { ...values, user_id: user.id };

            await TeamService.add(req);

            form.reset();
            refetchTeams();
            toast.success("Команда успешно добавлена");
        } catch (error) {
            console.error("AddTeamForm submission error", error);
            toast.error("Ошибка при добавлении команды.");
        }
    }

    return (
        <div className="w-[500px] max-w-full">
            <Form {...form}>
                <form
                    onSubmit={form.handleSubmit(onSubmit)}
                    className="space-y-8"
                >
                    <FormField
                        control={form.control}
                        name="name"
                        render={({ field }) => (
                            <FormItem>
                                <FormLabel>Название:</FormLabel>
                                <FormControl>
                                    <Input
                                        placeholder="Введи Название:"
                                        {...field}
                                    />
                                </FormControl>
                                <FormDescription>
                                    Название команды
                                </FormDescription>
                                <FormMessage />
                            </FormItem>
                        )}
                    />

                    <FormField
                        control={form.control}
                        name="description"
                        render={({ field }) => (
                            <FormItem>
                                <FormLabel>Описание:</FormLabel>
                                <FormControl>
                                    <Input
                                        placeholder="Введи Описание:"
                                        {...field}
                                    />
                                </FormControl>
                                <FormDescription>
                                    Описание команды
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

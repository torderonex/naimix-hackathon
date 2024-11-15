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

const formSchema = z.object({
    name: z.string().min(2, {
        message: "Имя должно содержать больше двух символов.",
    }),
    role: z.string().min(2, {
        message: "Роль должно содержать больше двух символов.",
    }),
    birthDate: z.coerce.date(),
    birthPlace: z.string().min(2, {
        message: "Место рождения должно содержать больше двух символов.",
    }),
});

export default function AddParticipantForm() {
    const form = useForm<z.infer<typeof formSchema>>({
        resolver: zodResolver(formSchema),
        defaultValues: {
            name: "",
            role: "",
            birthDate: new Date(),
            birthPlace: "",
        },
    });

    async function onSubmit(values: z.infer<typeof formSchema>) {
        try {
            console.log(values);
            // const response = await AuthService.login(values);

            // localStorage.setItem("auth_token", response.data.token);
            // checkAuth();
            // navigate("/");
        } catch (error) {
            console.error("AddParticipantForm submission error", error);
            toast.error(
                "Ошибка при добавлении команды. Проверьте данные и попробуйте еще раз."
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
                                <FormLabel>Название:</FormLabel>
                                <FormControl>
                                    <Input
                                        placeholder="Введите Название:"
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
                        name="birthDate"
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
                        name="birthPlace"
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

import { zodResolver } from "@hookform/resolvers/zod";
import { useForm } from "react-hook-form";
import { toast } from "sonner";
import { z } from "zod";
import {
    Card,
    CardContent,
    CardDescription,
    CardHeader,
    CardTitle,
} from "../ui/card";
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
import { useNavigate } from "react-router-dom";
import useUserStore from "@/store/user-store";
import AuthService from "@/services/auth-service";

const formSchema = z.object({
    username: z
        .string()
        .min(2, { message: "ФИО должно содержать больше двух символов." }),
    email: z
        .string()
        .email({ message: "Почта должна содержать больше двух символов." }),
    password: z
        .string()
        .min(8, { message: "Пароль должен содержать больше восьми символов." }),
});

export default function RegisterForm() {
    const form = useForm<z.infer<typeof formSchema>>({
        resolver: zodResolver(formSchema),
        defaultValues: {
            username: "",
            email: "",
            password: "",
        },
    });
    const navigate = useNavigate();
    const { checkAuth } = useUserStore();

    async function onSubmit(values: z.infer<typeof formSchema>) {
        try {
            const response = await AuthService.register(values);

            localStorage.setItem("auth_token", response.data.token);
            checkAuth();
            navigate("/");
        } catch (error) {
            console.error("RegisterForm submission error", error);
            toast.error(
                "Ошибка при регистрации. Проверь данные и попробуй еще раз."
            );
        }
    }

    return (
        <div className="flex justify-center">
            <Card className="w-[500px] max-w-full">
                <CardHeader>
                    <CardTitle>Регистрация</CardTitle>
                    <CardDescription>
                        Создай аккаунт, чтобы проверять совместимость.
                    </CardDescription>
                </CardHeader>
                <CardContent>
                    <Form {...form}>
                        <form
                            onSubmit={form.handleSubmit(onSubmit)}
                            className="space-y-8"
                        >
                            <FormField
                                control={form.control}
                                name="username"
                                render={({ field }) => (
                                    <FormItem>
                                        <FormLabel>ФИО:</FormLabel>
                                        <FormControl>
                                            <Input
                                                placeholder="Введи ФИО:"
                                                {...field}
                                            />
                                        </FormControl>
                                        <FormDescription>
                                            Ваше ФИО
                                        </FormDescription>
                                        <FormMessage />
                                    </FormItem>
                                )}
                            />

                            <FormField
                                control={form.control}
                                name="email"
                                render={({ field }) => (
                                    <FormItem>
                                        <FormLabel>Почта:</FormLabel>
                                        <FormControl>
                                            <Input
                                                placeholder="Введи Почту:"
                                                {...field}
                                            />
                                        </FormControl>
                                        <FormDescription>
                                            Ваша Почта
                                        </FormDescription>
                                        <FormMessage />
                                    </FormItem>
                                )}
                            />

                            <FormField
                                control={form.control}
                                name="password"
                                render={({ field }) => (
                                    <FormItem>
                                        <FormLabel>Пароль:</FormLabel>
                                        <FormControl>
                                            <Input
                                                placeholder="Введи пароль:"
                                                type="password"
                                                {...field}
                                            />
                                        </FormControl>
                                        <FormDescription>
                                            Ваш Пароль
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
                                    ? "Регистрация..."
                                    : "Зарегестрироваться"}
                            </Button>
                        </form>
                    </Form>
                </CardContent>
            </Card>
        </div>
    );
}

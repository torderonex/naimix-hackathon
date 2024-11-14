import { toast } from "sonner";
import { useForm, useFieldArray } from "react-hook-form";
import { zodResolver } from "@hookform/resolvers/zod";
import * as z from "zod";
import { Button } from "@/components/ui/button";
import { Form } from "@/components/ui/form";
import PersonInfo from "@/components/PersonInfo";

const formSchema = z.object({
    people: z.array(
        z.object({
            name: z.string().trim().min(1, { message: "Имя обязательно" }),
            birthDate: z.coerce.date(),
        })
    ),
});

export default function CompatibilityForm() {
    const form = useForm<z.infer<typeof formSchema>>({
        resolver: zodResolver(formSchema),
        defaultValues: {
            people: [
                { name: "", birthDate: new Date() },
                { name: "", birthDate: new Date() },
            ],
        },
    });

    const { fields, append, remove } = useFieldArray({
        control: form.control,
        name: "people",
    });

    function onSubmit(values: z.infer<typeof formSchema>) {
        try {
            console.log(values);
            toast(
                <pre className="mt-2 w-[340px] rounded-md bg-slate-950 p-4">
                    <code className="text-white">
                        {JSON.stringify(values, null, 2)}
                    </code>
                </pre>
            );
        } catch (error) {
            console.error("Form submission error", error);
            toast.error("Failed to submit the form. Please try again.");
        }
    }

    return (
        <Form {...form}>
            <form
                onSubmit={form.handleSubmit(onSubmit)}
                className="space-y-8 w-full sm:w-3xl max-w-full py-10 mx-auto"
            >
                <div className="overflow-x-auto py-4">
                    <div className="flex space-x-8 justify-center min-w-max">
                        {fields.map((item, index) => (
                            <div
                                key={item.id}
                                className="min-w-[300px] flex-shrink-0"
                            >
                                <PersonInfo
                                    index={index}
                                    remove={remove}
                                    field={form.control}
                                />
                            </div>
                        ))}
                    </div>
                </div>

                <div className="flex space-x-4 justify-center">
                    <Button
                        type="button"
                        variant={"outline"}
                        onClick={() =>
                            append({ name: "", birthDate: new Date() })
                        }
                    >
                        Добавить человека
                    </Button>
                    <Button type="submit">Проверить совместимость</Button>
                </div>
            </form>
        </Form>
    );
}

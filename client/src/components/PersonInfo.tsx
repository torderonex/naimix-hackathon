import {
    FormField,
    FormItem,
    FormLabel,
    FormControl,
    FormDescription,
    FormMessage,
} from "@/components/ui/form";
import { Input } from "@/components/ui/input";
import { DatetimePicker } from "@/components/ui/datetime-picker";
import { Button } from "@/components/ui/button";

interface PersonInfoProps {
    index: number;
    remove: (index: number) => void;
    field: any;
}

export default function PersonInfo({ index, remove, field }: PersonInfoProps) {
    return (
        <div className="space-y-8">
            <FormField
                control={field}
                name={`people[${index}].name`}
                render={({ field }) => (
                    <FormItem>
                        <FormLabel>Имя:</FormLabel>
                        <FormControl>
                            <Input placeholder="Введите имя:" {...field} />
                        </FormControl>
                        <FormDescription>
                            Имя человека, совместимость которого вы хотите
                            проверить.
                        </FormDescription>
                        <FormMessage />
                    </FormItem>
                )}
            />

            <FormField
                control={field}
                name={`people[${index}].birthDate`}
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
                        <FormDescription>ДД.ММ.ГГГГ | ЧЧ.ММ.CC</FormDescription>
                        <FormMessage />
                    </FormItem>
                )}
            />

            <Button
                type="button"
                onClick={() => remove(index)}
                variant={"link"}
                disabled={index <= 1}
            >
                Удалить человека
            </Button>
        </div>
    );
}

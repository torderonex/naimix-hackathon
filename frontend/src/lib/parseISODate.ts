export function parseISODate(dateString: string): Date {
    const date = new Date(dateString);
    console.log(date);

    // Проверяем, является ли дата валидной
    if (isNaN(date.getTime())) {
        throw new Error("Invalid date string");
    }

    return date; // Возвращаем объект Date
}

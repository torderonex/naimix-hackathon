import { useState } from "react";
import { useLocation } from "react-router-dom";
import { useParticipants } from "@/hooks/useParticipants";
import ParticipantsTable from "@/components/participants/ParticipantsTable";
import { Button } from "@/components/ui/button";
import ComparisonService from "@/services/comparison-service";
import ReactMarkdown from "react-markdown";
import { toast } from "sonner";
import { CircularProgressbar, buildStyles } from "react-circular-progressbar";
import "react-circular-progressbar/dist/styles.css";

const ParticipantsPage = () => {
    const location = useLocation();
    const id = location.pathname.split("/").pop();

    const { participants, loading, getAllTParticipants } = useParticipants(
        parseInt(id || "0")
    );

    const [birthDate, setBirthDate] = useState<string>("");
    const [compatibility, setCompatibility] = useState<string | null>(null);
    const [description, setDescription] = useState<string | null>(null);
    const [average, setAverage] = useState<number | null>(null);
    const [text, setText] = useState<string>("");
    const [percents, setPercents] = useState(null);

    const handleCalculateCompatibility = async () => {
        if (!birthDate) {
            toast.error("Пожалуйста, выбери дату рождения.");
            return;
        }

        if (!participants || participants.length < 1) {
            toast.error("В команде никого нет.");
            return;
        }

        try {
            toast("Идет рассчет...", {
                duration: 1000,
            });
            const formattedBirthDate = new Date(birthDate).toISOString();

            const resp = await ComparisonService.compareLot(
                participants,
                formattedBirthDate
            );

            const ps = resp.data;
            setPercents(ps.percents);
            const average = ps.avg;
            setAverage(average); // Set average state
            const description = ps.description;

            setCompatibility(`Общая совместимость с командой: ${average}%`);
            setDescription(description);
            toast("Рассчет получен.", {
                duration: 1000,
            });
        } catch (error) {
            console.error("Ошибка при расчете совместимости:", error);
            toast.error("Ошибка при расчете совместимости. Попробуй снова.");
        }
    };

    return (
        <div className="flex flex-col items-center px-16 py-8">
            <h2 className="text-2xl mb-6">
                Рассчитать совместимость нового кандидата
            </h2>

            <div className="mb-8 flex items-center space-x-4">
                <label htmlFor="birthDate" className="text-lg">
                    Дата рождения кандидата:
                </label>
                <input
                    id="birthDate"
                    type="date"
                    value={birthDate}
                    onChange={(e) => setBirthDate(e.target.value)}
                    className="border border-gray-300 rounded p-2"
                />
                <Button
                    onClick={handleCalculateCompatibility}
                    className="text-white px-4 py-2"
                >
                    Рассчитать совместимость всей команды
                </Button>
            </div>

            {compatibility && (
                <div className="mb-6 text-lg font-semibold text-customgreen">
                    {compatibility}
                </div>
            )}

            {average !== null && (
                <div className="mb-6 w-40 h-40">
                    <CircularProgressbar
                        value={average}
                        text={`${average}%`}
                        styles={buildStyles({
                            textColor:
                                average <= 35
                                    ? "#FF0000"
                                    : average <= 70
                                    ? "#FFA500"
                                    : "#0e6666",
                            pathColor:
                                average <= 35
                                    ? "#FF0000"
                                    : average <= 70
                                    ? "#FFA500"
                                    : "#0e6666",
                            trailColor: "#d6d6d6",
                        })}
                    />
                </div>
            )}

            {description && (
                <div className="mb-6 font-semibol max-w-5xl">{description}</div>
            )}

            <ParticipantsTable
                participants={participants}
                loading={loading}
                refetchParticipants={getAllTParticipants}
                comparisonPercents={percents}
                setText={setText}
                birthDate={birthDate}
            />
            {text !== "" && (
                <>
                    <h2 className="text-center my-8 text-2xl mt-10">
                        Проверка с выбранным участником команды
                    </h2>
                    <ReactMarkdown>{text}</ReactMarkdown>
                </>
            )}
        </div>
    );
};

export default ParticipantsPage;

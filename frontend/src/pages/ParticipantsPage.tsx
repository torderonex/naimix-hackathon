import React, { useState } from "react";
import { useLocation } from "react-router-dom";
import { useParticipants } from "@/hooks/useParticipants";
import ParticipantsTable from "@/components/participants/ParticipantsTable";
import { Button } from "@/components/ui/button";
import ComparisonService from "@/services/comparison-service";
import ReactMarkdown from 'react-markdown';
const ParticipantsPage = () => {
    const location = useLocation();
    const id = location.pathname.split("/").pop(); 
    
    const { participants, loading, getAllTParticipants } = useParticipants(parseInt(id || "0")); 
    
    const [birthDate, setBirthDate] = useState<string>(""); // Состояние для даты рождения
    const [compatibility, setCompatibility] = useState<string | null>(null); // Результат совместимости

    const [ text, setText] = useState<string> (""); 

    const [percents, setPercents] = useState(null);
    const handleCalculateCompatibility = async () => {
        if (!birthDate) {
            setCompatibility("Пожалуйста, выберите дату рождения.");
            return;
        }
    
        if (!participants || participants.length < 1) {
            setCompatibility("В команде никого нет.");
            return;
        }
    
        try {
            const formattedBirthDate = new Date(birthDate).toISOString(); // Преобразует в формат RFC3339 (подходит для Go)

            const resp = await ComparisonService.compareLot(participants, formattedBirthDate);
            // Извлекаем проценты из ответа
            const ps = resp.data
            setPercents(ps)
            // Считаем среднее значение
            const average = ps.reduce((acc: any, val: any) => acc + val) / ps.length;
            
            setCompatibility(`Общая совместимость с командой: ${average}%`);
        } catch (error) {
            console.error("Ошибка при расчете совместимости:", error);
            setCompatibility("Ошибка при расчете совместимости. Попробуйте снова.");
        }
    };
    return (
        <div className="flex flex-col items-center px-16 py-8">
            <h2 className="text-2xl mb-6">Участники команды</h2>
            
            <div className="mb-8 flex items-center space-x-4">
                <label htmlFor="birthDate" className="text-lg">Дата рождения:</label>
                <input
                    id="birthDate"
                    type="date"
                    value={birthDate}
                    onChange={(e) => setBirthDate(e.target.value)}
                    className="border border-gray-300 rounded p-2"
                />
                <Button
                    onClick={handleCalculateCompatibility}
                    className="text-white px-4 py-2  "
                >
                    Рассчитать совместимость
                </Button>
            </div>

            {compatibility && (
                <div className="mb-6 text-lg font-semibold text-green-600">
                    {compatibility}
                </div>
            )}

            <ParticipantsTable
                participants={participants}
                loading={loading}
                refetchParticipants={getAllTParticipants}
                comparisonPercents={percents}
                setText = {setText}
                birthDate={birthDate}
            />
            <ReactMarkdown>
            {text}
            </ReactMarkdown>
            </div>
    );
};

export default ParticipantsPage;

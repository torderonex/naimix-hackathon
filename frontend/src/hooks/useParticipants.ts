import { useState, useEffect } from "react";
import { toast } from "sonner";
import { Participant } from "@/types/participant";
import ParticipantService from "@/services/participant-service";

export function useParticipants(id: number) {
    const [participants, setParticipants] = useState<Participant[] | null>(
        null
    );
    const [loading, setLoading] = useState<boolean>(true);

    async function getAllTParticipants() {
        try {
            const response = await ParticipantService.getAllById(id);
            if (response.data === null) {
                setParticipants([]);
            } else {
                setParticipants(response.data);
            }
        } catch (err) {
            console.error("getAllTeams error", err);
            toast.error("Ошибка при получении всех команд");
            setParticipants([]);
        } finally {
            setLoading(false);
        }
    }

    useEffect(() => {
        getAllTParticipants();
    }, []);

    return { participants, loading, getAllTParticipants };
}

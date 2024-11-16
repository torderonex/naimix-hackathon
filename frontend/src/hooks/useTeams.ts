import { useState, useEffect } from "react";
import TeamService from "@/services/team-service";
import { toast } from "sonner";
import { Team } from "@/types/team";

export function useTeams(id : number) {
    const [teams, setTeams] = useState<Team[] | null>(null);
    const [loading, setLoading] = useState<boolean>(true);

    async function getAllTeams() {
        try {
            const response = await TeamService.getAllById(id);
            if (response.data === null) {
                setTeams([]);
            } else {
                setTeams(response.data);
            }
        } catch (err) {
            console.error("getAllTeams error", err);
            toast.error("Ошибка при получении всех команд");
            setTeams([]);
        } finally {
            setLoading(false);
        }
    }

    useEffect(() => {
        getAllTeams();
    }, []);

    return { teams, loading, getAllTeams };
}

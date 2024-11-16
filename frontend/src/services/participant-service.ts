import $api from "@/https/api";
import { Participant, ParticipantRequest } from "@/types/participant";
import { AxiosResponse } from "axios";

export default class ParticipantService {
    static async add(
        credentials: ParticipantRequest
    ): Promise<AxiosResponse<void>> {
        return $api.post("participant/add", credentials);
    }

    static async delete(id: number): Promise<AxiosResponse<void>> {
        return $api.delete(`participant/${id}`);
    }

    static async edit(
        credentials: ParticipantRequest,
        id: number
    ): Promise<AxiosResponse<void>> {
        return $api.put(`participant/${id}`, credentials);
    }

    static async getAll(): Promise<AxiosResponse<Participant[]>> {
        return $api.get("participant/all");
    }
}

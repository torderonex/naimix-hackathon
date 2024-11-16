import $api from "@/https/api";
import { TeamRequest, Team } from "@/types/team";
import { AxiosResponse } from "axios";

export default class TeamService {
    static async add(credentials: TeamRequest): Promise<AxiosResponse<void>> {
        return $api.post("team/add", credentials);
    }

    static async delete(id: number): Promise<AxiosResponse<void>> {
        return $api.delete(`team/${id}`);
    }

    static async edit(
        credentials: TeamRequest,
        id: number
    ): Promise<AxiosResponse<void>> {
        return $api.put(`team/${id}`, credentials);
    }

    static async getAllById(id : number): Promise<AxiosResponse<Team[]>> {
        return $api.get("team/user/" + id);
    }

    // static async getById(id: number): Promise<AxiosResponse<Team>> {
    //     return $api.get(`team/${id}`);
    // }
}

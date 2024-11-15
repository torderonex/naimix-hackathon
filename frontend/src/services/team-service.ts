import $api from "@/https/api";
import { AddTeamRequest } from "@/types/team";
import { AxiosResponse } from "axios";

export default class TeamService {
    static async add(credentials: AddTeamRequest): Promise<AxiosResponse<any>> {
        return $api.post("team/add", credentials);
    }
}

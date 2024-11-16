import $api from "@/https/api";
import { Participant } from "@/types/participant";
import { AxiosResponse } from "axios";

export default class ComparisonService {
    static async compare2(
        date1: string,
        date2: string
    ): Promise<AxiosResponse<any>> {
        return $api.post("compare/two", {
            date1,
            date2,
        });
    }

    static async compareLot(
        p: Participant[],
        date: string
    ): Promise<AxiosResponse<any>> {
        return $api.post("compare/lot", {
            patrticipants: p,
            birth_date: date,
        } as any);
    }
}

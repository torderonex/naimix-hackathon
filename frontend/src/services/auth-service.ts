import $api from "@/https/api";
import { AuthResponse, LoginRequest, RegisterRequest } from "@/types/auth";
import { User } from "@/types/user";
import { AxiosResponse } from "axios";

export default class AuthService {
    static async login(
        formData: LoginRequest
    ): Promise<AxiosResponse<AuthResponse>> {
        return $api.post("auth/login", formData);
    }

    static async register(
        formData: RegisterRequest
    ): Promise<AxiosResponse<AuthResponse>> {
        return $api.post("auth/signup", formData);
    }

    static async identity(): Promise<AxiosResponse<User>> {
        return $api.get("auth/identity");
    }
}

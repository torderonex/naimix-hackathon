export type ParticipantRequest = {
    name: string;
    team_id: number;
    role: string;
    birthdate: Date;
    birthplace: string;
    user_id: number;
};

export type Participant = {
    id: number;
    name: string;
    team_id: number;
    team_name: string;
    role: string;
    birthdate: string;
    birthplace: string;
    user_id: number;
};

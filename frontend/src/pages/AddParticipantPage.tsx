import AddParticipantForm from "@/components/participants/AddParticipantForm";
import ParticipantsTable from "@/components/participants/ParticipantsTable";
import { useParticipants } from "@/hooks/useParticipants";

export default function AddParticipantPage() {
    const { participants, loading, getAllTParticipants } = useParticipants();
    return (
        <div className="flex flex-col items-center px-16 py-8">
            <AddParticipantForm refetchParticipants={getAllTParticipants} />
            <ParticipantsTable
                participants={participants}
                loading={loading}
                refetchParticipants={getAllTParticipants}
            />
        </div>
    );
}

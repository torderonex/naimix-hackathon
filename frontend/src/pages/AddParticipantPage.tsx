import AddParticipantForm from "@/components/participants/AddParticipantForm";
import ParticipantsTable from "@/components/participants/ParticipantsTable";
import { useParticipantsCreator } from "@/hooks/useParticipantsByCreator";
import useUserStore from "@/store/user-store";

export default function AddParticipantPage() {
    const { user } = useUserStore();
    const { participants, loading, getAllTParticipants } =
        useParticipantsCreator(user!.id);
    return (
        <div className="flex flex-col items-center px-16 py-8">
            <AddParticipantForm refetchParticipants={getAllTParticipants} />
            <ParticipantsTable
                participants={participants}
                loading={loading}
                refetchParticipants={getAllTParticipants}
                comparisonPercents={null}
                setText={() => {}}
                birthDate=""
            />
        </div>
    );
}

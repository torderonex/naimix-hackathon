import AddParticipantForm from "@/components/participants/AddParticipantForm";
import ParticipantsTable from "@/components/participants/ParticipantsTable";

export default function AddParticipantPage() {
    return (
        <div className="flex flex-col items-center px-16 py-8">
            <AddParticipantForm />
            <ParticipantsTable />
        </div>
    );
}

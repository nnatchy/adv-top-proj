const ResultCard = ({
  studentId,
  candidateName,
  position,
  transactionId,
  isValid,
}) => {
  return (
    <div className="bg-white shadow-lg rounded-lg p-6 flex flex-col gap-4 w-[400px] max-w-full">
      <div className="text-lg font-semibold text-gray-800 break-words">
        Student ID: {studentId}
      </div>
      <div className="text-gray-600 text-lg">
        Voted For:{" "}
        <span className="font-semibold text-black">{candidateName}</span>
      </div>
      <div className="text-gray-600 text-lg">
        Position: <span className="font-semibold text-black">{position}</span>
      </div>
      <div className="text-gray-600 text-lg break-words">
        Transaction ID:{" "}
        <span className="font-mono text-blue-600">{transactionId}</span>
      </div>
      <div
        className={`text-lg font-bold ${
          isValid ? "text-green-600" : "text-red-600"
        }`}
      >
        {isValid ? "Valid Vote ✅" : "Invalid Vote ❌"}
      </div>
    </div>
  );
};

export default ResultCard;

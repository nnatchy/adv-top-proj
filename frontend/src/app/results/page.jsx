"use client";
import React, { useEffect, useState } from "react";
import { useRouter } from "next/navigation";
import { useDispatch, useSelector } from "react-redux";
import { Button } from "@mui/material";
import { motion } from "framer-motion";
import ResultCard from "../../components/common/ResultCard";
import { fetchVotingResultsApi } from "../../libs/votingLib"; // Import API
import { useSession } from "../../app/SessionContext"; // Import useSession

const ResultsPage = () => {
  const router = useRouter();
  const { userInfo } = useSelector((state) => state.user); // Access Redux state
  const [results, setResults] = useState([]); // Local state for results
  const [loading, setLoading] = useState(true); // Loading state
  const [error, setError] = useState(null); // Error state
  const [popupStatus, setPopupStatus] = useState(null); // Popup status
  const [voteSummary, setVoteSummary] = useState({}); // Summary of votes per year
  const { session } = useSession();

  const positions = {
    year_one: "1st Year",
    year_two: "2nd Year",
    year_three: "3rd Year",
    year_four: "4th Year",
    representative: "Representative",
  };


  useEffect(() => {
    if (!userInfo?.student_id) return; // Wait until userInfo is loaded

    const hasVotedYear = userInfo.hasVotedYear;
    const hasVotedRepresentative = userInfo.hasVotedRepresentative;

    if (!hasVotedYear || !hasVotedRepresentative) {
      setPopupStatus(401); // Unauthorized status
      setTimeout(() => {
        router.replace("/home"); // Redirect to home immediately
      }, 0); // Redirect immediately
    }
  }, [userInfo, router]);

  useEffect(() => {
    const fetchResults = async () => {
      setLoading(true);
      try {
        const fetchedResults = await fetchVotingResultsApi(); // Fetch results from API
        const formattedResults = formatResults(fetchedResults);
        setResults(formattedResults);
        calculateVoteSummary(formattedResults);
      } catch (err) {
        setError("Failed to fetch voting results. Please try again later.");
      } finally {
        setLoading(false);
      }
    };

    fetchResults();
  }, []);

  const formatResults = (results) => {
    return results.map((result) => ({
      studentId: result.student_id,
      candidateName: `${result.candidate_first_name} ${result.candidate_last_name}`,
      position: positions[result.candidate_position] || result.candidate_position.replace("_", " "), // Make position more readable
      transactionId: result.transaction_id,
      isValid: result.is_valid,
    }));
  };

  const calculateVoteSummary = (results) => {
    const summary = {};
    results.forEach((result) => {
      const position = result.position;
      if (!summary[position]) {
        summary[position] = { total: 0, valid: 0, invalid: 0 };
      }
      summary[position].total += 1;
      summary[position][result.isValid ? "valid" : "invalid"] += 1;
    });
    setVoteSummary(summary);
  };

  const handleBackToHome = () => {
    router.push("/home");
  };

  const handleCandidateResults = () => {
    router.push("/results/candidates");
  };

  if (popupStatus) {
    return (
      <div className="flex items-center justify-center h-screen bg-gray-100">
        <div className="text-center">
          <h1 className="text-red-500 text-3xl font-bold">
            Unauthorized Access!
          </h1>
          <p className="text-gray-700 mt-4">
            You need to vote for both your year and representative before
            accessing the results.
          </p>
        </div>
      </div>
    );
  }

  if (loading) {
    return <div className="text-center mt-8">Loading results...</div>;
  }

  if (error) {
    return <div className="text-center mt-8 text-red-500">{error}</div>;
  }

  return (
    <motion.div
      initial={{ opacity: 0 }}
      animate={{ opacity: 1 }}
      exit={{ opacity: 0 }}
      transition={{ duration: 0.5 }}
      className="px-6 py-8"
    >
      {/* Page Title */}
      <motion.div
        className="flex items-center justify-center"
        initial={{ y: -50, opacity: 0 }}
        animate={{ y: 0, opacity: 1 }}
        transition={{ duration: 0.8, delay: 0.1 }}
      >
        <h1 className="text-white text-4xl font-bold mb-6">Voting Results</h1>
      </motion.div>

      {/* Results Grid */}
      <motion.div
        className="grid grid-cols-1 sm:grid-cols-2 md:grid-cols-3 lg:grid-cols-4 gap-6"
        initial={{ opacity: 0 }}
        animate={{ opacity: 1 }}
        transition={{ duration: 0.8, delay: 0.3 }}
      >
        {results.map((result, index) => (
          <motion.div
            key={index}
            initial={{ scale: 0.9, opacity: 0 }}
            animate={{ scale: 1, opacity: 1 }}
            transition={{ duration: 0.5, delay: index * 0.1 }}
          >
            <ResultCard
              studentId={result.studentId}
              candidateName={result.candidateName}
              position={result.position}
              transactionId={result.transactionId}
              isValid={result.isValid}
            />
          </motion.div>
        ))}
      </motion.div>

      {/* Summary */}
      <motion.div
        className="bg-gray-800 text-white p-4 mt-8 rounded-lg"
        initial={{ opacity: 0, y: 50 }}
        animate={{ opacity: 1, y: 0 }}
        transition={{ duration: 0.8, delay: 0.6 }}
      >
        <h2 className="text-2xl font-bold text-center mb-4">
          Vote Summary Analysis
        </h2>
        <div className="grid grid-cols-1 md:grid-cols-3 lg:grid-cols-5 gap-4">
          {Object.entries(voteSummary).map(([position, stats]) => (
            <div
              key={position}
              className="bg-gray-700 p-4 rounded-md shadow-md text-center"
            >
              <h3 className="text-xl font-semibold capitalize">
                {position.replace("_", " ")}
              </h3>
              <p className="mt-2">
                <span className="font-bold">Total Votes:</span> {stats.total}
              </p>
              <p className="text-green-400 mt-1">
                <span className="font-bold">Valid Votes:</span> {stats.valid}
              </p>
              <p className="text-red-400 mt-1">
                <span className="font-bold">Invalid Votes:</span> {stats.invalid}
              </p>
            </div>
          ))}
        </div>
      </motion.div>

      {/* Buttons */}
      <motion.div
        className="mt-12 flex flex-col justify-center items-center gap-4"
        initial={{ opacity: 0 }}
        animate={{ opacity: 1 }}
        transition={{ duration: 0.8, delay: 0.8 }}
      >
        {/* Candidate Results Button */}
        <Button
          type="button"
          variant="contained"
          sx={{
            backgroundColor: "#007BFF",
            color: "#FFFFFF",
            fontSize: "20px",
            padding: "10px 20px",
            width: "18vw",
            height: "7vh",
            borderRadius: "8px",
            boxShadow: "none",
            textTransform: "none",
          }}
          onClick={handleCandidateResults}
        >
          View Candidate Results
        </Button>

        {/* Back to Home Button */}
        <Button
          type="button"
          variant="contained"
          sx={{
            backgroundColor: "#FF98A0",
            color: "#FFFFFF",
            fontSize: "20px",
            padding: "10px 20px",
            width: "18vw",
            height: "7vh",
            borderRadius: "8px",
            boxShadow: "none",
            textTransform: "none",
          }}
          onClick={handleBackToHome}
        >
          Back to Home
        </Button>
      </motion.div>
    </motion.div>
  );
};

export default ResultsPage;

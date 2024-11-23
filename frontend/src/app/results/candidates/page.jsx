"use client";
import React, { useEffect, useState } from "react";
import { useRouter } from "next/navigation";
import { Button } from "@mui/material";
import { motion } from "framer-motion";
import { fetchGetCandidatesRankingApi } from "../../../libs/votingLib"; // Import the API function

const CandidateResultsPage = () => {
  const router = useRouter();
  const [results, setResults] = useState({});
  const [loading, setLoading] = useState(true);
  const [error, setError] = useState(null);

  // Function to group and sort candidates by position
  const groupAndSortCandidates = (data) => {
    const groupedResults = {};
    data.forEach((item) => {
      const position = item.position.replace("_", " "); // Normalize position names
      if (!groupedResults[position]) {
        groupedResults[position] = [];
      }
      groupedResults[position].push({
        name: `${item.first_name} ${item.last_name}`,
        votes: item.votes,
      });
    });

    // Sort each group by descending vote count
    Object.keys(groupedResults).forEach((position) => {
      groupedResults[position].sort((a, b) => b.votes - a.votes);
    });

    return groupedResults;
  };

  useEffect(() => {
    const fetchCandidateResults = async () => {
      setLoading(true);
      try {
        const data = await fetchGetCandidatesRankingApi(); // Use the imported API function
        const sortedResults = groupAndSortCandidates(data);
        setResults(sortedResults);
      } catch (err) {
        setError("Failed to fetch candidate results. Please try again later.");
      } finally {
        setLoading(false);
      }
    };

    fetchCandidateResults();
  }, []);

  if (loading) {
    return <div className="text-center mt-8">Loading candidate results...</div>;
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
        <h1 className="text-white text-4xl font-bold mb-6">
          Candidate Results
        </h1>
      </motion.div>

      {/* Results Grid */}
      <motion.div
        className="grid grid-cols-1 sm:grid-cols-2 md:grid-cols-3 lg:grid-cols-5 gap-6"
        initial={{ opacity: 0 }}
        animate={{ opacity: 1 }}
        transition={{ duration: 0.8, delay: 0.3 }}
      >
        {Object.entries(results).map(([year, candidates], index) => (
          <motion.div
            key={index}
            className="bg-white p-4 rounded-lg shadow-lg"
            initial={{ scale: 0.9, opacity: 0 }}
            animate={{ scale: 1, opacity: 1 }}
            transition={{ duration: 0.5, delay: index * 0.1 }}
          >
            <h2 className="text-center font-bold text-lg mb-4 capitalize text-black">
              {year}
            </h2>
            <ul>
              {candidates.map((candidate, idx) => (
                <li
                  key={idx}
                  className={`flex justify-between text-gray-700 mb-2 ${
                    idx === 0
                      ? "text-yellow-500 font-bold"
                      : idx === 1
                      ? "text-gray-500 font-bold"
                      : idx === 2
                      ? "text-orange-500 font-bold"
                      : ""
                  }`}
                >
                  <span>
                    {idx + 1 === 1
                      ? "🥇"
                      : idx + 1 === 2
                      ? "🥈"
                      : idx + 1 === 3
                      ? "🥉"
                      : `${idx + 1}th`}{" "}
                    {candidate.name}
                  </span>
                  <span>{candidate.votes} votes</span>
                </li>
              ))}
            </ul>
          </motion.div>
        ))}
      </motion.div>

      {/* Buttons */}
      <motion.div
        className="mt-12 flex flex-col justify-center items-center gap-4"
        initial={{ opacity: 0 }}
        animate={{ opacity: 1 }}
        transition={{ duration: 0.8, delay: 0.6 }}
      >
        {/* Go Back to Results Button */}
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
          onClick={() => router.push("/results")}
        >
          Back to Results
        </Button>

        {/* Go to Home Button */}
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
          onClick={() => router.push("/home")}
        >
          Go to Home
        </Button>
      </motion.div>
    </motion.div>
  );
};

export default CandidateResultsPage;

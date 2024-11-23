"use client";

import React, { useEffect, useState } from "react";
import { useRouter } from "next/navigation";
import CandidateCard from "../../../components/common/CandidateCard";
import Colors from "../../../app/assets/styles/colors";
import { createCandidate, deleteCandidate } from "../../../libs/candidateLib";

export default function AdminHome() {
  const router = useRouter();

  const [candidates, setCandidates] = useState([]);
  const [selectedPosition, setSelectedPosition] = useState("year_one");
  const [loading, setLoading] = useState(false);
  const [error, setError] = useState(null);
  const [isCreateModalOpen, setCreateModalOpen] = useState(false);
  const [formData, setFormData] = useState({
    firstName: "",
    lastName: "",
    studentID: "",
    major: "CP",
  });

  const positions = [
    "year_one",
    "year_two",
    "year_three",
    "year_four",
    "representative",
  ];

  const majors = ["CP", "IE", "CE", "ME", "PE", "EE"];

  const fetchCandidates = async () => {
    setLoading(true);
    setError(null);

    try {
      const response = await fetch(
        `http://localhost:8010/api/user/v1/candidates?position=${selectedPosition}`
      );
      const data = await response.json();
      setCandidates(data);
    } catch (err) {
      setError("Failed to fetch candidates");
    } finally {
      setLoading(false);
    }
  };

  useEffect(() => {
    fetchCandidates();
  }, [selectedPosition]);

  const handleCreateCandidate = async () => {
    try {
      await createCandidate({
        FirstName: formData.firstName,
        LastName: formData.lastName,
        StudentID: formData.studentID,
        position: selectedPosition,
        major: formData.major,
      });
      fetchCandidates(); // Refresh candidates
      setCreateModalOpen(false);
      setFormData({ firstName: "", lastName: "", studentID: "", major: "CP" });
    } catch (error) {
      console.error("Failed to create candidate:", error);
    }
  };

  const handleDeleteCandidate = async (id) => {
    try {
      await deleteCandidate(id);
      fetchCandidates(); // Refresh candidates
    } catch (error) {
      console.error("Failed to delete candidate:", error);
    }
  };

  return (
    <div
      style={{
        backgroundColor: Colors.COLOR_PINK,
        minHeight: "100vh",
        color: Colors.COLOR_WHITE,
        padding: "20px",
        overflowY: "auto", // Enable vertical scrolling
      }}
    >
      <div className="absolute top-4 right-4">
        <button
          onClick={() => router.push("/")}
          className="px-4 py-2 bg-white rounded-md  transition"
          style={{
            color: Colors.COLOR_PINK,
          }}
        >
          Logout
        </button>
      </div>

      <h1 className="text-3xl font-bold mb-4">Admin Back Office</h1>

      {/* Dropdown for selecting position */}
      <div className="flex items-center mb-6">
        <select
          value={selectedPosition}
          onChange={(e) => setSelectedPosition(e.target.value)}
          className="px-4 py-2 bg-white text-black rounded-md mr-12"
        >
          {positions.map((position) => (
            <option key={position} value={position}>
              {position.replace("_", " ").toUpperCase()}
            </option>
          ))}
        </select>
        <button
          onClick={() => setCreateModalOpen(true)}
          className="px-4 py-2 bg-blue-500 text-white rounded-md transition"
        >
          Create Candidate
        </button>
      </div>

      {/* Candidate Cards */}
      <div className="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-4">
        {loading ? (
          <p>Loading...</p>
        ) : error ? (
          <p className="text-red-500">{error}</p>
        ) : candidates.length > 0 ? (
          candidates.map((candidate) => (
            <CandidateCard
              key={candidate.ID}
              id={candidate.ID}
              imageUrl="/img/profilemock.svg" // Replace with actual imageUrl if available
              firstName={candidate.FirstName}
              lastName={candidate.LastName}
              position={candidate.Position}
              major={candidate.Major}
              onDelete={() => handleDeleteCandidate(candidate.ID)}
            />
          ))
        ) : (
          <p>No candidates found for {selectedPosition.replace("_", " ")}</p>
        )}
      </div>

      {/* Create Candidate Modal */}
      {isCreateModalOpen && (
        <div className="fixed top-0 left-0 w-full h-full bg-black bg-opacity-50 flex items-center justify-center">
          <div className="bg-white p-6 rounded-lg text-black w-[350px]">
            <h2 className="text-lg font-bold mb-4">Create Candidate</h2>
            <input
              type="text"
              placeholder="First Name"
              value={formData.firstName}
              onChange={(e) =>
                setFormData({ ...formData, firstName: e.target.value })
              }
              className="block w-full p-2 mb-4 border rounded"
            />
            <input
              type="text"
              placeholder="Last Name"
              value={formData.lastName}
              onChange={(e) =>
                setFormData({ ...formData, lastName: e.target.value })
              }
              className="block w-full p-2 mb-4 border rounded"
            />
            <input
              type="text"
              placeholder="Student ID"
              value={formData.studentID}
              maxLength={10} // Restrict input to a maximum of 10 characters
              onChange={(e) =>
                setFormData({
                  ...formData,
                  studentID: e.target.value,
                })
              }
              className="block w-full p-2 mb-4 border rounded"
            />

            <select
              value={formData.major}
              onChange={(e) =>
                setFormData({ ...formData, major: e.target.value })
              }
              className="block w-full p-2 mb-4 border rounded"
            >
              {majors.map((major) => (
                <option key={major} value={major}>
                  {major}
                </option>
              ))}
            </select>
            <div className="flex justify-between pt-2">
              <button
                onClick={() => setCreateModalOpen(false)}
                className="px-4 py-2 bg-red-500 text-white rounded hover:bg-red-700 transition"
              >
                Cancel
              </button>
              <button
                onClick={handleCreateCandidate}
                className="px-4 py-2 bg-blue-500 text-white rounded hover:bg-blue-700 transition"
              >
                Create
              </button>
            </div>
          </div>
        </div>
      )}
    </div>
  );
}

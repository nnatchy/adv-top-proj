"use client";

import React, { useEffect, useState } from "react";
import { useRouter, useSearchParams } from "next/navigation";
import CandidateCardLarge from "../../../components/common/CandidateCardLarge";
import CandidateCardSkeleton from "../../../components/common/CandidateCardSkeleton";
import { searchCandidates } from "../../../libs/candidateLib";
import { useDispatch, useSelector } from "react-redux";
import { updateUserStatusThunk } from "../../../reducers/userReducer"; // Redux Thunk
import { Button } from "@mui/material";
import VotePopup from "../../../components/VotePopup";
import StatusPopup from "../../../components/StatusPopup";
import { getYearFromStudentId } from "../../../utils/getYear";
import { castVoteApi } from "../../../libs/votingLib"; // Import castVoteApi

export default function SearchPage() {
  const searchParams = useSearchParams();
  const position = searchParams.get("position") || "";
  const router = useRouter();
  const dispatch = useDispatch();
  const { userInfo } = useSelector((state) => state.user); // Redux state for user

  const [candidates, setCandidates] = useState([]);
  const [loading, setLoading] = useState(true);
  const [showContent, setShowContent] = useState(false);
  const [error, setError] = useState(null);
  const [isVotePopupOpen, setVotePopupOpen] = useState(false);
  const [isStatusPopupOpen, setStatusPopupOpen] = useState(false);
  const [popupStatus, setPopupStatus] = useState(null);
  const [selectedCandidate, setSelectedCandidate] = useState(null);
  const [isSubmitting, setIsSubmitting] = useState(false); // New state variable

  useEffect(() => {
    if (!userInfo?.student_id) return; // Wait until userInfo is loaded

    try {
      const studentYear = getYearFromStudentId(userInfo.student_id); // Get the user's year

      // Check if the user is accessing a valid position
      const validPosition =
        (!userInfo.hasVotedYear && studentYear === 1 && position === "year_one") ||
        (!userInfo.hasVotedYear && studentYear === 2 && position === "year_two") ||
        (!userInfo.hasVotedYear && studentYear === 3 && position === "year_three") ||
        (!userInfo.hasVotedYear && studentYear === 4 && position === "year_four") ||
        (!userInfo.hasVotedRepresentative && position === "representative");

      if (!validPosition) {
        console.warn(`Forbidden access to position: ${position}`);

        setPopupStatus(500);
        setTimeout(() => {
          router.replace("/home"); // Redirect to home after 0.5 seconds
        }, 500);
      }
    } catch (error) {
      console.error("Error validating user position:", error);
    }
  }, [userInfo, position, router]);

  useEffect(() => {
    if (!position) return;

    const fetchCandidates = async () => {
      setLoading(true);
      setShowContent(false);
      try {
        const response = await searchCandidates(position);
        setCandidates(response.data);
      } catch (error) {
        setError("Failed to fetch candidates");
      } finally {
        setLoading(false);
        setTimeout(() => {
          setShowContent(true);
        }, 500);
      }
    };

    fetchCandidates();
  }, [position]);

  const handleVoteClick = (candidate) => {
    if (!candidate) return;
    setSelectedCandidate(candidate);
    setVotePopupOpen(true);
  };

  const handleConfirmVote = async () => {
    if (isSubmitting) return; // Prevent double submission

    setIsSubmitting(true);
    setVotePopupOpen(false);

    try {
      // Ensure required data is available
      if (!selectedCandidate || !userInfo?.student_id) {
        throw new Error("Missing candidate or user information");
      }

      // Use the correct property name for candidate ID
      const candidateId = selectedCandidate.id; // Assuming 'id' is the correct property

      if (!candidateId) {
        throw new Error("Candidate ID is undefined");
      }

      // Prepare the vote payload
      const votePayload = {
        student_id: userInfo.student_id,
        candidate_id: candidateId,
      };

      console.log("Vote Payload:", votePayload); // For debugging

      // Call the API to cast the vote
      await castVoteApi(votePayload);

      // Update user status
      const updatePayload = {
        has_vote_year:
          userInfo.hasVotedYear || position !== "representative",
        has_vote_rep:
          userInfo.hasVotedRepresentative || position === "representative",
      };

      const walletAddress = userInfo.walletAddress;
      if (!walletAddress) {
        throw new Error("Wallet address is undefined");
      }

      await dispatch(
        updateUserStatusThunk({ walletAddress, updatePayload })
      );

      setPopupStatus(200); // Success
    } catch (error) {
      console.error("Failed to cast vote or update status:", error);
      setPopupStatus(500); // Error
    } finally {
      setIsSubmitting(false); // Reset the submitting state
      setStatusPopupOpen(true);
    }
  };

  const handleCloseStatusPopup = () => {
    setStatusPopupOpen(false);
    if (popupStatus === 200) {
      router.replace("/home"); // Redirect to home on success
    }
  };

  const handleBackToHome = () => {
    router.push("/home");
  };

  if (error) {
    return <div className="text-red-500 text-center mt-8">{error}</div>;
  }

  return (
    <div className="px-4 py-8">
      {/* Section Title */}
      <h1 className="text-black text-xl font-bold mb-6">Candidates Result</h1>

      {/* Candidate Cards */}
      <div className="space-y-4">
        {loading || !showContent
          ? Array.from({ length: 4 }).map((_, index) => (
              <CandidateCardSkeleton key={index} />
            ))
          : candidates.map((candidate) => {
              const positionMap = {
                year_one: "1st Year",
                year_two: "2nd Year",
                year_three: "3rd Year",
                year_four: "4th Year",
                representative: "Student Representative",
              };

              const readablePosition =
                positionMap[candidate.candidate_position] || "Unknown";

              const candidateId = candidate.id; // Assuming 'id' is the correct property

              return (
                <CandidateCardLarge
                  key={candidateId}
                  imageUrl={candidate.image_url || "/img/profilemock.svg"}
                  firstName={candidate.candidate_first_name}
                  lastName={candidate.candidate_last_name}
                  cp={readablePosition} // Assign the readable position
                  major={candidate.candidate_major || "Intania"} // Fallback value
                  onVote={() => handleVoteClick(candidate)} // Open confirmation popup
                />
              );
            })}
      </div>

      {/* Back to Home Button */}
      {!loading && showContent && (
        <div className="mt-12 flex justify-center">
          <Button
            type="submit"
            variant="contained"
            sx={{
              backgroundColor: "#FF98A0",
              color: "#FFFFFF",
              fontSize: "18px",
              padding: "12px 30px",
              borderRadius: "8px",
              boxShadow: "none",
              textTransform: "none",
            }}
            onClick={handleBackToHome} // Reset search placeholder and navigate to /home
          >
            Back to Home
          </Button>
        </div>
      )}

      <VotePopup
        isOpen={isVotePopupOpen}
        onClose={() => setVotePopupOpen(false)} // Close VotePopup
        candidateName={
          selectedCandidate
            ? `${selectedCandidate.candidate_first_name} ${selectedCandidate.candidate_last_name}`
            : "Unknown Candidate"
        }
        candidateImage={selectedCandidate?.image_url || "/img/profilemock.svg"}
        onConfirm={handleConfirmVote} // Trigger StatusPopup
        isSubmitting={isSubmitting} // Pass isSubmitting to disable the button if needed
      />

      {/* Status Popup */}
      <StatusPopup
        isOpen={isStatusPopupOpen}
        status={popupStatus}
        onClose={handleCloseStatusPopup}
      />
    </div>
  );
}

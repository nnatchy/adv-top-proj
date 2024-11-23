import axiosInstance from "../stores/axiosInstance";
import { VOTING_API_PATH } from "../constants";

// Fetch voting results from the backend
export const fetchVotingResultsApi = async () => {
  const userToken = localStorage.getItem("userToken");
  const response = await axiosInstance.get(`${VOTING_API_PATH}/results`, {
    headers: {
      "Content-Type": "application/json",
      Authorization: `Bearer ${userToken}`,
    },
  });
  return response.data; // Returns the array of voting results
};

// Fetch vote details for a specific student
export const fetchVoteDetailsApi = async (studentId) => {
  const userToken = localStorage.getItem("userToken");
  const response = await axiosInstance.get(`${VOTING_API_PATH}/details/${studentId}`, {
    headers: {
      "Content-Type": "application/json",
      Authorization: `Bearer ${userToken}`,
    },
  });
  return response.data; // Returns the vote details for the student
};

// Cast a vote
export const castVoteApi = async (votePayload) => {
  try {
    const userToken = localStorage.getItem("userToken");

    console.log("Casting Vote - Payload:", votePayload);

    const response = await axiosInstance.post(`${VOTING_API_PATH}/vote`, votePayload, {
      headers: {
        "Content-Type": "application/json",
        Authorization: `Bearer ${userToken}`,
      },
    });

    console.log("Casting Vote - Response:", response.data);

    return response.data; // Returns confirmation of the vote
  } catch (error) {
    console.error("Error casting vote:", error.response?.data || error.message);
    throw error;
  }
};


// End voting
export const endVotingApi = async () => {
  const userToken = localStorage.getItem("userToken");
  const response = await axiosInstance.post(`${VOTING_API_PATH}/end`, null, {
    headers: {
      "Content-Type": "application/json",
      Authorization: `Bearer ${userToken}`,
    },
  });
  return response.data; // Returns confirmation of voting end
};

// Fetch all votes from the contract
export const fetchAllVotesFromContractApi = async () => {
  const userToken = localStorage.getItem("userToken");
  const response = await axiosInstance.get(`${VOTING_API_PATH}/all-votes`, {
    headers: {
      "Content-Type": "application/json",
      Authorization: `Bearer ${userToken}`,
    },
  });
  return response.data; // Returns all votes from the contract
};

// Fetch all votes from the contract
export const fetchGetCandidatesRankingApi = async () => {
  const userToken = localStorage.getItem("userToken");
  const response = await axiosInstance.get(`${VOTING_API_PATH}/candidates-ranking`, {
    headers: {
      "Content-Type": "application/json",
      Authorization: `Bearer ${userToken}`,
    },
  });
  return response.data; // Returns all votes from the contract
};

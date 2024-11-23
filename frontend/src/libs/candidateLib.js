import axiosInstance from "../stores/axiosInstance";
import { CANDIDATE_API_PATH, USER_API_PATH } from "../constants";

export const searchCandidates = async (input) => {
  try {
    // Retrieve user token from localStorage
    // const userToken = localStorage.getItem("userToken");
    // if (!userToken) {
    //   throw new Error("User token not found in localStorage");
    // }

    // Make the API request
    const response = await axiosInstance.get(
      `${CANDIDATE_API_PATH}/search?position=${input}`,
      {
        headers: {
          "Content-Type": "application/json",
          // Authorization: `Bearer ${userToken}`,
        },
      }
    );

    // Return the response data
    return response;
  } catch (error) {
    console.error("Error in searchCandidates:", error);
    throw error;
  }
};

export const createCandidate = async (candidate) => {
  try {
    const response = await axiosInstance.post(
      `${USER_API_PATH}/candidates`,
      candidate,
      {
        headers: {
          "Content-Type": "application/json",
        },
      }
    );
    return response;
  } catch (error) {
    console.error("Error in createCandidate:", error);
    throw error;
  }
};

export const deleteCandidate = async (id) => {
  try {
    const response = await axiosInstance.delete(
      `${USER_API_PATH}/candidates/${id}`,
      {
        headers: {
          "Content-Type": "application/json",
        },
      }
    );
    return response;
  } catch (error) {
    console.error("Error in deleteCandidate:", error);
    throw error;
  }
};

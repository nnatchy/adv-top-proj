import axiosInstance from "../stores/axiosInstance";
import { USER_API_PATH } from "../constants";

// Fetch user information by wallet address
export const fetchUserInfo = async (wallet_address) => {
  const userToken = localStorage.getItem("userToken");
  const response = await axiosInstance.get(
    `${USER_API_PATH}/info?walletAddress=${wallet_address}`,
    {
      headers: {
        "Content-Type": "application/json",
        Authorization: `Bearer ${userToken}`,
      },
    }
  );
  return response.data; // Returns user information
};

// Update user status (has_vote_year, has_vote_rep)
export const updateUserStatus = async (wallet_address, updatePayload) => {
  const userToken = localStorage.getItem("userToken");
  const response = await axiosInstance.patch(
    `${USER_API_PATH}/status?walletAddress=${wallet_address}`,
    updatePayload, // Pass the request body here
    {
      headers: {
        "Content-Type": "application/json",
        Authorization: `Bearer ${userToken}`,
      },
    }
  );
  return response.data; // Returns confirmation message or updated data
};

// Admin Login
export const adminLogin = async (username, password) => {
  const response = await axiosInstance.post(
    `${USER_API_PATH}/admin/login`,
    { username, password }, // Send username and password in the request body
    {
      headers: {
        "Content-Type": "application/json",
      },
    }
  );
  return response.data; // Returns admin information or token
};

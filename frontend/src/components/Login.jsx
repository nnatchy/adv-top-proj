"use client";

import React, { useState, useEffect } from "react";
import { ethers } from "ethers";
import { useSession } from "./../app/SessionContext";
import Colors from "../app/assets/styles/colors";
import { USER_API_PATH } from "../constants";
import { useRouter } from "next/navigation"; // Add this import

const API_URL = process.env.NEXT_PUBLIC_API_URL;
const FIN_USER_API = API_URL + USER_API_PATH;

const Login = () => {
  const router = useRouter(); // Initialize router
  const [errorMessage, setErrorMessage] = useState(null);
  const [showForm, setShowForm] = useState(false);
  const [formData, setFormData] = useState({
    firstName: "",
    lastName: "",
    studentId: "",
    username: "",
  });

  const { session, updateSession, clearSession } = useSession();
  console.log(session);

  const checkActivationStatus = async (walletAddress) => {
    try {
      const response = await fetch(
        `${FIN_USER_API}/isActivate?walletAddress=${session.walletAddress}`
      );
      const data = await response.json();
      if (response.ok) {
        if (data.isActive && window.location.pathname !== "/home") {
          // router.push("/home"); // Redirect only if not already on the home page
        } else {
          setShowForm(true); // Show form if not active
          localStorage.setItem("showForm", true); // Persist state
        }
      } else {
        setErrorMessage(data.message || "Failed to check activation status");
      }
    } catch (error) {
      setErrorMessage("Error checking activation status");
    }
  };

  useEffect(() => {
    const storedShowForm = localStorage.getItem("showForm");
    if (storedShowForm === "true") {
      setShowForm(true);
    } else if (session.walletAddress) {
      checkActivationStatus(session.walletAddress);
    }
    // checkActivationStatus(session.walletAddress);
  }, [session.walletAddress]);

  const connectWallet = async () => {
    if (window.ethereum) {
      try {
        const provider = new ethers.providers.Web3Provider(window.ethereum);
        await provider.send("eth_requestAccounts", []);
        const signer = provider.getSigner();
        const address = await signer.getAddress();

        // Simulate sending address to backend for verification
        const response = await fetch(`${FIN_USER_API}/login`, {
          method: "POST",
          headers: {
            "Content-Type": "application/json",
          },
          body: JSON.stringify({ walletAddress: address }),
        });

        const data = await response.json();
        if (response.ok) {
          // Update session state
          updateSession({
            token: data.token,
            walletAddress: data.user.WalletAddress,
            id: data.user.ID,
            username: data.user.Username,
            role: data.user.Role,
          });

          if (!data.user.Username) {
            setShowForm(true); // Show the form if username is empty (inactive user)
            localStorage.setItem("showForm", "true"); // Persist showForm state
          }

          alert("Login successful");
        } else {
          setErrorMessage(data.message || "Login failed");
        }
      } catch (error) {
        setErrorMessage("Failed to connect wallet");
      }
    } else {
      setErrorMessage("MetaMask is not installed");
    }
  };

  const logout = () => {
    // Clear application session
    clearSession();
    setShowForm(false);
    localStorage.removeItem("showForm"); // Clear showForm state on logout

    alert("Logged out successfully. Please reconnect your wallet.");
  };

  const handleFormChange = (e) => {
    const { name, value } = e.target;
    setFormData((prevData) => ({
      ...prevData,
      [name]: value,
    }));
  };

  const handleFormSubmit = async (e) => {
    e.preventDefault();
    try {
      const response = await fetch(`${FIN_USER_API}/register`, {
        method: "POST",
        headers: {
          "Content-Type": "application/json",
          Authorization: `Bearer ${session.token}`, // Include token for authorization
        },
        body: JSON.stringify({
          walletAddress: session.walletAddress, // Get walletAddress from session
          username: formData.username,
          first_name: formData.firstName,
          last_name: formData.lastName,
          student_id: formData.studentId,
        }),
      });

      const data = await response.json();
      if (response.ok) {
        alert("User information updated successfully!");
        setShowForm(false); // Hide the form after successful submission
        localStorage.removeItem("showForm"); // Clear state from localStorage
      } else {
        setErrorMessage(data.message || "Failed to update user information");
      }
    } catch (error) {
      setErrorMessage("An error occurred while submitting the form");
    }
  };

  return (
    <div className="text-center">
      {session.token ? (
        <div>
          <p className="text-white mb-2">
            Connected as: {session.walletAddress}
          </p>
          <p className="mb-4">Username: {session.username}</p>
          {showForm ? (
            <form
              onSubmit={handleFormSubmit}
              className="bg-white p-4 rounded shadow-lg max-w-md mx-auto text-black "
            >
              <h2 className="text-2xl font-bold mb-4">Complete Your Profile</h2>
              <div className="mb-4">
                <label
                  className="block text-left text-black mb-2"
                  htmlFor="username"
                >
                  Username
                </label>
                <input
                  type="text"
                  id="username"
                  name="username"
                  value={formData.username}
                  onChange={handleFormChange}
                  className="w-full p-2 rounded bg-gray-100 text-black"
                  required
                />
              </div>
              <div className="mb-4">
                <label className="block text-left mb-2" htmlFor="firstName">
                  First Name
                </label>
                <input
                  type="text"
                  id="firstName"
                  name="firstName"
                  value={formData.firstName}
                  onChange={handleFormChange}
                  className="w-full p-2 rounded bg-gray-100 text-black"
                  required
                />
              </div>
              <div className="mb-4">
                <label className="block text-left mb-2" htmlFor="lastName">
                  Last Name
                </label>
                <input
                  type="text"
                  id="lastName"
                  name="lastName"
                  value={formData.lastName}
                  onChange={handleFormChange}
                  className="w-full p-2 rounded bg-gray-100 text-black"
                  required
                />
              </div>
              <div className="mb-4">
                <label className="block text-left mb-2" htmlFor="studentId">
                  Student ID
                </label>
                <input
                  type="text"
                  id="studentId"
                  name="studentId"
                  value={formData.studentId}
                  onChange={(e) => {
                    const value = e.target.value;
                    if (/^\d{0,10}$/.test(value)) {
                      handleFormChange(e);
                    }
                  }}
                  className="w-full p-2 rounded bg-gray-100 text-black "
                  required
                />
                {formData.studentId.length > 0 &&
                  formData.studentId.length !== 10 && (
                    <p className="text-red-500 mt-2">
                      Student ID must be exactly 10 digits
                    </p>
                  )}
              </div>

              <button
                type="submit"
                className="px-4 py-2 bg-blue-500 text-white rounded hover:bg-blue-700"
                // style={{
                //   backgroundColor: Colors.COLOR_PINK,
                // }}
              >
                Submit
              </button>
            </form>
          ) : (
            <></>
          )}
        </div>
      ) : (
        <button
          onClick={connectWallet}
          className="px-4 py-2  text-white rounded"
          style={{
            backgroundColor: Colors.COLOR_WHITE,
            color: Colors.COLOR_PINK,
          }}
        >
          Connect MetaMask
        </button>
      )}
      {errorMessage && <p className="mt-4 text-red-400">{errorMessage}</p>}
      <div>
        {session.token && (
          <div className="flex justify-center gap-4 mt-4">
            <button
              onClick={logout}
              className="px-4 py-2 bg-red-500 text-white rounded hover:bg-red-700"
            >
              Logout
            </button>
            <button
              onClick={() => router.push("/home")}
              className="px-4 py-2 bg-blue-500 text-white rounded hover:bg-blue-700"
            >
              Go to Home
            </button>
          </div>
        )}
        {!session.token && (
          <button
            onClick={() => router.push("/admin")}
            className="mt-8 px-4 py-2 transition duration-300 hover:underline"
          >
            or Sign in as an Administrator
          </button>
        )}
      </div>
    </div>
  );
};

export default Login;

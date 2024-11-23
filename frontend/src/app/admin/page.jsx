"use client";

import React, { useState } from "react";
import Colors from "../../app/assets/styles/colors";
import { useRouter } from "next/navigation";
import { useSession } from "../../app/SessionContext";
import { adminLogin } from "../../libs/userLib"; // Import the admin login API

export default function AdminLogin() {
  const [formData, setFormData] = useState({
    username: "",
    password: "",
  });
  const [errorMessage, setErrorMessage] = useState(null);
  const router = useRouter();
  const { session, updateSession, clearSession } = useSession(); // Access session context

  const handleInputChange = (e) => {
    const { name, value } = e.target;
    setFormData({
      ...formData,
      [name]: value,
    });
  };

  const handleSubmit = async (e) => {
    e.preventDefault();
    try {
      const data = await adminLogin(formData.username, formData.password); // Call API
      // Update session state for admin
      updateSession({
        token: null, // No token required for admin in this example
        walletAddress: null, // Admin does not have a wallet
        id: data.admin.ID,
        username: data.admin.Username,
        role: data.admin.Role,
      });

      router.push("/admin/home"); // Navigate to admin dashboard
    } catch (error) {
      setErrorMessage(
        error.response?.data?.message ||
          "Invalid credentials. Please try again."
      );
    }
  };

  const handleLogout = () => {
    clearSession(); // Clear session for admin
    router.push("/admin"); // Redirect to admin login
  };

  return (
    <div
      style={{
        backgroundColor: Colors.COLOR_PINK,
        minHeight: "100vh",
        display: "flex",
        flexDirection: "column",
        justifyContent: "center",
        alignItems: "center",
        color: Colors.COLOR_WHITE,
      }}
    >
      <form
        onSubmit={handleSubmit}
        style={{
          backgroundColor: Colors.COLOR_WHITE,
          padding: "30px 40px",
          borderRadius: "12px",
          boxShadow: "0 4px 8px rgba(0, 0, 0, 0.2)",
          maxWidth: "400px",
          width: "100%",
        }}
      >
        <h1 className="text-2xl font-bold mb-6 text-center text-black">
          Administrator Login
        </h1>
        {errorMessage && (
          <p className="text-red-500 mb-4 text-center">{errorMessage}</p>
        )}
        <div className="mb-4 text-black">
          <label className="block text-black mb-2" htmlFor="username">
            Username
          </label>
          <input
            type="text"
            id="username"
            name="username"
            value={formData.username}
            onChange={handleInputChange}
            className="w-full p-2 rounded bg-gray-100 border"
            required
          />
        </div>
        <div className="mb-4">
          <label className="block text-black mb-2" htmlFor="password">
            Password
          </label>
          <input
            type="password"
            id="password"
            name="password"
            value={formData.password}
            onChange={handleInputChange}
            className="w-full p-2 rounded bg-gray-100 text-black border"
            required
          />
        </div>
        <button
          type="submit"
          className="w-full px-4 py-2 rounded-lg bg-blue-500 text-white font-semibold hover:bg-blue-700 transition duration-300"
        >
          Submit
        </button>
      </form>
      <button
        onClick={() => router.push("/")}
        className="mt-8 px-4 py-2 transition duration-300 hover:underline"
      >
        get back to sign in with Metamask
      </button>
    </div>
  );
}

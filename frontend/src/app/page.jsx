// src/app/page.jsx
"use client";
import React from "react";
import Login from "../components/Login";
import { SessionProvider } from "./SessionContext";
import Colors from "../app/assets/styles/colors";
import { useRouter } from "next/navigation"; // Import useRouter

export default function Home() {
  const router = useRouter(); // Initialize router

  return (
    <SessionProvider>
      <div
        style={{
          backgroundColor: Colors.COLOR_PINK,
          minHeight: "100vh", // Full viewport height
          display: "flex",
          flexDirection: "column",
          justifyContent: "center",
          alignItems: "center",
          color: Colors.COLOR_WHITE, // White text
        }}
      >
        <h1 className="text-4xl font-bold mb-4">
          Welcome to Our Voting System
        </h1>
        <p className="mb-8 text-lg">
          Sign in with your MetaMask account to get started.
        </p>
        <div
          style={{
            backgroundColor: Colors.COLOR_LIGHT_GRAY,
            padding: "20px 40px",
            borderRadius: "12px",
            maxWidth: "500px",
            width: "100%",
          }}
        >
          <Login />
        </div>
      </div>
    </SessionProvider>
  );
}

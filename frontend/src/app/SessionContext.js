"use client";

// src/app/SessionContext.js
import React, { createContext, useContext, useState, useEffect } from "react";

// Define the SessionContext
const SessionContext = createContext();

export const SessionProvider = ({ children }) => {
  const [session, setSession] = useState({
    token: null,
    walletAddress: null,
  });
  const [loading, setLoading] = useState(true); // Add loading state

  useEffect(() => {
    const storedToken = localStorage.getItem("token");
    const storedWalletAddress = localStorage.getItem("walletAddress");
    setSession({
      token: storedToken || null,
      walletAddress: storedWalletAddress || null,
    });
    setLoading(false); // Mark session initialization as complete
  }, []);

  const updateSession = (newSession) => {
    if (newSession.token) {
      localStorage.setItem("token", newSession.token);
    } else {
      localStorage.removeItem("token");
    }

    if (newSession.walletAddress) {
      localStorage.setItem("walletAddress", newSession.walletAddress);
    } else {
      localStorage.removeItem("walletAddress");
    }

    setSession(newSession);
  };

  const clearSession = () => {
    localStorage.removeItem("token");
    localStorage.removeItem("walletAddress");
    setSession({ token: null, walletAddress: null });
  };

  return (
    <SessionContext.Provider
      value={{ session, loading, updateSession, clearSession }}
    >
      {children}
    </SessionContext.Provider>
  );
};

// Export the useSession hook for convenience
export const useSession = () => useContext(SessionContext);

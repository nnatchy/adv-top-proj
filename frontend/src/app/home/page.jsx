"use client";

import React, { useEffect } from "react";
import { useRouter } from "next/navigation";
import { Button, Typography, Box } from "@mui/material";
import Colors from "../../app/assets/styles/colors";
import { useDispatch, useSelector } from "react-redux";
import { fetchUserInfoThunk } from "../../reducers/userReducer"; // Thunk for fetching user info
import { useSession } from "../../app/SessionContext"; // Import useSession

export default function HomePage() {
  const router = useRouter();
  const dispatch = useDispatch();

  // Access user info and loading state from Redux
  const { userInfo, loading } = useSelector((state) => state.user);

  // Access session context
  const { session, loading: sessionLoading } = useSession();

  console.log("SESSION", session);

  // Fetch user info when walletAddress is available
  useEffect(() => {
    const walletAddress = session?.walletAddress;

    // Only proceed if walletAddress is defined
    if (walletAddress) {
      dispatch(fetchUserInfoThunk(walletAddress));
    }
  }, [dispatch, session?.walletAddress]);

  // Handle route change for reload
  useEffect(() => {
    const handleRouteChange = (url) => {
      if (url === "/home") {
        window.location.reload();
      }
    };

    router.events?.on("routeChangeComplete", handleRouteChange);

    return () => {
      router.events?.off("routeChangeComplete", handleRouteChange);
    };
  }, [router]);

  if (sessionLoading) {
    return (
      <div
        style={{
          height: "100vh",
          display: "flex",
          justifyContent: "center",
          alignItems: "center",
          fontSize: "18px",
          color: Colors.COLOR_GRAY,
        }}
      >
        Initializing session...
      </div>
    );
  }

  // Render loading state
  if (loading) {
    return (
      <div
        style={{
          height: "100vh",
          display: "flex",
          justifyContent: "center",
          alignItems: "center",
          fontSize: "18px",
          color: Colors.COLOR_GRAY,
        }}
      >
        Loading...
      </div>
    );
  }

  const hasVotedYear = userInfo?.hasVotedYear;
  const hasVotedRep = userInfo?.hasVotedRepresentative;

  return (
    <div
      style={{
        backgroundColor: Colors.COLOR_LIGHT_GRAY,
        height: "50vh",
        display: "flex",
        justifyContent: "center",
        alignItems: "center",
        flexDirection: "column",
      }}
    >
      <Box
        sx={{
          textAlign: "center",
          maxWidth: "100%",
          padding: "30px",
          borderRadius: "16px",
          boxShadow: "0 4px 12px rgba(0, 0, 0, 0.1)",
          borderColor: Colors.COLOR_BLACK,
          backgroundColor: Colors.COLOR_PINK, // Neutral background color
        }}
      >
        {hasVotedYear && hasVotedRep ? (
          <>
            <Typography
              variant="h4"
              sx={{
                fontWeight: "bold",
                marginBottom: "16px",
                color: Colors.COLOR_BLACK,
                fontFamily: "'Roboto', sans-serif",
              }}
            >
              You have already voted for all available options!
            </Typography>
            <Typography
              variant="h5"
              sx={{
                color: Colors.COLOR_BLACK,
                marginBottom: "24px",
                fontFamily: "'Roboto', sans-serif",
              }}
            >
              To see the results, please click the button below.
            </Typography>
            <Button
              variant="contained"
              onClick={() => router.push("/results")}
              sx={{
                backgroundColor: Colors.COLOR_BLACK,
                color: Colors.COLOR_WHITE,
                fontSize: "24px",
                width: "20vw",
                height: "8vh",
                borderRadius: "8px",
                textTransform: "none",
                "&:hover": {
                  backgroundColor: Colors.COLOR_WHITE,
                  color: Colors.COLOR_BLACK,
                },
              }}
            >
              Go to Results
            </Button>
          </>
        ) : (
          <Typography
            variant="h5"
            sx={{
              // fontWeight: "bold",
              // color: Colors.COLOR_BLACK,
              fontFamily: "'Roboto', sans-serif",
            }}
            className="text-white"
          >
            Please choose the type of candidate to vote.
          </Typography>
        )}
      </Box>
    </div>
  );
}

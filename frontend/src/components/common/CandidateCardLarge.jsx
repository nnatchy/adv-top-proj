"use client";
import React from "react";
import Image from "next/image";
import { Button } from "@mui/material";

const CandidateCardLarge = ({
  imageUrl,
  firstName,
  lastName,
  cp,
  major,
  onVote,
}) => {
  // Fallback to a placeholder image if the provided imageUrl is invalid
  const validImageUrl = imageUrl?.startsWith("https")
    ? imageUrl
    : "/img/profilemock.svg";

  return (
    <div
      className="w-[60%] bg-white shadow-lg rounded-lg p-8 flex items-center space-x-8"
      style={{
        margin: "20px auto", // Center horizontally with margin
      }}
    >
      {/* Candidate Image */}
      <Image
        src={validImageUrl}
        alt={`${firstName} ${lastName}`}
        width={120} // Larger size for 1920x1080
        height={120}
        className="object-cover rounded-full"
      />
      
      {/* Candidate Details */}
      <div className="flex-grow">
        <h3 className="text-2xl font-semibold text-black">{`${firstName} ${lastName}`}</h3>
        <p className="text-lg text-gray-500">{cp}</p>
        <p className="text-lg text-gray-500">{major}</p>
      </div>

      {/* Vote Button */}
      <Button
        variant="contained"
        sx={{
          backgroundColor: "#ee2e96",
          color: "white",
          height: "56px",
          width: "140px",
          borderRadius: "8px",
          fontSize: "1rem",
          textTransform: "none", // Disable uppercase text
        }}
        onClick={onVote} // Trigger the onVote action
      >
        VOTE
      </Button>
    </div>
  );
};

export default CandidateCardLarge;

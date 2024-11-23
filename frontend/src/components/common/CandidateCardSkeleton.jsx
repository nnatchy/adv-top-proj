import React from "react";
import { Skeleton } from "@mui/material";

const CandidateCardSkeleton = () => {
  return (
    <div
      className="w-[60%] bg-white shadow-lg rounded-lg p-8 flex items-center space-x-8"
      style={{
        margin: "20px auto", // Center horizontally with margin
      }}
    >
      {/* Skeleton for Image */}
      <Skeleton
        variant="circular"
        width={120}
        height={120}
        animation="wave"
        className="flex-shrink-0"
      />

      {/* Skeleton for Candidate Details */}
      <div className="flex-grow space-y-4">
        <Skeleton variant="text" width="60%" height={32} animation="wave" />
        <Skeleton variant="text" width="40%" height={24} animation="wave" />
        <Skeleton variant="text" width="40%" height={24} animation="wave" />
      </div>

      {/* Skeleton for Button */}
      <Skeleton
        variant="rectangular"
        width={140}
        height={56}
        animation="wave"
        className="rounded-lg"
      />
    </div>
  );
};

export default CandidateCardSkeleton;

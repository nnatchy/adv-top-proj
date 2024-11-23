"use client";

import React from "react";
import Image from "next/image";
import { FaTrash } from "react-icons/fa";

const positions = {
  year_one: "Year One",
  year_two: "Year Two",
  year_three: "Year Three",
  year_four: "Year Four",
  representative: "Representative",
};

const CandidateCard = ({
  id,
  imageUrl,
  firstName,
  lastName,
  position,
  major,
  onDelete,
}) => {
  return (
    <div
      className="w-full max-w-md bg-white shadow-lg rounded-lg p-4 flex items-center space-x-4 relative"
      style={{
        margin: "20px auto",
      }}
    >
      {/* Candidate Image */}
      <Image
        src={imageUrl}
        alt={`${firstName} ${lastName}`}
        width={80}
        height={80}
        className="object-cover rounded-full"
      />

      {/* Candidate Details */}
      <div className="flex-grow">
        <h3 className="text-lg font-semibold text-black">
          {`${firstName} ${lastName}`}
        </h3>
        <p className="text-sm text-gray-500">{`Position: ${
          positions[position] || "Unknown"
        }`}</p>
        <p className="text-sm text-gray-500">{`Major: ${major}`}</p>
      </div>

      {/* Delete Icon */}
      <button
        onClick={() => onDelete(id)}
        className="absolute top-2 right-2 text-red-500 hover:text-red-700 transition"
      >
        <FaTrash size={20} />
      </button>
    </div>
  );
};

export default CandidateCard;

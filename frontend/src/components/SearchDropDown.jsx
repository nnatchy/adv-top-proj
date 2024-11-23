"use client";
import React, { useState, useEffect } from "react";
import { useRouter } from "next/navigation";

const SearchDropdown = ({ positions, selectedPosition, onPositionSelect }) => {
  const [showDropdown, setShowDropdown] = useState(false);
  const router = useRouter();

  const toggleDropdown = () => setShowDropdown(!showDropdown);

  // Close dropdown when clicking outside
  useEffect(() => {
    const handleClickOutside = (event) => {
      if (!event.target.closest(".dropdown-container")) {
        setShowDropdown(false);
      }
    };

    if (showDropdown) {
      document.addEventListener("mousedown", handleClickOutside);
    }

    return () => {
      document.removeEventListener("mousedown", handleClickOutside);
    };
  }, [showDropdown]);

  const handlePositionSelect = (positionKey) => {
    onPositionSelect(positionKey);
    setShowDropdown(false);
    router.push(`/home/search?position=${encodeURIComponent(positionKey)}`);
  };

  return (
    <div className="relative dropdown-container px-96 pt-2">
      {/* Search Bar */}
      <div
        className="flex items-center bg-white shadow-md rounded-full p-2 cursor-pointer"
        onClick={toggleDropdown}
        style={{
          width: "100%", // Match the width of the dropdown
          boxSizing: "border-box", // Ensures padding doesn't affect width
        }}
      >
        <span
          style={{
            color: "#000",
            fontSize: "24px",
            marginLeft: "8px",
          }}
        >
          🔍
        </span>
        <input
          type="text"
          placeholder="Select Position"
          value={selectedPosition}
          readOnly
          className="text-black flex-grow bg-transparent border-none outline-none px-4 text-sm cursor-pointer"
        />
      </div>

      {/* Dropdown */}
      <div
        className={`absolute z-10 bg-white shadow-md rounded-md transition-all duration-300 ease-in-out ${
          showDropdown ? "opacity-100 max-h-64" : "opacity-0 max-h-0"
        } overflow-hidden`}
        style={{
          width: "47%", // Ensure the dropdown matches the search
        }}
      >
        <ul>
          {Object.entries(positions).map(([key, value]) => (
            <li
              key={key}
              className="px-4 py-2 hover:bg-gray-200 cursor-pointer text-black"
              onClick={() => handlePositionSelect(key)}
            >
              {value}
            </li>
          ))}
        </ul>
      </div>
    </div>
  );
};

export default SearchDropdown;

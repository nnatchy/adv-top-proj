"use client";

import React from "react";
import { motion, AnimatePresence } from "framer-motion";
import { Button } from "@mui/material";
import Image from "next/image";

const VotePopup = ({ isOpen, onClose, candidateName, candidateImage, onConfirm }) => {
  return (
    <AnimatePresence>
      {isOpen && (
        <>
          {/* Background overlay */}
          <motion.div
            className="fixed inset-0 bg-black opacity-50 z-10"
            initial={{ opacity: 0 }}
            animate={{ opacity: 0.5 }}
            exit={{ opacity: 0 }}
          />

          {/* Popup container */}
          <motion.div
            className="fixed inset-0 flex justify-center items-center z-20"
            initial={{ scale: 0.8, opacity: 0 }}
            animate={{ scale: 1, opacity: 1 }}
            exit={{ scale: 0.8, opacity: 0 }}
            transition={{ duration: 0.3 }}
          >
            <div className="bg-white rounded-lg shadow-lg w-[500px] p-8 flex flex-col items-center">
              {/* Candidate Image */}
              <Image
                src={candidateImage || "/img/profilemock.svg"}
                alt={candidateName}
                width={150}
                height={150}
                className="rounded-full object-cover mb-6"
              />

              {/* Title */}
              <h2 className="text-2xl font-bold text-gray-800 text-center mb-4">
                Confirm Your Vote
              </h2>

              {/* Candidate Details */}
              <p className="text-gray-600 text-center mb-8">
                Are you sure you want to vote for <strong>{candidateName}</strong>?
              </p>

              {/* Action Buttons */}
              <div className="flex justify-center space-x-4">
                <Button
                  variant="outlined"
                  onClick={onClose}
                  sx={{
                    color: "gray",
                    borderColor: "gray",
                    textTransform: "none",
                    padding: "10px 20px",
                  }}
                >
                  Cancel
                </Button>
                <Button
                  variant="contained"
                  onClick={onConfirm}
                  sx={{
                    backgroundColor: "#FF98A0",
                    color: "white",
                    textTransform: "none",
                    padding: "10px 20px",
                  }}
                >
                  Confirm
                </Button>
              </div>
            </div>
          </motion.div>
        </>
      )}
    </AnimatePresence>
  );
};

export default VotePopup;

"use client";

import React from "react";
import { motion, AnimatePresence } from "framer-motion";
import Image from "next/image";
import { Button } from "@mui/material";

const StatusPopup = ({ isOpen, status, onClose }) => {
  const isSuccess = status === 200;

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
            <div className="bg-white rounded-lg shadow-lg w-[500px] p-8 text-center">
              {/* Status Icon */}
              <div className="mb-6 flex justify-center">
                {isSuccess ? (
                  <Image
                    src="/img/check.svg"
                    alt="Success"
                    width={120} // Increased size
                    height={120}
                  />
                ) : (
                  <Image
                    src="/img/fail.svg"
                    alt="Error"
                    width={120} // Increased size
                    height={120}
                  />
                )}
              </div>

              {/* Title */}
              <h2 className="text-2xl font-bold text-gray-800 mb-4">
                {isSuccess ? "Submission Successful" : "Submission Failed"}
              </h2>

              {/* Message */}
              <p className="text-gray-600 mb-6 text-lg">
                {isSuccess
                  ? "Your vote has been successfully submitted."
                  : "An error occurred while submitting your vote. Please try again."}
              </p>

              {/* Action Button */}
              <Button
                variant="contained"
                sx={{
                  backgroundColor: isSuccess ? "#4caf50" : "#f44336",
                  color: "white",
                  padding: "10px 30px",
                  fontSize: "16px",
                  borderRadius: "8px",
                  textTransform: "none",
                }}
                onClick={onClose}
              >
                {isSuccess ? "Back to Home" : "Close"}
              </Button>
            </div>
          </motion.div>
        </>
      )}
    </AnimatePresence>
  );
};

export default StatusPopup;

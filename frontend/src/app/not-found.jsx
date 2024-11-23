"use client";
import React from "react";
import { useRouter } from "next/navigation";
import Colors from "./assets/styles/colors";

const NotFoundPage = () => {
  const router = useRouter();

  const goToHome = () => {
    router.push("/"); // Navigate back to home
  };

  return (
    <div
      className="flex flex-col items-center justify-center min-h-screen text-white px-12"
      style={{ backgroundColor: Colors.COLOR_PINK }}
    >
      <h1 className="text-4xl font-bold">Page Not Found</h1>
      <p className="text-lg mt-4">Oops! There is no page you're looking for</p>
      <button
        className="mt-8 px-6 py-3 bg-white text-black font-semibold rounded-lg hover:bg-gray-200 transition duration-300"
        onClick={goToHome}
      >
        Back to Home
      </button>
    </div>
  );
};

export default NotFoundPage;

"use client";
import { ReactNode, useEffect, useState } from "react";
import { usePathname, useSearchParams } from "next/navigation";
import Header from "../../components/common/Header";

export default function HomeLayout({ children }) {
  const pathname = usePathname();
  const isSearchPage = pathname.startsWith("/home/search"); // Check if it's the search page

  const [isSearchActive, setSearchActive] = useState(false);

  const handleSearchClick = () => {
    setSearchActive(true); // Activate search
  };

  const handleCloseSearch = () => {
    setSearchActive(false); // Deactivate search
  };

  return (
    <div className="w-screen h-screen flex flex-col bg-gray-50">
      {/* Header */}
      <Header
        toggleMenu={() => {}}
        handleSearchClick={handleSearchClick}
        animateHeader={!isSearchActive}
        searchValue=""
      />

      {/* Main Content */}
      <main className="flex-1 overflow-auto">
        {children}
      </main>
    </div>
  );
}

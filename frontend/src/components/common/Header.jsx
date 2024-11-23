"use client";
import { IoIosMenu } from "react-icons/io";
import Colors from "../../app/assets/styles/colors";
import { useState, useEffect } from "react";
import { usePathname } from "next/navigation";
import { useDispatch, useSelector } from "react-redux";
import Image from "next/image";
import SearchDropdown from "../SearchDropDown";
import { getYearFromStudentId } from "../../utils/getYear";
import { useSession } from "../../app/SessionContext";
import { useRouter } from "next/navigation"; // Add this import

const Header = ({ toggleMenu = () => {}, animateHeader }) => {
  const pathname = usePathname();
  const dispatch = useDispatch();
  const { clearSession } = useSession();
  const router = useRouter();

  // Access user info from the Redux store
  const { userInfo, loading } = useSelector((state) => state.user);
  const [isTextVisible, setTextVisible] = useState(false); // State to control text visibility
  const [filteredPositions, setFilteredPositions] = useState([]); // Filtered positions for the dropdown
  const [selectedPosition, setSelectedPosition] = useState(""); // State to track selected position
  const [fullName, setFullname] = useState("");

  // Map of positions for display text
  const positions = {
    year_one: "Year One",
    year_two: "Year Two",
    year_three: "Year Three",
    year_four: "Year Four",
    representative: "Representative",
  };

  const handleLogout = () => {
    clearSession(); // Clear session on logout
    alert("Logged out successfully.");
    router.push("/"); // Redirect to home or another page
  };

  useEffect(() => {
    setTextVisible(true); // Set the text to visible after the page renders
  }, []);

  useEffect(() => {
    if (!userInfo || !userInfo.student_id) {
      return;
    }

    try {
      const studentYear = getYearFromStudentId(userInfo.student_id);
      const availablePositions = [];

      if (!userInfo.hasVotedYear) {
        if (studentYear === 1) availablePositions.push("year_one");
        if (studentYear === 2) availablePositions.push("year_two");
        if (studentYear === 3) availablePositions.push("year_three");
        if (studentYear === 4) availablePositions.push("year_four");
      }

      // Add the representative option if the user hasn't voted for it
      if (!userInfo.hasVotedRepresentative) {
        availablePositions.push("representative");
      }
      const fullNameTmp =
        userInfo.first_name && userInfo.last_name
          ? `${userInfo.first_name} ${userInfo.last_name}`
          : "Anonymous User"; // Fallback name if user info is missing
      setFullname(fullNameTmp);
      setFilteredPositions(availablePositions); // Update filtered positions
    } catch (error) {
      console.error("Error filtering positions:", error);
    }
  }, [userInfo]); // Dependency array ensures it runs when userInfo updates

  useEffect(() => {
    const searchParams = new URLSearchParams(window.location.search);
    const positionFromQuery = searchParams.get("position");
    if (positionFromQuery && positions[positionFromQuery]) {
      setSelectedPosition(positions[positionFromQuery]);
    }
  }, [pathname]);

  return (
    <header
      className={`w-full py-4 px-6 transition-transform duration-500 ease-in-out ${
        animateHeader
          ? "translate-y-0 opacity-100"
          : "translate-y-[-150px] opacity-0"
      }`}
      style={{
        backgroundColor: Colors.COLOR_PINK,
      }}
    >
      <div className="">
        <div className="flex items-center justify-center w-full pt-8 ">
          <div className="flex flex-col justify-center items-center">
            <>
              <Image
                src="https://picsum.photos/seed/picsum/200/300"
                width={80} // Slightly larger image
                height={80}
                className="w-[80px] h-[80px] rounded-full object-cover bg-gray-300"
                alt="Profile"
              />
            </>
            <div className="flex flex-col ml-4 mt-2 items-center">
              <p className="text-white text-2xl font-bold">
                Engineer Voting System
              </p>
              <p className="text-gray-800 text-xl mr-1 font-bold">{fullName}</p>
            </div>
          </div>
        </div>
        <div className="absolute top-4 right-4">
          <button
            onClick={handleLogout}
            className="px-4 py-2 bg-white rounded-md transition"
            style={{
              color: Colors.COLOR_PINK,
            }}
          >
            Logout
          </button>
        </div>

        <hr className="border-t-1 border-white my-4" />

        {/* Use the SearchDropdown Component */}
        {!loading && userInfo && (
          <SearchDropdown
            positions={filteredPositions.reduce((acc, pos) => {
              acc[pos] = positions[pos]; // Map filtered positions to their display text
              return acc;
            }, {})}
            selectedPosition={selectedPosition}
            disabled={userInfo.has_vote_year && userInfo.has_vote_rep} // Disable dropdown if both votes are done
            onPositionSelect={(positionKey) =>
              setSelectedPosition(positions[positionKey])
            }
          />
        )}

        {pathname.startsWith("/home") && (
          <div className="flex justify-center flex-col">
            <div className="flex justify-center">
              <div
                className={`ml-4 mt-6 transition-all duration-1000 ease-in-out ${
                  isTextVisible
                    ? "opacity-100 translate-y-0"
                    : "opacity-0 -translate-y-10"
                }`}
              >
                <div className="flex items-center text-xl font-bold">
                  <p className="text-white mr-2">Hi</p>
                  <p className="text-white mr-2">{fullName}, </p>
                  <p className="text-gray-800 mr-2"> Welcome to</p>
                  <p className="text-white">Engineer Voting System</p>
                </div>
              </div>
            </div>
            <>
              <button
                onClick={() => router.push("/results")}
                className="mt-2 px-4 py-2 transition duration-300 hover:underline text-white"
              >
                already voted? go to result
              </button>
            </>
          </div>
        )}
      </div>
    </header>
  );
};

export default Header;

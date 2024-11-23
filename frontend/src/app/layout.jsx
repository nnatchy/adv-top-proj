import { Inter } from "next/font/google";
import "./globals.css";
import ReduxProvider from "../components/layout/ReduxProvider";
// import SessionProviders from "../components/layout/SessionProvider";
import { SessionProvider } from "./SessionContext";
import { getServerSession } from "next-auth";
// import { authOptions } from "./api/auth/[...nextauth]/route";

const inter = Inter({ subsets: ["latin"] });

export const metadata = {
  title: "AdvTopCom",
  description: "Generated by create next app",
};

export default async function RootLayout({ children }) {
  // const nextAuthsession = await getServerSession(authOptions);

  return (
    <html lang="en">
      <body>
        <ReduxProvider>
          <SessionProvider>
            {children}
          </SessionProvider>
            {/* {children} */}

        </ReduxProvider>
      </body>
    </html>
  );
}
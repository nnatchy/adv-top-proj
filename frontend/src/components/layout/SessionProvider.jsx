"use client";
import { SessionProvider } from "next-auth/react";
import { ReactNode } from "react";

const SessionProviders = ({ children, session }) => {
  return <SessionProvider session={session}>{children}</SessionProvider>;
};

export default SessionProviders;

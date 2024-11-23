"use client"
import { createContext, useContext, useState, useEffect } from 'react';
import { ethers } from 'ethers';

const EthereumContext = createContext();

export const EthereumProvider = ({ children }) => {
  const [provider, setProvider] = useState(null);

  useEffect(() => {
    const infuraProvider = new ethers.providers.InfuraProvider('sepolia', process.env.NEXT_PUBLIC_INFURA_PROJECT_ID);
    setProvider(infuraProvider);
  }, []);

  return (
    <EthereumContext.Provider value={{ provider }}>
      {children}
    </EthereumContext.Provider>
  );
};

export const useEthereum = () => useContext(EthereumContext);

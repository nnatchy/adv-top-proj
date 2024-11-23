"use client"
import { useState } from 'react';
import { setNumberInSimpleStorage } from '../libs/contracts';

export default function SetNumber() {
  const [value, setValue] = useState('');
  const [txHash, setTxHash] = useState('');

  const handleSetNumber = async () => {
    try {
      const data = await setNumberInSimpleStorage(value);
      setTxHash(data.transactionHash);
    } catch (error) {
      console.error('Error setting number:', error);
    }
  };

  return (
    <div>
      <h2 className="text-xl font-semibold">Set Number</h2>
      <input
        type="number"
        value={value}
        onChange={(e) => setValue(e.target.value)}
        className="border p-2 mr-2"
        placeholder="Enter number"
      />
      <button onClick={handleSetNumber} className="bg-blue-500 text-white p-2">
        Set Number
      </button>
      {txHash && <p>Transaction Hash: {txHash}</p>}
    </div>
  );
}

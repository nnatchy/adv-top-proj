"use client"
import { useState } from 'react';
import { getNumberFromSimpleStorage } from '../libs/contracts';

export default function GetNumber() {
  const [storedValue, setStoredValue] = useState(null);

  const handleGetNumber = async () => {
    try {
      const data = await getNumberFromSimpleStorage();
      setStoredValue(data.storedValue);
    } catch (error) {
      console.error('Error getting number:', error);
    }
  };

  return (
    <div>
      <h2 className="text-xl font-semibold">Get Number</h2>
      <button onClick={handleGetNumber} className="bg-green-500 text-white p-2">
        Get Number
      </button>
      {storedValue !== null && <p>Stored Value: {storedValue}</p>}
    </div>
  );
}

import SetNumber from '../../components/SetNumber';
import GetNumber from '../../components/GetNumber';
import { EthereumProvider } from '../../contexts/EthereumContext';

export default function SimpleStoragePage() {
  return (
    <EthereumProvider>
      <div className="p-8">
        <h1 className="text-2xl font-bold mb-4">SimpleStorage DApp</h1>
        <div className="grid grid-cols-1 md:grid-cols-2 gap-4">
          <SetNumber />
          <GetNumber />
        </div>
      </div>
    </EthereumProvider>
  );
}

import { usePrivy, useWallets } from "@privy-io/react-auth";
import { LoginWithEmail } from "./components";
import { useState, useEffect } from "react";
import { ethers } from "ethers";

const PYUSD_ADDRESS = "0x6c3ea9036406852006290770bedfcaba0e23a0e8";
const ERC20_ABI = [
  "function balanceOf(address owner) view returns (uint256)",
  "function decimals() view returns (uint8)"
];

function App() {
  const { ready, authenticated, user, logout } = usePrivy();
  const { wallets } = useWallets();

  // Balance states as string or null
  const [ethBalance, setEthBalance] = useState<string | null>(null);
  const [pyusdBalance, setPyusdBalance] = useState<string | null>(null);

  useEffect(() => {
    async function fetchBalances() {
      if (ready && authenticated && wallets && wallets.length > 0) {
        const provider = new ethers.BrowserProvider(await wallets[0].getEthereumProvider());
        const address = wallets[0].address;

        // ETH
        const ethWei = await provider.getBalance(address);
        setEthBalance(ethers.formatEther(ethWei));

        // PYUSD
        const pyusdContract = new ethers.Contract(PYUSD_ADDRESS, ERC20_ABI, provider);
        const rawPyusd = await pyusdContract.balanceOf(address);
        const decimals = await pyusdContract.decimals();
        setPyusdBalance(ethers.formatUnits(rawPyusd, decimals));
      }
    }
    fetchBalances();
  }, [ready, authenticated, wallets]);

  if (!ready) return <div>Loading...</div>;
  if (!(ready && authenticated)) return <LoginWithEmail />;

  return (
    <div>
      <div>Hello, {user?.email?.address}</div>
      <div>Wallet Address: {wallets && wallets[0] && wallets[0].address}</div>
      <div>ETH Balance: {ethBalance ?? "..."}</div>
      <div>PYUSD Balance: {pyusdBalance ?? "..."}</div>
      <button onClick={logout} disabled={!ready || !authenticated}>Log out</button>
    </div>
  );
}

export default App;

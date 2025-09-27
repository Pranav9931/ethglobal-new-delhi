import React, { useState, useEffect } from "react";
import { Buffer } from "buffer";
import { usePrivy, useWallets, useSessionSigners } from "@privy-io/react-auth";
import { LoginWithEmail } from "./components";
import { ethers } from "ethers";

// Polyfill Buffer globally for Privy SDK
window.Buffer = Buffer;

const PYUSD_ADDRESS = "0x6c3ea9036406852006290770bedfcaba0e23a0e8";
const ERC20_ABI = [
  "function balanceOf(address owner) view returns (uint256)",
  "function decimals() view returns (uint8)"
];

const backendSignerId = "jgnudaohe27wths41piqaizd";

function App() {
  const { ready, authenticated, user, logout } = usePrivy();
  const { wallets } = useWallets();
  const { addSessionSigners } = useSessionSigners();

  const [ethBalance, setEthBalance] = useState<string | null>(null);
  const [pyusdBalance, setPyusdBalance] = useState<string | null>(null);
  const [backendSignerAdded, setBackendSignerAdded] = useState(false);

  useEffect(() => {
    async function fetchBalances() {
      if (ready && authenticated && wallets && wallets.length > 0) {
        const provider = new ethers.BrowserProvider(await wallets[0].getEthereumProvider());
        const address = wallets[0].address;

        // Fetch ETH balance
        const ethWei = await provider.getBalance(address);
        setEthBalance(ethers.formatEther(ethWei));

        // Fetch PYUSD balance
        const pyusdContract = new ethers.Contract(PYUSD_ADDRESS, ERC20_ABI, provider);
        const rawPyusd = await pyusdContract.balanceOf(address);
        const decimals = await pyusdContract.decimals();
        setPyusdBalance(ethers.formatUnits(rawPyusd, decimals));
      }
    }
    fetchBalances();
  }, [ready, authenticated, wallets]);

  // Add backend signer session after wallet setup
  // useEffect(() => {
  //   async function addBackendSignerToWallet() {
  //     if (!backendSignerAdded && ready && authenticated && wallets && wallets.length > 0) {
  //       try {
  //         await addSessionSigners({
  //           address: wallets[0].address,
  //           signers: [{ signerId: backendSignerId, policyIds: [] }] // no policy restrictions
  //         });
  //         setBackendSignerAdded(true);
  //         console.log("Backend signer added successfully");
  //       } catch (error) {
  //         console.error("Failed to add backend signer:", error);
  //       }
  //     }
  //   }
  //   addBackendSignerToWallet();
  // }, [ready, authenticated, wallets, backendSignerAdded, addSessionSigners]);

  if (!ready) return <div>Loading...</div>;
  if (!authenticated) return <LoginWithEmail />;

  return (
    <div>
      <div>Hello, {user?.email?.address}</div>
      <div>Wallet Address: {wallets?.[0]?.address ?? "N/A"}</div>
      <div>ETH Balance: {ethBalance ?? "..."}</div>
      <div>PYUSD Balance: {pyusdBalance ?? "..."}</div>
      {/* <div>Backend Signer Added? {backendSignerAdded ? "YES" : "NO"}</div> */}
      <button onClick={logout} disabled={!ready || !authenticated}>Log out</button>
    </div>
  );
}

export default App;

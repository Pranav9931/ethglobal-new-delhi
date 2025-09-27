import {useState, useEffect} from 'react';
import {useCreateWallet, useLoginWithEmail, useWallets} from '@privy-io/react-auth';

const SESSION_KEY = 'privy_user_session';

export default function LoginWithEmail() {
  const [email, setEmail] = useState('');
  const [code, setCode] = useState('');
//   const [session, setSession] = useState(null);
  const [walletAddress, setWalletAddress] = useState('');

  const { wallets } = useWallets();

  useEffect(() => {
    const saved = localStorage.getItem(SESSION_KEY);
    if (saved) {
      const parsed = JSON.parse(saved);
    //   setSession(parsed);
      const wallet = parsed.user.linked_accounts.find((a: {type: string}) => a.type === 'wallet');
      if (wallet) setWalletAddress(wallet.address);
    }
  }, []);

  const {sendCode, loginWithCode} = useLoginWithEmail();
  const {createWallet} = useCreateWallet({
    onSuccess: ({Wallet}: any) => {
      const userData = {
        user: Wallet.user,
        token: Wallet.token,
        privy_access_token: Wallet.privy_access_token,
        refresh_token: Wallet.refresh_token,
      };
      localStorage.setItem(SESSION_KEY, JSON.stringify(userData));
    //   setSession(userData);
      setWalletAddress(Wallet.address);
      console.log('Created wallet', Wallet);
    },
    onError: err => {
    //   check if user has already a wallet..
        const wallet = wallets?.find(w => w.type === 'ethereum');
        if (wallet) {
            setWalletAddress(wallet.address);
            console.log('User already has a wallet', wallet);
        } else {
            console.error('Failed to create wallet', err);
        }
    }
  });

  const loginWithCodeAsync = async ({code}: {code: string}) => {
    try {
      await loginWithCode({code});
      await createWallet();
    } catch (error) {
      console.error('Failed to login with error', error);
    }
  };

  return (
    <div>
      {walletAddress ? (
        <div>Wallet Address: {walletAddress}</div>
      ) : (
        <>
          <input onChange={e => setEmail(e.target.value)} value={email} placeholder="Email" />
          <button onClick={() => sendCode({email})}>Send Code</button>
          <input onChange={e => setCode(e.target.value)} value={code} placeholder="Code" />
          <button onClick={() => loginWithCodeAsync({code})}>Login</button>
        </>
      )}
    </div>
  );
}

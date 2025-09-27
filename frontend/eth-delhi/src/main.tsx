import React from 'react';
import ReactDOM from 'react-dom/client';

import './index.css';

import {PrivyProvider} from '@privy-io/react-auth';

import App from './App';

const root = ReactDOM.createRoot(document.getElementById('root') as HTMLElement);

root.render(
  <React.StrictMode>
    <PrivyProvider
      appId="cmg17gsej004qla0dar9k36t8"
      // clientId="your-app-client-id"
      config={{
        // Create embedded wallets for users who don't have a wallet
        embeddedWallets: {
          ethereum: {
            createOnLogin: 'users-without-wallets'
          }
        }
      }}
    >
      <App />
    </PrivyProvider>
  </React.StrictMode>
);
/**
 * @format
 */

import { AppRegistry } from 'react-native';
import { QueryClient, QueryClientProvider, setConsole } from 'react-query';
import { RecoilRoot } from 'recoil';
import App from './src/App';
import { name as appName } from './app.json';
import RecoilOutside from 'recoil-outside';
import React from 'react';

export const queryClient = new QueryClient();
setConsole({
    log: console.log,
    warn: console.warn,
    error: console.warn,
});

const WrappedApp = () => (
    <>
        <RecoilRoot>
            <RecoilOutside />
            <QueryClientProvider client={queryClient}>
                <App />
            </QueryClientProvider>
        </RecoilRoot>
    </>
);

AppRegistry.registerComponent(appName, () => WrappedApp);

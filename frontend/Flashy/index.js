import { AppRegistry } from 'react-native';
import App from './App';
import { name as appName } from './app.json';
import { QueryClient, QueryClientProvider } from 'react-query';
import { RecoilRoot } from 'recoil';
import RecoilOutside from 'recoil-outside';

export const queryClient = new QueryClient();

const WrappedApp = () => {
    return (
        <>
            <RecoilRoot>
                <RecoilOutside />
                <QueryClientProvider client={queryClient}>
                    <App />
                </QueryClientProvider>
            </RecoilRoot>
        </>
    );
};

AppRegistry.registerComponent(appName, () => WrappedApp);

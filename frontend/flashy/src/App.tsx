import * as eva from '@eva-design/eva';
import { ApplicationProvider, IconRegistry } from '@ui-kitten/components';
import { EvaIconsPack } from '@ui-kitten/eva-icons';
import React from 'react';
import { StatusBar } from 'react-native';
import { useRecoilValue } from 'recoil';
import { Navigation } from './nav';
import { themeColorState } from './state/theme';
import { Settings } from 'react-native-fbsdk-next';

Settings.initializeSDK();

const App = () => {
    const themeColor = useRecoilValue(themeColorState);
    return (
        <>
            <StatusBar hidden />
            <IconRegistry icons={EvaIconsPack} />
            <ApplicationProvider {...eva} theme={themeColor}>
                <Navigation />
            </ApplicationProvider>
        </>
    );
};

export default App;

import { NavigationContainer } from '@react-navigation/native';
import { createNativeStackNavigator } from '@react-navigation/native-stack';
import { Settings, SCREENS, LogIn, Home, Learn, SignUp } from '../screens';
import React from 'react';

export type StackParams = {
    Home: undefined;
    LogIn: undefined;
    Learn: undefined;
    Settings: undefined;
    SignUp: undefined;
};

const Stack = createNativeStackNavigator<StackParams>();

export const Navigation = (): JSX.Element => {
    return (
        <NavigationContainer>
            <Stack.Navigator initialRouteName={SCREENS.LOG_IN}>
                <Stack.Screen
                    name={SCREENS.LOG_IN}
                    component={LogIn}
                    options={{ headerShown: false }}
                />
                <Stack.Screen name={SCREENS.SIGN_UP} component={SignUp} />
                <Stack.Screen name={SCREENS.HOME} component={Home} />
                <Stack.Screen name={SCREENS.LEARN} component={Learn} />
                <Stack.Screen name={SCREENS.SETTINGS} component={Settings} />
            </Stack.Navigator>
        </NavigationContainer>
    );
};

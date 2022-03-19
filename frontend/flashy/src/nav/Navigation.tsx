import { NavigationContainer } from '@react-navigation/native';
import { createNativeStackNavigator } from '@react-navigation/native-stack';
import { Settings, SCREENS, LogIn, Home, Learn } from '../screens';

export type StackParams = {
    Home: undefined;
    LogIn: undefined;
    Learn: undefined;
    Settings: undefined;
};

const Stack = createNativeStackNavigator<StackParams>();

export const Navigation = (): JSX.Element => {
    return (
        <NavigationContainer>
            <Stack.Navigator initialRouteName={SCREENS.LOG_IN}>
                <Stack.Screen name={SCREENS.LOG_IN} component={LogIn} />
                <Stack.Screen name={SCREENS.HOME} component={Home} />
                <Stack.Screen name={SCREENS.LEARN} component={Learn} />
                <Stack.Screen name={SCREENS.SETTINGS} component={Settings} />
            </Stack.Navigator>
        </NavigationContainer>
    );
};

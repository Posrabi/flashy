import { useNavigation } from '@react-navigation/native';
import { NativeStackNavigationProp } from '@react-navigation/native-stack';
import { Button, Icon, Input, Text } from '@ui-kitten/components';
import React from 'react';
import { StyleSheet, TouchableOpacity, SafeAreaView } from 'react-native';
import { LoginManager } from 'react-native-fbsdk-next';
import { useSetRecoilState } from 'recoil';
import { defaultUser } from '../api/defaults';
import EndpointsModule from '../api/users';
import { StackParams } from '../nav';
import { currentUser } from '../state/user';
import { SCREENS } from './constants';

type LogInScreenProp = NativeStackNavigationProp<StackParams, SCREENS.LOG_IN>;

export const LogIn = (): JSX.Element => {
    const [username, setUsername] = React.useState('');
    const [password, setPassword] = React.useState('');
    const [secureTextEntry, setSecureTextEntry] = React.useState(true);
    const [state, setState] = React.useState(true);
    const setUser = useSetRecoilState(currentUser);
    const nav = useNavigation<LogInScreenProp>();

    const EyeIcon = (props: any): JSX.Element => (
        <TouchableOpacity onPress={() => setSecureTextEntry(!secureTextEntry)}>
            <Icon {...props} name={!secureTextEntry ? 'eye' : 'eye-off'} />
        </TouchableOpacity>
    );

    const onLogIn = async (): Promise<void> => {
        try {
            const { user } = await EndpointsModule.LogIn({
                user_name: username,
                hash_password: password,
            });
            if (!user) {
                setState(false);
                console.log('no user return');
                return;
            }
            setUser(user);
            nav.navigate(SCREENS.HOME);
        } catch (e) {
            setState(false);
            console.error(e);
        }
    };

    const LogInWithFacebook = (): JSX.Element => {
        const logIn = () =>
            LoginManager.logInWithPermissions(['public_profile', 'email']).then(
                (result) => {
                    if (result.isCancelled) {
                        setState(false);
                        console.log('Log in cancelled');
                    } else {
                        console.log(
                            `Log in success with permissions: ${result.grantedPermissions?.toString()}`
                        );
                    }
                },
                (error) => {
                    setState(false);
                    console.error(error);
                }
            );
        return (
            <TouchableOpacity style={[styles.fields, styles.facebook]} onPress={logIn}>
                <Icon name="facebook" fill="white" width={30} height={30} />
                <Text style={styles.facebookText}>Sign in with Facebook</Text>
            </TouchableOpacity>
        );
    };

    return (
        <SafeAreaView style={styles.layout}>
            <Text style={styles.fieldsText}>Sign in with username and password</Text>
            <Input
                placeholder="Username"
                size="large"
                status={state ? 'basic' : 'danger'}
                value={username}
                style={styles.fields}
                onChangeText={setUsername}
            />
            <Input
                placeholder="Password"
                size="large"
                status={state ? 'basic' : 'danger'}
                value={password}
                style={styles.fields}
                secureTextEntry={secureTextEntry}
                onChangeText={setPassword}
                accessoryRight={EyeIcon}
            />
            {!state && (
                <Text style={[styles.caption, styles.error]} status="danger">
                    An error has occurred, please try again later.
                </Text>
            )}
            <Button status="info" style={[styles.fields]} onPress={onLogIn}>
                Sign In
            </Button>
            <Button
                style={[styles.fields]}
                status="info"
                appearance="outline"
                onPress={() => {
                    setUser(defaultUser());
                    nav.navigate(SCREENS.HOME);
                }}>
                Start without an account
            </Button>
            <Text style={styles.fieldsText}>Or</Text>
            <LogInWithFacebook />
            <Text style={styles.caption}>Don't have an account? Sign up here</Text>
            <Button
                status="primary"
                style={[styles.fields]}
                onPress={() => nav.navigate(SCREENS.SIGN_UP)}>
                Sign Up
            </Button>
        </SafeAreaView>
    );
};

const styles = StyleSheet.create({
    layout: {
        flex: 1,
        justifyContent: 'center',
        alignItems: 'center',
    },
    fields: {
        margin: 5,
        width: 250,
        height: 45,
        borderRadius: 5,
        justifyContent: 'center',
    },
    center: {
        textAlign: 'center',
        fontWeight: 'bold',
        fontSize: 17,
    },
    fieldsText: {
        fontWeight: 'bold',
        fontSize: 18,
        alignItems: 'center',
        textAlign: 'center',
        width: 250,
        margin: 10,
    },
    caption: {
        marginTop: 10,
        fontWeight: 'bold',
        fontSize: 15,
        alignItems: 'center',
        textAlign: 'center',
        width: 250,
    },
    error: {
        margin: 10,
    },
    facebookText: {
        textAlign: 'center',
        color: '#ffffff',
        fontWeight: 'bold',
        fontSize: 17,
    },
    facebook: {
        backgroundColor: '#385898',
        flexDirection: 'row',
        justifyContent: 'space-evenly',
        alignItems: 'center',
    },
});

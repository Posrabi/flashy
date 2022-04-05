import { useNavigation } from '@react-navigation/native';
import { NativeStackNavigationProp } from '@react-navigation/native-stack';
import { Button, Icon, Input, Text } from '@ui-kitten/components';
import React, { useEffect } from 'react';
import { StyleSheet, TouchableOpacity, SafeAreaView } from 'react-native';
import { AccessToken, LoginManager, Profile } from 'react-native-fbsdk-next';
import { useRecoilState } from 'recoil';
import { defaultUser } from '../api/defaults';
import EndpointsModule from '../api/users';
import { StackParams } from '../nav';
import {
    currentUser,
    getAuthTokenFromStorage,
    getFBAccessToken,
    getUserIDFromStorage,
    storeFBAccessToken,
    storeProfileURI,
    storeUser,
} from '../state/user';
import { CreateUserRequest } from '../types';
import { SCREENS } from './constants';

type LogInScreenProps = NativeStackNavigationProp<StackParams, SCREENS.LOG_IN>;

export const LogIn = (): JSX.Element => {
    const [username, setUsername] = React.useState('');
    const [password, setPassword] = React.useState('');
    const [secureTextEntry, setSecureTextEntry] = React.useState(true);
    const [state, setState] = React.useState(true);
    const [user, setUser] = useRecoilState(currentUser);
    const nav = useNavigation<LogInScreenProps>();
    const EyeIcon = (props: any): JSX.Element => (
        <TouchableOpacity onPress={() => setSecureTextEntry(!secureTextEntry)}>
            <Icon {...props} name={!secureTextEntry ? 'eye' : 'eye-off'} />
        </TouchableOpacity>
    );
    useEffect(() => {
        if (user.user_id) nav.navigate(SCREENS.HOME);
        else {
            (async () => {
                try {
                    const userID = await getUserIDFromStorage();
                    const authToken = await getAuthTokenFromStorage();
                    const facebookAccessToken = await getFBAccessToken();
                    if (facebookAccessToken && userID) {
                        const resp = await EndpointsModule.LogInWithFB({
                            user_id: userID,
                            facebook_access_token: facebookAccessToken,
                        });
                        if (resp.user) {
                            setUser(resp.user);
                            storeUser(resp.user.user_id, '');
                            storeFBAccessToken(resp.user.facebook_access_token);
                            nav.navigate(SCREENS.HOME);
                            return;
                        } else {
                            console.error('unable to find user');
                        }
                    }
                    if (userID && authToken) {
                        const resp = await EndpointsModule.GetUser({
                            user_id: userID,
                            // @ts-ignore
                            auth_token: authToken, // hack to pass auth_token to getUser endpoint.
                        });
                        if (resp.user) {
                            setUser(resp.user);
                            storeUser(resp.user.user_id, resp.user.auth_token);
                            nav.navigate(SCREENS.HOME);
                        } else {
                            console.error('unable to find user');
                        }
                    }
                } catch (e) {
                    console.error(e);
                }
            })();
        }
    }, []);

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
            storeUser(user.user_id, user.auth_token);
            nav.navigate(SCREENS.HOME);
        } catch (e) {
            setState(false);
            console.error(e);
        }
    };

    const LogInWithFacebook = (): JSX.Element => {
        const onFBLogIn = (): Promise<void> =>
            LoginManager.logInWithPermissions(['public_profile', 'email']).then(
                (result) => {
                    if (result.isCancelled) {
                        setState(false);
                        console.log('Log in cancelled');
                    } else {
                        setState(true);
                        AccessToken.getCurrentAccessToken().then((data) => {
                            if (data?.accessToken) {
                                Profile.getCurrentProfile().then(async (profile) => {
                                    try {
                                        const req: CreateUserRequest = {
                                            user: {
                                                name: profile?.name || 'N/A',
                                                email: profile?.email || '',
                                                facebook_access_token: data.accessToken.toString(),
                                                user_name: '',
                                                hash_password: '',
                                                auth_token: '',
                                                user_id: profile?.userID || '',
                                            },
                                        };
                                        const { user } = await EndpointsModule.CreateUser(req);
                                        if (user) {
                                            console.log(user.user_id);
                                            setUser(user);
                                            storeUser(user.user_id, '');
                                            storeFBAccessToken(user.facebook_access_token);
                                            nav.navigate(SCREENS.HOME);
                                        }
                                        storeProfileURI(profile?.imageURL || '');
                                    } catch (e) {
                                        setState(false);
                                        console.error(e);
                                    }
                                });
                            }
                        });
                    }
                },
                (error) => {
                    setState(false);
                    console.error(error);
                }
            );
        return (
            <TouchableOpacity style={[styles.fields, styles.facebook]} onPress={onFBLogIn}>
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
                autoCapitalize="none"
                size="large"
                status={state ? 'basic' : 'danger'}
                value={username}
                style={styles.fields}
                onChangeText={setUsername}
            />
            <Input
                placeholder="Password"
                size="large"
                autoCapitalize="none"
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

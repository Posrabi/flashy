import { useNavigation } from '@react-navigation/native';
import { NativeStackNavigationProp } from '@react-navigation/native-stack';
import { Button, Input, Icon, Text } from '@ui-kitten/components';
import React from 'react';
import { StyleSheet, TouchableOpacity } from 'react-native';
import { SafeAreaView } from 'react-native';
import { useSetRecoilState } from 'recoil';
import EndpointsModule from '../api/users';
import { StackParams } from '../nav';
import { currentUser } from '../state/user';
import { CreateUserRequest } from '../types';
import { SCREENS } from './constants';

type SignUpScreenProp = NativeStackNavigationProp<StackParams, SCREENS.SIGN_UP>;

export const SignUp = (): JSX.Element => {
    const [username, setUsername] = React.useState('');
    const [password, setPassword] = React.useState('');
    const [secure, setSecure] = React.useState(true);
    const [name, setName] = React.useState('');
    const [email, setEmail] = React.useState('');
    const [status, setStatus] = React.useState(true);
    const setUser = useSetRecoilState(currentUser);
    const nav = useNavigation<SignUpScreenProp>();

    const EyeIcon = (props: any): JSX.Element => (
        <TouchableOpacity onPress={() => setSecure(!secure)}>
            <Icon {...props} name={!secure ? 'eye' : 'eye-off'} />
        </TouchableOpacity>
    );
    const onRegister = async (): Promise<void> => {
        const req: CreateUserRequest = {
            user: {
                user_name: username,
                hash_password: password,
                name: name,
                email: email,
                auth_token: '',
                user_id: '',
            },
        };
        try {
            const { user } = await EndpointsModule.CreateUser(req);
            if (user) {
                setUser(user);
                // nav.navigate(SCREENS.HOME);
            } else {
                setStatus(false);
                throw new Error('create user endpoint does not return user');
            }
        } catch (e: unknown) {
            setStatus(false);
            console.error(e);
        }
    };

    return (
        <SafeAreaView>
            <Input
                size="medium"
                style={styles.fields}
                label="Name"
                placeholder="The name that will be display to your friends."
                onChangeText={setName}
                value={name}
            />
            <Input
                style={styles.fields}
                label="Email"
                placeholder="Your email."
                onChangeText={setEmail}
                value={email}
            />
            <Input
                style={styles.fields}
                label="Username"
                placeholder="Your username to sign in."
                onChangeText={setUsername}
                value={username}
            />
            <Input
                style={styles.fields}
                label="Password"
                placeholder="Your password."
                secureTextEntry={secure}
                onChangeText={setPassword}
                value={password}
                accessoryRight={EyeIcon}
            />
            <Button
                status={status ? 'primary' : 'danger'}
                style={styles.fields}
                onPress={onRegister}>
                Register Now
            </Button>
            {!status && (
                <Text style={[styles.fields, styles.errorText]} status="danger">
                    An error has occurred, please try again later.
                </Text>
            )}
        </SafeAreaView>
    );
};

const styles = StyleSheet.create({
    fields: {
        margin: 10,
        lineHeight: 30,
    },
    errorText: {
        textAlign: 'center',
    },
});

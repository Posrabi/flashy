import { useNavigation } from '@react-navigation/native';
import { NativeStackNavigationProp } from '@react-navigation/native-stack';
import { Button } from '@ui-kitten/components';
import React from 'react';
import { SafeAreaView, StyleSheet } from 'react-native';
import { useSetRecoilState } from 'recoil';
import { defaultUser } from '../api/defaults';
import { StackParams } from '../nav';
import { currentUser } from '../state/user';
import { SCREENS } from './constants';

type HomeScreenProps = NativeStackNavigationProp<StackParams, SCREENS.HOME>;
// TODO: home profile, friends, leaderboards
export const Home = (): JSX.Element => {
    const nav = useNavigation<HomeScreenProps>();
    const setUser = useSetRecoilState(currentUser);
    return (
        <SafeAreaView>
            <Button style={styles.button} onPress={() => nav.navigate(SCREENS.LEARN)}>
                Start learning new words now!
            </Button>
            <Button
                status="danger"
                style={styles.button}
                onPress={() => {
                    setUser(defaultUser());
                    nav.goBack();
                }}>
                Sign out
            </Button>
        </SafeAreaView>
    );
};

const styles = StyleSheet.create({
    button: {
        margin: 10,
    },
});

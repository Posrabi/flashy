import { useNavigation } from '@react-navigation/native';
import { NativeStackNavigationProp } from '@react-navigation/native-stack';
import { Button, Text, Icon } from '@ui-kitten/components';
import React from 'react';
import {
    SafeAreaView,
    StyleSheet,
    View,
    TouchableOpacity,
    ScrollView,
    NativeSyntheticEvent,
    NativeScrollEvent,
} from 'react-native';
import LinearGradient from 'react-native-linear-gradient';
import { useSetRecoilState } from 'recoil';
import { defaultUser } from '../api/defaults';
import { StackParams } from '../nav';
import { cardsCount, currentUser } from '../state/user';
import { SCREENS } from './constants';

type HomeScreenProps = NativeStackNavigationProp<StackParams, SCREENS.HOME>;
interface CardsCountModalProps {
    isVisible: Boolean;
}
// TODO: home profile, friends, leaderboards
export const Home = (): JSX.Element => {
    const nav = useNavigation<HomeScreenProps>();
    const setUser = useSetRecoilState(currentUser);
    const [renderCardsCountModal, setRenderCardsCountModal] = React.useState(false);
    const setCardsCount = useSetRecoilState(cardsCount);

    const onCardsScroll = (event: NativeSyntheticEvent<NativeScrollEvent>): void => {
        const cardsScrollIndex = Math.round(event.nativeEvent.contentOffset.y / 50);
        setCardsCount(cardsScrollIndex + 1);
    };

    const CardsCountModal = (props: CardsCountModalProps): JSX.Element => {
        if (!props.isVisible) return <></>;
        return (
            <View style={styles.modalContainer}>
                <View style={styles.modal}>
                    <TouchableOpacity
                        style={styles.modalClose}
                        onPress={() => setRenderCardsCountModal(false)}>
                        <Icon name="close-circle-outline" width={25} height={25} fill="black" />
                    </TouchableOpacity>
                    <Text style={styles.modalText}>Select how many cards you will need.</Text>
                    <View style={styles.cardsScroll}>
                        <ScrollView
                            onScroll={onCardsScroll}
                            showsVerticalScrollIndicator={false}
                            snapToInterval={50}
                            decelerationRate="fast">
                            {(() => {
                                let arr = [];
                                for (let i = 0; i <= 50; i++) {
                                    arr.push(
                                        <Text style={styles.cardsText} key={i}>
                                            {i}
                                        </Text>
                                    );
                                }
                                return arr;
                            })()}
                        </ScrollView>
                        <LinearGradient
                            colors={['rgba(255, 255, 255, 1)', 'rgba(255, 255, 255, 0)']}
                            style={styles.topGradient}
                        />
                        <LinearGradient
                            colors={['rgba(255, 255, 255, 0)', 'rgba(255, 255, 255, 1)']}
                            style={styles.bottomGradient}
                        />
                    </View>
                    <Button
                        status="success"
                        style={styles.confirmButton}
                        children={() => <Text style={styles.confirmText}>Confirm</Text>}
                        onPress={() => {
                            nav.navigate(SCREENS.LEARN);
                            setRenderCardsCountModal(false);
                        }}
                    />
                </View>
            </View>
        );
    };

    return (
        <SafeAreaView style={styles.container}>
            <Button style={styles.button} onPress={() => setRenderCardsCountModal(true)}>
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
            <CardsCountModal isVisible={renderCardsCountModal} />
        </SafeAreaView>
    );
};

const styles = StyleSheet.create({
    button: {
        margin: 10,
    },
    container: {
        flex: 1,
        display: 'flex',
        justifyContent: 'center',
        alignItems: 'center',
    },
    modalContainer: {
        position: 'absolute',
        backgroundColor: 'rgba(0,0,0,0.5)',
        top: 0,
        right: 0,
        left: 0,
        bottom: 0,
        display: 'flex',
        justifyContent: 'center',
        alignItems: 'center',
    },
    modal: {
        width: '65%',
        height: '45%',
        backgroundColor: '#ffffff',
        borderRadius: 15,
        elevation: 5,
        display: 'flex',
        alignItems: 'center',
    },
    modalText: {
        fontSize: 20,
        margin: 20,
        marginTop: 40,
        textAlign: 'center',
    },
    modalClose: {
        position: 'absolute',
        top: 0,
        left: 0,
        margin: 10,
    },
    cardsText: {
        textAlign: 'center',
        height: 50,
        fontSize: 20,
        fontWeight: 'bold',
    },
    cardsScroll: {
        height: 130,
        width: '80%',
        borderWidth: 2,
        borderColor: 'black',
    },
    topGradient: {
        position: 'absolute',
        width: '100%',
        top: 0,
        height: 70,
    },
    bottomGradient: {
        position: 'absolute',
        width: '100%',
        bottom: 0,
        height: 70,
    },
    confirmButton: {
        margin: 20,
        width: 125,
    },
    confirmText: {
        fontWeight: 'bold',
        fontSize: 18,
        color: '#ffffff',
    },
});

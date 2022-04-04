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
    Image,
    FlatList,
} from 'react-native';
import LinearGradient from 'react-native-linear-gradient';
import { useSetRecoilState, useRecoilState } from 'recoil';
import { defaultUser } from '../api/defaults';
import { LoadingScreen } from '../components/Loading';
import { StackParams } from '../nav';
import { useGetPhraseHistory } from '../state/phrase';
import { cardsCount, clearUser, currentUser } from '../state/user';
import { SCREENS } from './constants';

type HomeScreenProps = NativeStackNavigationProp<StackParams, SCREENS.HOME>;
interface CardsCountModalProps {
    isVisible: Boolean;
}
const miliToNano = 1000;
// TODO: friends, leaderboards
export const Home = (): JSX.Element => {
    const nav = useNavigation<HomeScreenProps>();
    const [user, setUser] = useRecoilState(currentUser);
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
                            decelerationRate="fast"
                            alwaysBounceVertical={true}
                            bounces={false}>
                            {(() => {
                                let arr = [];
                                for (let i = 0; i <= 50; i++) {
                                    arr.push(
                                        <Text style={styles.cardsText} key={i}>
                                            {i}
                                        </Text>
                                    );
                                }
                                arr.push(
                                    <Text style={styles.cardsText} key="end">
                                        {' '}
                                    </Text>
                                );
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

    const UserProfile = (): JSX.Element => {
        return (
            <View style={userStyles.profileContainer}>
                <View style={userStyles.pictureContainer}>
                    <Image
                        style={userStyles.profilePicture}
                        source={{
                            uri: 'https://scontent.fyhz1-1.fna.fbcdn.net/v/t1.6435-9/127454111_1288036738235233_8547489606110234618_n.jpg?_nc_cat=110&ccb=1-5&_nc_sid=09cbfe&_nc_ohc=E0GsZQPxtNsAX8vARAL&_nc_ht=scontent.fyhz1-1.fna&oh=00_AT9SyNoFjCxZj5wAtss6YP9YowxBm2UtSWAlTYU1xygJDw&oe=6261C468',
                        }}
                    />
                </View>

                <View style={userStyles.info}>
                    <Text style={userStyles.infoText}>{user.name}</Text>
                    <Text style={userStyles.infoText}>Completed: ... cards this month</Text>
                </View>
            </View>
        );
    };

    const History = (): JSX.Element => {
        const { isLoading, isError, error, data, isFetching } = useGetPhraseHistory(user.user_id);
        const [expanded, setExpanded] = React.useState(false);
        if (isError) console.error(error);
        if (expanded)
            return (
                <View style={historyStyles.expandedContainer}>
                    <View style={historyStyles.expandedModal}>
                        <TouchableOpacity
                            style={[
                                historyStyles.textContainer,
                                { borderTopLeftRadius: 15, borderTopRightRadius: 15 },
                            ]}
                            onPress={() => setExpanded(false)}>
                            <Text style={historyStyles.text}>History</Text>
                        </TouchableOpacity>
                        <FlatList
                            style={{ width: '100%' }}
                            showsVerticalScrollIndicator={false}
                            data={data}
                            renderItem={({ item }) => (
                                <View
                                    style={{
                                        flexDirection: 'row',
                                        alignItems: 'center',
                                        borderColor: 'black',
                                        margin: 2,
                                        marginVertical: 5,
                                    }}>
                                    <Text
                                        style={{
                                            width: '30%',
                                            marginHorizontal: 5,
                                            textAlign: 'center',
                                            fontWeight: 'bold',
                                        }}>
                                        {item.word}
                                    </Text>
                                    <Text
                                        style={{
                                            width: '40%',
                                            marginHorizontal: 5,
                                        }}>
                                        {item.sentence}
                                    </Text>
                                    <Text style={{ flex: 1, textAlign: 'center' }}>
                                        {new Date(
                                            item.phrase_time * miliToNano
                                        ).toLocaleDateString()}
                                    </Text>
                                </View>
                            )}
                        />
                    </View>
                </View>
            );
        return (
            <View style={historyStyles.container}>
                <TouchableOpacity
                    style={historyStyles.textContainer}
                    onPress={() => setExpanded(true)}>
                    <Text style={historyStyles.text}>History</Text>
                </TouchableOpacity>
                {isLoading || isFetching ? (
                    <LoadingScreen />
                ) : isError ? (
                    <View style={{ flex: 1, justifyContent: 'center' }}>
                        <Text status="danger">
                            An error has occurred while loading your history...
                        </Text>
                    </View>
                ) : (
                    <FlatList
                        style={{ width: '100%' }}
                        showsVerticalScrollIndicator={false}
                        data={data}
                        renderItem={({ item }) => (
                            <View
                                style={{
                                    flexDirection: 'row',
                                }}>
                                <Text
                                    style={{
                                        width: '30%',
                                        marginHorizontal: 5,
                                        textAlign: 'center',
                                        fontWeight: 'bold',
                                    }}>
                                    {item.word}
                                </Text>
                                <Text
                                    style={{
                                        width: '55%',
                                        marginHorizontal: 5,
                                    }}>
                                    {item.sentence}
                                </Text>
                            </View>
                        )}
                    />
                )}
            </View>
        );
    };

    return (
        <SafeAreaView style={styles.container}>
            <UserProfile />
            <History />
            <Button style={styles.button} onPress={() => setRenderCardsCountModal(true)}>
                Start learning now!
            </Button>
            <Button
                status="danger"
                style={styles.button}
                onPress={() => {
                    setUser(defaultUser());
                    clearUser();
                    nav.goBack();
                }}>
                <Icon name="log-out-outline" fill="white" width={25} height={25} />
            </Button>
            <CardsCountModal isVisible={renderCardsCountModal} />
        </SafeAreaView>
    );
};

const userStyles = StyleSheet.create({
    profileContainer: {
        width: '100%',
        flexDirection: 'row',
    },
    profilePicture: {
        width: 100,
        height: 100,
        resizeMode: 'cover',
    },
    pictureContainer: {
        overflow: 'hidden',
        margin: 30,
        height: 100,
        width: 100,
        borderRadius: 50,
        alignItems: 'center',
        justifyContent: 'center',
        elevation: 10,
    },
    info: {
        justifyContent: 'space-evenly',
    },
    infoText: {
        fontSize: 15,
    },
});

const historyStyles = StyleSheet.create({
    container: {
        height: 200,
        borderRadius: 10,
        borderColor: 'black',
        borderWidth: 2,
        width: '90%',
        alignItems: 'center',
        margin: 20,
    },
    textContainer: {
        borderTopLeftRadius: 8,
        borderTopRightRadius: 8,
        height: 35,
        width: '100%',
        justifyContent: 'center',
        alignItems: 'center',
        backgroundColor: '#5fad74',
    },
    text: {
        textAlign: 'center',
        fontWeight: 'bold',
        fontSize: 17,
        margin: 5,
    },
    expandedContainer: {
        position: 'absolute',
        height: '100%',
        width: '100%',
        justifyContent: 'center',
        alignItems: 'center',
        backgroundColor: 'rgba(0, 0, 0, 0.5)',
        zIndex: 1,
    },
    expandedModal: {
        height: '80%',
        width: '90%',
        backgroundColor: 'white',
        elevation: 10,
        borderRadius: 15,
        borderColor: 'black',
    },
});

const styles = StyleSheet.create({
    button: {
        margin: 10,
    },
    container: {
        flex: 1,
        display: 'flex',
        alignItems: 'center',
    },
    modalContainer: {
        position: 'absolute',
        backgroundColor: 'rgba(0,0,0,0.5)',
        width: '100%',
        height: '100%',
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

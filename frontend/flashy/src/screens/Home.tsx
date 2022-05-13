import { useNavigation } from '@react-navigation/native';
import { NativeStackNavigationProp } from '@react-navigation/native-stack';
import { Button, Text, Icon } from '@ui-kitten/components';
import React, { useEffect } from 'react';
import {
    SafeAreaView,
    StyleSheet,
    View,
    TouchableOpacity,
    ScrollView,
    NativeSyntheticEvent,
    NativeScrollEvent,
    FlatList,
    Image,
} from 'react-native';
import LinearGradient from 'react-native-linear-gradient';
import { useSetRecoilState, useRecoilState } from 'recoil';
import { defaultUser } from '../api/defaults';
import EndpointsModule from '../api/users';
import { LoadingScreen } from '../components/Loading';
import { StackParams } from '../nav';
import { useGetPhraseHistory } from '../state/phrase';
import { cardsCount, clearUser, currentUser, getProfileURI, useGetFriends } from '../state/user';
import { SCREENS } from './constants';

type HomeScreenProps = NativeStackNavigationProp<StackParams, SCREENS.HOME>;
interface CardsCountModalProps {
    isVisible: Boolean;
}
const miliToNano = 1000;
// TODO: FB login SDK integration.
export const Home = (): JSX.Element => {
    const [user, setUser] = useRecoilState(currentUser);
    const { isLoading, isError, error, data, isFetching } = useGetPhraseHistory(user.user_id);
    const nav = useNavigation<HomeScreenProps>();
    const [renderCardsCountModal, setRenderCardsCountModal] = React.useState(false);
    const setCardsCount = useSetRecoilState(cardsCount);
    const [profileURI, setProfileURI] = React.useState('');
    const onCardsScroll = (event: NativeSyntheticEvent<NativeScrollEvent>): void => {
        const cardsScrollIndex = Math.round(event.nativeEvent.contentOffset.y / 50);
        setCardsCount(cardsScrollIndex + 1);
    };
    const [friends, setFriends] = React.useState<Record<string, any>>({
        data: [],
        summary: { total_count: 0 },
    });
    useGetFriends((err, res) => {
        if (err) {
            console.log(err);
        } else if (res) {
            console.log(res);
            setFriends(res);
        }
    });

    useEffect(() => {
        (async () => {
            const uri = await getProfileURI();
            setProfileURI(uri);
        })();
        return setProfileURI('');
    }, []);

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
                    <Text style={styles.modalText}>Select the number of cards:</Text>
                    <View style={styles.cardsScroll}>
                        <ScrollView
                            onScroll={onCardsScroll}
                            showsVerticalScrollIndicator={false}
                            snapToInterval={50}
                            decelerationRate="fast"
                            alwaysBounceVertical={true}>
                            {(() => {
                                let arr = [];
                                for (let i = 0; i <= 50; i++) {
                                    arr.push(
                                        <View style={styles.cardsTextContainer} key={i}>
                                            <Text style={styles.cardsText}>{i}</Text>
                                        </View>
                                    );
                                }
                                arr.push(
                                    <View style={styles.cardsTextContainer} key="end">
                                        <Text style={styles.cardsText}> </Text>
                                    </View>
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
                        style={[styles.confirmButton, { backgroundColor: '#5fad74' }]}
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
                        source={
                            profileURI
                                ? {
                                      uri: profileURI,
                                  }
                                : require('../assets/blank-profile-picture-973460.png')
                        }
                    />
                </View>

                <View style={userStyles.info}>
                    <Text style={userStyles.infoText}>{user.name}</Text>
                    <Text style={userStyles.infoText}>
                        Completed:{' '}
                        {
                            // @ts-ignore
                            isError || isFetching || isFetching ? '...' : data.length
                        }{' '}
                        cards this month
                    </Text>
                </View>
            </View>
        );
    };

    const History = (): JSX.Element => {
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
                                            fontSize: 16,
                                        }}>
                                        {item.word}
                                    </Text>
                                    <Text
                                        style={{
                                            width: '35%',
                                            marginHorizontal: 5,
                                        }}>
                                        {item.sentence}
                                    </Text>
                                    {item.correct ? (
                                        <Icon
                                            name="checkmark-outline"
                                            width={25}
                                            height={25}
                                            fill="green"
                                        />
                                    ) : (
                                        <Icon
                                            name="close-outline"
                                            width={25}
                                            height={25}
                                            fill="red"
                                        />
                                    )}
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
                                    alignItems: 'center',
                                    justifyContent: 'center',
                                }}>
                                <Text
                                    style={{
                                        width: '25%',
                                        marginHorizontal: 5,
                                        textAlign: 'center',
                                        fontWeight: 'bold',
                                    }}>
                                    {item.word}
                                </Text>
                                <Text
                                    style={{
                                        width: '65%',
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

    const Leaderboard = (): JSX.Element => {
        const [expanded, setExpanded] = React.useState(false);
        if (expanded)
            return (
                <View style={historyStyles.expandedContainer}>
                    <View style={historyStyles.expandedModal}>
                        <TouchableOpacity
                            style={[
                                historyStyles.textContainer,
                                {
                                    borderTopLeftRadius: 15,
                                    borderTopRightRadius: 15,
                                    backgroundColor: '#f59b42',
                                },
                            ]}
                            onPress={() => setExpanded(false)}>
                            <Text style={historyStyles.text}>Leaderboard</Text>
                        </TouchableOpacity>
                        <FlatList
                            style={{ width: '100%' }}
                            showsVerticalScrollIndicator={false}
                            data={friends.data}
                            renderItem={({ item }) => (
                                <View
                                    style={{
                                        flexDirection: 'row',
                                        alignItems: 'center',
                                        borderColor: 'black',
                                        margin: 2,
                                        marginVertical: 5,
                                    }}>
                                    {/* <Text
                                        style={{
                                            width: '30%',
                                            marginHorizontal: 5,
                                            textAlign: 'center',
                                            fontWeight: 'bold',
                                            fontSize: 16,
                                        }}>
                                        {item.word}
                                    </Text>
                                    <Text
                                        style={{
                                            width: '35%',
                                            marginHorizontal: 5,
                                        }}>
                                        {item.sentence}
                                    </Text>
                                    {item.correct ? (
                                        <Icon
                                            name="close-outline"
                                            width={25}
                                            height={25}
                                            fill="red"
                                        />
                                    ) : (
                                        <Icon
                                            name="checkmark-outline"
                                            width={25}
                                            height={25}
                                            fill="green"
                                        />
                                    )}
                                    <Text style={{ flex: 1, textAlign: 'center' }}>
                                        {new Date(
                                            item.phrase_time * miliToNano
                                        ).toLocaleDateString()}
                                    </Text> */}
                                </View>
                            )}
                        />
                    </View>
                </View>
            );
        return (
            <View style={historyStyles.container}>
                <TouchableOpacity
                    style={[historyStyles.textContainer, { backgroundColor: '#f59b42' }]}
                    onPress={() => setExpanded(true)}>
                    <Text style={historyStyles.text}>Leaderboard</Text>
                </TouchableOpacity>
                {/* {isLoading || isFetching ? (
                    <LoadingScreen />
                ) : isError ? (
                    <View style={{ flex: 1, justifyContent: 'center' }}>
                        <Text status="danger">
                            An error has occurred while loading friends...
                        </Text>
                    </View>
                ) : ( */}
                <FlatList
                    style={{ width: '100%' }}
                    showsVerticalScrollIndicator={false}
                    data={friends.data}
                    renderItem={({ item }) => (
                        <View
                            style={{
                                flexDirection: 'row',
                            }}>
                            {/* <Text
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
                                </Text> */}
                        </View>
                    )}
                />
                {/* )} */}
            </View>
        );
    };

    return (
        <SafeAreaView style={styles.container}>
            <UserProfile />
            <History />
            <Leaderboard />
            <Button style={styles.button} onPress={() => setRenderCardsCountModal(true)}>
                Start now!
            </Button>
            <Button
                status="danger"
                style={styles.button}
                onPress={async () => {
                    try {
                        if (!user.user_id) return nav.goBack();
                        await EndpointsModule.LogOut({
                            user_id: user.user_id,
                        });
                        setUser(defaultUser());
                        clearUser();
                        nav.goBack();
                    } catch (e) {
                        console.error(e);
                    }
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
        width: 125,
        height: 125,
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
        margin: 15,
        marginBottom: 5,
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
        margin: 15,
        marginBottom: 5,
        width: 200,
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
        height: '47%',
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
    cardsTextContainer: {
        alignItems: 'center',
        justifyContent: 'center',
        height: 50,
    },
    cardsText: {
        textAlign: 'center',
        fontSize: 20,
        fontWeight: 'bold',
    },
    cardsScroll: {
        height: 150,
        width: '80%',
    },
    topGradient: {
        position: 'absolute',
        width: '100%',
        top: 0,
        height: 80,
    },
    bottomGradient: {
        position: 'absolute',
        width: '100%',
        bottom: 0,
        height: 80,
    },
    confirmButton: {
        margin: 20,
        width: 125,
        borderWidth: 0,
    },
    confirmText: {
        fontWeight: 'bold',
        fontSize: 18,
        color: '#ffffff',
    },
});

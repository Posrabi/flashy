import { useNavigation } from '@react-navigation/native';
import { NativeStackNavigationProp } from '@react-navigation/native-stack';
import { Input, Text, Icon, Button } from '@ui-kitten/components';
import React, { useRef, useMemo } from 'react';
import {
    SafeAreaView,
    StyleSheet,
    View,
    TouchableOpacity,
    PanResponder,
    Animated,
    GestureResponderHandlers,
} from 'react-native';
import { useRecoilState } from 'recoil';
import { StackParams } from '../nav';
import { cardsCount } from '../state/user';
import { SCREENS } from './constants';

interface HelpModalProps {
    isVisible: boolean;
}

interface BackModalProps {
    isVisible: boolean;
}

type LearnScreenProps = NativeStackNavigationProp<StackParams, SCREENS.LEARN>;

export const Learn = (): JSX.Element => {
    const [help, setHelp] = React.useState(false);
    const [back, setBack] = React.useState(false);
    const [cardCount, setCardCount] = useRecoilState(cardsCount);
    const nav = useNavigation<LearnScreenProps>();
    const pan = useRef(new Animated.ValueXY()).current;
    const panResponder = useMemo(
        () =>
            PanResponder.create({
                onMoveShouldSetPanResponder: () => true,
                onPanResponderMove: Animated.event([null, { dx: pan.x, dy: pan.y }], {
                    useNativeDriver: false,
                }),
                onPanResponderRelease: (_, gestureState) => {
                    if (gestureState.dy < -150) {
                        Animated.spring(pan, {
                            toValue: { x: gestureState.dx, y: -525 },
                            useNativeDriver: false,
                        }).start(() => {
                            setCardCount(cardCount - 1);
                            pan.setValue({ x: 0, y: 0 });
                        });
                    } else
                        Animated.spring(pan, {
                            toValue: { x: 0, y: 0 },
                            useNativeDriver: false,
                        }).start();
                },
            }),
        [cardCount]
    );

    const HelpModal = useRef((props: HelpModalProps): JSX.Element => {
        if (!props.isVisible) return <></>;
        return (
            <View style={styles.modalContainer}>
                <View style={styles.modal}>
                    <TouchableOpacity style={styles.modalClose} onPress={() => setHelp(false)}>
                        <Icon name="close-circle-outline" width={25} height={25} fill="black" />
                    </TouchableOpacity>
                    <Text style={styles.modalText}>
                        Type into the bottom text box the word you want to learn, then the top text
                        box the sentence.{'\n\n'}To move on to the next word, swipe the card up.
                    </Text>
                </View>
            </View>
        );
    }).current;

    const BackModal = useRef((props: BackModalProps): JSX.Element => {
        if (!props.isVisible) return <></>;
        return (
            <View style={styles.modalContainer}>
                <View style={styles.modal}>
                    <Text style={styles.modalText}>
                        Are you sure you want to go back? All of your progress will be lost.
                    </Text>
                    <View style={styles.modalButtonContainer}>
                        <Button
                            style={styles.backConfirm}
                            status="danger"
                            children={() => <Text style={styles.backConfirmText}>Yes</Text>}
                            onPress={() => {
                                nav.goBack();
                                setCardCount(1);
                            }}
                        />
                        <Button
                            style={styles.backConfirm}
                            children={() => <Text style={styles.backConfirmText}>No</Text>}
                            onPress={() => setBack(false)}
                        />
                    </View>
                </View>
            </View>
        );
    }).current;

    const CongratsModal = useRef((): JSX.Element => {
        return (
            <View style={styles.modalContainer}>
                <View style={styles.modal}>
                    <Text style={styles.modalText}>
                        Congratulations!{'\n\n'}You've just completed {cardCount} cards.
                    </Text>
                    <View style={styles.modalButtonContainer}>
                        <Button
                            style={{ width: 150 }}
                            children={() => (
                                <Text style={styles.backConfirmText} status="info">
                                    Return to home
                                </Text>
                            )}
                            onPress={() => {
                                nav.goBack();
                                setCardCount(1);
                            }}
                        />
                    </View>
                </View>
            </View>
        );
    }).current;

    return (
        <SafeAreaView style={styles.cardContainer}>
            <TouchableOpacity style={styles.back} onPress={() => setBack(true)}>
                <Icon name="backspace-outline" width={30} height={30} fill="black" />
            </TouchableOpacity>
            <TouchableOpacity style={styles.help} onPress={() => setHelp(true)}>
                <Icon name="question-mark-outline" width={30} height={30} fill="black" />
            </TouchableOpacity>
            {(() => {
                const arr = [];
                for (let i = 0; i < cardCount; i++) {
                    if (i == cardCount - 1) {
                        arr.push(
                            <Card
                                key={i}
                                top={true}
                                panHandler={panResponder.panHandlers}
                                pan={pan}
                            />
                        );
                    } else arr.push(<Card top={false} key={i} />);
                }
                return arr;
            })()}
            <HelpModal isVisible={help} />
            <BackModal isVisible={back} />
            {cardCount === 0 ? <CongratsModal /> : <></>}
        </SafeAreaView>
    );
};

interface CardProps {
    top: Boolean;
    panHandler?: GestureResponderHandlers;
    pan?: any;
}

const Card = (props: CardProps): JSX.Element => {
    const [sentence, setSentence] = React.useState('');
    const [word, setWord] = React.useState('');
    return (
        <Animated.View
            style={[
                styles.card,
                props.top
                    ? {
                          transform: [{ translateX: props.pan.x }, { translateY: props.pan.y }],
                      }
                    : {},
            ]}
            {...props.panHandler}>
            <Input
                value={sentence}
                onChangeText={setSentence}
                placeholder="Input your sentence here"
                style={styles.input}
                multiline={true}
                textStyle={styles.inputText}
            />
            <Input
                value={word}
                onChangeText={setWord}
                placeholder="Insert your word here"
                style={styles.input}
                textAlign="center"
                size="large"
            />
        </Animated.View>
    );
};

const styles = StyleSheet.create({
    cardContainer: {
        position: 'relative',
        height: '100%',
        width: '100%',
        display: 'flex',
        justifyContent: 'center',
        alignItems: 'center',
    },
    card: {
        position: 'absolute',
        backgroundColor: '#5fad74',
        width: 300,
        height: 300,
        display: 'flex',
        alignItems: 'center',
        elevation: 5,
        borderRadius: 25,
    },
    cardText: {
        color: '#ffffff',
        fontWeight: 'bold',
        fontSize: 25,
    },
    input: {
        width: '80%',
        margin: 20,
    },
    inputText: {
        minHeight: '60%',
    },
    help: {
        position: 'absolute',
        top: 0,
        right: 0,
        margin: 10,
        width: 40,
        height: 40,
        backgroundColor: '#ffffff',
        borderRadius: 20,
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
        width: '60%',
        height: '40%',
        backgroundColor: '#ffffff',
        borderRadius: 15,
        elevation: 5,
        display: 'flex',
        justifyContent: 'center',
        alignItems: 'center',
    },
    modalClose: {
        position: 'absolute',
        top: 0,
        left: 0,
        margin: 10,
    },
    modalText: {
        fontSize: 20,
        margin: 20,
        textAlign: 'center',
    },
    back: {
        position: 'absolute',
        top: 0,
        left: 0,
        margin: 10,
        width: 40,
        height: 40,
        backgroundColor: '#ffffff',
        borderRadius: 20,
        display: 'flex',
        justifyContent: 'center',
        alignItems: 'center',
    },
    modalButtonContainer: {
        width: '100%',
        justifyContent: 'space-evenly',
        display: 'flex',
        flexDirection: 'row',
        alignItems: 'center',
    },
    backConfirm: {
        width: 85,
    },
    backConfirmText: {
        textAlign: 'center',
        color: '#ffffff',
        fontWeight: 'bold',
    },
});

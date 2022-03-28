import { Phrase, User } from '../types';
import 'react-native-get-random-values';
import { v4 as uuidv4 } from 'uuid';

export const defaultUser = (): User => {
    return {
        user_name: 'guest',
        user_id: '',
        hash_password: '',
        name: 'Guest',
        email: '',
        auth_token: '',
    };
};

export const defaultPhrase: Phrase = {
    user_id: '',
    word: '',
    sentence: '',
    phrase_time: 0,
};

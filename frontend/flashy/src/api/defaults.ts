import { User } from '../types';
import 'react-native-get-random-values';
import { v4 as uuidv4 } from 'uuid';

export const defaultUser = (): User => {
    return {
        user_name: 'guest',
        user_id: uuidv4(),
        hash_password: '',
        name: 'Guest',
        email: '',
        auth_token: '',
    };
};

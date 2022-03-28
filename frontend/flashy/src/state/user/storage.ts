import EncryptedStorage from 'react-native-encrypted-storage';

export const storeUser = async (userID: string, authToken: string) => {
    try {
        await EncryptedStorage.setItem('user_id', userID);
        await EncryptedStorage.setItem('auth_token', authToken);
    } catch (error) {
        console.error(error);
    }
};

export const getUserIDFromStorage = async (): Promise<any> => {
    try {
        const auth_token = await EncryptedStorage.getItem('user_id');
        return auth_token;
    } catch (e) {
        console.error(e);
        return;
    }
};

export const getAuthTokenFromStorage = async (): Promise<any> => {
    try {
        const user_id = await EncryptedStorage.getItem('auth_token');
        return user_id;
    } catch (error) {
        console.error(error);
    }
};

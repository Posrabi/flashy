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

export const clearUser = async (): Promise<any> => {
    try {
        await EncryptedStorage.clear();
    } catch (error) {
        console.error(error);
    }
};

export const storeFBAccessToken = async (token: string): Promise<any> => {
    try {
        await EncryptedStorage.setItem('facebook_access_token', token);
    } catch (error) {
        console.error(error);
    }
};

export const getFBAccessToken = async (): Promise<any> => {
    try {
        const token = await EncryptedStorage.getItem('facebook_access_token');
        return token;
    } catch (error) {
        console.error(error);
    }
};

export const storeProfileURI = async (uri: string): Promise<any> => {
    try {
        await EncryptedStorage.setItem('profile_uri', uri);
    } catch (e) {
        console.error(e);
    }
};

export const getProfileURI = async (): Promise<any> => {
    try {
        const uri = await EncryptedStorage.getItem('profile_uri');
        return uri;
    } catch (e) {
        console.error(e);
    }
};

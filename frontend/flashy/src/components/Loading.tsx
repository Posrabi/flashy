import { ActivityIndicator, StyleSheet, View } from 'react-native';
import React from 'react';

export const LoadingScreen = (): JSX.Element => {
    return (
        <View style={styles.container}>
            <ActivityIndicator size="large" />
        </View>
    );
};

const styles = StyleSheet.create({
    container: {
        flex: 1,
        display: 'flex',
        alignItems: 'center',
        justifyContent: 'center',
    },
});

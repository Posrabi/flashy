import {
    GraphRequest,
    GraphRequestManager,
    AccessToken,
    GraphRequestCallback,
} from 'react-native-fbsdk-next';
import { QueryFunctionContext, useQuery, UseQueryResult } from 'react-query';

const getFriends = async (ctx: QueryFunctionContext): Promise<void> => {
    const [, callback] = ctx.queryKey as [string, GraphRequestCallback];
    const tokenData = (await AccessToken.getCurrentAccessToken()) || { accessToken: '' };
    const req = new GraphRequest(
        `me/friends`,
        {
            httpMethod: 'GET',
            accessToken: tokenData.accessToken.toString(),
        },
        callback
    );
    return new GraphRequestManager().addRequest(req).start();
};

export const useGetFriends = (callback: GraphRequestCallback): UseQueryResult<void> => {
    return useQuery<void, Error>(['getFriends', callback], getFriends);
};

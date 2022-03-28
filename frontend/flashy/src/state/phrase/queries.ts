import { QueryFunctionContext, useQuery, UseQueryResult } from 'react-query';
import EndpointsModule from '../../api/users';
import { GetPhrasesRequest, Phrase } from '../../types';

const days = 1000 * 3600 * 24;

const getPhrasesHistory = async (ctx: QueryFunctionContext): Promise<Phrase[]> => {
    const [, userID] = ctx.queryKey as [string, string];
    var end = new Date();
    const request: GetPhrasesRequest = {
        user_id: userID,
        start: end.getTime() - 30 * days,
        end: end.getTime(),
    };
    return (await EndpointsModule.GetPhrases(request)).phrases;
};

export const useGetPhraseHistory = (userID: string): UseQueryResult<Phrase[], Error> => {
    return useQuery<Phrase[], Error>(['getPhrasesHistory', userID], getPhrasesHistory, {
        refetchOnMount: false,
    });
};

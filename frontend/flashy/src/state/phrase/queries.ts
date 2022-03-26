import { QueryFunctionContext, useQuery, UseQueryResult } from 'react-query';
import EndpointsModule from '../../api/users';
import { GetPhrasesRequest, Phrase } from '../../types';

const getPhrasesHistory = async (ctx: QueryFunctionContext): Promise<Phrase[]> => {
    const [, userID] = ctx.queryKey as [string, string];
    var end = new Date(),
        start = end;
    start.setDate(start.getDate() - 30);
    const request: GetPhrasesRequest = {
        user_id: userID,
        start: start.getTime(),
        end: end.getTime(),
    };
    return (await EndpointsModule.GetPhrases(request)).phrases;
};

export const useGetPhraseHistory = (userID: string): UseQueryResult<Phrase[], Error> => {
    return useQuery<Phrase[], Error>(['getPhrasesHistory', userID], getPhrasesHistory);
};

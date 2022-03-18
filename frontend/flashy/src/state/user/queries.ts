import EndpointsModule from '../../api/users';
import { GetUserRequest, LogInRequest, User } from '../../types';
import { useQuery, QueryFunctionContext, UseQueryOptions, QueryObserverResult } from 'react-query';
import { defaultUser } from '../../api/defaults';

const getUser = async (ctx: QueryFunctionContext): Promise<User> => {
    const [, userID] = ctx.queryKey as [string, string];
    const req: GetUserRequest = { user_id: userID };
    const resp = await EndpointsModule.GetUser(req);

    return resp.user || defaultUser();
};

export const useGetUser = (
    userID: string,
    options?: UseQueryOptions<User, Error, User, string[]>
): QueryObserverResult<User, Error> => {
    return useQuery(['cache_user', userID], getUser, options);
};

const logIn = async (ctx: QueryFunctionContext): Promise<User> => {
    const [, username, password] = ctx.queryKey as [string, string, string];
    const req: LogInRequest = {
        user_name: username,
        hash_password: password,
    };
    const resp = await EndpointsModule.LogIn(req);

    return resp.user || defaultUser();
};

export const useLogIn = (
    username: string,
    password: string,
    options?: UseQueryOptions<User, Error, User, string[]>
): QueryObserverResult<User, Error> => {
    return useQuery(['log_in', username, password], logIn, options);
};

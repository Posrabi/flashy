import { useMutation, UseMutationOptions, UseMutationResult } from 'react-query';
import EndpointsModule from '../../api/users';
import { UpdateUserRequest, UpdateUserResponse } from '../../types';

const updateUser = async (req: UpdateUserRequest): Promise<UpdateUserResponse> => {
    return await EndpointsModule.UpdateUser(req);
};

export const useUpdateUser = (
    req: UpdateUserRequest,
    options?: UseMutationOptions<UpdateUserResponse, Error, UpdateUserRequest>
): UseMutationResult<UpdateUserResponse, Error, UpdateUserRequest> => {
    return useMutation(() => updateUser(req), options);
};

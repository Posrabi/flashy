import { useMutation, UseMutationOptions, UseMutationResult } from 'react-query';
import EndpointsModule from '../../api/users';
import {
    CreateUserRequest,
    CreateUserResponse,
    UpdateUserRequest,
    UpdateUserResponse,
    DeleteUserRequest,
    DeleteUserResponse,
    LogOutRequest,
    LogOutResponse,
} from '../../types';

const createUser = async (req: CreateUserRequest): Promise<CreateUserResponse> => {
    return await EndpointsModule.CreateUser(req);
};

export const useCreateUser = (
    options?: UseMutationOptions<CreateUserResponse, Error, CreateUserRequest>
): UseMutationResult<CreateUserResponse, Error, CreateUserRequest> => {
    return useMutation(createUser, options);
};

const updateUser = async (req: UpdateUserRequest): Promise<UpdateUserResponse> => {
    return await EndpointsModule.UpdateUser(req);
};

export const useUpdateUser = (
    options?: UseMutationOptions<UpdateUserResponse, Error, UpdateUserRequest>
): UseMutationResult<UpdateUserResponse, Error, UpdateUserRequest> => {
    return useMutation(updateUser, options);
};

const deleteUser = async (req: DeleteUserRequest): Promise<DeleteUserResponse> => {
    return await EndpointsModule.DeleteUser(req);
};

export const useDeleteUser = (
    options?: UseMutationOptions<DeleteUserResponse, Error, DeleteUserRequest>
): UseMutationResult<DeleteUserResponse, Error, DeleteUserRequest> => {
    return useMutation(deleteUser, options);
};

const logOut = async (req: LogOutRequest): Promise<LogOutResponse> => {
    return await EndpointsModule.LogOut(req);
};

export const useLogout = (
    options?: UseMutationOptions<LogOutResponse, Error, LogOutRequest>
): UseMutationResult<LogOutResponse, Error, LogOutRequest> => {
    return useMutation(logOut, options);
};

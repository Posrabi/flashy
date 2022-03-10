/* eslint-disable */
export const protobufPackage = "users.proto";

export interface User {
  user_name: string;
  hash_password: string;
  name: string;
  email: string;
  phone_number: string;
  auth_token: string;
  user_id: string;
}

export interface CreateUserRequest {
  user: User | undefined;
}

export interface CreateUserResponse {
  user: User | undefined;
}

export interface GetUserRequest {
  user_id: string;
}

export interface GetUserResponse {
  user: User | undefined;
}

export interface UpdateUserRequest {
  user: User | undefined;
}

export interface UpdateUserResponse {
  response: string;
}

export interface DeleteUserRequest {
  user_id: string;
  hash_password: string;
}

export interface DeleteUserResponse {
  response: string;
}

export interface LogInRequest {
  user_name: string;
  hash_password: string;
}

export interface LogInResponse {
  user: User | undefined;
}

export interface LogOutRequest {
  user_id: string;
}

export interface LogOutResponse {
  response: string;
}

export interface UsersAPI {
  CreateUser(request: CreateUserRequest): Promise<CreateUserResponse>;
  GetUser(request: GetUserRequest): Promise<GetUserResponse>;
  UpdateUser(request: UpdateUserRequest): Promise<UpdateUserResponse>;
  DeleteUser(request: DeleteUserRequest): Promise<DeleteUserResponse>;
  LogIn(request: LogInRequest): Promise<LogInResponse>;
  LogOut(request: LogOutRequest): Promise<LogOutResponse>;
}

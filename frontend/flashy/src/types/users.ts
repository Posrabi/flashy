/* eslint-disable */
export const protobufPackage = "users.proto";

export interface User {
  user_name: string;
  hash_password: string;
  name: string;
  email: string;
  auth_token: string;
  facebook_access_token: string;
  user_id: string;
}

export interface Phrase {
  user_id: string;
  word: string;
  sentence: string;
  phrase_time: number;
  correct: boolean;
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

export interface CreatePhraseRequest {
  phrase: Phrase | undefined;
}

export interface CreatePhraseResponse {
  response: string;
}

export interface GetPhrasesRequest {
  user_id: string;
  start: number;
  end: number;
}

export interface GetPhrasesResponse {
  phrases: Phrase[];
}

export interface DeletePhraseRequest {
  user_id: string;
  phrase_time: number;
}

export interface DeletePhraseResponse {
  response: string;
}

export interface LogInWithFBRequest {
  facebook_access_token: string;
  user_id: string;
}

export interface LogInWithFBResponse {
  user: User | undefined;
}

export interface UsersAPI {
  CreateUser(request: CreateUserRequest): Promise<CreateUserResponse>;
  GetUser(request: GetUserRequest): Promise<GetUserResponse>;
  UpdateUser(request: UpdateUserRequest): Promise<UpdateUserResponse>;
  DeleteUser(request: DeleteUserRequest): Promise<DeleteUserResponse>;
  LogIn(request: LogInRequest): Promise<LogInResponse>;
  LogOut(request: LogOutRequest): Promise<LogOutResponse>;
  CreatePhrase(request: CreatePhraseRequest): Promise<CreatePhraseResponse>;
  GetPhrases(request: GetPhrasesRequest): Promise<GetPhrasesResponse>;
  DeletePhrase(request: DeletePhraseRequest): Promise<DeletePhraseResponse>;
  LogInWithFB(request: LogInWithFBRequest): Promise<LogInWithFBResponse>;
}

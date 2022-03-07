/* eslint-disable */
import { util, configure, Writer, Reader } from "protobufjs/minimal";
import * as Long from "long";

export const protobufPackage = "users.proto";

export interface User {
  userName: string;
  hashPassword: string;
  name: string;
  email: string;
  phoneNumber: string;
  authToken: string;
  userId: string;
}

export interface CreateUserRequest {
  user: User | undefined;
}

export interface CreateUserResponse {
  user: User | undefined;
}

export interface GetUserRequest {
  userId: string;
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
  userId: string;
  hashPassword: string;
}

export interface DeleteUserResponse {
  response: string;
}

export interface LogInRequest {
  userName: string;
  hashPassword: string;
}

export interface LogInResponse {
  user: User | undefined;
}

export interface LogOutRequest {
  userId: string;
}

export interface LogOutResponse {
  response: string;
}

function createBaseUser(): User {
  return {
    userName: "",
    hashPassword: "",
    name: "",
    email: "",
    phoneNumber: "",
    authToken: "",
    userId: "",
  };
}

export const User = {
  encode(message: User, writer: Writer = Writer.create()): Writer {
    if (message.userName !== "") {
      writer.uint32(10).string(message.userName);
    }
    if (message.hashPassword !== "") {
      writer.uint32(18).string(message.hashPassword);
    }
    if (message.name !== "") {
      writer.uint32(26).string(message.name);
    }
    if (message.email !== "") {
      writer.uint32(34).string(message.email);
    }
    if (message.phoneNumber !== "") {
      writer.uint32(42).string(message.phoneNumber);
    }
    if (message.authToken !== "") {
      writer.uint32(50).string(message.authToken);
    }
    if (message.userId !== "") {
      writer.uint32(58).string(message.userId);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): User {
    const reader = input instanceof Reader ? input : new Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseUser();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.userName = reader.string();
          break;
        case 2:
          message.hashPassword = reader.string();
          break;
        case 3:
          message.name = reader.string();
          break;
        case 4:
          message.email = reader.string();
          break;
        case 5:
          message.phoneNumber = reader.string();
          break;
        case 6:
          message.authToken = reader.string();
          break;
        case 7:
          message.userId = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): User {
    return {
      userName: isSet(object.userName) ? String(object.userName) : "",
      hashPassword: isSet(object.hashPassword)
        ? String(object.hashPassword)
        : "",
      name: isSet(object.name) ? String(object.name) : "",
      email: isSet(object.email) ? String(object.email) : "",
      phoneNumber: isSet(object.phoneNumber) ? String(object.phoneNumber) : "",
      authToken: isSet(object.authToken) ? String(object.authToken) : "",
      userId: isSet(object.userId) ? String(object.userId) : "",
    };
  },

  toJSON(message: User): unknown {
    const obj: any = {};
    message.userName !== undefined && (obj.userName = message.userName);
    message.hashPassword !== undefined &&
      (obj.hashPassword = message.hashPassword);
    message.name !== undefined && (obj.name = message.name);
    message.email !== undefined && (obj.email = message.email);
    message.phoneNumber !== undefined &&
      (obj.phoneNumber = message.phoneNumber);
    message.authToken !== undefined && (obj.authToken = message.authToken);
    message.userId !== undefined && (obj.userId = message.userId);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<User>, I>>(object: I): User {
    const message = createBaseUser();
    message.userName = object.userName ?? "";
    message.hashPassword = object.hashPassword ?? "";
    message.name = object.name ?? "";
    message.email = object.email ?? "";
    message.phoneNumber = object.phoneNumber ?? "";
    message.authToken = object.authToken ?? "";
    message.userId = object.userId ?? "";
    return message;
  },
};

function createBaseCreateUserRequest(): CreateUserRequest {
  return { user: undefined };
}

export const CreateUserRequest = {
  encode(message: CreateUserRequest, writer: Writer = Writer.create()): Writer {
    if (message.user !== undefined) {
      User.encode(message.user, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): CreateUserRequest {
    const reader = input instanceof Reader ? input : new Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseCreateUserRequest();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.user = User.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): CreateUserRequest {
    return {
      user: isSet(object.user) ? User.fromJSON(object.user) : undefined,
    };
  },

  toJSON(message: CreateUserRequest): unknown {
    const obj: any = {};
    message.user !== undefined &&
      (obj.user = message.user ? User.toJSON(message.user) : undefined);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<CreateUserRequest>, I>>(
    object: I
  ): CreateUserRequest {
    const message = createBaseCreateUserRequest();
    message.user =
      object.user !== undefined && object.user !== null
        ? User.fromPartial(object.user)
        : undefined;
    return message;
  },
};

function createBaseCreateUserResponse(): CreateUserResponse {
  return { user: undefined };
}

export const CreateUserResponse = {
  encode(
    message: CreateUserResponse,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.user !== undefined) {
      User.encode(message.user, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): CreateUserResponse {
    const reader = input instanceof Reader ? input : new Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseCreateUserResponse();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.user = User.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): CreateUserResponse {
    return {
      user: isSet(object.user) ? User.fromJSON(object.user) : undefined,
    };
  },

  toJSON(message: CreateUserResponse): unknown {
    const obj: any = {};
    message.user !== undefined &&
      (obj.user = message.user ? User.toJSON(message.user) : undefined);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<CreateUserResponse>, I>>(
    object: I
  ): CreateUserResponse {
    const message = createBaseCreateUserResponse();
    message.user =
      object.user !== undefined && object.user !== null
        ? User.fromPartial(object.user)
        : undefined;
    return message;
  },
};

function createBaseGetUserRequest(): GetUserRequest {
  return { userId: "" };
}

export const GetUserRequest = {
  encode(message: GetUserRequest, writer: Writer = Writer.create()): Writer {
    if (message.userId !== "") {
      writer.uint32(10).string(message.userId);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): GetUserRequest {
    const reader = input instanceof Reader ? input : new Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseGetUserRequest();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.userId = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): GetUserRequest {
    return {
      userId: isSet(object.userId) ? String(object.userId) : "",
    };
  },

  toJSON(message: GetUserRequest): unknown {
    const obj: any = {};
    message.userId !== undefined && (obj.userId = message.userId);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<GetUserRequest>, I>>(
    object: I
  ): GetUserRequest {
    const message = createBaseGetUserRequest();
    message.userId = object.userId ?? "";
    return message;
  },
};

function createBaseGetUserResponse(): GetUserResponse {
  return { user: undefined };
}

export const GetUserResponse = {
  encode(message: GetUserResponse, writer: Writer = Writer.create()): Writer {
    if (message.user !== undefined) {
      User.encode(message.user, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): GetUserResponse {
    const reader = input instanceof Reader ? input : new Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseGetUserResponse();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.user = User.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): GetUserResponse {
    return {
      user: isSet(object.user) ? User.fromJSON(object.user) : undefined,
    };
  },

  toJSON(message: GetUserResponse): unknown {
    const obj: any = {};
    message.user !== undefined &&
      (obj.user = message.user ? User.toJSON(message.user) : undefined);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<GetUserResponse>, I>>(
    object: I
  ): GetUserResponse {
    const message = createBaseGetUserResponse();
    message.user =
      object.user !== undefined && object.user !== null
        ? User.fromPartial(object.user)
        : undefined;
    return message;
  },
};

function createBaseUpdateUserRequest(): UpdateUserRequest {
  return { user: undefined };
}

export const UpdateUserRequest = {
  encode(message: UpdateUserRequest, writer: Writer = Writer.create()): Writer {
    if (message.user !== undefined) {
      User.encode(message.user, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): UpdateUserRequest {
    const reader = input instanceof Reader ? input : new Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseUpdateUserRequest();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.user = User.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): UpdateUserRequest {
    return {
      user: isSet(object.user) ? User.fromJSON(object.user) : undefined,
    };
  },

  toJSON(message: UpdateUserRequest): unknown {
    const obj: any = {};
    message.user !== undefined &&
      (obj.user = message.user ? User.toJSON(message.user) : undefined);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<UpdateUserRequest>, I>>(
    object: I
  ): UpdateUserRequest {
    const message = createBaseUpdateUserRequest();
    message.user =
      object.user !== undefined && object.user !== null
        ? User.fromPartial(object.user)
        : undefined;
    return message;
  },
};

function createBaseUpdateUserResponse(): UpdateUserResponse {
  return { response: "" };
}

export const UpdateUserResponse = {
  encode(
    message: UpdateUserResponse,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.response !== "") {
      writer.uint32(10).string(message.response);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): UpdateUserResponse {
    const reader = input instanceof Reader ? input : new Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseUpdateUserResponse();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.response = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): UpdateUserResponse {
    return {
      response: isSet(object.response) ? String(object.response) : "",
    };
  },

  toJSON(message: UpdateUserResponse): unknown {
    const obj: any = {};
    message.response !== undefined && (obj.response = message.response);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<UpdateUserResponse>, I>>(
    object: I
  ): UpdateUserResponse {
    const message = createBaseUpdateUserResponse();
    message.response = object.response ?? "";
    return message;
  },
};

function createBaseDeleteUserRequest(): DeleteUserRequest {
  return { userId: "", hashPassword: "" };
}

export const DeleteUserRequest = {
  encode(message: DeleteUserRequest, writer: Writer = Writer.create()): Writer {
    if (message.userId !== "") {
      writer.uint32(10).string(message.userId);
    }
    if (message.hashPassword !== "") {
      writer.uint32(18).string(message.hashPassword);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): DeleteUserRequest {
    const reader = input instanceof Reader ? input : new Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseDeleteUserRequest();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.userId = reader.string();
          break;
        case 2:
          message.hashPassword = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): DeleteUserRequest {
    return {
      userId: isSet(object.userId) ? String(object.userId) : "",
      hashPassword: isSet(object.hashPassword)
        ? String(object.hashPassword)
        : "",
    };
  },

  toJSON(message: DeleteUserRequest): unknown {
    const obj: any = {};
    message.userId !== undefined && (obj.userId = message.userId);
    message.hashPassword !== undefined &&
      (obj.hashPassword = message.hashPassword);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<DeleteUserRequest>, I>>(
    object: I
  ): DeleteUserRequest {
    const message = createBaseDeleteUserRequest();
    message.userId = object.userId ?? "";
    message.hashPassword = object.hashPassword ?? "";
    return message;
  },
};

function createBaseDeleteUserResponse(): DeleteUserResponse {
  return { response: "" };
}

export const DeleteUserResponse = {
  encode(
    message: DeleteUserResponse,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.response !== "") {
      writer.uint32(10).string(message.response);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): DeleteUserResponse {
    const reader = input instanceof Reader ? input : new Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseDeleteUserResponse();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.response = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): DeleteUserResponse {
    return {
      response: isSet(object.response) ? String(object.response) : "",
    };
  },

  toJSON(message: DeleteUserResponse): unknown {
    const obj: any = {};
    message.response !== undefined && (obj.response = message.response);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<DeleteUserResponse>, I>>(
    object: I
  ): DeleteUserResponse {
    const message = createBaseDeleteUserResponse();
    message.response = object.response ?? "";
    return message;
  },
};

function createBaseLogInRequest(): LogInRequest {
  return { userName: "", hashPassword: "" };
}

export const LogInRequest = {
  encode(message: LogInRequest, writer: Writer = Writer.create()): Writer {
    if (message.userName !== "") {
      writer.uint32(10).string(message.userName);
    }
    if (message.hashPassword !== "") {
      writer.uint32(18).string(message.hashPassword);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): LogInRequest {
    const reader = input instanceof Reader ? input : new Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseLogInRequest();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.userName = reader.string();
          break;
        case 2:
          message.hashPassword = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): LogInRequest {
    return {
      userName: isSet(object.userName) ? String(object.userName) : "",
      hashPassword: isSet(object.hashPassword)
        ? String(object.hashPassword)
        : "",
    };
  },

  toJSON(message: LogInRequest): unknown {
    const obj: any = {};
    message.userName !== undefined && (obj.userName = message.userName);
    message.hashPassword !== undefined &&
      (obj.hashPassword = message.hashPassword);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<LogInRequest>, I>>(
    object: I
  ): LogInRequest {
    const message = createBaseLogInRequest();
    message.userName = object.userName ?? "";
    message.hashPassword = object.hashPassword ?? "";
    return message;
  },
};

function createBaseLogInResponse(): LogInResponse {
  return { user: undefined };
}

export const LogInResponse = {
  encode(message: LogInResponse, writer: Writer = Writer.create()): Writer {
    if (message.user !== undefined) {
      User.encode(message.user, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): LogInResponse {
    const reader = input instanceof Reader ? input : new Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseLogInResponse();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.user = User.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): LogInResponse {
    return {
      user: isSet(object.user) ? User.fromJSON(object.user) : undefined,
    };
  },

  toJSON(message: LogInResponse): unknown {
    const obj: any = {};
    message.user !== undefined &&
      (obj.user = message.user ? User.toJSON(message.user) : undefined);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<LogInResponse>, I>>(
    object: I
  ): LogInResponse {
    const message = createBaseLogInResponse();
    message.user =
      object.user !== undefined && object.user !== null
        ? User.fromPartial(object.user)
        : undefined;
    return message;
  },
};

function createBaseLogOutRequest(): LogOutRequest {
  return { userId: "" };
}

export const LogOutRequest = {
  encode(message: LogOutRequest, writer: Writer = Writer.create()): Writer {
    if (message.userId !== "") {
      writer.uint32(18).string(message.userId);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): LogOutRequest {
    const reader = input instanceof Reader ? input : new Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseLogOutRequest();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 2:
          message.userId = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): LogOutRequest {
    return {
      userId: isSet(object.userId) ? String(object.userId) : "",
    };
  },

  toJSON(message: LogOutRequest): unknown {
    const obj: any = {};
    message.userId !== undefined && (obj.userId = message.userId);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<LogOutRequest>, I>>(
    object: I
  ): LogOutRequest {
    const message = createBaseLogOutRequest();
    message.userId = object.userId ?? "";
    return message;
  },
};

function createBaseLogOutResponse(): LogOutResponse {
  return { response: "" };
}

export const LogOutResponse = {
  encode(message: LogOutResponse, writer: Writer = Writer.create()): Writer {
    if (message.response !== "") {
      writer.uint32(10).string(message.response);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): LogOutResponse {
    const reader = input instanceof Reader ? input : new Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseLogOutResponse();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.response = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): LogOutResponse {
    return {
      response: isSet(object.response) ? String(object.response) : "",
    };
  },

  toJSON(message: LogOutResponse): unknown {
    const obj: any = {};
    message.response !== undefined && (obj.response = message.response);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<LogOutResponse>, I>>(
    object: I
  ): LogOutResponse {
    const message = createBaseLogOutResponse();
    message.response = object.response ?? "";
    return message;
  },
};

export interface UsersAPI {
  CreateUser(request: CreateUserRequest): Promise<CreateUserResponse>;
  GetUser(request: GetUserRequest): Promise<GetUserResponse>;
  UpdateUser(request: UpdateUserRequest): Promise<UpdateUserResponse>;
  DeleteUser(request: DeleteUserRequest): Promise<DeleteUserResponse>;
  LogIn(request: LogInRequest): Promise<LogInResponse>;
  LogOut(request: LogOutRequest): Promise<LogOutResponse>;
}

export class UsersAPIClientImpl implements UsersAPI {
  private readonly rpc: Rpc;
  constructor(rpc: Rpc) {
    this.rpc = rpc;
    this.CreateUser = this.CreateUser.bind(this);
    this.GetUser = this.GetUser.bind(this);
    this.UpdateUser = this.UpdateUser.bind(this);
    this.DeleteUser = this.DeleteUser.bind(this);
    this.LogIn = this.LogIn.bind(this);
    this.LogOut = this.LogOut.bind(this);
  }
  CreateUser(request: CreateUserRequest): Promise<CreateUserResponse> {
    const data = CreateUserRequest.encode(request).finish();
    const promise = this.rpc.request(
      "users.proto.UsersAPI",
      "CreateUser",
      data
    );
    return promise.then((data) => CreateUserResponse.decode(new Reader(data)));
  }

  GetUser(request: GetUserRequest): Promise<GetUserResponse> {
    const data = GetUserRequest.encode(request).finish();
    const promise = this.rpc.request("users.proto.UsersAPI", "GetUser", data);
    return promise.then((data) => GetUserResponse.decode(new Reader(data)));
  }

  UpdateUser(request: UpdateUserRequest): Promise<UpdateUserResponse> {
    const data = UpdateUserRequest.encode(request).finish();
    const promise = this.rpc.request(
      "users.proto.UsersAPI",
      "UpdateUser",
      data
    );
    return promise.then((data) => UpdateUserResponse.decode(new Reader(data)));
  }

  DeleteUser(request: DeleteUserRequest): Promise<DeleteUserResponse> {
    const data = DeleteUserRequest.encode(request).finish();
    const promise = this.rpc.request(
      "users.proto.UsersAPI",
      "DeleteUser",
      data
    );
    return promise.then((data) => DeleteUserResponse.decode(new Reader(data)));
  }

  LogIn(request: LogInRequest): Promise<LogInResponse> {
    const data = LogInRequest.encode(request).finish();
    const promise = this.rpc.request("users.proto.UsersAPI", "LogIn", data);
    return promise.then((data) => LogInResponse.decode(new Reader(data)));
  }

  LogOut(request: LogOutRequest): Promise<LogOutResponse> {
    const data = LogOutRequest.encode(request).finish();
    const promise = this.rpc.request("users.proto.UsersAPI", "LogOut", data);
    return promise.then((data) => LogOutResponse.decode(new Reader(data)));
  }
}

interface Rpc {
  request(
    service: string,
    method: string,
    data: Uint8Array
  ): Promise<Uint8Array>;
}

type Builtin =
  | Date
  | Function
  | Uint8Array
  | string
  | number
  | boolean
  | undefined;

export type DeepPartial<T> = T extends Builtin
  ? T
  : T extends Array<infer U>
  ? Array<DeepPartial<U>>
  : T extends ReadonlyArray<infer U>
  ? ReadonlyArray<DeepPartial<U>>
  : T extends {}
  ? { [K in keyof T]?: DeepPartial<T[K]> }
  : Partial<T>;

type KeysOfUnion<T> = T extends T ? keyof T : never;
export type Exact<P, I extends P> = P extends Builtin
  ? P
  : P & { [K in keyof P]: Exact<P[K], I[K]> } & Record<
        Exclude<keyof I, KeysOfUnion<P>>,
        never
      >;

// If you get a compile-error about 'Constructor<Long> and ... have no overlap',
// add '--ts_proto_opt=esModuleInterop=true' as a flag when calling 'protoc'.
if (util.Long !== Long) {
  util.Long = Long as any;
  configure();
}

function isSet(value: any): boolean {
  return value !== null && value !== undefined;
}

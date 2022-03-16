package com.flashy;

import java.io.PrintWriter;
import java.io.StringWriter;
import java.lang.ref.WeakReference;
import java.text.MessageFormat;
import java.net.InetAddress;

import com.facebook.react.bridge.NativeModule;
import com.facebook.react.bridge.ReactApplicationContext;
import com.facebook.react.bridge.ReactContext;
import com.facebook.react.bridge.ReactContextBaseJavaModule;
import com.facebook.react.bridge.ReactMethod;
import com.facebook.react.bridge.Promise;
import com.facebook.react.bridge.ReadableMap;
import com.facebook.react.bridge.WritableMap;
import com.facebook.react.bridge.Arguments;
import io.grpc.StatusRuntimeException;
import io.grpc.ManagedChannelBuilder;
import io.grpc.ManagedChannel;
import android.os.AsyncTask;

public class EndpointsModule extends ReactContextBaseJavaModule {
  private final ManagedChannel channel;
  EndpointsModule(ReactApplicationContext context){
    super(context);
    channel = ManagedChannelBuilder.forAddress("10.0.2.2", 8080).usePlaintext().build();
    
  }

  @Override
  public String getName() {
    return "EndpointsModule";
  }

  @ReactMethod 
  public void CreateUser(ReadableMap request, Promise promise) {
    try {
      UsersProto.CreateUserRequest req = UsersProto.CreateUserRequest.newBuilder()
        .setUser(EntityConverter.convertJSUserToEntity(
          request.hasKey("user") ? request.getMap("user") : Arguments.createMap()
        ))
        .build();
      UsersProto.CreateUserResponse resp = (UsersProto.CreateUserResponse) new GrpcCall(new CreateUserRunnable(req), channel).execute()
        .get();
      promise.resolve(EntityConverter.convertUserEntityToJS(resp.getUser()));
    } catch (Exception e) {
      promise.reject(e);
    }
  }

  @ReactMethod
  public void GetUser(ReadableMap request, Promise promise) {
    try {
      UsersProto.GetUserRequest req = UsersProto.GetUserRequest.newBuilder()
        .setUserId(request.hasKey("user_id") ? request.getString("user_id") : "")
        .build();
      UsersProto.GetUserResponse resp = (UsersProto.GetUserResponse) new GrpcCall(new GetUserRunnable(req), channel).execute()
        .get();
      promise.resolve(EntityConverter.convertUserEntityToJS(resp.getUser()));
    } catch (Exception e) {
      promise.reject(e);
    }
  }

  @ReactMethod
  public void UpdateUser(ReadableMap request, Promise promise) {
    try {
      UsersProto.UpdateUserRequest req = UsersProto.UpdateUserRequest.newBuilder()
        .setUser(EntityConverter.convertJSUserToEntity(
          request.hasKey("user") ? request.getMap("user") : Arguments.createMap()
        ))
        .build();
      UsersProto.UpdateUserResponse resp = (UsersProto.UpdateUserResponse) new GrpcCall(new UpdateUserRunnable(req), channel).execute()
        .get();
      promise.resolve(resp.getResponse()); // String
    } catch (Exception e) {
      promise.reject(e);
    }
    
  }

  @ReactMethod
  public void DeleteUser(ReadableMap request, Promise promise) {
    try {
      UsersProto.DeleteUserRequest req = UsersProto.DeleteUserRequest.newBuilder()
        .setUserId(request.hasKey("user_id") ? request.getString("user_id") : "")
        .setHashPassword(request.hasKey("hash_password") ? request.getString("hash_password") : "")
        .build();
      UsersProto.DeleteUserResponse resp = (UsersProto.DeleteUserResponse) new GrpcCall(new DeleteUserRunnable(req), channel).execute()
        .get();
      promise.resolve(resp.getResponse()); // String
    } catch (Exception e) {
      promise.reject(e);
    }
    
  }

  @ReactMethod
  public void LogIn(ReadableMap request, Promise promise) {
    try {
      UsersProto.LogInRequest req = UsersProto.LogInRequest.newBuilder()
        .setUserName(request.hasKey("user_name") ? request.getString("user_name") : "")
        .setHashPassword(request.hasKey("hash_password") ? request.getString("hash_password") : "")
        .build();
      UsersProto.LogInResponse resp = (UsersProto.LogInResponse) new GrpcCall(new LogInRunnable(req), channel).execute()
        .get();
      promise.resolve(EntityConverter.convertUserEntityToJS(resp.getUser()));
    } catch (Exception e) {
      promise.reject(e);
    }
    
  }

  @ReactMethod
  public void LogOut(ReadableMap request, Promise promise) {
    try {
      UsersProto.LogOutRequest req = UsersProto.LogOutRequest.newBuilder()
        .setUserId(request.hasKey("user_id") ? request.getString("user_id") : "")
        .build();
      UsersProto.LogOutResponse resp = (UsersProto.LogOutResponse) new GrpcCall(new LogOutRunnable(req), channel).execute()
        .get();
      promise.resolve(resp.getResponse());
    } catch (Exception e) {
      promise.reject(e);
    }
   
  }

  private static class GrpcCall extends AsyncTask<Void, Void, Object> {
    private final GrpcRunnable grpcRunnable;
    private final ManagedChannel channel;

    GrpcCall(GrpcRunnable grpcRunnable, ManagedChannel channel) {
      this.grpcRunnable = grpcRunnable;
      this.channel = channel;
    }

    @Override
    protected Object doInBackground(Void... params) {
      try {
        return grpcRunnable.run(UsersAPIGrpc.newBlockingStub(channel), UsersAPIGrpc.newStub(channel));
      } catch (Exception e) {
        StringWriter sw = new StringWriter();
        PrintWriter pw = new PrintWriter(sw);
        e.printStackTrace(pw);
        pw.flush();
        System.out.println(sw);
        return null;
      }
    }
  }

  private interface GrpcRunnable {
    Object run(UsersAPIGrpc.UsersAPIBlockingStub blockingStub, UsersAPIGrpc.UsersAPIStub asyncStub) throws Exception;
  }

  private static class CreateUserRunnable implements GrpcRunnable {
    private final UsersProto.CreateUserRequest request;
    CreateUserRunnable(UsersProto.CreateUserRequest request) {
      this.request = request;
    }
      @Override
      public UsersProto.CreateUserResponse run(UsersAPIGrpc.UsersAPIBlockingStub blockingStub, UsersAPIGrpc.UsersAPIStub asyncStub) throws Exception {
        return createUser(request, blockingStub);
      }

      private UsersProto.CreateUserResponse createUser(UsersProto.CreateUserRequest request, UsersAPIGrpc.UsersAPIBlockingStub blockingStub) throws StatusRuntimeException {
        return blockingStub.createUser(request);
      }
    }

  private static class GetUserRunnable implements GrpcRunnable {
    private final UsersProto.GetUserRequest request;
    GetUserRunnable(UsersProto.GetUserRequest request) {
      this.request = request;
    }
      @Override
      public UsersProto.GetUserResponse run(UsersAPIGrpc.UsersAPIBlockingStub blockingStub, UsersAPIGrpc.UsersAPIStub asyncStub) throws Exception {
        return getUser(request, blockingStub);
      }
      private UsersProto.GetUserResponse getUser(UsersProto.GetUserRequest request, UsersAPIGrpc.UsersAPIBlockingStub blockingStub) throws StatusRuntimeException {
        return blockingStub.getUser(request);
      }
  }

  private static class UpdateUserRunnable implements GrpcRunnable {
    private final UsersProto.UpdateUserRequest request;
    UpdateUserRunnable(UsersProto.UpdateUserRequest request) {
      this.request = request;
    }

      @Override
      public UsersProto.UpdateUserResponse run(UsersAPIGrpc.UsersAPIBlockingStub blockingStub, UsersAPIGrpc.UsersAPIStub asyncStub) throws Exception {
        return updateUser(request, blockingStub);
      }
      private UsersProto.UpdateUserResponse updateUser(UsersProto.UpdateUserRequest request, UsersAPIGrpc.UsersAPIBlockingStub blockingStub) throws StatusRuntimeException {
        return blockingStub.updateUser(request);
      }   
  }

  private static class DeleteUserRunnable implements GrpcRunnable {
    private final UsersProto.DeleteUserRequest request;
    DeleteUserRunnable(UsersProto.DeleteUserRequest request) {
      this.request = request;
    }

      @Override
      public UsersProto.DeleteUserResponse run(UsersAPIGrpc.UsersAPIBlockingStub blockingStub, UsersAPIGrpc.UsersAPIStub asyncStub) throws Exception {
        return deleteUser(request, blockingStub);
      }
      private UsersProto.DeleteUserResponse deleteUser(UsersProto.DeleteUserRequest request, UsersAPIGrpc.UsersAPIBlockingStub blockingStub) throws StatusRuntimeException {
        return blockingStub.deleteUser(request);
      }
  }

  private static class LogInRunnable implements GrpcRunnable {
    private final UsersProto.LogInRequest request;
    LogInRunnable(UsersProto.LogInRequest request) {
      this.request = request;
    }

      @Override
      public UsersProto.LogInResponse run(UsersAPIGrpc.UsersAPIBlockingStub blockingStub, UsersAPIGrpc.UsersAPIStub asyncStub) throws Exception {
        return logIn(request, blockingStub);
      }
      private UsersProto.LogInResponse logIn(UsersProto.LogInRequest request, UsersAPIGrpc.UsersAPIBlockingStub blockingStub) throws StatusRuntimeException {
        return blockingStub.logIn(request);
      }
  }

  private static class LogOutRunnable implements GrpcRunnable {
    private final UsersProto.LogOutRequest request;
    LogOutRunnable(UsersProto.LogOutRequest request) {
      this.request = request;
    }

      @Override
      public UsersProto.LogOutResponse run(UsersAPIGrpc.UsersAPIBlockingStub blockingStub, UsersAPIGrpc.UsersAPIStub asyncStub) throws Exception {
        return logOut(request, blockingStub);
      }
      private UsersProto.LogOutResponse logOut(UsersProto.LogOutRequest request, UsersAPIGrpc.UsersAPIBlockingStub blockingStub) throws StatusRuntimeException {
        return blockingStub.logOut(request);
      }
      
  }

  private static class EntityConverter {
    protected static WritableMap convertUserEntityToJS(UsersProto.User user) {
      WritableMap map = Arguments.createMap();

      map.putString("user_id", user.getUserId());
      map.putString("user_name", user.getUserName());
      map.putString("hash_password", user.getHashPassword());
      map.putString("name", user.getName());
      map.putString("email", user.getEmail());
      map.putString("auth_token", user.getAuthToken());

      return map;
    }
    protected static UsersProto.User convertJSUserToEntity(ReadableMap user) {
      return UsersProto.User.newBuilder()
        .setUserName(user.hasKey("user_name") ? user.getString("user_name") : "")
        .setAuthToken(user.hasKey("auth_token") ? user.getString("auth_token") : "")
        .setEmail(user.hasKey("email") ? user.getString("email") : "")
        .setHashPassword(user.hasKey("hash_password") ? user.getString("hash_password") : "")
        .setUserId(user.hasKey("user_id") ? user.getString("user_id") : "")
        .setName(user.hasKey("name") ? user.getString("name") : "")
        .build();
    }
  }
}


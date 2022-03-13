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
  public void CreateUser(UsersProto.CreateUserRequest request, Promise promise) {
    try {
      UsersProto.CreateUserResponse resp = (UsersProto.CreateUserResponse) new GrpcCall(new CreateUserRunnable(request), channel).execute().get();
      promise.resolve(resp);
    } catch (Exception e) {
      promise.reject(e);
    }
  }

  @ReactMethod
  public void GetUser(UsersProto.GetUserRequest request, Promise promise) {
    try {
      UsersProto.GetUserResponse resp = (UsersProto.GetUserResponse) new GrpcCall(new GetUserRunnable(request), channel).execute().get();
      promise.resolve(resp);
    } catch (Exception e) {
      promise.reject(e);
    }
  }

  @ReactMethod
  public void UpdateUser(UsersProto.UpdateUserRequest request, Promise promise) {
    try {
      UsersProto.UpdateUserResponse resp = (UsersProto.UpdateUserResponse) new GrpcCall(new UpdateUserRunnable(request), channel).execute().get();
      promise.resolve(resp);
    } catch (Exception e) {
      promise.reject(e);
    }
    
  }

  @ReactMethod
  public void DeleteUser(UsersProto.DeleteUserRequest request, Promise promise) {
    try {
      UsersProto.DeleteUserResponse resp = (UsersProto.DeleteUserResponse) new GrpcCall(new DeleteUserRunnable(request), channel).execute().get();
      promise.resolve(resp);
    } catch (Exception e) {
      promise.reject(e);
    }
    
  }

  @ReactMethod
  public void LogIn(UsersProto.LogInRequest request, Promise promise) {
    try {
      UsersProto.LogInResponse resp = (UsersProto.LogInResponse) new GrpcCall(new LogInRunnable(request), channel).execute().get();
      promise.resolve(resp);
    } catch (Exception e) {
      promise.reject(e);
    }
    
  }

  @ReactMethod
  public void LogIn(UsersProto.LogOutRequest request, Promise promise) {
    try {
      UsersProto.LogOutResponse resp = (UsersProto.LogOutResponse) new GrpcCall(new LogOutRunnable(request), channel).execute().get();
      promise.resolve(resp);
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
        StringBuffer logs = new StringBuffer();
        appendLogs(logs, "*** CreateUser: ");
        UsersProto.CreateUserResponse response = blockingStub.createUser(request);
        UsersProto.User user = response.getUser();
        appendLogs(logs, "name={0} email={1} username={2} hashPassword={3} authToken={5} userID={6} phoneNumber={7}", 
        user.getName(), user.getEmail(), user.getUserName(), user.getHashPassword(), user.getAuthToken(), user.getUserId(), user.getPhoneNumber());
        System.out.println(logs.toString());
        return response;
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
        StringBuffer logs = new StringBuffer();
        appendLogs(logs, "*** GetUser: ");
        UsersProto.GetUserResponse response = blockingStub.getUser(request);
        UsersProto.User user = response.getUser();
        appendLogs(logs, "name={0} email={1} username={2} hashPassword={3} authToken={5} userID={6} phoneNumber={7}", 
        user.getName(), user.getEmail(), user.getUserName(), user.getHashPassword(), user.getAuthToken(), user.getUserId(), user.getPhoneNumber());
        System.out.println(logs.toString());
        return response;
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
        StringBuffer logs = new StringBuffer();
        appendLogs(logs, "*** UpdateUser: ");
        UsersProto.UpdateUserResponse response = blockingStub.updateUser(request);
        String msg = response.getResponse();
        appendLogs(logs, "status={0}", msg );
        System.out.println(logs.toString());
        return response;
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
        StringBuffer logs = new StringBuffer();
        appendLogs(logs, "*** DeleteUser: ");
        UsersProto.DeleteUserResponse response = blockingStub.deleteUser(request);
        String msg = response.getResponse();
        appendLogs(logs, "status={0}", msg );
        System.out.println(logs.toString());
        return response;
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
        StringBuffer logs = new StringBuffer();
        appendLogs(logs, "*** GetUser: ");
        UsersProto.LogInResponse response = blockingStub.logIn(request);
        UsersProto.User user = response.getUser();
        appendLogs(logs, "name={0} email={1} username={2} hashPassword={3} authToken={5} userID={6} phoneNumber={7}", 
        user.getName(), user.getEmail(), user.getUserName(), user.getHashPassword(), user.getAuthToken(), user.getUserId(), user.getPhoneNumber());
        System.out.println(logs.toString());
        return response;
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
        StringBuffer logs = new StringBuffer();
        appendLogs(logs, "*** DeleteUser: ");
        UsersProto.LogOutResponse response = blockingStub.logOut(request);
        String msg = response.getResponse();
        appendLogs(logs, "status={0}", msg );
        System.out.println(logs.toString());
        return response;
      }
      
  }

  private static void appendLogs(StringBuffer logs, String msg, Object... params) {
    if (params.length > 0) {
      logs.append(MessageFormat.format(msg, params));
    } else {
      logs.append(msg);
    }
    logs.append("\n");
  }

}


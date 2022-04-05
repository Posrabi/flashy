package com.flashy;

import java.io.PrintWriter;
import java.io.StringWriter;
import java.lang.ref.WeakReference;
import java.text.MessageFormat;
import java.net.InetAddress;
import java.util.List;

import com.facebook.react.bridge.NativeModule;
import com.facebook.react.bridge.ReactApplicationContext;
import com.facebook.react.bridge.ReactContext;
import com.facebook.react.bridge.ReactContextBaseJavaModule;
import com.facebook.react.bridge.ReactMethod;
import com.facebook.react.bridge.Promise;
import com.facebook.react.bridge.ReadableMap;
import com.facebook.react.bridge.WritableMap;
import com.facebook.react.bridge.WritableArray;
import com.facebook.react.bridge.Arguments;
import io.grpc.StatusRuntimeException;
import io.grpc.ManagedChannelBuilder;
import io.grpc.ManagedChannel;
import io.grpc.Metadata;
import io.grpc.stub.MetadataUtils;
import android.os.AsyncTask;

public class EndpointsModule extends ReactContextBaseJavaModule {
  private final ManagedChannel channel;
  EndpointsModule(ReactApplicationContext context){
    super(context);
    channel = ManagedChannelBuilder.forAddress("localhost", 8080).usePlaintext().build();
    
  }

  @Override
  public String getName() {
    return "EndpointsModule";
  }

  @ReactMethod 
  public void CreateUser(ReadableMap request, Promise promise) {
    try {
      UsersProto.CreateUserRequest req = UsersProto.CreateUserRequest.newBuilder()
        .setUser( EntityConverter.convertJSUserToEntity( 
          request.hasKey("user") ? request.getMap("user") : Arguments.createMap()
          ))
        .build();
      UsersProto.CreateUserResponse resp = (UsersProto.CreateUserResponse) new GrpcCall(new CreateUserRunnable(req), channel, null).execute()
        .get();
      if (resp == null) {
        promise.reject(new NullPointerException("no response"));
      }
      promise.resolve(EntityConverter.createJSResponseWithUser(resp.getUser()));
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
      Metadata metadata = new Metadata();
      Metadata.Key<String> key = Metadata.Key.of("JWTToken", Metadata.ASCII_STRING_MARSHALLER);
      metadata.put(key, request.hasKey("auth_token") ? request.getString("auth_token") : "");
      UsersProto.GetUserResponse resp = (UsersProto.GetUserResponse) new GrpcCall(new GetUserRunnable(req), channel, metadata).execute()
        .get();
      if (resp == null) {
        promise.reject(new NullPointerException("no response"));
      }
      promise.resolve(EntityConverter.createJSResponseWithUser(resp.getUser()));
    } catch (Exception e) {
      promise.reject(e);
    }
  }

  @ReactMethod
  public void UpdateUser(ReadableMap request, Promise promise) {
    try {
      UsersProto.UpdateUserRequest req = UsersProto.UpdateUserRequest.newBuilder()
        .setUser( EntityConverter.convertJSUserToEntity( 
          request.hasKey("user") ? request.getMap("user") : Arguments.createMap()
          ))
        .build();
      UsersProto.UpdateUserResponse resp = (UsersProto.UpdateUserResponse) new GrpcCall(new UpdateUserRunnable(req), channel, null).execute()
        .get();
      if (resp == null) {
        promise.reject(new NullPointerException("no response"));
      }
      promise.resolve(EntityConverter.createJSResponseWithStatus(resp.getResponse()));
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
      UsersProto.DeleteUserResponse resp = (UsersProto.DeleteUserResponse) new GrpcCall(new DeleteUserRunnable(req), channel, null).execute()
        .get();
      if (resp == null) {
        promise.reject(new NullPointerException("no response"));
      }
      promise.resolve(EntityConverter.createJSResponseWithStatus(resp.getResponse()));
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
      UsersProto.LogInResponse resp = (UsersProto.LogInResponse) new GrpcCall(new LogInRunnable(req), channel, null).execute()
        .get();
      if (resp == null) {
        promise.reject(new NullPointerException("no response"));
      }
      promise.resolve(EntityConverter.createJSResponseWithUser(resp.getUser()));
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
      UsersProto.LogOutResponse resp = (UsersProto.LogOutResponse) new GrpcCall(new LogOutRunnable(req), channel, null).execute()
        .get();
      if (resp == null) {
        promise.reject(new NullPointerException("no response"));
      }
      promise.resolve(EntityConverter.createJSResponseWithStatus(resp.getResponse()));
    } catch (Exception e) {
      promise.reject(e);
    }
   
  }

  @ReactMethod
  public void CreatePhrase(ReadableMap request, Promise promise) {
    try {
      UsersProto.CreatePhraseRequest req = UsersProto.CreatePhraseRequest.newBuilder()
        .setPhrase( EntityConverter.convertJSPhraseToEntity(
          request.hasKey("phrase") ? request.getMap("phrase") : Arguments.createMap()
          ))
        .build();
      UsersProto.CreatePhraseResponse resp = (UsersProto.CreatePhraseResponse) new GrpcCall(new CreatePhraseRunnable(req), channel, null).execute()
        .get();
      if (resp == null) {
        promise.reject(new NullPointerException("no response"));
      }
      promise.resolve(EntityConverter.createJSResponseWithStatus(resp.getResponse()));
    } catch (Exception e) {
      promise.reject(e);
    }
  }

  @ReactMethod
  public void GetPhrases(ReadableMap request, Promise promise) {
    try {
      UsersProto.GetPhrasesRequest req = UsersProto.GetPhrasesRequest.newBuilder()
        .setUserId(request.hasKey("user_id") ? request.getString("user_id") : "")
        .setStart(request.hasKey("start") ? (long) request.getDouble("start") : 0)
        .setEnd(request.hasKey("end") ? (long) request.getDouble("end") : 0)
        .build();
      UsersProto.GetPhrasesResponse resp = (UsersProto.GetPhrasesResponse) new GrpcCall(new GetPhrasesRunnable(req), channel, null).execute()
        .get();
      if (resp == null) {
        promise.reject(new NullPointerException("no response"));
      }
      promise.resolve(EntityConverter.createJSResponseWithPhrases(resp.getPhrasesList()));
    } catch (Exception e) {
      promise.reject(e);
    }
  }

  @ReactMethod
  public void DeletePhrase(ReadableMap request, Promise promise) {
    try {
      UsersProto.DeletePhraseRequest req = UsersProto.DeletePhraseRequest.newBuilder()
        .setUserId(request.hasKey("user_id") ? request.getString("user_id") : "")
        .setPhraseTime(request.hasKey("phrase_time") ? (long) request.getDouble("phrase_time") : 0)
        .build();
      UsersProto.DeletePhraseResponse resp = (UsersProto.DeletePhraseResponse) new GrpcCall(new DeletePhraseRunnable(req), channel, null).execute()
        .get();
      if (resp == null) {
        promise.reject(new NullPointerException("no response"));
      }
      promise.resolve(EntityConverter.createJSResponseWithStatus(resp.getResponse()));
    } catch (Exception e) {
      promise.reject(e);
    }
  }

  @ReactMethod
  public void LogInWithFB(ReadableMap request, Promise promise) {
    try {
      UsersProto.LogInWithFBRequest req = UsersProto.LogInWithFBRequest.newBuilder()
        .setUserId(request.hasKey("user_id") ? request.getString("user_id") : "")
        .setFacebookAccessToken(request.hasKey("facebook_access_token") ? request.getString("facebook_access_token") : "")
        .build();
      UsersProto.LogInWithFBResponse resp = (UsersProto.LogInWithFBResponse) new GrpcCall(new LogInWithFBRunnable(req), channel, null).execute()
        .get();
      if (resp == null) {
        promise.reject(new NullPointerException("no response"));
      }
      promise.resolve(EntityConverter.createJSResponseWithUser(resp.getUser())); 
    } catch (Exception e) {
      promise.reject(e);
    }
  }

  private static class GrpcCall extends AsyncTask<Void, Void, Object> {
    private final GrpcRunnable grpcRunnable;
    private final ManagedChannel channel;
    private final Metadata metadata;

    GrpcCall(GrpcRunnable grpcRunnable, ManagedChannel channel, Metadata metadata) {
      this.grpcRunnable = grpcRunnable;
      this.channel = channel;
      this.metadata = metadata;
    }

    @Override
    protected Object doInBackground(Void... params) {
      try {
        if (metadata == null) {
          return grpcRunnable.run(UsersAPIGrpc.newBlockingStub(channel), UsersAPIGrpc.newStub(channel));
        } else {
          return grpcRunnable.run(UsersAPIGrpc.newBlockingStub(channel).withInterceptors(MetadataUtils.newAttachHeadersInterceptor(metadata)),
             UsersAPIGrpc.newStub(channel).withInterceptors(MetadataUtils.newAttachHeadersInterceptor(metadata)));
        }
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

  private static class CreatePhraseRunnable implements GrpcRunnable {
    private final UsersProto.CreatePhraseRequest request;
    CreatePhraseRunnable(UsersProto.CreatePhraseRequest request) {
      this.request = request;
    }

      @Override
      public UsersProto.CreatePhraseResponse run(UsersAPIGrpc.UsersAPIBlockingStub blockingStub, UsersAPIGrpc.UsersAPIStub asyncStub) throws Exception {
        return createPhrase(request, blockingStub);
      }

      private UsersProto.CreatePhraseResponse createPhrase(UsersProto.CreatePhraseRequest request, UsersAPIGrpc.UsersAPIBlockingStub blockingStub) throws StatusRuntimeException {
        return blockingStub.createPhrase(request);
      }
  }

  private static class GetPhrasesRunnable implements GrpcRunnable {
    private final UsersProto.GetPhrasesRequest request;
    GetPhrasesRunnable(UsersProto.GetPhrasesRequest request) {
      this.request = request;
    }

      @Override
      public UsersProto.GetPhrasesResponse run(UsersAPIGrpc.UsersAPIBlockingStub blockingStub, UsersAPIGrpc.UsersAPIStub asyncStub) throws Exception {
        return getPhrases(request, blockingStub);
      }

      private UsersProto.GetPhrasesResponse getPhrases(UsersProto.GetPhrasesRequest request, UsersAPIGrpc.UsersAPIBlockingStub blockingStub) throws StatusRuntimeException {
        return blockingStub.getPhrases(request);
      }
  }

  private static class DeletePhraseRunnable implements GrpcRunnable {
    private final UsersProto.DeletePhraseRequest request;
    DeletePhraseRunnable(UsersProto.DeletePhraseRequest request) {
      this.request = request;
    }

      @Override
      public UsersProto.DeletePhraseResponse run(UsersAPIGrpc.UsersAPIBlockingStub blockingStub, UsersAPIGrpc.UsersAPIStub asyncStub) throws Exception {
        return deletePhrase(request, blockingStub);
      }

      private UsersProto.DeletePhraseResponse deletePhrase(UsersProto.DeletePhraseRequest request, UsersAPIGrpc.UsersAPIBlockingStub blockingStub) throws StatusRuntimeException {
        return blockingStub.deletePhrase(request);
      }
  }

  private static class LogInWithFBRunnable implements GrpcRunnable {
    private final UsersProto.LogInWithFBRequest request;
    LogInWithFBRunnable(UsersProto.LogInWithFBRequest request) {
      this.request = request;
    }
      @Override
      public UsersProto.LogInWithFBResponse run(UsersAPIGrpc.UsersAPIBlockingStub blockingStub, UsersAPIGrpc.UsersAPIStub asyncStub) throws Exception {
        return logInWithFB(request, blockingStub);
      }

      private UsersProto.LogInWithFBResponse logInWithFB(UsersProto.LogInWithFBRequest request, UsersAPIGrpc.UsersAPIBlockingStub blockingStub) throws StatusRuntimeException {
        return blockingStub.logInWithFB(request);
      }

  }

  private static class EntityConverter {
    // Types
    protected static WritableMap convertUserEntityToJS(UsersProto.User user) {
      WritableMap map = Arguments.createMap();

      map.putString("user_id", user.getUserId());
      map.putString("user_name", user.getUserName());
      map.putString("hash_password", user.getHashPassword());
      map.putString("name", user.getName());
      map.putString("email", user.getEmail());
      map.putString("facebook_access_token", user.getFacebookAccessToken());
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
        .setFacebookAccessToken(user.hasKey("facebook_access_token") ? user.getString("facebook_access_token") : "")
        .build();
    }

    protected static WritableMap convertPhraseEntityToJS(UsersProto.Phrase phrase) {
      WritableMap map = Arguments.createMap();

      map.putString("user_id", phrase.getUserId());
      map.putString("word", phrase.getWord());
      map.putString("sentence", phrase.getSentence());
      map.putDouble("phrase_time", phrase.getPhraseTime());
      map.putBoolean("correct", phrase.getCorrect());

      return map;
    }

    protected static UsersProto.Phrase convertJSPhraseToEntity(ReadableMap phrase) {
      return UsersProto.Phrase.newBuilder()
        .setUserId(phrase.hasKey("user_id") ? phrase.getString("user_id") : "")
        .setWord(phrase.hasKey("word") ? phrase.getString("word") : "")
        .setSentence(phrase.hasKey("sentence") ? phrase.getString("sentence") : "")
        .setPhraseTime(phrase.hasKey("phrase_time") ? (long) phrase.getDouble("phrase_time") : 0)
        .setCorrect(phrase.getBoolean("correct"))
        .build();
    }

    // Responses
    protected static WritableMap createJSResponseWithUser(UsersProto.User user) {
      WritableMap jsUser = convertUserEntityToJS(user);
      WritableMap resp = Arguments.createMap();

      resp.putMap("user", jsUser);

      return resp;
    }

    protected static WritableMap createJSResponseWithStatus(String status) {
      WritableMap resp = Arguments.createMap();

      resp.putString("response", status );

      return resp;
    }

    protected static WritableMap createJSResponseWithPhrases(List<UsersProto.Phrase> phrases) {
      WritableMap resp = Arguments.createMap();
      WritableArray arr = Arguments.createArray();

      for (UsersProto.Phrase phrase : phrases) {
        arr.pushMap(convertPhraseEntityToJS(phrase));
      }
      resp.putArray("phrases", arr);

      return resp;
    }
  }
}


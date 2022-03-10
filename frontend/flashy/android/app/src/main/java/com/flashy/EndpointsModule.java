package com.flashy;

import com.facebook.react.bridge.NativeModule;
import com.facebook.react.bridge.ReactApplicationContext;
import com.facebook.react.bridge.ReactContext;
import com.facebook.react.bridge.ReactContextBaseJavaModule;
import com.facebook.react.bridge.ReactMethod;



public class EndpointsModule extends ReactContextBaseJavaModule {
  EndpointsModule(ReactApplicationContext context) {
    super(context);
  }

  @Override
  public String getName() {
    return "Endpoints";
  }

  @ReactMethod 
  public void CreateUser() {

  }

  @ReactMethod
  public void GetUser() {

  }

  @ReactMethod
  public void UpdateUser() {

  }

  @ReactMethod
  public void DeleteUser() {

  }

  @ReactMethod
  public void LogIn() {

  }

  @ReactMethod
  public void LogOut() {

  }
}

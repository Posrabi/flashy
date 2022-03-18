import { NativeModules } from 'react-native';
import { UsersAPI } from '../types/users';

const { EndpointsModule } = NativeModules;

export default EndpointsModule as UsersAPI;

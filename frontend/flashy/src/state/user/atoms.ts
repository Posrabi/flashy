import { atom, RecoilState } from 'recoil';
import { defaultUser } from '../../api/defaults';
import { User } from '../../types';

export const currentUser: RecoilState<User> = atom<User>({
    key: 'currentUser',
    default: defaultUser(),
});

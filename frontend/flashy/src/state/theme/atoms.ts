import * as eva from '@eva-design/eva';
import { atom, RecoilState } from 'recoil';

export type EvaTheme = typeof eva.light | typeof eva.dark;

export const themeColorState: RecoilState<EvaTheme> = atom({
    key: 'themeColor',
    default: eva.dark,
});

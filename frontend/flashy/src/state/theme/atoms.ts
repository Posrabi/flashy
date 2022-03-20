import * as eva from '@eva-design/eva';
import { Appearance } from 'react-native';
import { atom, RecoilState } from 'recoil';

export type EvaTheme = typeof eva.light | typeof eva.dark;

export const themeColorState: RecoilState<EvaTheme> = atom<EvaTheme>({
    key: 'themeColor',
    default: Appearance.getColorScheme() === 'dark' ? eva.dark : eva.light,
});

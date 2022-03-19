import { SetterOrUpdater, useSetRecoilState } from 'recoil';
import { EvaTheme, themeColorState } from './atoms';

export const toggleThemeColor = (themeColor: EvaTheme): void => {
    const setTheme: SetterOrUpdater<EvaTheme> = useSetRecoilState(themeColorState);
    return setTheme(themeColor);
};

/**
 *
 * @param array any array
 * @returns the shuffled array
 */
export const arrayShuffler = (array: Array<any>): Array<any> => {
    let currentIndex = array.length,
        randomIndex;

    // While there remain elements to shuffle...
    while (currentIndex != 0) {
        // Pick a remaining element...
        randomIndex = Math.floor(Math.random() * currentIndex);
        currentIndex--;

        // And swap it with the current element.
        [array[currentIndex], array[randomIndex]] = [array[randomIndex], array[currentIndex]];
    }

    return array;
};

const substringPositionsArray = (strArray: string[], search: string): Array<number> => {
    const result = [];
    for (let i = 0; i < strArray.length; i++) {
        if (strArray[i] === search) {
            result.push(i);
        }
    }
    return result;
};

/**
 *
 * @param str string to search
 * @param search string to search for
 * @returns boolean of whether search is in str
 */
export const stringSearcher = (str: string, search: string): Array<number> => {
    return substringPositionsArray(str.split(' '), search);
};

export const stringHider = (str: string, replace: string): string => {
    const strArray = str.split(' ');
    const indices = substringPositionsArray(strArray, replace);
    for (let i = 0; i < indices.length; i++) {
        strArray[indices[i]] = '.....';
    }
    console.log(strArray);
    let res = '';
    for (const str of strArray) {
        res += str + ' ';
    }
    return res.slice(0, res.length - 1);
};

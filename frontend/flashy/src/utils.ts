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

/**
 *
 * @param str string to search
 * @param search string to search for
 * @returns boolean of whether search is in str
 */
export const stringSearcher = (str: string, search: string): Boolean => {
    let index = str.indexOf(search);
    if (index !== -1) {
        if (
            (index === 0 && str[index + search.length] === ' ') ||
            (index === str.length - search.length && str[index - 1] === ' ') ||
            (str[index - 1] === ' ' && str[index + search.length] === ' ')
        )
            return true;
        return false;
    }
    return false;
};

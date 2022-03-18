module.exports = {
    root: true,
    parser: '@typescript-eslint/parser',
    parserOptions: {
        ecmaVersion: 2020,
        sourceType: 'module',
        ecmaFeatures: {
            jsx: true,
        },
    },
    plugins: ['eslint-plugin-tsdoc', 'react-hooks'],
    settings: {
        react: {
            version: 'detect',
        },
    },
    extends: [
        '@react-native-community',
        'universe',
        'prettier',
        'prettier/@typescript-eslint',
        'prettier/react',
        'plugin:@typescript-eslint/recommended',
        'plugin:prettier/recommended', // Enables eslint-plugin-prettier and eslint-config-prettier. This will display prettier errors as ESLint errors. Make sure this is always the last configuration in the extends array.
    ],
    rules: {
        'func-style': ['error', 'expression'],
        '@typescript-eslint/typedef': [
            'error',
            {
                arrowParameter: true,
                parameter: true,
            },
        ],
        '@typescript-eslint/no-inferrable-types': [
            'error',
            {
                ignoreParameters: true,
            },
        ],
        // React 17 no longer requires react in scope for JSX
        'react/jsx-uses-react': 'off',
        'react/react-in-jsx-scope': 'off',
        'no-shadow': 'off',
        '@typescript-eslint/no-shadow': ['error'],
        'react-hooks/rules-of-hooks': 'error', // Checks rules of Hooks
        'react-hooks/exhaustive-deps': [
            'warn',
            {
                additionalHooks: '(useRecoilCallback|useRecoilTransaction_UNSTABLE)', // Checks effect dependencies
                // "tsdoc/syntax": "warn"
            },
        ],
    },
};

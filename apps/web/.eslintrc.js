module.exports = {
  parser: '@typescript-eslint/parser',
  plugins: ['@typescript-eslint', 'prettier', 'import', 'unused-imports'],
  ignorePatterns: ['react-app-env.d.ts', '.eslintrc.js', 'docker/**', 'dist/**', 'node_modules/**'],
  extends: [
    'plugin:@typescript-eslint/recommended',
    'plugin:prettier/recommended',
    'plugin:import/warnings',
    'plugin:import/typescript',
    'prettier'
  ],
  rules: {
    'import/named': ['warn'],
    'import/imports-first': ['error', 'absolute-first'],
    'import/newline-after-import': ['warn', {count: 1}],
    'import/no-deprecated': ['error'],
    'import/no-dynamic-require': ['warn'],
    'import/no-unused-modules': ['warn'],
    'import/order': ['warn'],
    'no-console': 'warn',
    'unused-imports/no-unused-imports': 'error',
  },
}

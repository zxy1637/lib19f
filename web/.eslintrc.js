/**
 * @type {import('eslint').Linter.Config}
 */
module.exports = {
  env: {
    browser: true,
    es2021: true,
    node: true,
  },
  root: true,
  extends: [
    "eslint:recommended",
    "plugin:vue/vue3-essential",
    "plugin:@typescript-eslint/recommended",
    "prettier",
  ],
  overrides: [],
  parser: "vue-eslint-parser",
  parserOptions: {
    parser: "@typescript-eslint/parser",
    ecmaVersion: "latest",
    sourceType: "module",
    // project: "./tsconfig.json",
  },
  plugins: ["vue", "@typescript-eslint", "simple-import-sort"],
  rules: {
    "sort-vars": ["error", { ignoreCase: false }],
    "simple-import-sort/imports": "error",
    "simple-import-sort/exports": "error",
    "vue/multi-word-component-names": [
      "error",
      {
        ignores: ["index", "error"],
      },
    ],
  },
};

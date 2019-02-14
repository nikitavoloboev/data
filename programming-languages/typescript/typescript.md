# [TypeScript](https://github.com/Microsoft/TypeScript)

[TS Handbook](https://www.typescriptlang.org/docs/handbook/basic-types.html) and [this article](https://toddmotto.com/typescript-introduction) are great intros.

## Notes

- TypeScript can figure the return type out by looking at the return statements, so we can also optionally leave this off in many cases
- You really shouldn’t use arrow function as class members. Arrow functions are duplicated in memory for each class instance, while proper method members are written once in your class’s prototype and shared between all instances.
- [Unlike classes, an interface is a virtual structure that only exists within the context of TypeScript. The TypeScript compiler uses interfaces solely for type-checking purposes.](https://toddmotto.com/classes-vs-interfaces-in-typescript)
- `any` means “ignore the type”, `unknown` means “we don’t know the type”, which aren’t the same thing, and for better code you’d rather not know the type than ignore it.
- You can’t use `unknown` as a simple replacement for `any`, that wouldn’t make sense. `unkwnown` carries the implication that you should not use this variable other than passing it around. Or testing its type to get rid of the “unknown” state.

## Links

- [React & Redux in TypeScript - Static Typing Guide](https://github.com/piotrwitek/react-redux-typescript-guide#readme)
- [Opening Remarks by Anders Hejlsberg (2018)](https://www.youtube.com/watch?v=wpgKd-rwnMw)
- [Ultimate React Component Patterns with Typescript 2.8](https://levelup.gitconnected.com/ultimate-react-component-patterns-with-typescript-2-8-82990c516935)
- [Very opinionated front end boilerplate](https://ts-react-boilerplate.js.org/)
- [Babel 7 + TypeScript](http://artsy.github.io/blog/2017/11/27/Babel-7-and-TypeScript/)
- [Classes vs interfaces in TS](https://toddmotto.com/classes-vs-interfaces-in-typescript)
- [StREST](https://github.com/eykhagen/strest#readme) - Set up tests for REST in seconds with YAML.
- [Fork TS Checker Webpack Plugin](https://github.com/Realytics/fork-ts-checker-webpack-plugin) - Webpack plugin that runs typescript type checker on a separate process.
- [Tsickle](https://github.com/angular/tsickle) - TypeScript to Closure Translator.
- [Set of TSLint rules used on some Microsoft projects](https://github.com/Microsoft/tslint-microsoft-contrib)
- [TypeSearch](https://microsoft.github.io/TypeSearch/) - Search for TS types.
- [TypeScript ESLint](https://github.com/typescript-eslint/typescript-eslint) - Monorepo for all the tooling which enables ESLint to support TypeScript.
- [TypeScript With Babel: A Beautiful Marriage (2018)](https://iamturns.com/typescript-babel/)
- [ts-node-dev](https://github.com/whitecolor/ts-node-dev) - Restarts target node process when any of required files changes (as standard node-dev) but shares Typescript compilation process between restarts.
- [When to use `never` and `unknown` in Typescript (2019)](https://blog.logrocket.com/when-to-use-never-and-unknown-in-typescript-5e4d6c5799ad)
- [bibliography on Gradual Typing](https://github.com/samth/gradual-typing-bib#readme)
- [dtslint](https://github.com/Microsoft/dtslint) - Tests a TypeScript declaration file for style and correctness.
- [Clean Code concepts adapted for TypeScript](https://github.com/labs42io/clean-code-typescript)
- [SonarTS](https://github.com/SonarSource/SonarTS) - Static code analyzer for TypeScript.
- [TypeScript TSLint Language Service Plugin](https://github.com/Microsoft/typescript-tslint-plugin)
- [TypeScript Guidelines](https://github.com/unional/typescript-guidelines#readme) - Guideline to focus on how to write effective TypeScript with minimal effort.
- [Practical Advanced TypeScript](https://egghead.io/courses/practical-advanced-typescript)
- [typescript-json-schema](https://github.com/YousefED/typescript-json-schema) - Generate json-schemas from your Typescript sources.
- [tslint-config-security](https://github.com/webschik/tslint-config-security) - TSLint security rules.

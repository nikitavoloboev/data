# [TypeScript](https://github.com/Microsoft/TypeScript)

[TS Handbook](https://www.typescriptlang.org/docs/handbook/basic-types.html) and [this article](https://toddmotto.com/typescript-introduction) are great intros.

## Notes

- TypeScript can figure the return type out by looking at the return statements, so we can also optionally leave this off in many cases
- You really shouldn’t use arrow function as class members. Arrow functions are duplicated in memory for each class instance, while proper method members are written once in your class’s prototype and shared between all instances.
- [Unlike classes, an interface is a virtual structure that only exists within the context of TypeScript. The TypeScript compiler uses interfaces solely for type-checking purposes.](https://toddmotto.com/classes-vs-interfaces-in-typescript)
- `any` means “ignore the type”, `unknown` means “we don’t know the type”, which aren’t the same thing, and for better code you’d rather not know the type than ignore it.
- You can’t use `unknown` as a simple replacement for `any`, that wouldn’t make sense. `unkwnown` carries the implication that you should not use this variable other than passing it around. Or testing its type to get rid of the “unknown” state.
- Tuples are very useful for simulating multiple returns like React hooks, and other combined values that don't need string object keys

## Links

- [TypeScript Language Specification](https://github.com/Microsoft/TypeScript/blob/master/doc/spec.md#readme)
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
- [Clean Code concepts adapted for TypeScript](https://github.com/labs42io/clean-code-typescript#readme)
- [SonarTS](https://github.com/SonarSource/SonarTS) - Static code analyzer for TypeScript.
- [TypeScript TSLint Language Service Plugin](https://github.com/Microsoft/typescript-tslint-plugin)
- [TypeScript Guidelines](https://github.com/unional/typescript-guidelines#readme) - Guideline to focus on how to write effective TypeScript with minimal effort.
- [Practical Advanced TypeScript](https://egghead.io/courses/practical-advanced-typescript)
- [typescript-json-schema](https://github.com/YousefED/typescript-json-schema) - Generate json-schemas from your Typescript sources.
- [tslint-config-security](https://github.com/webschik/tslint-config-security) - TSLint security rules.
- [TypeScript Definition Style Guide](https://github.com/sindresorhus/typescript-definition-style-guide#readme)
- [Conditional Type Checks](https://github.com/dsherret/conditional-type-checks) - Types for testing TypeScript types.
- [quicktype](https://github.com/quicktype/quicktype) - Generate types and converters from JSON, Schema, and GraphQL.
- [Collection of essential TypeScript types](https://github.com/sindresorhus/type-fest)
- [Be Super with TypeScript and Jared Palmer episode (2019)](https://overcast.fm/+N_6KEDwFo)
- [typesync](https://github.com/jeffijoe/typesync) - Install missing TypeScript typings for dependencies in your package.json.
- [TSdx](https://github.com/palmerhq/tsdx) - Zero-config CLI for TypeScript package development.
- [Programming TypeScript book Boris Cherny (2019)](https://www.oreilly.com/library/view/programming-typescript/9781492037644/)
- [OneFraction](https://github.com/TrillCyborg/onefraction) - React/ApolloGraphQL/Node/Mongo demo written in Typescript.
- [TypeScript 3.0: The unknown Type (2019)](https://mariusschulz.com/blog/typescript-3-0-the-unknown-type)
- [Lobsters: What have you learned from adopting Typescript into an existing JS codebase? (2019)](https://lobste.rs/s/3ucfhp/what_have_you_learned_from_adopting)
- [React+TypeScript Cheatsheets](https://github.com/typescript-cheatsheets/react-typescript-cheatsheet#readme)
- [Using Typescript to make invalid states irrepresentable](http://www.javiercasas.com/articles/typescript-impossible-states-irrepresentable)
- [High-level notes about TypeScript](https://github.com/orta/typescript-notes)

## Images

![](https://i.imgur.com/APrrI2V.png)

> TypeScript’s type hierarchy

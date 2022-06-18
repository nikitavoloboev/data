---
title: Redux
---

# [Redux](https://redux.js.org)

The Gist of Redux is a follows:

- The whole state of your app is stored in an object tree inside a single store.
- The only way to change the state tree is to emit an action, an object describing what happened.
- To specify how the actions transform the state tree, you write pure reducers.

There are [better approaches to manage state in React now](https://twitter.com/kamyshev_dev/status/1441736479240122372). [Sometimes Redux can be useful](https://www.reddit.com/r/reactjs/comments/s4hxd2/how_do_you_know_if_a_project_will_need_redux/) perhaps.

## OSS Redux apps

- [NextJS + Redux template](https://github.com/iShibi/nextjs-template)

## Notes

- OwnPops is something your component receives from the parent. StateProps is something your component receives from the Redux store.
- [In the app a user activity generates an event. The event handler calls the dispatch() function that is sending the current state and an action (object) to the rootReducer(). The action object contains the relevant data for the requested change of state slice. The rootReducer() will interpret the action.type, process the data and generate a new state. After the store has received the new state, it triggers the re-render of the React Redux app. It also triggers the execution of all listener functions that are registered with the subscribe() method to the store. Furthermore, all components that are subscribed with connect(mapStateToProps) to the store now receive the new state data as defined in mapStateToProps().](https://medium.com/@javascript_7596/react-redux-concept-workflow-cheatsheet-be00e3ffa853)
- You bind all the transformations on your UI to a single object. All stuff that happens on the UI is a consequence of a transformation on that object. So:
  1. You just gained undo/redo, if you track alterations on your object.
  2. You know the state of your application at all times, just looking at the object. You can also keep it and save it for later use.
- [Avoid using Immutable.js in Redux apps. Use Redux Toolkit & Immer instead.](https://twitter.com/acemarke/status/1213573285314809856)
- [Actions are events, not commands. An action says "this thing happened", and then anything in the store (reducers, middlwares) is allowed to do something with that knowledge. An action does not say "please do this thing".](https://www.reddit.com/r/reduxjs/comments/ed3609/why_redux_advises_against_using_actions_as_setters/)

## Links

- [React Redux — Concept, Workflow & Cheatsheet (2017)](https://medium.com/@javascript_7596/react-redux-concept-workflow-cheatsheet-be00e3ffa853)
- [Redesigning Redux](https://hackernoon.com/redesigning-redux-b2baee8b8a38)
- [React Redux Style guide](https://github.com/iraycd/React-Redux-Styleguide)
- [Mark Erikson - The Fundamentals of Redux](https://www.youtube.com/watch?v=ewelU8WHXQ4&index=7&list=PLRvKvw42Rc7OWK5s-YGGFSmByDzzgC0HP)
- [Redux Persist](https://github.com/rt2zz/redux-persist) - Persist and rehydrate a Redux store.
- [The Suspense is Killing Redux](https://medium.com/@ryanflorence/the-suspense-is-killing-redux-e888f9692430)
- [Getting started with Redux course](https://egghead.io/courses/getting-started-with-redux)
- [Erik Rasmussen | Abstracting Form State with Redux Form (2016)](https://www.youtube.com/watch?v=eDTi7lYR1VU)
- [A Dummy’s Guide to Redux and Thunk in React](https://medium.com/@stowball/a-dummys-guide-to-redux-and-thunk-in-react-d8904a7005d3)
- [redux-actions](https://github.com/redux-utilities/redux-actions) - Flux Standard Action utilities for Redux.
- [Redux Ecosystem Links](https://github.com/markerikson/redux-ecosystem-links) - Categorized list of Redux-related addons, libraries, and utilities.
- [typesafe-actions](https://github.com/piotrwitek/typesafe-actions) - Typesafe Action Creators for Redux / Flux Architectures (in TypeScript).
- [TypeScript FSA](https://github.com/aikoven/typescript-fsa) - Type-safe action creator utilities.
- [Robodux](https://github.com/neurosnap/robodux) - Remove repetitive tasks from Redux.
- [redux-react-hook](https://github.com/facebookincubator/redux-react-hook) - React hook for accessing mapped state and dispatch from a Redux store.
- [Connected React Router](https://github.com/supasate/connected-react-router) - Redux binding for React Router.
- [SoundCloud Redux](https://github.com/r-park/soundcloud-redux) - Basic SoundCloud API client built with React, Redux, and Redux Saga.
- [Logux](https://github.com/logux/logux) - Instead of sending HTTP requests (e.g., AJAX, REST, and GraphQL) it synchronizes log of operations between client, server, and other clients through WebSocket.
- [React Hooks implementation for Redux](https://github.com/epeli/redux-hooks)
- [Scaling a react/redux codebase for multiple platforms (2019)](https://erock.io/scaling-js-codebase-multiple-platforms/)
- [Redux Loop](https://github.com/redux-loop/redux-loop) - Library that ports Elm's effect system to Redux.
- [Redux side effects and me (2017)](https://medium.com/magnetis-backstage/redux-side-effects-and-me-89c104a4b149)
- [Redux without the sanity checks in a single file. Don't use this, use normal Redux. :-)](https://gist.github.com/gaearon/ffd88b0e4f00b22c3159)
- [redux-bundler](https://github.com/HenrikJoreteg/redux-bundler) - Compose a Redux store out of smaller bundles of functionality.
- [Redux Starter Kit](https://github.com/reduxjs/redux-starter-kit) - Simple set of tools to make using Redux easier.
- [Redux DevTools](https://github.com/reduxjs/redux-devtools) - DevTools for Redux with hot reloading, action replay, and customizable UI.
- [Redux-ORM](https://github.com/redux-orm/redux-orm) - Small, simple and immutable ORM to manage relational data in your Redux store.
- [dva](https://github.com/dvajs/dva) - React and redux based, lightweight and elm-style framework. (Inspired by elm and choo).
- [reactive-react-redux](https://github.com/dai-shi/reactive-react-redux) - React Redux binding with React Hooks and Proxy.
- [redux-logic](https://github.com/jeffbski/redux-logic) - Redux middleware for organizing all your business logic. Intercept actions and perform async processing.
- [Re-reselect](https://github.com/toomuchdesign/re-reselect) - Enhance Reselect selectors with deeper memoization and cache management.
- [Redux First History](https://github.com/salvoravida/redux-first-history) - Redux history binding support react-router - @reach/router.
- [Our Redux Anti Pattern](https://rangle.slides.com/yazanalaboudi/deck#/) - Guide to predictable scalability.
- [Reasons to learn Redux as a JavaScript Developer (2019)](https://www.robinwieruch.de/redux-javascript) ([HN](https://news.ycombinator.com/item?id=21926659))
- [Idiomatic Redux](https://blog.isquaredsoftware.com/series/idiomatic-redux/)
- [Is redux really a good idea? (2020)](https://www.reddit.com/r/reactjs/comments/epxavs/is_redux_really_a_good_idea/)
- [How To Not Have A Mess with React Hooks & Redux (2020)](https://orizens.com/blog/how-to-not-have-a-mess-with-react-hooks-and-redux/)
- [React + Redux + Comlink = Off-main-thread (2019)](https://surma.dev/things/react-redux-comlink/)
- [Redux Batch](https://github.com/manaflair/redux-batch) - Enhance your Redux store to support batched actions.
- [Redux Toolkit](https://redux-toolkit.js.org/) - Official, opinionated, batteries-included toolset for efficient Redux development. ([Code](https://github.com/reduxjs/redux-toolkit))
- [Redux Shouldn’t Hold All Of Your Data (2019)](https://michaelwashburnjr.com/2019/12/09/stop-storing-data-redux/)
- [Redux Thunk](https://github.com/reduxjs/redux-thunk) - Thunk middleware for Redux.
- [React context vs Redux in 2020](https://gist.github.com/slikts/57ff1acdb6f5b2ea075b701d1daf896d)
- [Entire Redux in Web Worker](https://github.com/dai-shi/redux-in-worker)
- [Event Driven Redux](https://github.com/dmmulroy/talks/blob/master/event-driven-redux/slides.pdf)
- [redux-act](https://github.com/pauldijou/redux-act) - Opinionated lib to create actions and reducers for Redux. The main goal is to use actions themselves as references inside the reducers rather than string constants.
- [Redux-First Router](https://github.com/faceyspacey/redux-first-router) - Think of your app in terms of states, not routes or components. Connect your components and just dispatch Flux Standard Actions.
- [Redux Devtools: tips & tricks for faster debugging (2020)](https://blog.logrocket.com/redux-devtools-tips-tricks-for-faster-debugging/)
- [Value of Redux Thunks (2020)](https://www.reddit.com/r/reactjs/comments/fmpcou/are_thunks_obsolete/fl5dtvn/?context=3)
- [How UI-driven State Increases Accidental Complexity (2020)](https://evgenii.info/ui-driven-state/) ([HN](https://news.ycombinator.com/item?id=22680369))
- [HN: Redux – Not Dead Yet (2020)](https://news.ycombinator.com/item?id=22822198)
- [Redux Actuator](https://github.com/molefrog/redux-actuator) - Communicate between components through Redux store.
- [Notes on redux design](https://twitter.com/buildsghost/status/1255756148084367361)
- [Dynadux](https://github.com/aneldev/dynadux) - Advanced and simpler Stores based on dispatch and Reducers. Alternative to Redux.
- [Why I Stopped Using Redux (2020)](https://dev.to/g_abud/why-i-quit-redux-1knl)
- [Redux Overview and Concepts](https://redux.js.org/tutorials/essentials/part-1-overview-concepts) ([HN](https://news.ycombinator.com/item?id=23950153))
- [The evil pattern of Redux that reduces boilerplate (2020)](https://www.albertgao.xyz/2020/09/22/evil-pattern-of-redux-that-reduces-boilerplate/)
- [Redux In Web Workers: Off-Main-Thread Redux Reducers and Middleware (2020)](https://blog.axlight.com/posts/redux-in-worker-off-main-thread-redux-reducers-and-middleware/)
- [Refactoring a Redux app to use Recoil (2020)](https://blog.logrocket.com/refactoring-redux-app-to-use-recoil/)
- [Repluggable](https://github.com/wix/repluggable) - Pluggable micro frontends in React+Redux apps.
- [Redux is half of a pattern (2020)](https://dev.to/davidkpiano/redux-is-half-of-a-pattern-1-2-1hd7)
- [RTK Query](https://rtk-query-docs.netlify.app/) - Data fetching and caching addon for Redux Toolkit. ([Code](https://github.com/rtk-incubator/rtk-query)) ([Tweet](https://twitter.com/acemarke/status/1333863983330299904))
- [Functors - Redux (2020)](https://functional.christmas/2020/8)
- [Let’s Learn Modern Redux! (with Mark Erikson) (2021)](https://www.youtube.com/watch?v=9zySeP5vH9c) ([Code](https://github.com/learnwithjason/lets-learn-redux-toolkit))
- [Stop using the “container/presentational” pattern in Redux (2021)](https://medium.com/nmc-techblog/why-you-should-stop-using-the-container-presentational-pattern-in-redux-29b112406128) ([Reddit](https://www.reddit.com/r/reactjs/comments/ox6cwk/why_you_should_stop_using_the/))
- [Typed Redux Saga](https://github.com/agiledigital/typed-redux-saga) - Better TypeScript typing in redux-saga.
- [Why Redux isn't good any more (2021)](https://twitter.com/kamyshev_dev/status/1441736479240122372)
- [localfirst/state](https://github.com/local-first-web/state) - Automatically replicated Redux store that gives your app offline capabilities and secure peer-to-peer synchronization superpowers.
- [Is Redux still recommended in 2021?](https://www.reddit.com/r/reactjs/comments/pwfubd/is_redux_still_recommended_in_2021/)
- [Redux Toolkit Query vs React Query (2021)](https://www.youtube.com/watch?v=LDS1ll93P-s) ([Tweet](https://twitter.com/acemarke/status/1450219978813288448))
- [use-redux-effect](https://github.com/Qeepsake/use-redux-effect) - Powerful React hook that subscribes to Redux store events.
- [Should you use Redux still (2021)](https://twitter.com/traversymedia/status/1466034407609974786)
- [Confidently Testing Redux Applications with Jest & TypeScript (2021)](https://egghead.io/courses/confidently-testing-redux-applications-with-jest-typescript-16e17d9b)
- [How do you know if a project will need Redux? (2022)](https://www.reddit.com/r/reactjs/comments/s4hxd2/how_do_you_know_if_a_project_will_need_redux/)
- [Should we be teaching Redux in 2022?](https://www.reddit.com/r/reactjs/comments/squatd/should_we_be_teaching_redux_in_2022/)
- [Redux Toolkit v1.8: new "listener" side effects middleware](https://github.com/reduxjs/redux-toolkit/releases/tag/v1.8.0) ([Reddit](https://www.reddit.com/r/reactjs/comments/t3lpx9/redux_toolkit_v18_new_listener_side_effects/))
- [Redux is pointless when using react query / Apollo client / swr](https://twitter.com/zachcodes/status/1501573905374400512)
- [Writing Redux Reducers in Rust (2022)](https://fiberplane.dev/blog/writing-redux-reducers-in-rust/) ([Reddit](https://www.reddit.com/r/programming/comments/txi3py/writing_redux_reducers_in_rust/))

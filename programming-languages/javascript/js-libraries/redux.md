# [Redux](https://redux.js.org)

It is common to use React with a Flux pattern to manage the state of the application. Redux is one of the most used and well supported state managers to be used (but not exclusively) with React. The Gist of Redux is a follows:

- The whole state of your app is stored in an object tree inside a single store.
- The only way to change the state tree is to emit an action, an object describing what happened.
- To specify how the actions transform the state tree, you write pure reducers.

## Notes

- OwnPops is something your component receives from the parent. StateProps is something your component receives from the Redux store.
- [In the app a user activity generates an event. The event handler calls the dispatch() function that is sending the current state and an action (object) to the rootReducer(). The action object contains the relevant data for the requested change of state slice. The rootReducer() will interpret the action.type, process the data and generate a new state. After the store has received the new state, it triggers the re-render of the React Redux app. It also triggers the execution of all listener functions that are registered with the subscribe() method to the store. Furthermore, all components that are subscribed with connect(mapStateToProps) to the store now receive the new state data as defined in mapStateToProps().](https://medium.com/@javascript_7596/react-redux-concept-workflow-cheatsheet-be00e3ffa853)
- You bind all the transformations on your UI to a single object. All stuff that happens on the UI is a consequence of a transformation on that object. So:
  1. You just gained undo/redo, if you track alterations on your object.
  2. You know the state of your application at all times, just looking at the object. You can also keep it and save it for later use.

## Links

- [React Redux — Concept, Workflow & Cheatsheet (2017)](https://medium.com/@javascript_7596/react-redux-concept-workflow-cheatsheet-be00e3ffa853)
- [Redesigning Redux](https://hackernoon.com/redesigning-redux-b2baee8b8a38)
- [React Redux Style guide](https://github.com/iraycd/React-Redux-Styleguide#readme)
- [Mark Erikson - The Fundamentals of Redux](https://www.youtube.com/watch?v=ewelU8WHXQ4&index=7&list=PLRvKvw42Rc7OWK5s-YGGFSmByDzzgC0HP)
- [Redux Persist](https://github.com/rt2zz/redux-persist#readme) - Persist and rehydrate a Redux store.
- [The Suspense is Killing Redux](https://medium.com/@ryanflorence/the-suspense-is-killing-redux-e888f9692430)
- [Getting started with Redux course](https://egghead.io/courses/getting-started-with-redux)
- [Erik Rasmussen | Abstracting Form State with Redux Form (2016)](https://www.youtube.com/watch?v=eDTi7lYR1VU)
- [A Dummy’s Guide to Redux and Thunk in React](https://medium.com/@stowball/a-dummys-guide-to-redux-and-thunk-in-react-d8904a7005d3)
- [redux-actions](https://github.com/redux-utilities/redux-actions) - Flux Standard Action utilities for Redux.
- [Redux Ecosystem Links](https://github.com/markerikson/redux-ecosystem-links#readme) - Categorized list of Redux-related addons, libraries, and utilities.
- [typesafe-actions](https://github.com/piotrwitek/typesafe-actions) - Typesafe Action Creators for Redux / Flux Architectures (in TypeScript).
- [TypeScript FSA](https://github.com/aikoven/typescript-fsa) - Type-safe action creator utilities.
- [Robodux](https://github.com/neurosnap/robodux) - Remove repetitive tasks from Redux.
- [redux-react-hook](https://github.com/facebookincubator/redux-react-hook) - React hook for accessing mapped state and dispatch from a Redux store.
- [Connected React Router](https://github.com/supasate/connected-react-router) - Redux binding for React Router.
- [SoundCloud Redux](https://github.com/r-park/soundcloud-redux) - Basic SoundCloud API client built with React, Redux, and Redux Saga.
- [Logux](https://github.com/logux/logux) - Instead of sending HTTP requests (e.g., AJAX, REST, and GraphQL) it synchronizes log of operations between client, server, and other clients through WebSocket.
- [React Hooks implementation for Redux](https://github.com/epeli/redux-hooks)
- [Scaling a react/redux codebase for multiple platforms (2019)](https://erock.io/scaling-js-codebase-multiple-platforms/)

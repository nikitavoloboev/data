# [Redux](https://redux.js.org)

Redux holds the state of your web app. It is one big JavaScript object that contains all the necessary information to render your app at some point in time.

You bind all the transformations on your UI to a single object. All stuff that happens on the UI is a consequence of a transformation on that object.

So:

1. You just gained undo/redo, if you track alterations on your object.
2. You know the state of your application at all times, just looking at the object. You can also keep it and save it for later use.

## Notes

- OwnPops is something your component receives from the parent. StateProps is something your component receives from the Redux store.
- [In the app a user activity generates an event. The event handler calls the dispatch() function that is sending the current state and an action (object) to the rootReducer(). The action object contains the relevant data for the requested change of state slice. The rootReducer() will interpret the action.type, process the data and generate a new state. After the store has received the new state, it triggers the re-render of the React Redux app. It also triggers the execution of all listener functions that are registered with the subscribe() method to the store. Furthermore, all components that are subscribed with connect(mapStateToProps) to the store now receive the new state data as defined in mapStateToProps().](https://medium.com/@javascript_7596/react-redux-concept-workflow-cheatsheet-be00e3ffa853)

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
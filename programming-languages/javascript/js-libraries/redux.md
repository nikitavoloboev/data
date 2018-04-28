# [Redux](https://redux.js.org)
Redux holds the state of your web app. It is one big javascript object that contains all the necessary information to render your app at some point in time.

You bind all the transformations on your UI to a single object. All stuff that happens on the UI is a consequence of a transformation on that object.

So:
1. You just gained undo/redo, if you track alterations on your object.
2. You know the state of your application at all times, just looking at the object. You can also keep it and save it for later use.

## Links
- [Redesigning redux](https://hackernoon.com/redesigning-redux-b2baee8b8a38)
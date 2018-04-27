# [React](https://reactjs.org)
The core premise for React is that UIs are simply a projection of data into a different form of data.

## Notes
- Immediate mode simply means you specify what to redraw on every frame, there is no caching unless you specify it. And you basically redraw whenever some state changes (in a game this is going to be at frame rate).
	- In React, when some state changes, you respecify the DOM for components whose state has changed, but asynchronously the library determines how to make the DOM update more efficient on the next frame redraw.
- Retained-mode means you modify the scene graph (aka DOM) using imperative statements, it is difficult to keep your UI in synch with your models. With immediate-mode, you simply create a function f(m) over your model m to render it on each frame rate (which also often involves imperative instructions affecting the frame buffer, but the buffer can be cleared on each frame so who cares).
	- Retained-mode caches by default (often in opaque ways), which was the whole point (only re-render parts of the scene that have changed). You can roll your own caching for immediate mode, usually via some kind of invalidation scheme (use image for a node if nothing changed for this component, otherwise call that node's re-render method). On the other side, projects like React takes the retained-mode DOM and make it look more like an immediate-mode abstraction without sacrificing so much performance.

## Links
- [A practical understanding of Flux](http://sircmpwn.github.io/2015/07/20/A-practical-understanding-of-Flux.html)
- [Advanced React Component Patterns](https://egghead.io/courses/advanced-react-component-patterns)
- [Introduction to The Beginner's Guide to ReactJS](https://egghead.io/lessons/react-introduction-to-the-beginner-s-guide-to-reactjs)
- [React - Basic Theoretical Concepts](https://github.com/reactjs/react-basic#readme)
- [Didact](https://github.com/hexacta/didact) - DIY guide to build your own React.
- [ReactiFlux Learning](https://www.reactiflux.com/learning/)
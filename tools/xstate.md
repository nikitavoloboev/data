# [XState](https://xstate.js.org/)

I [model state machines by first listing events (press/do something), tasks (side effects, split by services (long running and you care about them ending) and actions (sync, take no time at all)) together with lifecycle (when it's active)](https://www.youtube.com/watch?v=wykDyFwr8Lk). Services can spawn events.

## Notes

- [When should you use XState? When you have something complex to build, with: 1. Many side effects, especially long-running ones like call recording. 2. Sequences of things that MUST happen in a certain order (choose devices, record, preview). Either dead simple (useState) or everything else (XState).](https://twitter.com/mpocock1/status/1451486390911373314)
- [I use XState anytime I have to write my own useState and it is more than a single boolean per component, or anytime I'd use useEffect.](https://twitter.com/triangulo_dev/status/1451533181329547304)

## Links

- [XState Docs](https://xstate.js.org/docs/)
- [Series of examples showing how to model application state with statechart using xstate](https://github.com/coodoo/xstate-examples)
- [XState Visualizer](https://stately.ai/viz) - Visualizer for XState machines. ([Code](https://github.com/statelyai/xstate-viz))
- [Introduction to State Machines Using XState (2021)](https://egghead.io/courses/introduction-to-state-machines-using-xstate)
- [State Machines & XState Workshop](https://github.com/davidkpiano/frontend-masters-xstate-v2)
- [Why XState?](https://kentcdodds.com/calls/01/29/why-x-state) ([Tweet](https://twitter.com/mpocock1/status/1451486390911373314))
- [XState: the Visual Future of State Management (2021)](https://www.youtube.com/watch?v=4Mue0Qr_apE)
- [xstate-awaitable-send](https://github.com/sebinsua/xstate-awaitable-send) - Fire an event into an XState `Machine` and then wait for it to resolve.
- [Frontend Masters React + XState workshop (2020)](https://frontendmasters.com/workshops/xstate-react/) ([Code](https://github.com/davidkpiano/frontend-masters-react-workshop))
- [7guis-xstate-stream](https://github.com/with-heart/7guis-xstate-stream) - Solve The 7 Tasks of 7GUIs using xstate.
- [xstate-parser](https://github.com/statelyai/xstate-parser) - Experiments in building a tool for XState to parse basic information about machines declared inside files.
- [Why XState is THE State Management Tool (2022)](https://clevertech.biz/insights/why-x-state-is-the-state-management-tool)
- [XState Docs Speedrun - 2022](https://www.youtube.com/watch?v=2eurRx-tR-I)
- [xState+react-query-integration - CodeSandbox](https://codesandbox.io/s/u53ey)
- [XState React Template - CodeSandbox](https://codesandbox.io/s/xstate-react-template-3t2tg?file=/src/index.js)
- [Modelling Statecharts: A step-by-step guide (2022)](https://www.youtube.com/watch?v=wykDyFwr8Lk)
- [Designing State Machines #1 - Fetch Logic](https://www.youtube.com/watch?v=Km8QGtRr1z8)

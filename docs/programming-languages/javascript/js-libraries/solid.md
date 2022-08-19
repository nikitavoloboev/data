---
title: SolidJS
---

# [SolidJS](https://www.solidjs.com/)

> Simple and performant reactivity for building user interfaces

Using Solid for all my web apps and static sites. Love how it takes all the best parts of [React](../js-libraries/react/react.md) (component tree model with one way data flow & JSX) but [does it better](https://www.youtube.com/watch?v=UhGV8yYnvQE) due to [fine grained reactivity ](https://dev.to/ryansolid/a-hands-on-introduction-to-fine-grained-reactivity-3ndf) & avoiding VDOM thus being [much faster than React](https://twitter.com/nikitavoloboev/status/1528479450828087299).

[Ryan Carniato](https://twitter.com/RyanCarniato) does [great YouTube videos on web/solidJS](https://www.youtube.com/c/RyanCarniato9/videos) ([SolidJS: Past, Present, Future](https://www.youtube.com/watch?v=dOgnQ9SuKYk) is insightful, especially part about [future focus](https://www.youtube.com/watch?v=dOgnQ9SuKYk&t=7555s)). [Solid Start](https://github.com/solidjs/solid-start) is great starter template. [solid-primitives/fetch](https://www.npmjs.com/package/@solid-primitives/fetch) is great alternative to [React Query](https://react-query.tanstack.com/).

[Hope UI](https://hope-ui.com/) is nice component library. [Motion One for Solid](https://motion.dev/solid/quick-start) is great for animations.

My [personal site](https://github.com/nikitavoloboev/nikitavoloboev) & [LA](https://github.com/learn-anything/learn-anything/) are OSS & built with Solid.

## OSS apps

- [Solid Hacker News](https://github.com/solidjs/solid-hackernews)
- [Solid Storefront](https://github.com/zaiste/solid-storefront) - Open-source e-commerce storefront in Solid.js with TypeScript, GraphQL, URQL & Tailwind CSS.
- [Chrome Extension Boilerplate with SolidJS + Vite + TypeScript](https://github.com/fuyutarow/solid-chrome-extension-template)

## Notes

- [To me the biggest difference between React and Solid is that things that can change are wrapped in signals in Solid, and in dependencies arrays in React.](https://twitter.com/fabiospampinato/status/1528537000504184834)
- [Components are rendered only once and are there pretty much to organize the code. The built-in components (Show, For) are awesome and should be used instead of following React-type patterns (multiple dynamic returns, map(), etc).](https://www.youtube.com/watch?v=Ilf34WjMBkU)

## Links

- [SolidJS New Docs](https://docs.solidjs.com/) ([Code](https://github.com/solidjs/solid-docs-next))
- [solid-nanostores](https://github.com/nanostores/solid) - Global state management in Solid using Nano Stores.
- [Solid Lib Starter](https://github.com/amoutonbrady/solid-lib-starter)
- [Solid Three, Custom Renderers, and SolidStart w/ Nikhil Saraf (2022)](https://www.youtube.com/watch?v=lsWXyyEsw7E)
- [Solid App Router](https://github.com/solidjs/solid-app-router) - Universal router for Solid inspired by Ember and React Router.
- [Let's do an AMA (2022)](https://www.youtube.com/watch?v=8_YiKUb6DW8)
- [Solid Command Palette](https://github.com/itaditya/solid-command-palette) - UI Library for Command Palette in SolidJS webapps.
- [Solid Start](https://github.com/solidjs/solid-start) - Solid's official starter.
- [A Solid option for building UIs (2022)](https://overcast.fm/+Id5U6EB5M)
- [Solid Select](https://github.com/thisbeyond/solid-select) - Select component for Solid.
- [Solid Docs Code](https://github.com/solidjs/solid-docs)
- [Solid Labels](https://github.com/LXSMNSYC/solid-labels) - Simple, reactive labels for SolidJS.
- [Solid Headless](https://github.com/LXSMNSYC/solid-headless) - Headless UI for SolidJS.
- [Streaming SolidJS: Fine-Grained Reactivity (2021)](https://www.youtube.com/watch?v=b9e7VXs_A4s)
- [solid-firebase](https://github.com/wobsoriano/solid-firebase) - Solid hooks for Firebase.
- [Solid Primitives](https://github.com/davedbase/solid-primitives)
- [solid-ui](https://github.com/solid/solid-ui) - User Interface widgets and utilities for Solid.
- [vite-plugin-solid](https://github.com/solidjs/vite-plugin-solid) - Simple integration to run solid-js with Vite.
- [Solid MultiSelect](https://github.com/DigiChanges/solid-multiselect) - Multi-select implementation for Solid.
- [Solid.js feels like what I always wanted React to be](https://typeofnan.dev/solid-js-feels-like-what-i-always-wanted-react-to-be/) ([HN](https://news.ycombinator.com/item?id=30508524))
- [Solid Starter Kit](https://github.com/one-aalam/solid-starter-kit)
- [Solid Styled](https://github.com/LXSMNSYC/solid-styled) - Reactive stylesheets for SolidJS.
- [Solid Search for Community Solid Server](https://github.com/ontola/solid-search-community-server) - Adds full-text search to Community Solid Server. Powered by atomic-server.
- [Blitz](https://github.com/dimensionhq/blitz) - Fast, simple and efficient state management for SolidJS apps.
- [solid-boundaries](https://github.com/everweij/solid-boundaries) - Utility to track the boundaries of html-elements in SolidJS.
- [i18next for Solid](https://github.com/mbarzda/solid-i18next)
- [Solid FLIP](https://github.com/otonashixav/solid-flip) - Lightweight transition library for Solid.
- [Solid Playground](https://playground.solidjs.com/) - Quickly discover what the solid compiler will generate from your JSX template. ([Code](https://github.com/solidjs/solid-playground))
- [solid-utils](https://github.com/amoutonbrady/solid-utils) - Ultimate companion of all your solid-js applications.
- [Solid Slider](https://github.com/davedbase/solid-slider) - Carousel/slider implementation for SolidJS.
- [SolidHack 2022](https://hack.solidjs.com/) ([Code](https://github.com/solidjs/solidhack-submissions))
- [solid-spring](https://github.com/Aslemammad/solid-spring) - Port of react-spring, for SolidJS.
- [Solid Ninja Keys](https://github.com/wobsoriano/solid-ninja-keys) - Add cmd+k interface to your solid site.
- [SolidJS plugin for Preview.js](https://github.com/fwouts/previewjs-solid-plugin)
- [Solid JSX](https://github.com/high1/solid-jsx) - Use mdx or xdm with solid-js.
- [Solid-Blocks](https://github.com/atk/solid-blocks) - UI building blocks for SolidJS.
- [Solid + Netlify Edge functions powered by Deno HN Demo](https://github.com/solidjs/solid-hackernews/tree/netlify-edge) ([Tweet](https://twitter.com/RyanCarniato/status/1516510395787218944))
- [Solid Aria](https://github.com/solidjs-community/solid-aria) - Library of high-quality primitives that help you build accessible user interfaces with SolidJS.
- [Motion One for Solid](https://motion.dev/solid/quick-start) - Combines Solid's amazing performance with declarative WAAPI animations, independent transforms, springs, stagger, and more. ([Tweet](https://twitter.com/motiondotdev/status/1520025108403863554))
- [Solid Services](https://github.com/Exelord/solid-services) - Solid.js library adding a services layer for global shared state.
- [Solid Proxies](https://github.com/Exelord/solid-proxies) - Solid.js library adding signaling to built-in non-primitives.
- [Solid DND](https://github.com/thisbeyond/solid-dnd) - Lightweight, performant, extensible drag and drop toolkit for Solid. ([Web](https://solid-dnd.com/)) ([Web Code](https://github.com/thisbeyond/solid-dnd-site))
- [Turbo Solid](https://github.com/StudioLambda/TurboSolid) - Lightweight asynchronous data management for solid.
- [Solid in 100 Seconds (2022)](https://www.youtube.com/watch?v=hw3Bx5vxKl0)
- [solid-primitives/fetch](https://www.npmjs.com/package/@solid-primitives/fetch) - Creates a primitive to support abortable HTTP requests. If any reactive request options changes, the request is aborted automatically.
- [Praises of SolidJS (2022)](https://twitter.com/andrewgreenh/status/1526908199386988544)
- [Let's Learn SolidJS (2021)](https://www.youtube.com/watch?v=ZZ-a7B761Ds)
- [Hope UI](https://github.com/fabien-ml/hope-ui) - SolidJS component library you've hoped for. ([Docs](https://hope-ui.com/docs/getting-started))
- [Solid DSL Experiments](https://github.com/solidjs-community/solid-dsl)
- [Solid Meta](https://github.com/solidjs/solid-meta) - Asynchronous SSR-ready Document Head management for Solid based on React Head.
- [Early return in Solid](https://twitter.com/Huxpro/status/1526711254928068608)
- [The Quest for ReactiveScript (2021)](https://dev.to/this-is-learning/the-quest-for-reactivescript-3ka3)
- [solid-auto-animate](https://github.com/LXSMNSYC/solid-auto-animate) - SolidJS bindings for FormKit's AutoAnimate.
- [Solid Dev Tools](https://github.com/thetarnav/solid-devtools) - Reactivity Debugger & Devtools Chrome Extension for SolidJS.
- [solid-register](https://github.com/atk/solid-register) - Solid.js execution environment for Node.js. Allows running and testing Solid.js browser code in Node.js.
- [solid-marked](https://github.com/LXSMNSYC/solid-marked) - MDX/Markdown compiler for SolidJS.
- [Solid Native](https://github.com/MrFoxPro/solid-nativescript-experiments) - NativeScript bindings for Solid.JS.
- [Solid Templates (using Vite)](https://github.com/solidjs/templates)
- [Solid Heroicons](https://github.com/amoutonbrady/solid-heroicons)
- [SolidJS - YouTube](https://www.youtube.com/channel/UCXsRnrbzIX8KHdf86PE241Q/videos)
- [Carbon Components (SolidJS Port)](https://github.com/mosheduminer/carbon-components-solid)
- [SolidJS Effector](https://community.effector.dev/core/prerelease-version-of-effector-solid-is-now-available-for-public-use-eei)
- [Solid Native](https://github.com/tjjfvi/solid-native) - Expands the Solid ecosystem to include development of native mobile applications.
- [typesafe-i18n Solid](https://github.com/ivanhofer/typesafe-i18n/tree/main/packages/adapter-solid)
- [Solid-Use](https://github.com/LXSMNSYC/solid-use) - Collection of SolidJS utilities.
- [rollup-preset-solid](https://github.com/amoutonbrady/rollup-preset-solid) - Small opinionated preset for rollup to bundle your solid libraries with rollup.
- [SolidJS: Past, Present, Future (2022)](https://www.youtube.com/watch?v=dOgnQ9SuKYk)
- [Solid JS and How its Different | Ryan Carniato (2022)](https://www.youtube.com/watch?v=O4sgwuMQns0)
- [Solid Virtual Scroll](https://github.com/Supertigerr/solid-virtual-scroll)
- [Solid Cached Resource](https://github.com/yonathan06/solid-cached-resource) - Cache Solidjs resources by key (derived from the resource source).
- [SUID](https://github.com/swordev/suid) - Port of Material-UI (MUI) built with SolidJS.
- [Solid Hotkeys](https://github.com/alekangelov/solid-hotkeys) - Hotkeys library for solid that adds only 2 event listeners.
- [How to Deploy SolidJS (2022)](https://dev.to/brittneypostma/how-to-deploy-solidjs-4hoi)
- [Solid Service API](https://github.com/solidjs/solid-service-api) - Code that powers Solid Service API on Cloudflare Workers.
- [Opinionated Vite Starter Template](https://github.com/olgam4/bat)
- [Solid Custom Scrollbars](https://github.com/diragb/solid-custom-scrollbars)
- [SolidJS: Reactivity Unchained – Ryan Carniato (2022)](https://www.youtube.com/watch?v=UhGV8yYnvQE)
- [JavaScript UI Compilers: Comparing Svelte and Solid (2019)](https://ryansolid.medium.com/javascript-ui-compilers-comparing-svelte-and-solid-cbcba2120cea)
- [Thinking Granular: How is SolidJS so Performant? (2021)](https://dev.to/ryansolid/thinking-granular-how-is-solidjs-so-performant-4g37)
- [Solid Valtio](https://github.com/wobsoriano/solid-valtio) - State management in Solid with valtio.
- [ReScript SolidJS](https://github.com/Fattafatta/rescript-solidjs)
- [Solid Cache](https://github.com/LXSMNSYC/solid-cache) - Resource caching in SolidJS.
- [SolidJS Package Monorepo Starter](https://github.com/solidjs-community/monorepo-starter)
- [Nested Routing, Parallelized Data Fetching, and SolidJS w/ Ryan Turnquist (2022)](https://www.youtube.com/watch?v=-TeXU9UZ_4w)
- [Solid Floating UI](https://github.com/LXSMNSYC/solid-floating-ui)
- [Solid Future Working Group](https://github.com/solidjs/solid-workgroup/discussions)
- [Rigidity](https://github.com/LXSMNSYC/rigidity) - SSR framework for SolidJS. ([Tweet](https://twitter.com/lxsmnsyc/status/1552312144951005185))
- [Rallax](https://github.com/thetarnav/rallax) - Parallax library for SolidJS.
- [Solid-Dexie](https://github.com/faassen/solid-dexie) - Dexie integration for Solid.
- [SolidJS bindings for atomic-router](https://github.com/atomic-router/solid)
- [Introducing SolidStart w/ Nikhil Saraf (2022)](https://www.youtube.com/watch?v=eAwuPvRXNdY)
- [Solid Toast](https://github.com/ardeora/solid-toast) - Customizable Toast Notifications for SolidJS.
- [Solid-Query](https://github.com/ardeora/solid-query) - Tanstack Query Adapter For Solid JS.
- [Solid Jest](https://github.com/solidjs/solid-jest) - Jest preset for SolidJS.
- [Solid Testing Library](https://github.com/solidjs/solid-testing-library) - Simple and complete Solid testing utilities that encourage good testing practices.
- [Solid-Pebble](https://github.com/lxsmnsyc/solid-pebble) - State management library for SolidJS.

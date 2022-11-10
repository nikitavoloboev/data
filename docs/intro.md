---
slug: /
title: My Knowledge Wiki
---

# [My Knowledge Wiki 🌿](https://wiki.nikiv.dev)

This is my personal wiki where I share [everything I know](sharing/everything-I-know.md) about this world in form of an [online wiki](other/wiki-workflow.md) built with [Docusaurus](tools/docusaurus.md) on [GitHub](https://github.com/nikitavoloboev/knowledge).

If this is your first time visiting this wiki, take a look [here](#getting-started) as it describes this wiki, its structure and goals.

Below are my 1,000+ notes & their relations visualized.

![](https://i.imgur.com/SbhfKGm.png)

![](https://i.imgur.com/ODDdwde.png)

This wiki is large. 100,000+ lines of markdown large (calculated with [loc](https://github.com/cgag/loc))

```
~/Dropbox/Write/knowledge
❯ loc
--------------------------------------------------------------------------------
 Language             Files        Lines        Blank      Comment         Code
--------------------------------------------------------------------------------
 Markdown               1000        95000         5000            0        100000
```

## Getting started

This garden is quite literally my digital brain. It includes my thoughts, notes and links on topics I care about.

I [update the notes](other/wiki-workflow.md) daily. I use the website and the [search plugin](https://github.com/nikitavoloboev/alfred-my-mind) I built for it together with the search bar above in right corner (or CMD+K hotkey) to query the content inside.

The content being markdown files [found here](https://github.com/nikitavoloboev/knowledge/tree/main/docs), which after every commit to the [GitHub repo](https://github.com/nikitavoloboev/knowledge) builds the website using [Docusaurus](tools/docusaurus.md) and publishes it to [wiki.nikiv.dev](https://wiki.nikiv.dev) from which you are likely reading this page.

The way I structure each page often looks the same and follows a structure. The sidebar you see on the left is sorted alphabetically. It's nested too. Many top level folders are nested quite deep. For example take [programming languages](programming-languages/programming-languages.md), it's top level folder (you can see [all the folders and structure of them here](https://github.com/nikitavoloboev/knowledge/blob/main/SUMMARY.md)). All of these pages are part of programming languages:

![](https://i.imgur.com/wBj77ch.png)

Take [SolidJS](programming-languages/javascript/js-libraries/solid.md) page.

![](https://i.imgur.com/2DdQKxl.png)

As you can see, SolidJS is part of [JS libraries](programming-languages/javascript/js-libraries/js-libraries.md) (big file) which is part of [JavaScript](programming-languages/javascript/javascript.md) and JavaScript as a language is part of top level [programming languages](programming-languages/programming-languages.md) I mentioned above. There are tons more folders like this. The way I include what's part of what is subject to my interpretation so it might be confusing to some.

What I would suggest if this is the first time you came across this wiki and genuinely want to learn something new is to scroll the sidebar on the left or search for something, more likely then not, I wrote something about it.

Lots of this content will one day live or co exist with [Learn Anything](ideas/learn-anything.md) project I work on.

If you use mac, [Alfred My Mind](https://github.com/nikitavoloboev/alfred-my-mind) plugin is incredible. I plan to make both the plugin and this setup with the wiki reusable by anyone.

Most often I share my life on [Twitter](https://twitter.com/nikitavoloboev) (most everything), [Instagram](https://www.instagram.com/nikitavoloboev/) (photos, stories from life, Q&A's) and also [Telegram](https://t.me/niki_log) (raw and more private thoughts/photos/updates I don't necessarily want to push to all my follower's feeds).

I also want to learn to write well written [articles](sharing/my-articles.md) and make high quality [YouTube videos](https://www.youtube.com/channel/UCEKqrUfr_FMKIO9XSJS4vDw/videos).

And of course I love [writing code](sharing/my-github.md) that solves various problems I have and I share it all on [GitHub](https://github.com/nikitavoloboev). I also love [Reddit](https://reddit.com/user/nikivi), [HN](https://news.ycombinator.com/user?id=nikivi) & [Lobsters](https://lobste.rs/u/nikivi) as far as communities go.

[Here](sharing/sharing.md) you can find all the things I made and shared thus far.

That was long. 😺

I tried to write it as both a reference you can skim to get the best info on how to use this wiki or ideally inspire you to make a wiki of your own. There are [many great wikis out there](other/wiki-workflow.md#similar-wikis-i-liked) and the list grows with every day. [All the tools are there](other/wiki-workflow.md).

## Must read pages

There's 1,000+ pages in here and lots of them I think are interesting. But these few can be nice starts for a read:

- [Solving Problems](research/solving-problems.md) describes my approach to find problems to solve and solving them effectively. 
- [Karabiner](macOS/apps/karabiner/karabiner.md) as it describes a genuinely life changing tool that makes updating this wiki and operating my mac at fast speeds possible
- [VSCode](text-editors/vs-code/vs-code.md) & [Sublime Text](text-editors/sublime-text/sublime-text.md) for similar reason as Karabiner, life changing tools. I use Sublime Text for [editing this wiki](other/wiki-workflow.md) and VSCode for writing [code](programming/programming.md). Mostly [TypeScript](programming-languages/typescript/typescript.md) and [Go](programming-languages/go/go.md) now.
- [Happiness](life/happiness.md) includes my lessons I learned for how to live a happier life
- [Focusing](focusing/focusing.md) as it's a very important page for me. It links to other important to me pages: [Rules](focusing/rules.md) & [Habits](focusing/habits.md). These 2 pages are essentially a summary of my life as far as my value system / lessons learned / habits go.
- [Relationships](relationships/relationships.md) is another important to me area in life as I want to be someone friends think fondly of and it contains my thoughts on that
- [My Workflow](sharing/my-workflow.md) includes a summary of what I use in life.
- [Looking Back](looking-back/looking-back.md) as it's my life diary. It includes a lot of how I felt, what I thought about and did at any point in time starting from 2017 when I first started writing a journal.

There is a lot more out there but I think above is a good start. Hope you find some of it interesting. ♥️

As far as tech goes, my current obsessions, [goals](focusing/goals.md) and [ideas](ideas/ideas.md) for how to live a [better future](future/future.md) are [learning](education/learning.md) more of [Solid](programming-languages/javascript/js-libraries/solid.md) and [Go](programming-languages/go/go.md)/[Swift](programming-languages/swift/swift.md) to build [amazing experiences](design/design-inspiration.md).

## Content Structure

As mentioned, I like to keep a sort of structure for pages. It's roughly:

1. Title with optional description to describe topic.
2. My thoughts on the topic.
3. Subtopics - Various subtopics related to the main topic. i.e. [Alfred](macOS/apps/alfred/alfred.md) page has a few subtopics.
4. Notes - Notes on the topic as well as things I found interesting on the internet regarding the topic. I often give a link of where I got things from.
5. Code - Code snippets.
6. Links - Links related to the topic.

## Helping improve the wiki

If you find some mistake, especially if something that I say is plain wrong, please fork [this repository](https://github.com/nikitavoloboev/knowledge) and make a PR with correct changes. Or [open an issue](https://github.com/nikitavoloboev/knowledge/issues/new) saying what you think is wrong, asking questions or making suggestions. Any feedback if it's constructive is welcome. So as any suggestion to improve the content. There is also GitHub [file search](https://github.com/nikitavoloboev/knowledge/find/main) & [content search](https://github.com/nikitavoloboev/knowledge/search?q=karabiner&unscoped_q=karabiner) available that's quite neat.

## Grow your own digital garden 🌱

As I mentioned, these wikis are also called [digital gardens](https://joelhooks.com/digital-garden). There is a [DigitalGardens community on Reddit](https://www.reddit.com/r/DigitalGardens/) which discusses tools to maintain & create these wikis. As well as how to use them to develop [new transformative tools for thought](https://numinous.productions/ttft/).

I collected a list of [wikis I liked the most](other/wiki-workflow.md#similar-wikis-i-liked) for inspiration. For example, [Devine's Wiki](https://wiki.xxiivv.com/site/home.html) is a great rabbit hole of fascinating info. Devine is one half of [100 Rabbits](https://100r.co/site/home.html) which is fascinating couple that [lives on a boat and does art](https://www.youtube.com/watch?v=BW32yUEymvU). I want to live a life like this one day but alas.

I ♥️ meeting new people so feel free to [reach out](https://twitter.com/nikitavoloboev). I also hope you share what you know with the world as a form of a wiki like this too. It truly had a life changing effect on my life.

## Thank you

You can support me on [GitHub](https://github.com/sponsors/nikitavoloboev) or look into [other projects](https://nikiv.dev/projects) I shared. ♥️

[![CC4](https://img.shields.io/badge/license-CC4-0a0a0a.svg?style=flat&colorA=0a0a0a)](https://creativecommons.org/licenses/by/4.0/) [![Twitter](http://bit.ly/nikitatweet)](https://twitter.com/nikitavoloboev)

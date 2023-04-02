# 2023 March

![](https://images.nikiv.dev/broke-again-23.jpeg)

28 next month. Can't afford food. Got [fired for 8th time](../../health/depression.md) this month.

Nothing new under the sun. Except I hoped my life would get better when I was younger. 😹

## Inlang

Worked for a total of 3 weeks at [Inlang](https://inlang.com/). It's great company/team but way too smart for me.

My job was supposed to be in taking over the Git SDK. The [thesis of Inlang](https://www.youtube.com/watch?v=CZr6A5gwmFs) is that you can use git as the database to solve many problems that a lot of existing solutions only try to hotfix. Such as localizations.

For that you need a great Git SDK that builds features on top like live collaboration, fast loading of git repo details on the web and more.

I joined and tried to evaluate [libgit2](https://libgit2.org/), specifically [WASM Git](https://github.com/petersalomonsen/wasm-git) as potential alternative to at the time used [isomorphic-git](https://isomorphic-git.org/).

Eventually that led to me having to write an [RFC on choosing between libgit2 and isomorphic-git for Git SDK](https://github.com/inlang/inlang/pull/455). I loved the experience of writing the RFC and will certainly try doing it more for big decisions around my projects.

As I was writing the RFC and trying to make WASM Git work as a package with Rollup together with extending isomorphic-git to do shallow clones (only fetching git metadata + file blobs as you need them), it was clear to Inlang, I was not a fit. They needed someone who knows Git better.

It happens. Life goes on and all. I did learn a lot about git in these 3 weeks.

I actually wrote code daily with only few breaks to walk to/from library where I worked at. Used [ChatGPT](../../machine-learning/chatgpt.md) heavily to debug problems or learn and it was enjoyable.

In summary, [everything in git is built on top of hashes](https://www.youtube.com/watch?v=ig5E8CcdM9g). Files are hashed. Commits are are hashes of the hashed committed files. Branches are hashes of commits.

Also [this talk on using Git](https://www.youtube.com/watch?v=4EOZvow1mk4) was amazing.

In summary, I did love working the 3 weeks there. I learned a lot and it was my first ever opportunity to work on an open source project and actually get paid to do it. I love the team and the vision to build git based apps. And wish them the best.

I love how they use Discord for all communication with nice well purposed channels. The textual checkins were great. I worked in companies that did morning video checkins and thought it was unnecessary. Textual checkins are perfect enough honestly.

Also Inlang is [built using Solid](https://github.com/inlang/inlang) and I will be using quite a few patterns from the code in my own future projects. Specifically the way Solid stores are used as the editor loads and renders. Everything is a signal.

If I wasn't as bad of a programmer, it'd be a dreamy place to work at for sure. But alas.

## Retro

Don't know what to say. I am trying. Maybe I am not smart enough and should exercise more or maybe programming is not for me.

I remember once, on my 2nd job I ever got, when I got fired. CEO called me to his office to talk to me personally. He said I should go back to school and programming is not for me. I should reconsider this profession and choose something else. Because I will never earn money in this industry with my brain. Was an interesting walk from his office afterwards.

That was the 2nd time I was fired. First job I got, i was fired for not being good enough. Before the first job, I applied to [YC](https://www.ycombinator.com/) with [LA project](../../ideas/learn-anything.md) and got rejected 3 times. I got the message then and stopped applying.

After that fun talk with CEO who told me to reconsider programming as profession I got fired 6 more times. The 8th one this month in 2023.

Maybe I should just give up indeed?

There is a certain sense of calmness in accepting that you are who you are. I am not smart or great in any way. The fact that I can't afford food or to pay tax is nothing extraordinary too. Many people are in same situation as I am and are making by. Why should I complain?

I will die eventually anyway, why suffer more beyond the suffering that is unavoidable in the moment.

I want to be truthful to myself especially in these look backs. These are the raw facts and thoughts. I need to get out of this mess that I have been for past 28 years of my life.

I need to build paid products and do it daily. Every minute of my days before something comes of it.

I don't want to get another job and get fired again. It's not good for anyone. I am okay with nearly starving. Human body doesn't need a lot of food to make it through. As for tax, I will try figure out something. Maybe I will close this wiki and all the open source projects I have. And will never do open source again because it brought me nothing in monetary value and is just suicidal to keep going.

## KusKus

For next month, the goal is to build a todo app I always wanted to build. There's many todo apps. But this one will be different. It even has a name, KusKus.

How it looks now:

![](https://images.nikiv.dev/kuskus-start-23.png)

> Started building the app with my younger brother. Solid code can be so nice to work with.

The app will have hotkeys similar to [2Do](../../macOS/apps/2do.md). Will have integration with GitHub issues similar to now dead [Ship](https://www.realartists.com/blog/ship-20.html). And it will be AI integrated.

For example, you write a task, press a hotkey and GPT will create subtasks for you. There will be live collaboration, nice design and one killer feature for why I currently use [Obsidian](../../tools/obsidian.md).

![](https://images.nikiv.dev/obsidian-todo-setup-23.png)

You can write todos into a text area with vim bindings. Just write things and indent as if it's a text file. And with hotkey, it will turn the text area into actual todos with due dates and all. [TaskPaper](https://www.taskpaper.com/) is similar to this in style.

Essentially I want to build a todo app I always wanted to exist.

The todo app and [LA](../../ideas/learn-anything.md). 2 projects that will take all my time going forward.

I also want to build my own [ChatGPT](../../machine-learning/chatgpt.md) app. I now use [this Tauri desktop wrapper](https://github.com/lencx/ChatGPT). It's nice. But I want ability to pipe results into GPT from selected text with pre prompts. And I want to build a better service for sharing GPT prompts than the existing [ShareGPT](https://sharegpt.com/).

I truly don't care about even outcomes now. I set [goals](../../focusing/goals.md) and I try reach them. And I fail and fail until I reach them. Failure forces you to stay real and not live in dream lands of where you think you're going or even where you think you are right now.

Only thing I personally should start doing though is changing my approach a bit. I feel like someone who bashes his head against the wall expecting it to break after a certain point but it never breaks.

Will be using [Solid](../../programming-languages/javascript/js-libraries/solid.md) + [Grafbase](../../networking/graphql/grafbase.md) to build things out. Also [Ark](https://github.com/chakra-ui/ark) with some [Tailwind](../../front-end/css/tailwind-css.md) seems nice for UI.

## Back to Spain

Went back to Spain this month too to see my family and 🐶. Love them a lot.

![](https://images.nikiv.dev/rachel-outside-march23.jpeg)

> Rachel, my 15 year old adorable labrador

I once checked how long labradors live on average and I should not have done that. Apparently its only 10-12 years.

Thus I try to spend majority of my time to make life as nice possible for her. We go on longer walks and generally try to have a lot of fun.

I also try to not skimp on quality dog food despite my dire financial situation. I would rather not eat than give lesser quality food for dogs.

## Watched

I did get to watch [Scenes from a Marriage](https://trakt.tv/shows/scenes-from-a-marriage-2021) with my partner and loved it. It reaffirmed to me that most issues in relationships stem from miscommunication or lack of it. Also showed well how being hurt really bad once by someone closes you of to ever be fragile again with that person.

## Listened

![](https://images.nikiv.dev/listened-march-23.png)

[We Came And Left](https://open.spotify.com/track/2bKVbvI4F0FktcUXrQ9vIe) is great little tune.

Most of the songs I listen to now are instrumental.

---
title: Twitter
---

# [Twitter](https://twitter.com/)

Use [Simplified Twitter](https://github.com/brunolemos/simplified-twitter) extension with [couple css tweaks](#css-code-to-clean-up-twitter) to hide distracting elements when I browse Twitter on [Safari](../web/browsers/safari.md) in [separate window](https://twitter.com/nikitavoloboev/status/1601959862648582145) (with `jk` [keyboard binds](https://help.twitter.com/en/using-twitter/how-to-tweet) to move between tweets, `l` to like, `b` to bookmark, `r` to reply). And official app on iOS.

I love Twitter as it's currently still the best place to stream your thoughts to public and have others interact with it in real time. Nothing yet comes close to the network and UX of Twitter for doing this. Except maybe [Elk](https://elk.zone) client for [Mastodon](../social-networks/mastodon.md).

I have trust Twitter will improve despite questionable past decisions. They will [make mistakes, roll things back, learn and adapt](https://twitter.com/esthercrawford/status/1590386711179464705). My main concern with Twitter is the lack of transparency around decision making involving suspensions, the amount of bots spamming notifications and DMs and that the code is not open source.

I'd also love it if they could make bookmarking tweets fast on mobile. Ideally by allowing users to swap the view count on tweets in the feed with a bookmark icon similar to what you see when you open a tweet in full view. I often use bookmarks for interesting things I want to see in detail later and it takes too long time to bookmark now. 😞

Also the fact that Twitter still shows [your own tweets on your timeline feed](https://twitter.com/mountain_ghosts/status/1628509280625197059) is insane to me. Would also love it if you could make the badge on notification icon to [only show the number for new replies](https://twitter.com/nikitavoloboev/status/1629085529265393668).

I'd also love for something like [Nostr](../social-networks/nostr.md) or [Farcaster](../social-networks/farcaster.md) protocols to power a service like Twitter with all the UX niceness that Twitter has. [Decentralization](../networking/decentralization.md) has its issues but I think in the long term, a true public web square should be in the power of the community.

I love the idea of freedom of speech, but not freedom of reach whereby anyone can have a voice but the platform won't boost hurtful or perhaps verifiably incorrect rhetoric. Twitter's [Community Notes](https://twitter.com/CommunityNotes) is great in that regard. Will see how Twitter fares at making this work.

[Mastodon](../social-networks/mastodon.md) is interesting alternative but needs smoother onboarding, client UX and ability to have 'public likes' for it to ever make me want to use it more. I get a lot of my content from seeing likes of people I admire. Not to mention that with Mastodon, you still have to follow the rules of the server you joined, otherwise you get suspended with all content deleted same as Twitter. That said, [Elk client](https://github.com/elk-zone/elk) is gorgeous. And shows how much innovation can happen if all development was not done behind closed doors.

I am trying to build my own [app with a network of people that's based on sharing your learnings and ideas](../ideas/learn-anything.md). You would be able to share notes, ideas in similar way I do here in this wiki but more streamlined. With everything being [open source](https://github.com/learn-anything/learn-anything).

I dislike that Twitter gives you no tools to actually build some kind of knowledge system off your own or other's tweets. A dream Twitter like network in my eyes would treat knowledge as first class thing, from which you can share parts or all of your knowledge with all or 'publish/tweet' parts of it out to your followers. But also have a place to catalogue, share and collaborate on ideas together.

Twitter will always have a special place and it will be hard to beat it at what it does best despite the fact that it's not transparent and not open source. But will see. TikTok showed that it's possible to disrupt even such giant networks if you position/market your product well.

As mentioned before, Twitter not being open source is horrible. I find it [frustrating that Twitter iOS client doesn't have the ability to long press `like` icon to bookmark](https://twitter.com/nikitavoloboev/status/1590403314264334338). If the code was open source, I would have gladly made PR for this already. For some reason non official Twitter clients have worse UX feed content so moving from official clients is a non solution for me. It'd also be great if when you visit Twitter profile, it would tell you instantly if that person is added to a list of yours.

In many ways it's a matter of time that Twitter gets replaced with an open source transparent protocol that mimics all the UX features of Twitter. Innovate or die.

[Official Twitter TS SDK](https://github.com/twitterdev/twitter-api-typescript-sdk) is nice for building things on top of Twitter.

[How does Twitter work](https://www.youtube.com/watch?v=z6xslDMimME) is nice overview of [Twitter tech](https://twitter.com/_dml/status/1592619959947243520).

[Bird SQL](https://www.perplexity.ai/sql) is nice way to search Twitter.

## CSS code to clean up Twitter

I use [Simplified Twitter](https://github.com/brunolemos/simplified-twitter) to clean up many things.

On top of the extension, I use [Cascadea](https://cascadea.app/) extension to on Safari to hide some remaining annoying/distracting elements with CSS selectors. Code for it below.

It hides [badges on the notification icon](https://twitter.com/nikitavoloboev/status/1629050086700339200), [view counts on tweets](https://twitter.com/nikitavoloboev/status/1629080048568344576), ['new tweets' banner](https://twitter.com/nikitavoloboev/status/1629077068045508608) and few more things.

Final result should look [like this](https://twitter.com/nikitavoloboev/status/1629105196356739073).

```css
[aria-label="1 unread items"] {
  display: none;
}
[aria-label="2 unread items"] {
  display: none;
}
[aria-label="3 unread items"] {
  display: none;
}
[aria-label="4 unread items"] {
  display: none;
}
[aria-label="5 unread items"] {
  display: none;
}
[aria-label="6 unread items"] {
  display: none;
}
[aria-label="7 unread items"] {
  display: none;
}
[aria-label="undefined unread items"] {
  display: none;
}
[aria-label="Home"] {
  display: none;
}
[aria-label="New Tweets are available. Push the period key to go to the them."] {
  display: none;
}
article[data-testid="tweet"] div:has(> a[aria-label*="View Tweet analytics"]) {
  display: none;
}
```

## Interesting hash tags

- [#WomensArt](https://twitter.com/hashtag/WomensArt)
- [#VisibleWomen](https://twitter.com/hashtag/VisibleWomen)
- [#коточетверг](https://twitter.com/hashtag/%D0%BA%D0%BE%D1%82%D0%BE%D1%87%D0%B5%D1%82%D0%B2%D0%B5%D1%80%D0%B3)

## Notes

- Can search over my own Twitter profile with Nitter [here](https://nitter.net/nikitavoloboev/search?f=tweets).
- [To prevent usernames and URLs from being converted to links, interrupt them with a zero-width space.](https://twitter.com/markdalgleish/status/1340897088486273034)
- [Normalize reading the whole tweet/tweet thread before replying](https://twitter.com/sarah_edo/status/1370806530497376257)
- [Twitter stops being fun after a certain number of followers because you can’t tweet anything without a barrage of opinions/projections.](https://twitter.com/notdanilu/status/1378770853169627137)
- [Twitter is a lot like prison. The best move on your first day is to pick a fight with the biggest account on here.](https://twitter.com/nikitabier/status/1391049141124427782)
- [Each day on twitter there is one main character. The goal is to never be it.](https://twitter.com/maplecocaine/status/1080665226410889217)
- [How to become a tech twitter influencer in five easy steps: 1. be relatable 2. keep a regular posting schedule 3. engage with your followers 4. go on podcasts!](https://twitter.com/ctrlshifti/status/1476086483245101056)
- [Twitter has useful "from people you follow" search filter](https://twitter.com/Gavmn/status/1507371911310807043)
- [New Twitter CEO Checklist: 1. Open-source the algorithm. 2. Eliminate all bots. 3. Restore free speech.](https://twitter.com/DavidSacks/status/1515363131249213444)
- [The essential truth of every social network is that the product is content moderation, and everyone hates the people who decide how content moderation works. Content moderation is what Twitter makes — it is the thing that defines the user experience.](https://twitter.com/simonw/status/1586017662726856704)

## Links

- [How I use Twitter (2018)](https://krausefx.com/blog/how-i-use-twitter)
- [Twitter Lists](https://github.com/AndySparks/captains-log/blob/master/resources/reading-lists/twitter-lists.md)
- [Instant Twitter Search](https://twitter-search.vercel.app/) - Instantly search across your entire Twitter history with a beautiful UI powered by Algolia. ([Code](https://github.com/transitive-bullshit/twitter-search))
- [TWINT](https://github.com/twintproject/twint) - Twitter Intelligence Tool.
- [Everyone can build a Twitter audience (2020)](https://gumroad.com/l/twitter-audience/launch)
- [Nitter](https://nitter.net/) - Free and open source alternative Twitter front-end focused on privacy. ([Code](https://github.com/zedeus/nitter)) ([HN](https://news.ycombinator.com/item?id=28294099))
- [Treeverse](https://github.com/paulgb/Treeverse) - Extension for navigating burgeoning Twitter conversations.
- [Twitter as a tool to learn (2020)](https://twitter.com/eriktorenberg/status/1260352115102248961)
- [Twitter Followers](https://github.com/ConradIrwin/twitter-followers) - Tool to download all a user's twitter followers.
- [Build an open source tool to export your Twitter followers](https://github.com/balajis/twitter-export)
- [Get twitter avatar for a username](https://github.com/siddharthkp/twitter-avatar)
- [Thread Reader](https://threadreaderapp.com/) - Transform Twitter Threads into a Readable Page.
- [Debubble](https://debubble.me/) - Publishing tool that facilitates a debate between two Twitter users. ([HN](https://news.ycombinator.com/item?id=23728499))
- [Cheap Twitter Bots, Done Quick](https://cheapbotsdonequick.com/) - Bots are written in Tracery, a tool for writing generative grammars.
- [How to get your @brand on Twitter if taken (2020)](https://www.indiehackers.com/post/how-to-get-your-brand-on-twitter-if-taken-3e9c449974)
- [Tweet Seeker](https://2017.lionheartsw.com/software/tweet-seeker/) - Best way to find your tweets on Twitter.
- [Vicariously](https://vicariously.io/) - Easiest way to create Twitter lists based on the follows of other users.
- [Twitter API Tutorial (2018)](http://socialmedia-class.org/twittertutorial.html)
- [Twuota](https://twuota.poxmedia.net/) - Twitter timeline statistics.
- [On Using Twitter (2020)](https://medium.com/@emilymenonbender/on-using-twitter-84fbd80c8919) ([HN](https://news.ycombinator.com/item?id=24076739))
- [Twitter Crush](https://woaini.page/) - Find twitter user you admire the most.
- [Top Tweets Bot](https://toptweetsbot.appspot.com/) - Creates and maintains a collection of your top tweets.
- [The.Rip](https://the.rip/) - Unroll Twitter threads into markdown. ([Tweet](https://twitter.com/benzguo/status/1294832715049517058))
- [Advanced Search on Twitter](https://github.com/igorbrigadir/twitter-advanced-search)
- [Tracking significant changes to the Twitter API or platform as a whole](https://github.com/igorbrigadir/twitter-history)
- [Twitter Developers Forum](https://twittercommunity.com/latest)
- [First hands on the new Twitter API (2020)](https://dev.to/bearer/first-hands-on-the-new-twitter-api-44e9)
- [HuggingTweets](https://github.com/borisdayma/huggingtweets) - Train a model to generate tweets. .
- [ilo](https://ilo.so/) - Better Twitter analytics.
- [toptweets.by](https://toptweets.by/) - Use toptweets.by / <user\> to see the top tweets by that user. ([Code](https://github.com/michaelrbock/besttweets)) ([Tweet](https://twitter.com/michaelrbock/status/1316952215354048513))
- [Awesome Twitter Data](https://github.com/shaypal5/awesome-twitter-data) - List of Twitter datasets and related resources.
- [Gardening Your Twitter: Growing Your Followers (2020)](https://steipete.com/posts/growing-your-twitter-followers/)
- [twitter-text](https://github.com/twitter/twitter-text) - Collection of libraries and conformance tests to standardize parsing of Tweet text. ([In Rust](https://github.com/egg-mode-rs/egg-mode-text))
- [The Art of Twitter: A Guide To Building Your Twitter Account](https://gumroad.com/l/TwitterGuide)
- [Bookmarklet that redirects you to original tweet of video that someone else used in their tweet](https://twitter.com/wongmjane/status/1202293089395568640)
- [Gramtion](https://github.com/lRomul/gramtion) - Twitter bot for generating photo descriptions (alt text).
- [All My Tweets](https://www.allmytweets.net) - View all your tweets on one page.
- [twitterscraper](https://github.com/taspinar/twitterscraper) - Scrape Twitter for Tweets.
- [Twit Thread](https://github.com/adblanc/twit-thread) - Utility module for twitter bots based on Twit, Twitter API client for node.
- [Twistory](https://twistory.ml/) - Shows you "On This Day" Tweets. ([Article](https://shkspr.mobi/blog/2020/11/introducing-on-this-day-in-twistory/))
- [Twitter Inspiration Handbook](https://marketingexamples.com/handbook/twitter-inspiration)
- [Use Netlify Functions and the Twitter API v2 as a CMS for Your Gatsby Blog (2020)](https://paulie.dev/posts/2020/11/gatsby-netlify-twitter/) ([HN](https://news.ycombinator.com/item?id=25186006))
- [Twurl](https://github.com/twitter/twurl) - OAuth-enabled curl for the Twitter API.
- [I Made a Self-Quoting Tweet (2020)](https://oisinmoran.com/quinetweet) ([HN](https://news.ycombinator.com/item?id=25244872))
- [Tools for fighting abuse on Twitter](https://github.com/travisbrown/cancel-culture) - Contains some low-tech tools designed to help you make Twitter a nicer place for yourself.
- [DownloadThisVideo](https://thisvid.space/) - Download videos and GIFs off Twitter. ([Code](https://github.com/shalvah/DownloadThisVideo))
- [Typefully](https://typefully.app/) - Write & publish great tweets, without distractions. ([HN](https://news.ycombinator.com/item?id=25358108))
- [Twitter Monitor](https://github.com/ezolla/Twitter-Monitor) - Monitor Twitter Accounts w/ Discord Webhooks.
- [Twemex](https://twemex.app/) - Sidebar for Twitter. ([Beta](https://www.notion.so/Twemex-Beta-8a33819331cc41c4970bc73ea242f0d2))
- [Valuables](https://v.cent.co/) - Buy tweets. Autographed by your favorite creators.
- [GraphQL Tweetletter](https://tweets.dgraph.io/) - Collection of the most interesting and popular tweets, delivered to your inbox every week.
- [Building Lucerne, a Twitter experience tailored to me (2020)](https://thesephist.com/posts/lucerne/)
- [twitter-to-sqlite](https://github.com/dogsheep/twitter-to-sqlite) - Save data from Twitter to a SQLite database.
- [How to Gradually Exit Twitter (2020)](https://balajis.com/how-to-gradually-exit-twitter/)
- [Twitter CLI](https://github.com/sferik/t) - Command-line power tool for Twitter.
- [Simple and unlimited twitter scraper with python and without authentication](https://github.com/Altimis/Scweet)
- [fleets](https://github.com/karan/fleets) - Automatically delete tweets, retweets, and favorites. ([Another version](https://github.com/caarlos0/twitter-cleaner))
- [Block Party](https://www.blockpartyapp.com/) - Filter out unwanted @mentions from Twitter, and continue to use Twitter as normal.
- [Trendsmap](https://www.trendsmap.com/) - Twitter Trending Hashtags and Topics.
- [Tweetpik](https://tweetpik.vercel.app/) - Convert Tweets to Images.
- [TwitRSS.me](https://github.com/ciderpunx/twitrssme) - Tool to make Twitter timelines and searches into RSS feeds.
- [Twitter Account Data Analysis](https://github.com/engali94/Twitter-Account-Analyzer) - Using various Python libraries such as Pandas, tweetPy, JSON ans matplotLib to take a sneak peek on your Twitter account using Google Colab.
- [Twitter Birdwatch Guide](https://twitter.github.io/birdwatch/) ([Code](https://github.com/twitter/birdwatch))
- [Twitter improves API usage for researchers (2021)](https://blog.twitter.com/developer/en_us/topics/tips/2021/enabling-the-future-of-academic-research-with-the-twitter-api.html) ([HN](https://news.ycombinator.com/item?id=25927803))
- [Twitter’s Birdwatch is Fundamentally Flawed (2021)](https://soatok.blog/2021/01/27/twitters-birdwatch-is-fundamentally-flawed/)
- [Spoonbill](http://spoonbill.io/) - See profile changes from the people you follow on Twitter or other social networks.
- [fllwrs](http://fllwrs.com/) - Keep track of who follows and unfollows you on twitter.
- [React Static Tweets](https://github.com/transitive-bullshit/react-static-tweets) - Extremely fast static renderer for tweets.
- [Quitting Twitter (2021)](https://blog.nindalf.com/posts/quitting-twitter/) ([HN](https://news.ycombinator.com/item?id=26267529))
- [Summary of Daniel Vassallo's Twitter Course (2021)](https://coursemaker.org/blog/summary-daniel-vassallo-twitter-course/)
- [Great tweets saved by users](<https://twitter.com/search?q=%22save%20thread%22%20(%40readwiseio)%20min_faves%3A3&src=typed_query>)
- [Twitter is taking on Clubhouse, Substack and Patreon with new products (2021)](https://www.theverge.com/22319527/twitter-kayvon-beykpour-interview-consumer-product-decoder)
- [CrowdFox](https://www.crowdfox.io/) - Turn Twitter followers into paying customers.
- [Exploring Tweets with SQLite and WASM](https://divu.in/experiments/wasm/twitter-sqlite) ([HN](https://news.ycombinator.com/item?id=26675984))
- [Growth Tweets](https://www.growthtweets.link/) - Weekly-updated repository of tweets that matter.
- [Best of Twitter - Alexey Guzey](https://guzey.com/best-of-twitter/) - Send 10-20 of the most interesting tweets I see every week.
- [The power of Twitter (2021)](https://twitter.com/gregisenberg/status/1378395998859493378)
- [Archive Team: A Smattering of Tweets](https://archive.org/details/archiveteam_twitter) ([HN](https://news.ycombinator.com/item?id=26703203))
- [Telegram bot that forwards Tweets](https://github.com/franciscod/telegram-twitter-forwarder-bot)
- [Goodbye](https://github.com/ahmetb/goodbye) - Notify yourself when someone unfollows you on Twitter.
- [Twitlistr](https://www.twitlistr.com/) - Manage Twitter Lists With Ease.
- [Atsh.io](https://www.atsh.io/) - GUI for Advanced Twitter Search Queries.
- [10 Advanced Twitter Features](https://twitter.com/dickiebush/status/1376914220411985926)
- [Twitter Spaces, a few weeks in (2021)](http://dtrace.org/blogs/bmc/2021/05/02/twitter-spaces-a-few-weeks-in/)
- [BrandBird Studio](https://www.brandbird.app/) - Smart image editor for SaaS founders who want beautiful graphics without effort. ([Twitter](https://twitter.com/BrandBirdApp))
- [Twitter Clone](https://github.com/HotPotatoC/twitter-clone) - Twitter Clone developed using Go + Vue 3 + Vite + TailwindCSS + PostgreSQL + Redis.
- [TwitterStream](https://github.com/Fallenstedt/twitter-stream) - Stream Tweets With Twitter's v2 Streaming API using Go.
- [snap-tweet](https://github.com/privatenumber/snap-tweet) - Command-line tool to capture clean and simple tweet snapshots.
- [Twitter Developer Platform](https://developer.twitter.com/en) - Use Cases, Tutorials, & Documentation.
- [Static Tweets with MDX and Next.js (2021)](https://blog.maximeheckel.com/posts/static-tweets-with-mdx-nextjs/)
- [Broadcast](https://github.com/daneden/Broadcast) - Write-only Twitter client. It lets you post tweets and reply to your last tweet.
- [Rebuilding Twitter’s Public API (2021)](https://www.youtube.com/watch?v=axFSaH_-IMw)
- [Twitter API v2 sample code](https://github.com/twitterdev/Twitter-API-v2-sample-code) - Sample code for the Twitter API early access endpoints (Python, Java, Ruby, and Node.js).
- [BlackMagic.so](https://blackmagic.so/) - Twitter Growth Tool. ([Sidebar](https://blackmagic.so/sidebar/)) ([Tweet](https://twitter.com/tdinh_me/status/1469555195596054529))
- [#buildinpublic](https://buildinpublic.com/) - Discover what people are building in public.
- [Twitter's new visual design system](https://twitter.com/ashlie/status/1425506038506147840) ([Tweet](https://twitter.com/TwitterDesign/status/1425505308563099650))
- [TidyTweets](https://tidytweets.org/) - Tidy up your Following list on Twitter. ([Code](https://github.com/rlueder/tidytweets))
- [twilens](https://github.com/ciffelia/twilens) - Full-text search for your tweets.
- [RSS 2 Twitter](https://github.com/umputun/rss2twitter) - Publish RSS updates to Twitter.
- [Kizie](https://kizie.co/) - New way to interact with Twitter. ([HN](https://news.ycombinator.com/item?id=28549960))
- [Feeds Mage](https://www.feedsmage.com/) - Scans your Twitter follows to discover blogs and newsletters.
- [Minimal Twitter](https://github.com/ThomasWang/minimal-twitter) - Minimal Theme for the new Twitter UI.
- [How many accounts block you on Twitter](https://blolook.osa-p.net/index.html?lang=en)
- [Thread Hunt](https://threadhunt.xyz/) - Quality Twitter threads from undiscovered creators.
- [Static Tweet Next.js](https://github.com/lfades/static-tweet) - Completely customizable static tweet for Next.js applications. ([Demo](https://static-tweet.vercel.app/))
- [Twitter audit trail backup](https://github.com/ahmetb/twitter-audit-log) - Backs up my follower list, following list, blocked accounts list and muted accounts list periodically using GitHub Actions.
- [retweeters](https://github.com/craftweg/retweeters) - Micro-CLI powered by NodeJS for fetching the retweeters of a tweet.
- [Brain Marks](https://github.com/mikaelacaron/brain-marks) - Categorize your Twitter bookmarks, being created during the Big Brain Hacks Hackathon.
- [A small DOCUMERICA Twitter bot (2021)](https://blog.yossarian.net/2021/10/25/A-small-documerica-twitter-bot)
- [TweetShelf](https://www.tweetshelf.com/) - Recommendations from your Twitter friend.
- [Good Twitter follows (2021)](https://twitter.com/aaditsh/status/1453745795862257672)
- [Tweet Hunter](https://tweethunter.io/) - Grow Your Twitter Audience.
- [Awesome Twitter Communities](https://github.com/caarlos0/awesome-twitter-communities)
- [TwitterFOMO.dev](https://twitterfomo.dev/) - Best WebDev Tweets. ([Code](https://github.com/tomdohnal/twitter-fomo))
- [How Twitter got research right (2021)](https://www.theverge.com/2021/11/19/22790174/twitter-research-transparency-published-findings)
- [Search your all tweets](https://github.com/azu/mytweets)
- [Twitter App Ideas](https://app.simplenote.com/publish/jlwK77) ([Tweet](https://twitter.com/swyx/status/1462534363535626244))
- [Transform your twitter following into list](https://github.com/guillaumeLamanda/twitter-follow2list)
- [Twitter Has a New CEO; What About a New Business Model? (2021)](https://stratechery.com/2021/twitter-has-a-new-ceo-what-about-a-new-business-model/) ([HN](https://news.ycombinator.com/item?id=29391876))
- [Twitter Is About to Get Way Worse (2021)](https://bariweiss.substack.com/p/twitter-is-about-to-get-way-worse) ([Tweet](https://twitter.com/jordanbpeterson/status/1466132542571134982))
- [Twitter Clone in Rust](https://github.com/evoxmusic/twitter-clone-rust)
- [Twitter Query Language](https://gist.github.com/andyclarkemedia/3b4e062a45323138bd28ec52d80eb7b1)
- [Twitter, the Intimacy Machine (2021)](https://ravenmagazine.org/magazine/twitter-the-intimacy-machine/)
- [Twttr](https://github.com/sreechar/twttr) - Streamlines your twitter surfing process by prioritizing user-friendly features. ([Web](https://twttr.sreecha.me/))
- [A Guide to Twitter](https://tasshin.com/blog/a-guide-to-twitter/) ([HN](https://news.ycombinator.com/item?id=29752379))
- [Nazuna](https://github.com/Sazzo/nazuna) - Download Twitter videos using your terminal.
- [Twitter Stats Collection and Analysis](https://github.com/jschuur/twitter-stats) - Self-hosted performance tracking for your Twitter followers and tweets.
- [Pikaso](https://pikaso.me/) - Get a clean and clutter-free screenshot of any tweet.
- [Twift](https://github.com/daneden/Twift) - Asynchronous Twitter Swift library based on the Twitter v2 API.
- [Simplified Twitter](https://github.com/brunolemos/simplified-twitter) - Remove distractions from the new Twitter layout. Browser extension.
- [Tracking the far right on Twitter](https://github.com/travisbrown/twitter-watch)
- [Annotating the Tweebank Corpus on Named Entity Recognition and Building NLP Models for Social Media Analysis (2022)](https://arxiv.org/abs/2201.07281) ([Code](https://github.com/social-machines/TweebankNLP))
- [Turn your tweets/threads into a blog and RSS feed](https://typefully.com/profile) ([HN](https://news.ycombinator.com/item?id=30440835))
- [Rust Retweet Bot](https://github.com/multimeric/RustLangRetweet) - Rust bot that runs periodically on AWS Lambda and retweets any Tweets matching a query.
- [Fritter](https://fritter.cc/) - Privacy-friendly Twitter frontend for mobile devices. ([Code](https://github.com/jonjomckay/fritter))
- [Refined Twitter Lite](https://github.com/giuseppeg/refined-twitter-lite) - Adds UI improvements and new useful features to Twitter Lite.
- [twprs](https://github.com/travisbrown/twprs) - Twitter profile tools for Rust.
- [How to Download Twitter Spaces That Aren't Yours](https://www.swyx.io/download-twitter-spaces/)
- [Bookmarks lookup and Notion](https://github.com/twitterdev/Bookmarks-Notion-Notebook) ([Tweet](https://twitter.com/suhemparack/status/1507100523534565376))
- [Jack Dorsey has regrets about building Twitter (2022)](https://twitter.com/jack/status/1510314535671922689) ([HN](https://news.ycombinator.com/item?id=30899950))
- [Thread Helper](https://threadhelper.com/) - Serendipity engine on the Twitter sidebar. ([Code](https://github.com/threadhelper/threadhelper))
- [Tweetscape](https://github.com/rooteco/tweetscape) - Supercharged Twitter Feed.
- [Twitter API v2 code](https://github.com/sofiapinto/Twitter-API-v2-code) - Scripts that can be used to retrieve data from Twitter API v2.
- [Stabletweet](https://github.com/markjaquith/stabletweet) - Cloudflare Worker proxy for checking whether a tweet exists.
- [Advanced Twitter Search Tips](https://twitter.com/TessaRDavis/status/1512402324102291467)
- [Better TweetDeck](https://better.tw/) - Take TweetDeck to the next level. ([Code](https://github.com/eramdam/BetterTweetDeck))
- [Twitter Bookmarks Search WebExtension](https://github.com/flybayer/twitter-bookmarks-search)
- [Twitter Card Image Generator](https://github.com/Ladicle/tcardgen) - Generate Twitter card (OGP) images for your blog posts.
- [twintent](https://github.com/sheepla/twintent) - Small instant tweeting command.
- [Twitter API open evolution](https://github.com/twitterdev/open-evolution) - Open evolution proposals for the Twitter API.
- [Twitter Scraper](https://github.com/bisguzar/twitter-scraper) - Scrape the Twitter Frontend API without authentication.
- [Twitter Graph](https://github.com/eleurent/twitter-graph) - Fetch and visualize the graph of your Twitter friends and followers.
- [Titter](https://github.com/dabit3/titter) - Decentralized Twitter prototype built with GraphQL, Next.js, Ceramic, Arweave, and Bundlr.
- [Back to the Future of Twitter (2022)](https://stratechery.com/2022/back-to-the-future-of-twitter/) ([Reddit](https://www.reddit.com/r/slatestarcodex/comments/u6rfsi/back_to_the_future_of_twitter/))
- [Twitter API v2 Plugin Token Refresher](https://github.com/alkihis/twitter-api-v2-plugin-token-refresher)
- [Real-time twitter graph of your friends](https://github.com/Nican/Furland)
- [How Twitter's Algorithmic Feed Works (2022)](https://transitivebullsh.it/oss-twitter-algorithm-part-1) ([HN](https://news.ycombinator.com/item?id=31115755))
- [HN: Twitter accepts Musk's $43B offer (2022)](https://news.ycombinator.com/item?id=31153277)
- [HN: Twitter The Algorithm](https://news.ycombinator.com/item?id=31160546)
- [Parse Tweet Archive Python Script](https://github.com/mshea/Parse-Twitter-Archive)
- [Elon Conquers The Twitterverse (2022)](https://bariweiss.substack.com/p/elon-conquers-the-twitterverse?s=w) ([Tweet](https://twitter.com/micsolana/status/1518980601197608960))
- [Search your Twitter Bookmarks](https://github.com/twitterdev/bookmarks-search)
- [Why Twitter is a Musk-Have (2022)](https://blog.lopp.net/why-twitter-musk-have/)
- [Twitter Financial information](https://investor.twitterinc.com/financial-information/Additional-Information-and-Where-To-Find-It/default.aspx)
- [Twitter Helper](https://github.com/ErikBjare/twitter-helper) - Figure out who you're following but not interacting with (no likes, replies, retweets).
- [“Free Speech” Ought to Mean More than Mocking Trans People (2022)](https://www.thebulwark.com/free-speech-ought-to-mean-more-than-mocking-trans-people/) ([Tweet](https://twitter.com/NGrossman81/status/1526524194858704897))
- [Oxide and Friends Twitter Spaces Notes](https://github.com/oxidecomputer/twitter-spaces)
- [How you can reach 0 to 1000 followers on Twitter quickly](https://devvnomad.notion.site/How-you-can-reach-0-to-1000-followers-on-Twitter-quickly-928cef4051e2454290549f53303e791a)
- [Awesome Modern Twitter API](https://github.com/andypiper/awesome-modern-twitter-api)
- [Twitter web has nice shortcuts](https://twitter.com/ossia/status/1533448928242417666)
- [Tweak New Twitter](https://github.com/insin/tweak-new-twitter/) - Browser extension which removes algorithmic content from Twitter, hides news & trends, lets you control which shared tweets appear on your timeline. ([HN](https://news.ycombinator.com/item?id=34499615))
- [Static twitter embed](https://ianmuchina.com/blog/12-tweet-embed/) ([Lobsters](https://lobste.rs/s/owbvhi/static_twitter_embed))
- [Twitter Video Downloader](https://twitter-video-download.com/)
- [memory.lol](https://github.com/travisbrown/memory.lol) - Tiny web service that provides historical information about social media accounts. ([Private](https://github.com/travisbrown/memory.lol-private))
- [ornithology](https://github.com/jonhoo/ornithology) - Tool that parses your Twitter archive and highlights interesting data from it.
- [Who Unfollowed Me on Twitter App](https://t.co/ntMla1yp4e)
- [Twitter BotStarter template](https://github.com/luvi/botstarter)
- [Tweetic](https://github.com/zernonia/tweetic) - Convert Tweets to Static HTML. ([Web](https://www.tweetic.io/))
- [ListFollowers.com](https://listfollowers.com/) - Fetch a list of someone's follows, followers, or mutuals, or combine two such lists from two users, and export the results to CSV or JSON. ([Tweet](https://twitter.com/mechanical_monk/status/1550055763317506050))
- [Quick Tweet Search Extension](https://github.com/taro-28/quick-tweet-search) - Chrome Extension to quickly search Tweet by Twitter user.
- [Doing give-aways on Twitter with GitHub actions](https://github.com/alexeygrigorev/twitter-raffles)
- [Loading Twitter Birdwatch into SQLite for analysis with Datasette](https://til.simonwillison.net/twitter/birdwatch-sqlite)
- [Take back your Twitter feed](https://mute.guru/) - Take back your twitter feed, by instantly muting all engagement farmers. ([Code](https://github.com/m1guelpf/mute.guru))
- [sauron](https://github.com/chshersh/sauron) - CLI tool that fetches info from Twitter and analyses it.
- [Tweeety](https://github.com/Jamelle-Boose/tweeety) - App to stream Tweets in real-time using Twitter API v2.
- [TwitterToNitter](https://github.com/no-gravity/TwitterToNitter) - Bookmarklet that shows the current Twitter page on Nitter. On every click it choses a random Nitter instance.
- [Twitter branching threads](https://twitter.com/jordanmoore/status/1574827170492551207)
- [FixTweet](https://github.com/dangeredwolf/FixTweet) - Embed Twitter videos, polls, translations, & more on Discord and Telegram.
- [ModernDeck](https://moderndeck.app/) - Beautiful, powerful Twitter client for desktop. ([Code](https://github.com/dangeredwolf/ModernDeck))
- [How to analyze edited Tweets with the Twitter API v2 using Python (2022)](https://dev.to/suhemparack/how-to-analyze-edited-tweets-with-the-twitter-api-v2-using-python-odj)
- [Nitter Instances Uptime](https://xnaas.github.io/nitter-instances/) - Automated uptime monitoring of Nitter instances. ([Code](https://github.com/xnaas/nitter-instances))
- [Great devs on Twitter to follow](https://twitter.com/swyx/status/1580660708550680576)
- [Twitter Scraper](https://github.com/mahrtayyab/tweety)
- [Twitter Stream](https://github.com/tesaguri/twitter-stream-rs) - Rust library for listening on Twitter Streaming API.
- [Pipitor](https://github.com/tesaguri/pipitor) - Twitter bot that listens on WebSub/Twitter streams and (re)Tweet the updates in real time.
- [Elon Musk Has Taken Twitter: Day Zero (2022)](https://www.piratewires.com/p/elon-musk-has-taken-twitter-day-zero)
- [Twitter is moving its tech stack to Bluesky to run as a protocol not a company](https://twitter.com/davetroy/status/1586166535592509440)
- [Ask HN: What social media site could replace Twitter? (2022)](https://news.ycombinator.com/item?id=33416498)
- [Ask HN: How would you design an alternative Twitter (2022)](https://news.ycombinator.com/item?id=33419574)
- [Tell HN: The issues of Twitter are not a technical problem to solve (2022)](https://news.ycombinator.com/item?id=33421489)
- [From Google to Twitter (2022)](https://ma.nu/blog/from-google-to-twitter) ([HN](https://news.ycombinator.com/item?id=33473249))
- [Twittoons](https://twittoons.com/) - Cartoons about life at Twitter.
- [Publish Twitter Archive Site](https://github.com/siakaramalegos/pesos-tweets-11ty)
- [Git as a janky Twitter replacement](https://github.com/diracdeltas/tweets) ([Tweet](https://twitter.com/bcrypt/status/1588416861552582657)) ([HN](https://news.ycombinator.com/item?id=33488996))
- [Ask HN: Do you believe there's really an alternative to Twitter? (2022)](https://news.ycombinator.com/item?id=33490217)
- [TwitterHD](https://github.com/DavidBuchanan314/TwitterHD) - Userscript that forces twitter to always load images and videos in full resolution.
- [Veryfied](https://github.com/evilsocket/veryfied) - Browser extension for Twitter that will mark pre-Musk era verified accounts.
- [Twitter's Birdwatch archived docs](https://github.com/twitter/new-repo)
- [Twitter Open Source](https://opensource.twitter.dev/) - Identifying projects we've released, organizations we support, and the work we do to support open source. ([Code](https://github.com/twitter/opensource-website))
- [TwitterOSS Metrics](https://github.com/twitter/metrics) - Generates periodic reports based on the health of Twitter Open Source projects.
- [All verified users on Twitter as of Nov 1st, 2022](https://github.com/IntiDC/twitter-verified)
- [Elon asking for what Twitter users find annoying on Twitter now (2022)](https://twitter.com/elonmusk/status/1590383937284870145)
- [Python code to parse a Twitter archive and output in various ways](https://github.com/timhutton/twitter-archive-parser)
- [Eight Dollars](https://github.com/wseagar/eight-dollars) - Browser extension that shows twitter blue vs real verified users.
- [twdl](https://github.com/dogancelik/twdl) - Tool for downloading media of individual tweets.
- [Tracking the far right on Twitter](https://github.com/travisbrown/twitter-watch)
- [How to Search Your Entire Twitter Archive With One Line of jq (2022)](https://www.ianwootten.co.uk/2022/11/14/how-to-search-your-entire-twitter-archive-with-one-line-of-jq/)
- [Why Dave LaMacchia stayed at Twitter for 9+ years (2022)](https://twitter.com/_dml/status/1592619959947243520)
- [Twitter's amazing engineering work over years (2022)](https://twitter.com/danluu/status/1592774269733601281)
- [George Hotz | Programming | so how does twitter work? | API | GraphQL | requests | Backend | Scala (2022)](https://www.youtube.com/watch?v=z6xslDMimME) ([Tweet](https://twitter.com/realGeorgeHotz/status/1593109753579786240))
- [The Infrastructure Behind Twitter: Scale (2017)](https://blog.twitter.com/engineering/en_us/topics/infrastructure/2017/the-infrastructure-behind-twitter-scale)
- [Twitter Favorites(Likes) Archive](https://github.com/15cm/twitter-favorites-archive) - Series of scripts to archive metadata and medias of your Twitter Favorites(Likes).
- [stweet](https://github.com/markowanga/stweet) - Advanced python library to scrap Twitter (tweets, users) from unofficial API.
- [Backup Twitter data now](https://www.reddit.com/r/DataHoarder/comments/yy7tig/backup_twitter_now_multiple_critical_infra_teams/)
- [Converting Your Twitter Archive to Markdown (2022)](https://matthiasott.com/notes/converting-your-twitter-archive-to-markdown)
- [How Twitter might play out, the ramifications, and what Elon’s calculus might be](https://twitter.com/LBacaj/status/1593705261939802112)
- [unflwrs](https://github.com/Syfaro/unflwrs) - Tool to track Twitter followers.
- [Twitter Architecture 2022 vs 2012](https://twitter.com/alexxubyte/status/1594008281340530688)
- [Things Twitter should improve](https://twitter.com/growing_daniel/status/1594365574850654210)
- [tweetback Twitter Archive](https://github.com/tweetback/tweetback) - Take ownership of your Twitter data.
- [Why Twitter Didn’t Go Down: From a Real Twitter SRE (2022)](https://matthewtejo.substack.com/p/why-twitter-didnt-go-down-from-a) ([HN](https://news.ycombinator.com/item?id=33701371)) ([Lobsters](https://lobste.rs/s/4mupjz/why_twitter_didn_t_go_down_from_real)) ([Reddit](https://www.reddit.com/r/programming/comments/z24gya/why_twitter_didnt_go_down_from_a_real_twitter_sre/))
- [Twitter Clone in Vue](https://github.com/madebyfabian/twitter-clone)
- [Save Your Threads](https://github.com/harvard-lil/archive.social) - High-fidelity capture of Twitter threads as sealed PDFs.
- [Improving Twitter Search ideas](https://twitter.com/realGeorgeHotz/status/1594790788168294404)
- [Tool to backup your twitter likes](https://yodapunk.gumroad.com/l/twitter-backup) ([Tweet](https://twitter.com/YodaPunk/status/1595418496762105856))
- [dewey](https://getdewey.co/) - Organize & Share Twitter Bookmarks.
- [Tweet Admin](https://tweetadmin.com/) - Do more with Twitter.
- [Twitter, when the wall came down (2022)](http://dtrace.org/blogs/bmc/2022/11/05/twitter-when-the-wall-came-down/)
- [TwitterSearchTokenTest](https://github.com/T-Ev/TwitterSearchTokenTest)
- [Stork](https://github.com/robertoszek/pleroma-bot) - Mirror your favorite Twitter accounts in the Fediverse, so you can follow their updates from the comfort of your favorite instance.
- [Trendagram](https://github.com/advaithhl/Trendagram) - Telegram bot for scheduled Twitter trends.
- [George Hotz | Exploring Twitter Open Source Code | Scala (2022)](https://www.youtube.com/watch?v=nvtoOxNfDQo)
- [TwitVault](https://github.com/terhechte/twitvault) - Easily Archive and Search Your Twitter Data with our Syncable Desktop App. ([Web](https://terhechte.github.io/twitvault/))
- [Download Twitter Video](https://github.com/egoist/download-twitter-video) - Easiest way to download any Twitter video.
- [Handle.horse](https://handle.horse/) - Get notified when your favorite Twitter @handle becomes available.
- [Twitter Media](https://media.okikio.dev/) - Enter a Tweet URL to download the video/image in it. ([Code](https://github.com/okikio/twitter-media))
- [Render tweet into beautiful markdown](https://github.com/silentroach/tweet.md)
- [The Fifth Estate (2022)](https://www.piratewires.com/p/the-fifth-estate)
- [The Twitter Files, Part Six (2022)](https://twitter.com/mtaibbi/status/1603857534737072128) ([HN](https://news.ycombinator.com/item?id=34020654))
- [Add empty space to links banned by Twitter to get around the ban](https://twitter.com/molly0xFFF/status/1603837343093428237)
- [Twitter's anti-Mastodon filter evasion](https://infosec.exchange/@postmodern/109523637731779949) ([HN](https://news.ycombinator.com/item?id=34036265))
- [tweetback/canonical](https://github.com/tweetback/tweetback-canonical) - Package to resolve twitter URLs to new canonically hosted twitter backups.
- [Bird SQL](https://www.perplexity.ai/sql) - Twitter search interface translates natural language into SQL. ([Tweet](https://twitter.com/perplexity_ai/status/1603441221753372673))
- [Make your own simple, public, searchable Twitter archive](https://tinysubversions.com/twitter-archive/make-your-own/) ([Code](https://github.com/dariusk/twitter-archiver)) ([Tweet](https://twitter.com/tinysubversions/status/1604552603630649344)) ([HN](https://news.ycombinator.com/item?id=34056378))
- [Minimal Theme for Twitter](https://typefully.com/minimal-twitter)
- [Nostr Protocol](https://github.com/nostr-protocol/nostr) - Truly censorship-resistant alternative to Twitter that has a chance of working.
- [Faster Twitter Embeds (2022)](https://blog.kevmo314.com/faster-twitter.html) ([Or you can use Nitter embed](https://nitter.fly.dev/mhevery/status/1606438382561026049/embed))
- [Twitter Graph](https://twit.deta.dev/) - Like Google Trends but for Twitter.
- [Embed Tweets on your site with minimal JS, thanks to Qwik](https://github.com/mhevery/qwik-twitter)
- [Python Twitter Examples](https://github.com/ideoforms/python-twitter-examples) - Examples of using Python for Twitter social data mining, using the python-twitter-tools framework.
- [Convert Twitter threads from archive to Markdown files](https://github.com/duncanmalashock/tweetparse)
- [Production Twitter on One Machine: 100Gbps NICs and NVMe are fast (2023)](https://thume.ca/2023/01/02/one-machine-twitter/) ([HN](https://news.ycombinator.com/item?id=34291191))
- [Twitter performance prototype](https://github.com/trishume/twitterperf) - Rust prototype of handling the full production load of Twitter's core timeline collation on a single core by only doing the very basics in-memory.
- [Hive.one](https://hive.one/) - Find reputable Twitter accounts.
- [Caching at Twitter with Yao Yue (2023)](https://softwareengineeringdaily.com/2023/01/12/caching-at-twitter-with-yao-yue/)
- [Twitter Blue data](https://github.com/travisbrown/blue)
- [Twitter Bio Generator](https://www.twitterbio.com/) ([Code](https://github.com/Nutlope/twitterbio))
- [BirdsiteLIVE](https://github.com/NicolasConstant/BirdsiteLive) - ActivityPub bridge from Twitter.
- [Twitter Video Tools](https://github.com/code-yeongyu/TwitterVideoTools) - Multi-processing supported twitter video downloader.
- [Intercepting t.co links using DNS rewrites (2023)](https://djharper.dev/post/2023/01/29/intercepting-t.co-links-using-dns-rewrites/) ([HN](https://news.ycombinator.com/item?id=34571448)) ([Lobsters](https://lobste.rs/s/zvpabd/intercepting_t_co_links_using_dns))
- [Bird.makeup](https://bird.makeup/) - Twitter to ActivityPub bridge. ([HN](https://news.ycombinator.com/item?id=34748669)) ([Code](https://sr.ht/~cloutier/bird.makeup/))
- [TwitterBackup](https://github.com/one-among-us/TwitterBackup) - Backup other users' twitter accounts with Twitter API.
- [next-tweet](https://github.com/vercel-labs/next-tweet) - Embedded and static tweet for Next.js applications.
- [abbrevia.me](https://abbrevia.me/) - Abbreviate users' information from their latest tweets. ([Code](https://github.com/heedrox/abbreviame))
- [Twitter's Recommendation Algorithm (2023)](https://blog.twitter.com/engineering/en_us/topics/open-source/2023/twitter-recommendation-algorithm) ([HN](https://news.ycombinator.com/item?id=35391433)) ([Code](https://github.com/twitter/the-algorithm))

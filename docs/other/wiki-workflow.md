# My workflow in writing and maintaining this wiki

The entire wiki is open source on [GitHub](https://github.com/nikitavoloboev/knowledge). Everything is [authored with markdown](https://github.com/nikitavoloboev/knowledge/tree/main/docs). On every commit, [Vercel](https://vercel.com) that is hooked up to the GitHub repo runs [Docusaurus](https://docusaurus.io) and new version of wiki is published.

The order of items in sidebar is alphabetical. Previously I used [GitBook](https://www.gitbook.com/) for publishing and that used [SUMMARY.md](https://github.com/nikitavoloboev/knowledge/blob/main/SUMMARY.md) file for custom order of items in sidebar.

I write and edit all the markdown files in this wiki from [Sublime Text](../text-editors/sublime-text/sublime-text.md) with [Vim mode](https://github.com/guillermooo/Six).

I use [this Alfred workflow](https://github.com/nikitavoloboev/small-workflows#personal-workflows) to quickly search through all the `.md` files that are found in this wiki.

![](https://i.imgur.com/t8oHfjO.png)

This workflow also searches through all folders in this wiki and lets me create new folders inside those folders or create new entries (md files) inside the folders.

I also wrote a [comment that goes in more detail over the ways I update the wiki](https://lobste.rs/s/e5lx5p/what_note_taking_team_wiki_personal_wiki#c_vczr1n).

I use [Alfred My Mind](https://github.com/nikitavoloboev/alfred-my-mind) to search for wiki entry to open/share. I plan to improve this system a lot going forward.

## Editing on iOS

I open the [docs part of the GitHub repo](https://github.com/nikitavoloboev/knowledge/tree/main/docs) folder in [Obsidian Mobile](https://obsidian.md/mobile). This lets me edit the wiki on the go. Most times I write messages to myself on Telegram and transfer it later to my wiki on mac.

## Making notes on books/papers

When reading books/papers, I annotate PDF files with [Preview](<https://en.wikipedia.org/wiki/Preview_(macOS)>) and ePub files with [Books app](https://www.apple.com/apple-books/). Once I read the book/paper I transfer the annotations to the markdown file and connect the file to wiki. I often write a review on Goodreads and add the review and notes I made to [books](../books/books.md).

This system needs improvement though and I want to try using [Readwise](https://readwise.io/) and more automation for both reading things more efficiently and taking notes.

## Using this wiki

Now that I can edit this wiki and extend it at the speed of a thought. I can open any of the wiki entries in seconds by searching for the file I need with [Alfred My Mind](https://github.com/nikitavoloboev/alfred-my-mind).

![](https://i.imgur.com/vRtsNbq.png)

And anyone who owns Alfred can download the workflow and search through the wiki too. Outside of that, GitBook also provides a pretty neat search in the wiki itself. In there you can make textual queries too and files where the queried text is will be shown.

## Wiki reference

I try to structure everything inside this wiki as understandable and readable by anyone. Sometimes I use abbreviations, so if you find anything confusing in this wiki, read here and if it's still confusing, [ask a question on GitHub](https://github.com/nikitavoloboev/knowledge/issues/new?assignees=&labels=question&template=question.md)

#### Abbreviations

- LA = [Learn Anything](../ideas/learn-anything.md)
- KM = [Keyboard Maestro](../macOS/apps/keyboard-maestro/keyboard-maestro.md). Will often reference [KM macros I made](../macOS/apps/keyboard-maestro/km-macros.md).
- repo = GitHub repository
- bound to .. = may say things like `bound to v + r` which means that I created an action that activates when I hold `v` key and then press `r`. For how that works read [this](../macOS/apps/karabiner/karabiner.md).

#### Years in Links

Nearly every page in the wiki has a `## Links` section. Those are essentially my bookmarks. I add a year in brackets to links like **(2017)** which stands for the year the resource be it article/video/.. was published in.

## [Alfred My Mind](https://github.com/nikitavoloboev/alfred-my-mind)

I wrote an Alfred workflow that as the name suggests, acts as my second brain that I use to quickly access anything I have [indexed for myself](../knowledge/knowledge.md). I share it because I wish other people took this idea and made their own _Alfred My Mind_ or a _Knowledge wiki_ but for their own use. Sharing knowledge and tools like that is very powerful and makes a big difference. As I use and reuse other people's work and knowledge many times a day. It never hurts to add to the mix.

I want to visualize knowledge and thoughts and ideas I have and share most of these things with other people through this wiki, [the longer articles I write](../sharing/my-articles.md), the [YouTube videos I make](https://www.youtube.com/channel/UCEKqrUfr_FMKIO9XSJS4vDw/videos) and of course finally through the [code I write](../sharing/my-github.md) to make useful tools, apps and interactive visualizations that all can use and enjoy.

## Future improvements

I want to:

1. Have better access to notes/links. [Alfred My Mind](https://github.com/nikitavoloboev/alfred-my-mind) is lacking. I want faster access.
2. The graph you see on [main page](../intro.md) with nodes and connections (generated by [Obsidian](../tools/obsidian.md)), I want it on the web and interactive and up to date.
3. Improve the search in here with Docusaurus. Don't like it currently and some pages like [rust libraries](../programming-languages/rust/rust-libraries/rust-libraries.md) are not even being indexed for being too big.
4. Be more active in sharing my knowledge in smaller tidbits. [Simon Willison’s Weblog](https://simonwillison.net/) is amazing example of this. So many TILs on so many topics. I want all my useful knowledge and references indexed digitally.

## Similar wikis I liked

- [Devine Lu Linvega](https://wiki.xxiivv.com/site/home.html) ([Code](https://github.com/XXIIVV/Oscean)) ([Index](https://wiki.xxiivv.com/site/index.html))
- [100 Rabbits](https://100r.co/site/home.html) ([Code](https://github.com/hundredrabbits/100r.co))
- [Gwern](https://www.gwern.net/) - Can also suffix URLs with `.page` to get the source code. ([Code](https://github.com/gjord/gwern.net))
- [Paul Copplestone](https://paul.copplest.one/knowledge/) ([Code](https://github.com/kiwicopple/paul.copplest.one))
- [Andyʼs working notes](https://notes.andymatuschak.org/About_these_notes) ([HN](https://news.ycombinator.com/item?id=24288480)) ([Obsidian version](https://publish.obsidian.md/andymatuschak/_Start+Here))
- [Yoshua Wuyts](https://github.com/yoshuawuyts/notes)
- [Oleg Kiselyov](http://okmij.org/ftp/) - Lots of stuff on FP.
- [Richard Litt](https://github.com/RichardLitt/knowledge)
- [Darshan Chaudhary](https://github.com/darshanime/notes)
- [Josh Avanier](https://joshavanier.github.io/wiki/notes.html)
- [Brennan Letkeman](https://ltkmn.gitbook.io/brendex/)
- [Andy Sparks](https://github.com/AndySparks/captains-log)
- [Exobrain of Dima Gerasimov](https://beepb00p.xyz/exobrain/) ([Code](https://github.com/karlicoss/exobrain))
- [Krzysztof Kowalczyk](https://blog.kowalczyk.info/)
- [Phil Eaton's notes](http://notes.eatonphil.com/)
- [Blue Book](https://lyz-code.github.io/blue-book/) ([Code](https://github.com/lyz-code/blue-book/))
- [Jethro Kuan](https://braindump.jethro.dev/zettels/) ([Code](https://github.com/jethrokuan/braindump))
- [C2](http://wiki.c2.com/)
- [Allen](https://allenleein.github.io/1930/01/01/knowledge-base.html) ([Code](https://github.com/allenleein/knowledge-base))
- [Full-stack web development notes](https://github.com/8483/notes)
- [Max Masnick](https://maxmasnick.com/kb/)
- [Uncurled](https://un.curl.dev/) - Wiki by Daniel Stenberg. ([Code](https://github.com/bagder/uncurled))
- [Leon Bambrick](https://til.secretgeek.net/) ([Code](https://github.com/secretGeek/today-i-learned))
- [Ben Lynn](http://www-cs-students.stanford.edu/~blynn/)
- [Harrison's Wiki](https://wiki.harrison.dev/)
- [Greg's wiki](https://mywiki.wooledge.org/EnglishFrontPage)
- [Brandur's Fragments](https://brandur.org/fragments)
- [Rachel Brindle](https://knowledge.rachelbrindle.com/) ([Code](https://github.com/younata/personal_knowledge))
- [Ioannis Kourouklides](https://wiki.kourouklides.com/wiki/Main_Page)
- [Dmitrii Gerasimov](https://beepb00p.xyz/) ([Code](https://github.com/karlicoss/beepb00p))
- [Flavio Copes](https://flaviocopes.com/)
- [Ellie's wiki](https://ellie.wiki/) ([Code](https://github.com/elliebike/wiki))
- [Max Stoiber's notes](https://notes.mxstbr.com/Evergreen_notes)
- [Tom Critchlow](https://tomcritchlow.com/wiki/)
- [panthema.net](http://bingmann.github.io/) - Diverse collection of interesting ideas.
- [Buster Benson](https://busterbenson.com/piles/)
- [Burke Libbey's notes](https://notes.burke.libbey.me/)
- [The Commonplace Book](https://www.ralphrudd.com/)
- [Leandro Ardissone](https://knowledge.lardissone.now.sh/#about-me)
- [Nicola's decentralized-research](https://github.com/nicola/decentralized-research)
- [Timothy Andrew](https://timothyandrew.net/learning/wiki/)
- [Wayan Jimmy](https://wayanjimmy-notebook.netlify.com/)
- [Wesley Moore](https://linkedlist.org/)
- [Josh Branchaud's TIL](https://github.com/jbranchaud/til)
- [Lee Byron's TIL](https://leebyron.com/til/) ([Code](https://github.com/leebyron/til))
- [Slides of Chen Hui Jing](https://github.com/huijing/slides)
- [Hongyi Shen's notes](https://github.com/wilbeibi/NotesIndex)
- [Shreyas Minocha](https://wiki.shreyasminocha.me/index.html) ([Code](https://github.com/shreyasminocha/wiki))
- [NicoHood](https://github.com/NicoHood/NicoHood.github.io/wiki)
- [Sanyam Kapoor](https://www.sanyamkapoor.com/kb)
- [Blake Robbins](https://blakeir.com/notes)
- [Bismuth Garden](https://bismuth.garden/)
- [Alopex Networks wiki](https://wiki.alopex.li/_categories)
- [Michael Mellinger math notes](https://github.com/melling/MathAndScienceNotes)
- [Bookmarks, tweets, cheatsheets, one-liners](https://github.com/ccampean/almanacs)
- [Bayle Shanks](http://www.bayleshanks.com/)
- [Noah Trenaman](https://blog.noahtren.com/note/3acea2e1/)
- [Matous Dzivjak](https://github.com/matoous/wiki)
- [Kevin S Lin](https://www.kevinslin.com/)
- [Alex's Zettelkasten](https://notes.alexkehayias.com/)
- [Yizhou Shan](http://lastweek.io/)
- [The Refined Mind](https://refinedmind.co/)
- [Azlen Elza](https://notes.azlen.me/g3tibyfv/) ([Code](https://github.com/azlen/azlen.me)) ([Home site](https://azlen.me/))
- [Vlad Patryshev](https://github.com/vpatryshev/wowiki/wiki)
- [Brady Joslin's TIL](https://til.bradyjoslin.com/_introduction/)
- [Rosie Campbell's Notes](https://rosiecampbell.me/notes)
- [Michael Lubinsky's notes and code snippets](http://mlubinsky.github.io/) ([Code](https://github.com/mlubinsky/mlubinsky.github.com))
- [Chris Nkinthorn](https://nkintc.gitbook.io/brainless/) ([Code](https://github.com/nkintc/nkintc.github.io))
- [TIL by Stefan Judis](https://www.stefanjudis.com/today-i-learned/)
- [David Gasquez](https://publish.obsidian.md/davidgasquez/Handbook)
- [Abigail Africa's Notes](https://www.notion.so/Notes-3a56fef72dc94bc28b0804efc789913b)
- [Sean Breckenridge](https://exobrain.sean.fish/) ([Code](https://github.com/seanbreckenridge/exobrain))
- [Nat Eliason notes](https://www.nateliason.com/notes)
- [Bharat’s Digital Garden](https://bharatkalluri.com/notes)
- [Chrisman Brown](https://github.com/chrisman/knowledge/wiki)
- [Sridhar Ratnakumar](https://www.srid.ca/z-index.html)
- [Timothy Andrew](https://timothyandrew.net/learning/)
- [Dennis Ideler's notes](https://github.com/dideler/notes)
- [nchrs.xyz](https://nchrs.xyz/site/home.html)
- [Anne-Laure Le Cunff](https://www.mentalnodes.com/)
- [Paul Batchelor](https://pbat.ch/wiki/)
- [Notes of InvertedPassion](https://notes.invertedpassion.com/_Start+here_)
- [Alex's Garden](https://publish.obsidian.md/alexander/1_Home/%F0%9F%8F%A0+Home)
- [Yang Zhang CS notes](https://github.com/yang/notes)
- [Quinn Casey](https://quinncasey.com/projects/digital-garden/)
- [Ten Digits](https://tendigits.space/)
- [Alex Forencich's Wiki](http://www.alexforencich.com/wiki/en/start)
- [Luciano Strika](https://strikingloo.github.io/wiki/)
- [Bennett Hardwick](https://bennetthardwick.com/wiki/)
- [Micheels Learning Notes](https://learning-notes.mistermicheels.com/about/about/) ([Code](https://github.com/mistermicheels/learning-notes))
- [Neil's Noodlemaps](https://commonplace.doubleloop.net/)
- [Jesse Squires TIL](https://jessesquires.github.io/TIL/) ([Code](https://github.com/jessesquires/TIL))
- [Piotr Gaczkowski](https://garden.doomhammer.info/)
- [chl's org file notes](https://github.com/even4void/org)
- [Cosma Shalizi](http://bactra.org/) ([Notebooks](http://bactra.org/notebooks/))
- [Roy Niang](https://royniang.com/home.html) ([Code](https://git.sr.ht/~royniang/koikoi))
- [Rob Haisfields's Roam](https://roamresearch.com/#/app/Rob-Haisfield-Thinking-in-Public/page/XtzimCS2z)
- [Hyperfine Village](https://roamresearch.com/#/app/hyperfinelabs/page/TYt89wtA7)
- [Ben Lynn Notes](https://crypto.stanford.edu/pbc/notes/)
- [Nick Belzer](https://notes.nickbelzer.me/) ([Code](https://github.com/nbelzer/notes))
- [Mike's Mind](https://mind.miketannenbaum.com/) ([Code](https://github.com/MikeTannenbaum/my-public-mind))
- [David Seah](https://davidseah.gitbook.io/davidseah/) ([Code](https://github.com/davidseah/knowledgebank))
- [John Ohno](http://www.lord-enki.net/) ([Misc Files](https://github.com/enkiv2/misc))
- [Allan's Digital Garden](https://publish.obsidian.md/allanmacgregor/Meta/Index)
- [Dercuano](https://gitlab.com/kragen/dercuano) - Although not freely available on the web, have to download a folder.
- [Aurelio](https://github.com/nobitagit/knowledge)
- [Thomasorus's Garden](https://thomasorus.com/home.html)
- [Yenly Ma TIL](https://til.yenly.wtf/)
- [Nathan's Manuals](https://github.com/sandersn/manual)
- [Magnus's Brain](https://kmaasrud.com/brain/)
- [Abhinav's Notes](https://notes.abhinavsarkar.net/)
- [Hugo Cisneros](https://hugocisneros.com/notes/)
- [McCoy R. Becker](https://femtomc.github.io/)
- [Rosano Garden](https://rosano.hmm.garden/)
- [Matt Titmus TIL](https://github.com/clockworksoul/today-i-learned)
- [Scott Spence's Garden](https://scottspence.com/garden)
- [Uncertainty Mindset Wiki](https://uncertaintymindset.org/UM1/index.html)
- [Kormyen's Memex](https://kormyen.github.io/memex/)
- [Chotrin's Wiki](https://chotrin.tilde.institute/)
- [Ten Digits](https://tendigits.space/site/index.html)
- [Sridhar Ratnakumar](https://notes.srid.ca/) ([Code](https://github.com/srid/notes.srid.ca))
- [Uzay-G's Wiki](https://knowledge.uzpg.me/)
- [Clinton Boys's Garden](http://mtsolitary.com/)
- [Brett Abramczyk's Wiki](https://github.com/babramczyk/wiki)
- [Artyom's tech notes](https://tek.brick.do/)
- [D. Bohdan's Wiki](https://dbohdan.com/wiki/sitemap)
- [Albert Zeyer](https://github.com/albertz/wiki)
- [Fabian Bösiger's Docs](https://github.com/fabianboesiger/documents)
- [Juliette Rapala](https://github.com/jrapala/wiki)
- [Daniel Imfeld's Notes](https://imfeld.dev/notes)
- [Rikard Hjort's Wiki](https://github.com/hjorthjort/wiki)
- [Anthony's second brain](https://garden.anthonyamar.fr/Welcome+in+my+mind+%F0%9F%A7%A0)
- [Hrishikesh Bhaskaran's Wiki](https://wiki.stultus.in/)
- [Jason Yuchen's Notes](https://github.com/JasonYuchen/notes)
- [neeasade's notes](https://notes.neeasade.net/)
- [Ferreira's Wiki](https://github.com/slowernews/notebook)
- [マリウス](https://xn--gckvb8fzb.com/)
- [Kirill Gorbachyonok's Wiki](https://github.com/japanese-goblinn/knowledge-base)
- [Zac's digital garden](https://zacjones.dev/) ([Code](https://github.com/zacjones93/zacs-digital-garden))
- [Robb Knight's Wiki](https://intersect.rknight.me/) ([Code](https://github.com/rknightuk/intersect))
- [Edward Yerburgh's CS Notes](https://notes.eddyerburgh.me/) ([Code](https://github.com/eddyerburgh/notes))
- [Nimor's Zettelkasten](https://nimor111.github.io/notebook/) ([Code](https://github.com/Nimor111/notebook))
- [Gordon Brander's Patterns](http://gordonbrander.com/pattern/)
- [EyebrowHairs](https://www.eyebrowhairs.com/) ([Code](https://github.com/EyebrowHairs/garden))
- [Pepe Doval's Wiki](https://github.com/pepellou/knowledge)
- [Jacob Chvatal](https://github.com/jakeisnt/wiki) ([Web](https://anagora.org/@jakeisnt/wiki))
- [masayume](https://github.com/masayume/DigitalGarden)
- [Kate Bowie](https://katebowies-garden.netlify.app/)
- [Binny's garden](https://notes.binnyva.com/)
- [David Winterbottom's TIL](https://til.codeinthehole.com/)
- [Alex Schroeder: Diary](https://alexschroeder.ch/)
- [Well Ordered Wiki](https://wiki.wellorder.net/wiki/)
- [minimal.click](https://minimal.click/)
- [SUDOGAMI](https://sudogami.com/site/home.html)
- [Map of Computing](https://github.com/ComputingTeachers/mapOfComputing)
- [Simon Willison: TIL](https://til.simonwillison.net/) ([Code](https://github.com/simonw/til))
- [Brian DeVries's Wiki](https://github.com/techCarpenter/garage-wiki)
- [Nick Groenen's digital garden](https://notes.nick.groenen.me/)
- [Brandon's Digital Garden](https://wiki.brandontoner.ca/)
- [Param Singh's Notes](https://notes.param.codes/)
- [Matous Dzivjak](https://wiki.dzx.cz/)
- [Evan's Notes](https://peppyhare.github.io/r/) ([Code](https://github.com/PeppyHare/r))
- [Allen Lee](https://www.haskell.computer/) ([Code](https://github.com/ale0sx/brains))
- [Keyvan Akbary's Notes](https://keyvanakbary.github.io/learning-notes/) ([Code](https://github.com/keyvanakbary/learning-notes))
- [Wojciech Matuszewski](https://github.com/WojciechMatuszewski/programming-notes)
- [Eduardo Julian's bookmarks](https://github.com/LuxLang/lux/tree/master/documentation/bookmark)
- [Raphael Sprenger](https://garden.raphaelsprenger.de/index.html)
- [Личная вики](https://ndrewnee.gitbook.io/wiki/)
- [Boris Mann](https://bmannconsulting.com/)
- [Yu-An Chan's wiki](https://alanchan.netlify.app/index.html)
- [The Wise Opportunist](https://opportunist.luseeds.com/)
- [Constantin Campean's Almanacs](https://github.com/costinEEST/almanacs)
- [Allen Lee's Digital Brain](https://github.com/ale0sx/brains)
- [Ruben Rodrigues's Wiki](https://github.com/rubrodc/ruben-2.0)
- [Szymon Kaliski's Notes](https://szymonkaliski.com/notes)
- [Vegard Stikbakke's TILs](https://til.vegardstikbakke.com/) ([Code](https://github.com/vegarsti/til))
- [Dev Notes](https://github.com/mebusy/notes)
- [Rahul Rajeev](https://garden.rahulrajeev.net/)
- [Rickard Dag's Log](https://devlog.willcodefor.beer/) ([Code](https://github.com/believer/devlog))
- [Owen's Digigarden](https://publish.obsidian.md/geobo/Welcome+to+my+digital+garden+%F0%9F%8C%B1)
- [Charlie T](https://blog.charlietrochlil.com/) ([Code](https://github.com/rhymeswithvocal/digital-garden))
- [Imron Rosyadi](https://irosyadi.gitbook.io/irosyadi/) ([Code](https://github.com/irosyadi/gitbook))
- [Mineral Existence](https://mineralexistence.com/wiki.html) ([Code](https://github.com/flber/Mineral-Existence))
- [List of Exobrain Blogs and Digital Gardens](https://irosyadi.gitbook.io/irosyadi/app/exobrain-digital-garden)
- [Meta knowledge](https://github.com/RichardLitt/knowledge) - More wikis.
- [Webring](https://webring.xxiivv.com/) - Attempt to inspire artists & developers to build their own website and share traffic among each other. ([Code](https://github.com/XXIIVV/webring))
- [Curated list of awesome Public Zettelkastens](https://github.com/KasperZutterman/Second-Brain)
- [Best of Digital Gardens of GitHub](https://github.com/lyz-code/best-of-digital-gardens)

## Wiki software

- [Docusaurus](https://docusaurus.io)
- [GitBook](https://www.gitbook.com/)
- [Oscean](https://github.com/XXIIVV/Oscean) - Flow-based serverless wiki.
- [WeeWiki](https://github.com/PaulBatchelor/weewiki) - Wee little wiki engine used to generate personal wikis and mind maps.
- [Dnote](https://github.com/dnote/dnote) - Simple personal knowledge base.
- [Instiki](https://github.com/parasew/instiki) - Basic wiki clone so pretty and easy to set up, you’ll wonder if it’s really a wiki.

## Notes

- [Don't waste energy chasing fancy tools and methodologies without already having a simple workflow in place. Once you have a good idea of what works for you, then introduce tools designed to make your life easier.](https://news.ycombinator.com/item?id=24251068)
- [Reading without taking notes is useless.](https://twitter.com/anthilemoon/status/1261991953593401346)
- [Feed your personal wiki into GTP-3 to automatically answer emails/..](https://twitter.com/maccaw/status/1441773166079184896)
- [People are obsessed with migrating their notes and todos from one tool to another. The obsession with meta work as an excuse for not doing actual work is such an odd one.](https://twitter.com/ninarstepanov/status/1450515900822523907)
- [A tragically underutilized fact in productivity software today is that most people's entire textual datasets for a lifetime can fit in modern PCs' RAM. Just load it up & search it in memory. We don't need to send everything across the planet. Things can be /so much/ faster.](https://twitter.com/thesephist/status/1455279982322749440)
- [I learned today that a group of students used a Google doc to take lecture notes-- they all took notes simultaneously in a collective file.](https://twitter.com/mckellogs/status/811339472205910016)

## Links

- [Does anyone else keep their own knowledge wiki?](https://lobste.rs/s/ord0rg/does_anyone_else_keep_their_own_knowledge)
- [Show HN: Everything I Know Wiki (2019)](https://news.ycombinator.com/item?id=19468993)
- [Ask HN: Do you keep a personal knowledge repository? (2019)](https://news.ycombinator.com/item?id=20007108)
- [Patterns of personal knowledge bases](https://github.com/zadam/trilium/wiki/Patterns-of-personal-knowledge-base)
- [Ask HN: How do you share/organize knowledge at work and life? (2019)](https://news.ycombinator.com/item?id=21310030)
- [Ask HN: How did you build up your personal knowledge base? (2019)](https://news.ycombinator.com/item?id=21332957)
- [Note-taking strategy 2019](https://cpbotha.net/2019/09/21/note-taking-strategy-2019/)
- [Smart Sync Workshop: David Perell + Tiago Forte (2019)](https://www.youtube.com/watch?v=lNJ33ImlZzs)
- [Notion Office Hours: Tiago Forte (2019)](https://www.youtube.com/watch?v=sDNooHDj2Dk)
- [How to annotate literally everything](https://beepb00p.xyz/annotating.html) - Comprehensive overview of existing tools, strategies and thoughts on interacting with your data. ([HN](https://news.ycombinator.com/item?id=21635012))
- [Digital Tools I Wish Existed (2019)](https://jborichevskiy.com/posts/digital-tools/) ([HN](https://news.ycombinator.com/item?id=21659876))
- [The sad state of personal data and infrastructure (2019)](https://beepb00p.xyz/sad-infra.html)
- [What tools do you use to maintain a personal log/journal? (2020)](https://lobste.rs/s/peevtw/what_tools_do_you_use_maintain_personal)
- [Managing my personal knowledge base (2020)](https://tkainrad.dev/posts/managing-my-personal-knowledge-base/) ([HN](https://news.ycombinator.com/item?id=22000791))
- [my](https://github.com/karlicoss/my) - Python interface into my life.
- [What data on myself I collect and why? (2020)](https://beepb00p.xyz/my-data.html) ([HN](https://news.ycombinator.com/item?id=26045797))
- [Zettelkästen? (2019)](https://clerestory.netlify.com/zk/)
- [Nototo](https://www.nototo.app/) - All your notes. On one map. ([HN](https://news.ycombinator.com/item?id=22087780))
- [BookStack](https://www.bookstackapp.com/) - Simple, self-hosted, easy-to-use platform for organising and storing information. ([HN](https://news.ycombinator.com/item?id=29851834))
- [Oscean](https://wiki.xxiivv.com/site/oscean.html) - Static wiki engine written entirely in C, designed to be deployed from low-power devices with gcc as its only dependecy.
- [Building personal search infrastructure for your knowledge and code (2019)](https://beepb00p.xyz/pkm-search.html) ([HN](https://news.ycombinator.com/item?id=22160572))
- [Against unnecessary databases (2020)](https://beepb00p.xyz/unnecessary-db.html)
- [Promnesia](https://github.com/karlicoss/promnesia) - Another piece of your extended mind. ([Article](https://beepb00p.xyz/promnesia.html)) ([HN](https://news.ycombinator.com/item?id=23668507))
- [My productivity app for the past 12 years has been a single .txt file (2020)](https://jeffhuang.com/productivity_text_file/) ([HN](https://news.ycombinator.com/item?id=22276184)) ([Lobsters](https://lobste.rs/s/ettc1n/my_productivity_app_is_single_txt_file))
- [Contextualise](https://github.com/brettkromkamp/contextualise) - Simple and flexible tool particularly suited for organising information-heavy projects and activities consisting of unstructured and widely diverse data and information resources. ([Web](https://contextualise.dev/))
- [How To Take Smart Notes With Org-mode (2020)](https://blog.jethro.dev/posts/how_to_take_smart_notes_org/) ([HN](https://news.ycombinator.com/item?id=22337681))
- [Copernic](http://copernic.space/) - Aims to make practical cooperation around the creation, publication, storage, re-use and maintenance of knowledge bases, and in general structured data that are bigger than memory.
- [Zettelkasten Method](https://zettelkasten.de/posts/overview/)
- [A Text Renaissance (2020)](https://www.ribbonfarm.com/2020/02/24/a-text-renaissance/) ([HN](https://news.ycombinator.com/item?id=22442027))
- [Ask HN: Good ways to capture institutional knowledge? (2020)](https://news.ycombinator.com/item?id=22454333)
- [Ask HN: How to Take Good Notes? (2020)](https://news.ycombinator.com/item?id=22473209)
- [How To Take Smart Notes: 10 Principles to Revolutionize Your Note-Taking and Writing (2020)](https://fortelabs.co/blog/how-to-take-smart-notes/) ([HN](https://news.ycombinator.com/item?id=22341518))
- [Hackpad](https://github.com/dropbox/hackpad) - Web-based realtime wiki, based on the open source EtherPad collaborative document editor.
- [Human Programming Interface (2020)](https://lobste.rs/s/vzkb7a/human_programming_interface_python) - My life in a Python package.
- [Emvi](https://emvi.com/) - Knowledge management platform for companies and teams.
- [Human Programming Interface: My life in a Python package (2020)](https://beepb00p.xyz/hpi.html) ([HN](https://news.ycombinator.com/item?id=23101869))
- [Ask HN: What do you use to keep track of bookmarks/notes/snippets? (2020)](https://news.ycombinator.com/item?id=22778123)
- [Zettelkasten — How One German Scholar Was So Freakishly Productive (2019)](https://writingcooperative.com/zettelkasten-how-one-german-scholar-was-so-freakishly-productive-997e4e0ca125) ([Lobsters](https://lobste.rs/s/syoikp/zettelkasten_how_one_german_scholar_was))
- [Memex](https://github.com/WorldBrain/Memex) - Browser Extension to full-text search your browsing history & bookmarks. ([Web](https://worldbrain.io/)) ([HN](https://news.ycombinator.com/item?id=23227186)) ([Interview with Oliver Sauter](https://www.youtube.com/watch?v=XH2nbgEQ78M))
- [Memex Mobile](https://github.com/WorldBrain/Memex-Mobile) - Mobile app for Memex.
- [Thread on digital gardens, personal wikis, and experimental knowledge systems (2020)](https://twitter.com/Mappletons/status/1250532315459194880)
- [Dogsheep](https://dogsheep.github.io/) - Collection of tools for personal analytics using SQLite and Datasette. ([Article](https://datasette.substack.com/p/dogsheep-personal-analytics-with)) ([GitHub](https://github.com/dogsheep))
- [Lobsters: How do you take notes and organize your knowledge? (2020)](https://lobste.rs/s/7catij/how_do_you_take_notes_organize_your)
- [A Short History of Bi-Directional Links (2020)](https://maggieappleton.com/bidirectionals)
- [Digital Gardening](https://github.com/MaggieAppleton/digital-gardeners) - Resources, links, projects, and ideas for gardeners tending their digital notes on the public interwebs.
- [Neuron Zettelkasten](https://neuron.zettel.page/) - Command-line based system for managing your Zettelkasten. ([Code](https://github.com/srid/neuron)) ([Lobsters](https://lobste.rs/s/i1uylr/neuron_0_6_released_future_proof_note)) ([1.0 release lobsters](https://lobste.rs/s/ve9plk/neuron_1_0_released_future_proof_note))
- [Neuron Template](https://github.com/srid/neuron-template) - How to publish your own neuron site.
- [How can we build an extension of your mind? (2020)](https://vanschneider.com/how-can-we-build-an-extension-of-your-mind)
- [mmap.it](https://www.mmap.it/) - Map knowledge into memory with seamless search and note taking. ([Code](https://github.com/pdepip/mmap.it))
- [Luhmann’s Zettelkasten — A Productivity Tool That Works Like Your Brain (2019)](https://emvi.com/blog/luhmanns-zettelkasten-a-productivity-tool-that-works-like-your-brain-N9Gd2G4aPv)
- [Building a digital garden (2019)](https://tomcritchlow.com/2019/02/17/building-digital-garden/) - How I built myself a simple wiki using folders and files and published via Jekyll.
- [monotome](https://github.com/cblgh/monotome) - Personal knowledge base system. markdown markup, runs in the browser.
- [Digital Gardeners](https://nesslabs.com/digital-gardeners) - Small Telegram group for people actively using & building digital gardens. ([Shared notes](https://github.com/digitalgardeners/notes))
- [Thinking in maps: from the Lascaux caves to knowledge graphs (2020)](https://nesslabs.com/thinking-in-maps)
- [Trilium Notes](https://github.com/zadam/trilium) - Hierarchical note taking application with focus on building large personal knowledge bases. ([HN](https://news.ycombinator.com/item?id=23335759)) ([HN](https://news.ycombinator.com/item?id=32130401))
- [Ask HN: What is your favorite way to read online content? (2020)](https://news.ycombinator.com/item?id=23340031)
- [Digital garden Jekyll template](https://github.com/maximevaillancourt/digital-garden-jekyll-template)
- [Mathematicians, how do you keep your notes? Why? (2020)](https://www.reddit.com/r/math/comments/gt0adz/mathematicians_how_do_you_keep_your_notes_why/)
- [Stop Taking Regular Notes; Use a Zettelkasten Instead (2020)](https://eugeneyan.com/2020/04/05/note-taking-zettelkasten/) ([HN](https://news.ycombinator.com/item?id=23386630))
- [Personal Wiki for Vim](https://github.com/vimwiki/vimwiki) ([HN](https://news.ycombinator.com/item?id=23402014))
- [Zettelkasten note-taking in 10 minutes (2020)](https://blog.viktomas.com/posts/slip-box/) ([HN](https://news.ycombinator.com/item?id=23445742))
- [Networked Notebooks Catalogue](https://github.com/prathyvsh/networked-notebooks) - Collection of networked notebooks that is slowly taking shape on the Internet.
- [Building a Second Brain: The Illustrated Notes (2020)](https://maggieappleton.com/basb) ([HN](https://news.ycombinator.com/item?id=23514371))
- [Collected Notes](https://collectednotes.com/) - Simplest, and most powerful note-taking blogging platform. ([Code](https://news.ycombinator.com/item?id=23514682))
- [Kumu](https://kumu.io/) - Makes it easy to organize complex data into relationship maps that are beautiful to look at and a pleasure to use.
- [Relanote](https://relanote.com/) - Connect your notes into a knowledge graph.
- [Milanote](https://milanote.com/) - Tool for organizing creative projects.
- [mymind](https://mymind.com/) - Extension for your mind.
- [Foam](https://github.com/foambubble/foam) - Personal knowledge management and sharing system for VSCode. ([Foam Workspace](https://mathieudutour.github.io/foam-gatsby-template/)) ([Lobsters](https://lobste.rs/s/xufl8d/foam_personal_knowledge_management)) ([Web](https://foambubble.github.io/foam/)) ([HN](https://news.ycombinator.com/item?id=23666950)) ([HN 2](https://news.ycombinator.com/item?id=25760066))
- [Supernotes](https://supernotes.app/) - Collaborative note-taking app. ([HN](https://news.ycombinator.com/item?id=25790365)) ([HN](https://news.ycombinator.com/item?id=30440275))
- [Gthnk](http://www.gthnk.com/) - Personal Journal.
- [Growing the Evergreens (2020)](https://maggieappleton.com/evergreens)
- [DigitalGardens subreddit](https://www.reddit.com/r/DigitalGardens/)
- [Gardener](https://github.com/BharatKalluri/gardener) - Command line tool to help you manage your Knowledge management system / Digital Garden.
- [On Bookmarks and Note Taking (2020)](https://chrisman.github.io/11.html) ([Lobsters](https://lobste.rs/s/eg45vr/on_bookmarks_note_taking))
- [Using VimWiki for Note Taking (2020)](https://thelinell.com/using-vimwiki/) ([Lobsters](https://lobste.rs/s/gynavy/using_vimwiki_for_note_taking))
- [Artificial Brain Networked notebook apps](https://www.notion.so/Artificial-Brain-Networked-notebook-app-a131b468fc6f43218fb8105430304709)
- [Ask HN: How to Take Notes? (2020)](https://news.ycombinator.com/item?id=23844490)
- [histre](https://histre.com/) - Effortless Knowledge Base.
- [Notebooks (2020)](https://blog.yoshuawuyts.com/notebooks/)
- [Digital gardens (2020)](https://schof.co/digital-garden/)
- [Zettlr](https://www.zettlr.com/) - A Markdown Editor for the 21st Century. ([Why is Zettlr Open Source?](https://www.zettlr.com/post/why-zettlr-open-source)) ([Code](https://github.com/Zettlr/Zettlr))
- [CodiMD](https://github.com/codimd/server) - Self-hosted, real-time collaborative markdown notes. ([HN](https://news.ycombinator.com/item?id=23997361))
- [Logseq](https://logseq.com/) - Local-first knowledge base which syncs using Github. ([Code](https://github.com/logseq/logseq)) ([HN](https://news.ycombinator.com/item?id=25090176)) ([Awesome](https://github.com/logseq/awesome-logseq)) ([Plugin Samples](https://github.com/logseq/logseq-plugin-samples))
- [Sharing knowledge in a remote team, across timezones (2020)](https://erickhun.com/posts/sharing-knowledge-in-a-remote-team/) ([HN](https://news.ycombinator.com/item?id=24021103))
- [Foam: One Month In (2020)](https://jevakallio.github.io/notes/foam-one-month-in)
- [If you take notes, which note-taking tool do you use? (2020)](https://twitter.com/anthilemoon/status/1291646162919993344)
- [Building my personal knowledge repository (2020)](https://notes.nickbelzer.me/topics/2020/05/12/building-my-personal-knowledge-repository.html)
- [An Iterative Approach to Notes (2020)](https://medium.com/@bytebase/an-iterative-approach-to-notes-f1c2a28c4d29) ([HN](https://news.ycombinator.com/item?id=24108466))
- [Where is your notebook? (2020)](https://bowero.nl/blog/2020/08/16/where-is-your-notebook/) ([Lobsters](https://lobste.rs/s/lmh8gj/where_is_your_notebook))
- [Archivy](https://github.com/Uzay-G/archivy) - Self Hosted Knowledge Base embedded into your filesystem. ([HN](https://news.ycombinator.com/item?id=24199419))
- [How to choose the right note-taking app (2020)](https://nesslabs.com/how-to-choose-the-right-note-taking-app) ([Tweet](https://twitter.com/anthilemoon/status/1296047396976316416))
- [Using Obsidian to manage goals, tasks, notes, and software dev knowledge base (2020)](https://joshwin.imprint.to/post/how-i-use-obsidian-to-manage-my-goals-tasks-notes-and-software-development-knowledge-base) ([HN](https://news.ycombinator.com/item?id=24206727))
- [Hode](https://github.com/JeffreyBenjaminBrown/hode) - Hypergraph editor. ([Explanation](https://news.ycombinator.com/item?id=24261725))
- [Remembering what you Read: Zettelkasten vs. P.A.R.A (2020)](https://www.zainrizvi.io/blog/remembering-what-you-read-zettelkasten-vs-para/) ([HN](https://news.ycombinator.com/item?id=24251068))
- [Scrapbox](https://scrapbox.io/) - Knowledge base built for infinite ideas.
- [RWX.GG](https://rwx.gg/) - Progressive Knowledge App.
- [Git-based Wiki (2020)](https://www.bit-101.com/blog/2020/09/git-based-wiki/) ([HN](https://news.ycombinator.com/item?id=24351298))
- [DWiki](https://utcc.utoronto.ca/~cks/space/dwiki/DWiki)
- [Cerveau](https://www.srid.ca/689c4a39.html) - Future-proof web app for notes. ([Lobsters](https://lobste.rs/s/xajgsn/cerveau_future_proof_web_app_for_notes))
- [A Meta-Layer for Notes (2020)](https://julian.digital/2020/09/04/a-meta-layer-for-notes/)
- [A Hierarchy First Approach to Note Taking](https://www.kevinslin.com/notes/3dd58f62-fee5-4f93-b9f1-b0f0f59a9b64.html)
- [How to set up your own digital garden (2020)](https://nesslabs.com/digital-garden-set-up)
- [Cosy](https://cosy.af/) - Internal company hub & workplace search.
- [Digital Gardens list](https://docs.google.com/spreadsheets/d/1KtEjnuZEHxUmoiA37_MMM4OFyQcbwVUaLBFa12P8cnU/edit#gid=0)
- [Kb](https://github.com/gnebbia/kb) - Text-oriented minimalist command line knowledge base manager. ([HN](https://news.ycombinator.com/item?id=24506280))
- [List of Zettelkasten software (2020)](https://www.reddit.com/r/Zettelkasten/comments/flygc4/lets_build_a_list_of_zettelkasten_software/)
- [What {note taking|team wiki|personal wiki|external brain} tool do you use? (2020)](https://lobste.rs/s/e5lx5p/what_note_taking_team_wiki_personal_wiki)
- [TiddlyWiki](https://tiddlywiki.com/) - Non-linear personal web notebook. ([Code](https://github.com/Jermolene/TiddlyDesktop))
- [Notes with TiddlyWiki](http://beza1e1.tuxen.de/tiddlywiki_notes.html)
- [Rekowl](https://rekowl.com/) - You personal knowledge library.
- [nb](https://github.com/xwmx/nb) - CLI plain-text notes & bookmarks with Git, sync, encryption, and more. ([HN](https://news.ycombinator.com/item?id=24709393)) ([Docs](https://xwmx.github.io/nb/))
- [NoteApps](https://www.noteapps.info/) - Encyclopedia of note taking apps. ([Launch Notes](https://www.noteapps.info/launch_note)) ([Lobsters](https://lobste.rs/s/uii8of/noteapps_info_launch_notes))
- [Organising knowledge with multi-level content: Making knowledge easier to understand, remember and communicate (2018)](https://www.francismiller.com/organising_knowledge_paper.pdf)
- [Zetteldeft](https://github.com/EFLS/zetteldeft) - Emacs package for a Zettelkasten system.
- [Luhmann's Original Zettelkasten Digitalized (2020)](https://niklas-luhmann-archiv.de/bestand/zettelkasten/zettel/ZK_1_NB_1_1_V) ([HN](https://news.ycombinator.com/item?id=24794569))
- [Walling](https://walling.app/) - Visual Walls to Collect Ideas & Plot out Projects.
- [How do I work with Obsidian on Mobile? (2020)](https://forum.obsidian.md/t/how-do-i-work-with-obsidian-on-mobile/471)
- [Introduction to the Zettelkasten Method (2020)](https://zettelkasten.de/introduction/) ([HN](https://news.ycombinator.com/item?id=24916536))
- [How I remember everything I learn (2020)](https://dev.to/aurelio/how-i-remember-everything-i-learn-19mi)
- [Digital gardens let you cultivate your own little bit of the internet (2020)](https://www.technologyreview.com/2020/09/03/1007716/digital-gardens-let-you-cultivate-your-own-little-bit-of-the-internet/) ([HN](https://news.ycombinator.com/item?id=24996780))
- [Standard Notes](https://standardnotes.org/) - Free, open-source, and completely encrypted notes app. ([Code](https://github.com/standardnotes/web)) ([HN](https://news.ycombinator.com/item?id=26838992)) ([App](https://app.standardnotes.org/)) ([Awesome](https://github.com/jonhadfield/awesome-standard-notes))
- [Linked knowledge greatest links](https://scrapbox.io/benfoden/Linked_knowledge_greatest_links)
- [Information Storage and Archiving strategies (2020)](https://strikingloo.github.io/wiki-articles/information-storage/archiving)
- [zettel](https://github.com/hackstream/zettel) - Notes organizer - based on Zettelkasten methodology. Written in Go.
- [Foamy NextJS Starter](https://github.com/yenly/foamy-nextjs) - Basic Foam + NextJS with MDX starter for building a digital garden.
- [Awesome Knowledge Management](https://github.com/brettkromkamp/awesome-knowledge-management)
- [Memex](https://github.com/kormyen/memex) - Simple bookmarks and notes.
- [POC on how a Memex could potentially work](https://github.com/steve-1820/memex)
- [Agora](https://anagora.org/) - Distributed, goal-oriented social network centered around a cooperatively built and maintained knowledge graph. ([HN](https://news.ycombinator.com/item?id=25573523))
- [Wiki setup with mdBook](https://news.ycombinator.com/item?id=23363574)
- [Exomind](https://github.com/appaquet/exomind) - Personal knowledge management tool hosted on your own personal cloud.
- [Miraheze](https://miraheze.org/) - Community-centric, ad free, and locally controlled wiki hosting platform funded 100% by donations.
- [North Notes](https://northnotes.app/) - Native, infinitely nest-able notes with multi-dimensional tagging.
- [The difference between note-taking and note-making](https://nesslabs.com/from-note-taking-to-note-making) ([HN](https://news.ycombinator.com/item?id=25789073))
- [Notes on wikis](https://github.com/jakeisnt/wiki/blob/master/pages/wiki.org)
- [cadmus](https://github.com/RyanGreenup/cadmus) - Shell Scripts to Facilitate Effective Note Taking.
- [Memo App](https://usememo.com/) - Take Smarter notes with GitHub Gists. ([Code](https://github.com/btk/memo))
- [Totallib](https://totallib.com/) - Note-taking for better thinking, augmented by AI.
- [Growing Digital Gardens w/ Maggie Appleton (2021)](https://www.youtube.com/watch?t=141&v=Afq3Y2I2WqM)
- [em](https://github.com/cybersemics/em) - Beautiful, minimalistic note-taking app for personal sensemaking.
- [Notemarks](https://notemarks.app/) - Git based labeling app to manage notes, documents, and bookmarks. ([Code](https://github.com/notemarks/notemarks))
- [Method for maximizing the value of networked note-taking systems, using Obsidian (2021)](https://www.quora.com/q/tophersblog/The-connection-between-Zettelkasten-and-systems-theory-is-a-vital-one-to-consider-In-line-with-August-Bradleys-insigh)
- [Deconstructing digital gardens (2020)](https://vivqu.com/blog/2020/10/18/digital-gardens/)
- [Serenity Notes](https://www.serenity.re/en/notes) - E2E encrypted notes. ([HN](https://news.ycombinator.com/item?id=26567497)) ([Code](https://github.com/SerenityNotes/serenity-notes-clients)) ([Backend Code](https://github.com/SerenityNotes/serenity-notes-backend))
- [A Brief History & Ethos of the Digital Garden (2021)](https://maggieappleton.com/garden-history)
- [Thread on 2-way links (2021)](https://twitter.com/elzr/status/1378821500246065154)
- [Gatsby Garden](https://github.com/binnyva/gatsby-garden) - Digital Garden Theme for Gatsby. Gatsby Garden lets you create a static HTML version of your markdown notes.
- [Mem](https://get.mem.ai/) - Capture and access information from anywhere. Simple as Apple Notes — powered by a collaborative graph database. ([Intro](https://www.youtube.com/watch?v=hWvQUvWUCeU))
- [Building a Memex](https://github.com/adri/memex) - Search of my personal data.
- [Ask HN: How do you organize your knowledge? (2021)](https://news.ycombinator.com/item?id=26935838)
- [Reflect](https://reflect.app/) - Keep track of your notes, books, and meetings. ([GitHub](https://github.com/team-reflect))
- [KnowAssist](https://www.knowassist.com/) - Good knowledge management.
- [Gatsby Brain](https://github.com/aengusmcmillin/gatsby-theme-brain) - Use this theme to integrate Roam Research inspired bidirectional linking of notes into a network on your site.
- [Notea](https://cinwell.com/notea/) - Self-hosted note-taking app stored on S3. ([Code](https://github.com/QingWei-Li/notea))
- [Codex OS](https://twitter.com/codexeditor) - Creating the knowledge worker's OS.
- [Merlot](https://github.com/thesephist/merlot) - Web based Markdown writing app built with isomorphic Ink and Torus.
- [Hypernotes](https://hypernotes.zenkit.com/) - Platform for collaboration and project management.
- [Innos Note](https://innos.io/)
- [How NOT to take Smart Notes (2021)](https://www.reddit.com/r/ObsidianMD/comments/ng9pqg/how_not_to_take_smart_notes/)
- [rsms's memex](https://github.com/rsms/memex) - Software for managing my digital information, like tweets.
- [Notable](https://notable.app/) - Markdown-based note-taking app.
- [emanote](https://github.com/srid/emanote) - Spiritual successor to neuron, based on Ema. Create beautiful websites like wikis. ([Web](https://note.ema.srid.ca/))
- [Joplin](https://joplinapp.org/) - Open source note taking and to-do application with synchronisation capabilities. ([Code](https://github.com/laurent22/joplin))
- [Dedicated hardware for note taking](https://twitter.com/jessmartin/status/1365290998243409922)
- [Nota](https://nota.md/) - Pro writing app designed for local Markdown files. ([Reddit](https://www.reddit.com/r/macapps/comments/o0gu2j/nota_pro_writing_app_designed_for_local_markdown/)) ([Issues](https://github.com/notaapp/nota/issues))
- [Note Taking in 2021](http://blog.dornea.nu/2021/06/13/note-taking-in-2021/) ([HN](https://news.ycombinator.com/item?id=27513008))
- [Fuzzynote (fzn)](https://github.com/Sambigeara/fuzzynote) - Terminal-based, CRDT-backed, collaborative note-taking. ([HN](https://news.ycombinator.com/item?id=27606604))
- [My approach to taking notes in meetings (2021)](https://blog.witful.com/how-i-take-notes-in-meetings/) ([HN](https://news.ycombinator.com/item?id=27664419))
- [Incremental Note-Taking (2021)](https://thesephist.com/posts/inc/) ([HN](https://news.ycombinator.com/item?id=27667538)) ([Code](https://github.com/thesephist/inc))
- [How to turn notes into action (2021)](https://twitter.com/Julian/status/1418292810780790792)
- [NotePlan](https://noteplan.co/) - Your tasks, notes, and calendar. All linked in one place.
- [How I Make Drafts Work for Me (2021)](https://heydingus.net/blog/2021/7/how-i-make-drafts-work-for-me)
- [Awesome Digital Gardens](https://github.com/kyrose/awesome-digital-gardens)
- [Mura Notes](https://mura-notes.com/) - Anonymous, Lightweight and Shareable Notes and Todos. ([HN](https://news.ycombinator.com/item?id=27984366))
- [The Productivists Forum](https://www.theproductivists.club/) - Group of personal knowledge management (PKM) and productivity enthusiasts.
- [Ideaflow](https://www.ideaflow.io/) - Notebook that augments your intelligence.
- [Notabase](https://notabase.io/) - Personal knowledge base for networked thinking. ([Code](https://github.com/churichard/notabase))
- [The open calendar, task and note space is a mess (2021)](https://stevenvanbael.com/open-calendar-task-space-is-a-mess) ([HN](https://news.ycombinator.com/item?id=28358463))
- [The (re-)emergence of digital gardens (2021)](https://distroid.substack.com/p/distroid-issue-13-digital-gardens)
- [Ask HN: How do you take notes throughout your work day? (2021)](https://news.ycombinator.com/item?id=28508811)
- [The 5 Rs of Note Taking](https://aliabdaal.com/the-5-rs-of-note-taking/)
- [Noteship](https://noteship.com/) - Local-first notes app for Mac, with data longevity in mind. ([HN](https://news.ycombinator.com/item?id=28617596))
- [zk](https://github.com/mickael-menu/zk) - Plain text note-taking assistant.
- [Why you need a "WTF Notebook"](https://www.simplermachines.com/why-you-need-a-wtf-notebook/)
- [Taio](https://taio.app/) - Markdown editor and text actions for iPhone, iPad, and Mac. ([Twitter](https://twitter.com/TaioApp))
- [digital-garden.dev](https://digital-garden.dev/) - Next.js powered digital garden starter. ([Code](https://github.com/inadeqtfuturs/garden))
- [JINGO](https://github.com/claudioc/jingo) - Git based wiki engine written for node.js, with a decent design, a search capability and a good typography.
- [Wiki.js](https://js.wiki/) - Modern, lightweight and powerful wiki app built on NodeJS. ([Code](https://github.com/Requarks/wiki))
- [Gollum](https://github.com/gollum/gollum) - Simple, Git-powered wiki with a sweet API and local frontend. ([HN](https://news.ycombinator.com/item?id=28730764))
- [Meta](https://projectmeta.app/) - Note-taking tool for visual learning.
- [Ikke](https://github.com/laffra/Ikke) - Index and search your personal data quickly and privately.
- [Reflect Academy](https://reflect.academy/) - Learn how to take great notes. ([Tweet](https://twitter.com/maccaw/status/1452657502441648141))
- [Remnote](https://www.remnote.com/) - All-in-One Tool for Thinking and Learning. ([Twitter](https://twitter.com/rem_note))
- [The Science of Tools for Thought and Cognition Augmentation Software (2021)](https://moritz.digital/blog/cas) ([Video](https://www.youtube.com/watch?v=zVuaOECnUJE))
- [OpenMemex](https://github.com/austinvhuang/openmemex) - Open source, local-first knowledge platform.
- [mdzk](https://github.com/mdzk-rs/mdzk) - Plain text Zettelkasten system that is based on the mdBook API.
- [Good examples of wikis with custom software (2021)](https://twitter.com/LauraDeming/status/1457101797076586496?s=20)
- [Quartz](https://github.com/jackyzha0/quartz) - Host your own second brain and digital garden for free. ([Web](https://quartz.jzhao.xyz/))
- [How to build a second brain as a software developer (2021)](https://aseemthakar.com/how-to-build-a-second-brain-as-a-software-developer/) ([HN](https://news.ycombinator.com/item?id=29188418))
- [ff](https://github.com/ff-notes/ff) - Distributed note taker and task manager.
- [Learn In Public](https://www.swyx.io/learn-in-public/)
- [unigraph-dev](https://github.com/unigraph-dev/unigraph-dev) - Local-first and universal graph database, knowledge engine, and workspace for your life.
- [ummm](https://www.ummm.co/) - Organize your thoughts. ([HN](https://news.ycombinator.com/item?id=29304438))
- [Keep a Knowledge Log](https://bruno-oliveira.github.io/techblog/Keep-a-log/) ([HN](https://news.ycombinator.com/item?id=29436562))
- [Early conceptarium workflows](https://paulbricman.com/reflections/early-conceptarium-workflows)
- [Mi](https://github.com/elchead/misou) - Personal search engine.
- [/tap](https://www.tatatap.com/) - Powerful and customizable note-taking system. ([HN](https://news.ycombinator.com/item?id=29663233))
- [Zetten](https://github.com/hrvolapeter/zetten) - Note taking app - Firebase + SwiftUI.
- [Ask HN: Why the Obsession with Note Taking? (2022)](https://news.ycombinator.com/item?id=30098219)
- [Gains I'm seeing from my second brain tool](https://joeldare.com/gains-im-seeing-from-my-second-brain-tool.html) ([HN](https://news.ycombinator.com/item?id=30151963))
- [Workflows in my note app Skrift (2022)](https://harry.vangberg.name/posts/2021-12-30-note-workflows-in-skrift/) ([HN](https://news.ycombinator.com/item?id=30167563))
- [Ter](https://github.com/kkga/ter) - Tiny wiki-style site builder with Zettelkasten flavor.
- [Lurnby](https://github.com/Roznoshchik/Lurnby) - Tool for active reading and personal knowledge management. ([Web](https://www.lurnby.com/)) ([HN](https://news.ycombinator.com/item?id=30299847))
- [Windi](https://windi.app/) - Knowledge management and sharing platform based on short notes. ([HN](https://news.ycombinator.com/item?id=30302690))
- [Awesome Knowledge Management](https://github.com/brettkromkamp/awesome-knowledge-management)
- [AI note garden: link suggestions (2022)](https://smus.com/ai-note-garden-linker/)
- [My Notebook System](http://ratfactor.com/notes) ([HN](https://news.ycombinator.com/item?id=30395238))
- [Write plain text files](https://sive.rs/plaintext) ([HN](https://news.ycombinator.com/item?id=30521545))
- [Plaintext Productivity](https://plaintext-productivity.net/) ([HN](https://news.ycombinator.com/item?id=30745524))
- [Programmable Notes](https://maggieappleton.com/programmatic-notes) ([HN](https://news.ycombinator.com/item?id=30814964))
- [zk](https://github.com/sirupsen/zk) - Zettelkasten on the command-line.
- [xit](https://xit.jotaen.net/) - Plain-text file format for todos and check lists. ([Spec](https://github.com/jotaen/xit)) ([HN](https://news.ycombinator.com/item?id=30879327))
- [chronicle-etl](https://github.com/chronicle-app/chronicle-etl) - CLI toolkit for extracting and working with your digital history.
- [25+ years of personal knowledge management (2022)](https://dsebastien.net/blog/2022-04-03-25-years-of-personal-knowledge-management) ([HN](https://news.ycombinator.com/item?id=30903940))
- [Heptabase](https://heptabase.com/) - Note-taking tool for visual learning.
- [My Org Roam Notes Workflow (2021)](https://hugocisneros.com/blog/my-org-roam-notes-workflow/) ([HN](https://news.ycombinator.com/item?id=30965343))
- [todo.txt format](https://github.com/todotxt/todo.txt) - Complete primer on the whys and hows of todo.txt.
- [Harika](https://github.com/quolpr/harika) - Offline-first, performance-focused note taking app for organizing your knowledge database.
- [Digital Garden theme for Hugo](https://github.com/apvarun/digital-garden-hugo-theme)
- [Lepiter: Knowledge Mgmt and Multi-Lang Notebooks and Moldable Dev](https://lepiter.io/feenk/introducing-lepiter--knowledge-management--e2p6apqsz5npq7m4xte0kkywn/)
- [Zettelkasten, Linking Your Thinking, and Nick Milo's Search for Ground (2022)](https://writing.bobdoto.computer/zettelkasten-linking-your-thinking-and-nick-milos-search-for-ground/) ([HN](https://news.ycombinator.com/item?id=31507132))
- [GitNoter](https://github.com/git-noter/gitnoter) - Open source, markdown-based, self-hosted note taking web app. ([Web](https://gitnoter.com/login)) ([HN](https://news.ycombinator.com/item?id=31580045))
- [Ask HN: Why are there no good note taking apps (2022)](https://news.ycombinator.com/item?id=31645045)
- [SiYuan](https://b3log.org/siyuan/en/) - Local-first personal knowledge management system, support fine-grained block-level reference and Markdown WYSIWYG. ([Code](https://github.com/siyuan-note/siyuan))
- [Building a Memex (with Ruby) by Andrew Louis (2018)](https://www.youtube.com/watch?v=DFWxvQn4cf8)
- [Zettelkasten 101 Series: What Is a Fleeting Note?](https://writing.bobdoto.computer/what-is-a-fleeting-note/) ([HN](https://news.ycombinator.com/item?id=32225169))
- [Ask HN: What is a sustainable methodology for taking notes of your learning? (2022)](https://news.ycombinator.com/item?id=32211734)
- [Ask HN: What's your note-taking methodology? (2022)](https://news.ycombinator.com/item?id=32252957)
- [Affine](https://affine.pro/) - Next-Gen Knowledge Base to Replace Notion & Miro. ([Code](https://github.com/toeverything/AFFiNE))
- [Curio](https://www.zengobi.com/curio/) - Note Taking, Mind Mapping, Brainstorming. ([HN](https://news.ycombinator.com/item?id=32432724))
- [I wish I could organize my thoughts (2022)](https://drewdevault.com/2022/08/10/Organizing-my-thoughts.html) ([HN](https://news.ycombinator.com/item?id=32427588))

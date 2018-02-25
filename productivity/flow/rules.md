# Rules
Some simple rules to follow for myself doing things.

## Life
- Simplify as much as possible. Never repeat myself.
- Remember the [lessons I learned](../../sharing/lessons-learned.md). So as to not repeat the mistakes I have made.
- Always learn new things.
- Exercise daily.
- Build ideas.
- Follow good nutrition. No sugar, drink water.
- Share everything I know and made.
- No food/item wastage.
- Be mindful of my actions and thoughts.
- Minimise distractions.
- Follow these rules.

## Routine
- Write. (1+ article / month)
- Read a book (1+ hour / day)

## Nutrition
- No processed sugar. No food with extra sugar in it.
- No drinks other than water, coffee and tea. No added sugar in the drinks.
- Always have water/tea next to me when working on something. Fill up the water/tea if it empties.
- No meat or dairy products. Health and [ethical reasons](../../life/compassion.md).
- Focus on Whole Food, Plant-Based diet.

## Writing
- Use proper punctuation and capitalise things appropriately.
- When adding __:__, don't add space beforehand.
- When making tasks (2Do, Trello, GitHub issues), use imperative tense.
- Remove unnecessary and _filler_ words.
	- If something can be said in less words, say it in less words.
	- Always try to be clear and to the point.
- Both in code and prose. Say more with less and never repeat myself.
- Minimise using word _here_.
	- When wanting to reference a certain URL, I don't need to mention _and you can see it here_. I can most likely just change the thing I want to show and wrap that in URL. It looks cleaner this way and saves space.

## Markdown
- Don't leave unnecessary whitespace between headings and text.
- Don't need to add __:__ to the end of text when adding image below as it is implicit.
- Only use HTML to adjust the image if the original size is not what I need. Don't adjust images blindly and check how the final image looks rendered compared to the text.
- Use __Bold__ for strong emphasis. Use _Italic_ for soft emphasis. And `code quotes` for actual code, variables, path names and quoting inline (in same sentence).

## Wiki
- Describe links on the same line with __-__ separator (if needs more description, then sub dash).
- Don't repeat myself or anyone else. If something has already been said either by me or someone else, link to it (especially in this wiki).
- Interlink entries between each other within wiki where necessary. Don't over do it.
- Try to use other services for logging and other things where possible. Like in the case of [Letterboxd](https://letterboxd.com/NikitaVoloboev/) and its Lists feature. It's better to make lists there and link to the lists from wiki. Find the right tool for the job. The wiki should act as a glue between things.
- Keep all files lowercased.
- If I am watching a talk, it is best not to jot down notes from the talk itself but capture ideas and note down these ideas in the wiki, sometimes referencing the talk where the idea or insight was taken from for context.
- Think about structuring the content and entries as I add new entries. Should a new category be added for the topic (thing) and under what category should the entry be placed.
- If I move a file anywhere in the wiki, make sure to search for the path to the file in the entire wiki and change the references to the previous file with new one.

## GitHub
- Add appropriate topics to each repo.
- Make many experiments and share them.
- Add credits to repos and acknowledge work of people I am using at end of repos. Under __# Credits__.
- Use relative links instead of hard wired links. For example in GitHub you can use `../../issues/` to reference issues of current repo from readme.

## GitHub README
- Use HTML for rendering images if I need control on their size and position.
- Add appropriate _quote_ to briefly describe the repo (often is the same as description but with links). Don't repeat what is said in the quote in the description underneath the quote.
- If the README is long enough, add a __Contents__ reference table. Can omit Contents heading.
- Don't add anything that is not needed and focus on the content.
- Add a way to visually show the workflow/library in action where appropriate.
- Don't link to Imgur images as links, render them in the GitHub README itself.
- In screenshots of Alfred prompts and other prompts, show the cursor.
- Add __Related__ section to most repos linking to similar projects either of my own or of other. Not all projects need this.
- Don't add Thank You clause in _idea_ stage projects.
- Don't add _Say Thanks_ badge on curated lists.

## Adding resources (links)
- When adding resources and bookmarks. Priorotise adding them to LA curated lists first and LA itself (not Trello). The goal of the wiki is to mostly contain my own notes and references.
	- The resources should be moderated and extended by the community through clear guidelines.

## Git
- Always initialise new projects with Git.
- Do atomic commits and write [proper appropriate commit messages](https://chris.beams.io/posts/git-commit/).
- Start commits uppercased and prefix them with an action (Add/Remove/Fix/Update). Use [imperative tense](https://pbs.twimg.com/media/DBPQbTrXkAA4v-H.jpg).
	- On larger projects, prefix commits by scope for easier search and index.
- Don't use emojis in commit messages.
- Commit changes and push before making a release. This way the commits come bundled together with release as change log.
- When making changes to other projects, respect their Git workflow and commit style.

## Dotfiles
- Comment all my aliases briefly saying what they do. Comment all zsh functions. Be mindful that I can then search through these functions and aliases later.

## Code
- Write code to be read by other people. Prefer to be explicit over implicit.
- Delimit files that have spaces with __-__ instead of __\___.
- Write a small comment on top of the file to describe the purpose of the file. Where appropriate.
- Use switch cases over long ifs.
- Function name should describe what it does, not when it's run. Good function names summarize what happens inside.

## Commenting
- In comments, don't write __Will ..__. Just write what it will do. Remove unnecessary (filler) words.
- Uppercase comments for consistency.
- Use imperative tense. Instead of _# Outputs_, use _# Output_.
- Only add one space after the code for inline comments. Don't space them out.
- Comment over my code (be mindful of not repeating things, sometimes code speaks for itself).

## Keyboard Maestro
- All macro names are lowercased. Unless it is necessary.
- When doing macros to _focus_ on something, name the macro like _sidebar focus_ rather than _focus on sidebar_.

## Karabiner
- Try to make semantic mappings and chunk mappings by theme or context.
- Prefix variable names by action (i.e. EDIT/OPEN/ALFRED). Each variable name should be descriptive.

## Alfred
- Keyword triggers should be one string (no spaces). Try to make keyword be semantically appropriate to workflow and easy to memorise and type.
- Name external triggers with spaces between words (lowercased).
- Release all my workflows (that have no sensitive information) either in [small workflows](https://github.com/nikitavoloboev/small-workflows) or standalone if they contain quite a bit of code.
- Name all modifier triggers (uppercased).

## MindNode
- Use one tree structure for organisation especially when sharing maps (makes the map much easier to follow).

## Social Networks
- Only add/follow people I am really interested in or know.

## Sharing files
- Temporary shares with [Transfer](https://transfer.sh).
- More permanent shares are shared with Dropbox. Don't delete items from _shares_ folder unless I delete the file somewhere myself.
- If the files are stored on GitHub somewhere, I can use [Rawgit](http://rawgit.com/) to get a downloadable link of the raw file.

## Sharing my work
- Share the link to the project on the appropriate website/forum/subreddit.

## Releasing workflows with OneUpdater
- Don't forget to update version number and then also commit the workflow so info.plist on GitHub has the latest version.

## SnippetsLab
- Lowercase and prefix all snippets. Notes can be uppercased.
- Upload all non personal snippets as Gists.
- Create smart group for each new prefix I add.
- Add correct language syntax to every snippet.
- Prefer to use comments in the code itself rather than notes (easier to follow).
- Map command line tools in the curated list. Only add commands where I do want to add some optional commands as snippet to SnippetsLab.
- Attach the URL to snippets as note as a first thing always if URL is needed. Add other notes underneath the URL.

## Planning
- Plan the next day fully in advance in the calendar. Follow through with the plan but adapt where necessary.

## 2Do
- Only contain things I really need to get done in _Today_.
- All 2Do tasks need to be actionable.
- Work on tasks based on priority (starred completed first).
- Don't include context in the task title since I can do add it as a tag instead.

## Ideas
- If the idea is not private. Share it in Trello (hopefully soon in Crafting Ideas).

## Capitalising things
- Capitalise everything where appropriate (Sentences / Headings / Alfred workflows / GitHub descriptions / GitHub issues and PRs / Comments in SnippetsLab / Main dashed points / Tweets / ..).

## Ship app
- Prioritise repos in the sidebar based on priority/urgency.

## [Thinking](../../research/solving-problems.md) map
- Only have actual things I am struggling to work through there that need slicing (breaking down of smaller tasks).
- Keep the list to maximum of 5 main nodes in there. If there is more, break them down and move them to either 2Do or GitHub issues and only keep things I am working on.

## Reading books
- Make notes as I read books.
- Review each book I read on Goodreads and add it to an appropriate category. Add the review to [books](../books/books.md) after.

## Productivity
- Work to complete the task I have at hand.
- Set deadlines for all things. Adjust plan according to the deadlines.
- Deadlines set are of utmost importance. Do my best to complete the deadlines I have set for today before venturing in trying to do anything else. Learn how to deal and set appropriate deadlines with time.
- Limit distractions. Strive to increase signal to noise ratio.
- Always turn [Focus](https://heyfocus.com) when starting to work and make things and fully focus on the task/s at hand. Adjust the focus list to block even more distractions and tune my workflow to only do the task/s I have.
- Live by my calendar. Assign tasks and blocks of time in advance and work through each of the blocks. Adjust when necessary.
- Always have water next to me.
- Work away from home ideally.

## Anki
- Don't make a card of a fact or something that you can easily look up.

## Product Hunt
- Add products I like to an appropriate collection.

## Vim
- Comment over all mappings in vimrc itself.
- Only add things that are essential to me and what I actually use.

## Spotify
- Add all songs I like to my [Likes](https://open.spotify.com/user/nikitavoloboev/playlist/0ERn0U4qZIKC8Dy7RrMMsn?si=s7begtjFTM-vaAp0ZplMaQ) playlist.
	- Add all songs I _really_ like to [Favourites](spotify:user:nikitavoloboev:playlist:7j0M4e0nxFtsrLREfcj2qk) playlist.
	- Add songs to other playlists accordingly.

## Focus
- Be very aggressive in cutting any kind of distractions when working.
- Look at where I spend my time in Timing and cut those out with Focus and scripts.

## Licensing
- License my own code under MIT.
- License curated lists and other things under [CC0](https://creativecommons.org/publicdomain/zero/1.0/).

## Prefixing
- Prefer to have the prefix reference the meaning in some way for easier mental mapping of prefix to meaning.
- Name and prefix things with focus on clarity. The semantics of prefixes should ideally be obvious. And naming should be appropriate.

## Sleep
- Prefer to sleep and wake up at the same time.
	- 5:30AM - Wake up.
	- 10:00PM - Sleep.
- Sleep for ~ 7 hours each day.

## Timing
- Minimise _Distracting_ time and aim at __< 1 hour Distracted time__ a day. Ideally less. Block everything that is _Distracting_ when working.
- Mark all time past in Timing. Adjust filters and projects to what kind of things I am working on.
- Only mark non distracting blocks with tasks. Those blocks are ones that matter.

## Day One
- Mark all memorable moments and experiences in Day One.
- Mark memorable moments throughout the day and put them in __Photo a Day__ and __Travel__ journals.

## Tools
- Only use what is necessary. Experiment with things but never keep things you are not actually using. That mostly includes plugins and modifications.
	 - Try to establish a good mental model of what you need and what you use as you apprach modifying your workflow and tools.

## Images
- Keep screenshots and interesting images saved in Pixave.

## Knowledge indexing
- Index knowledge that is actionable in case I might need to reference it later.

## macOS
- Only run apps I am actually using. Quit everything else.
- Delete apps I am not using and have no need over. Only use what I need or may need.
- Have full control over what is going on my system to a reasonable degree. Don't use/install things I do not actually need.

## Email
- Do emails in burst. Bring my inbox to 0 every day. Action on each email.
- Archive emails if I want them to be searchable later. Delete emails otherwise.
- Star emails that are important and I may need to reference late.
- Sort newsletters automatically into groups for later viewings.
- Only show relevant emails to me in Inbox.
- Label all emails accordingly (if it is appropriate).

## Looking back
- Don't say _Want to_. Instead say _Will_ in _Changes_. This applies more generally too, try be more proactive with anything I do. Intention counts.

## Finance
- Track every purchase made with YNAB. Putting it in its respective category.
- Adhere to my budget that I have set for each of the categories. If there is no budget left for eating out. Don't eat out.

## Videos
- Watch talks and tutorials at 1.25+ speed. Unless it hurts coprehension than watch it at slower speeds.

## Information
- Don't save unnecessary information. This mostly concerns tracking data and the like. More data is not always better. Quality and purposeful data is.
- Make the data actionable in some way. Either as a way that you can query it later and read or look back upon in the future (as is case with Day One) or as a way you can immediately take action on in form of analysis of it.
	- Knowing what goals you have and the reasons why you are collecting or using certain kind of data is important and saves a lot of time in the future when you want to _minimise_ and _simplify_ things in life.

## Trello
- No need to create _Doing now_ and _Do next_ blocks. I can use _watching_ feature for that and seperate cards instead by topic. Use _tags_ to give priority to tasks. (In my case _Important_, _Interesting_ and _Next_ are always present tags in most boards).
- Share boards publically that contain no sensetive or private information.
- For boards like _Learning_ and _Books_, no need to add links to cards as the links to these things should already be added in LA curated lists which are easily queriable.
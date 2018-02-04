# Rules
Some simple rules to follow for myself doing things.

## Nutrition
- No processed sugar. Includes sugary carbonated beverages. Drink water!
- No food with extra sugar in it.
- No meat or dairy products. Health and [ethical reasons](../../life/compassion.md).
- No drinks other than water, coffee and tea. No added sugar in the drinks.

## Wiki
- Describe links on the same line with `-` separator (if needs more description, then sub dash).
- Don't repeat myself or anyone else. If something has already been said either by me or someone else, link to it (especially in this wiki).
- Interlink entries between each other within wiki where necessary. Don't over do it.
- Try to use other services for logging and other things where possible. Like in the case of [Letterboxd](https://letterboxd.com/NikitaVoloboev/) and its Lists feature. It's better to make lists there and link to the lists from wiki. Find the right tool for the job. The wiki should act as a glue between things.
- Keep all files lowercased.
- If I am watching a talk, it is best not to jot down notes from the talk itself but capture ideas and note down these ideas in the wiki, sometimes referencing the talk where the idea or insight was taken from for context.

## Adding resources (links)
- When adding resources and bookmarks. Priorotise adding them to LA curated lists first and LA itself (not Trello). The goal of the wiki is to mostly contain my own notes and references.
	- The resources should be moderated and extended by the community through clear guidelines.

## Git
- Always initialise new projects with Git.
- Do atomic commits and write [proper appropriate commit messages](https://chris.beams.io/posts/git-commit/).
- Start commits uppercased and prefix them with an action (Add/Remove/Fix/Update). Use imperative tense.
- Don't use emojis in commit messages.

## Dotfiles
- Comment all my aliases briefly saying what they do. Comment all zsh functions. Be mindful that I can then search through these functions and aliases later.

## Code
- Write code to be read by other people. Prefer to be explicit over implicit.
- Delimit files that have spaces with `-` instead of `_`.

## Commenting
- In comments, don't write `Will ..`. Just write what it will do. Remove unnecessary (filler) words.
- Uppercase comments for consistency.
- Use imperative tense. Instead of `# Outputs`, use `# Output`.
- Only add one space after the code for inline comments. Don't space them out.
- Comment over my code (be mindful of not repeating things, sometimes code speaks for itself).

## Keyboard Maestro
- All macro names are lowercased. Unless it is necessary.

## Karabiner
- Document any changes I make in Karabiner map instantly in the map (update where necessary).
- Try to make semantic mappings and chunk mappings by theme or context.
- Prefix variable names by action (i.e. EDIT/OPEN/ALFRED).

## Markdown
- Don't leave unnecessary whitespace between headings and text.
- Don't need to add `:` to the end of text when adding image below as it is implicit.
- Only use HTML to adjust the image if the original size is not what I need. Don't adjust images blindly and check how the final image looks rendered compared to the text.

## GitHub
- Add appropriate topics to each repo.
- Make many experiments and share them.
- Add credits to repos and acknowledge work of people I am using at end of repos (`Credits`).
- Use relative links instead of hard wired links. For example in GitHub you can use `../../issues/` to reference issues of current repo from readme.
- Lowercase license, contributing and readme files. There is no need for them to be uppercased.

## GitHub readme
- Use HTML for rendering images if I need control on their size and position.
- Add appropriate _quote_ to briefly describe the repo (often is the same as description but with links). Don't repeat what is said in the quote in the description underneath the quote.
- If the readme is long enough, add a `Contents` table.
- Don't add anything that is not needed and focus on the content.
- Add a way to visually show the workflow/library in action where appropriate.
- Don't link to Imgur images as links, render them in the GitHub readme itself.
- In screenshots of Alfred prompts and other prompts, show the cursor.
- Add `Related` section to most repos linking to similar projects either of my own or of other. Not all projects need this.
- Don't add Thank You clause in `idea` stage projects.
- If the project has many ways to contribute, create a `contributing.md` file and link to it.

## Alfred
- Keyword triggers should be one string (no spaces). Try to make keyword be semantically appropriate to workflow and easy to memorise and type.
- Name external triggers with spaces between words (lowercased).
- Release all my workflows (that have no sensitive information) either in [small workflows](https://github.com/nikitavoloboev/small-workflows) or standalone if they contain quite a bit of code.
- Name all modifier triggers (uppercased).

## MindNode
- Use one tree structure for organisation especially when sharing maps (makes the map much easier to follow).

## Writing
- Use proper punctuation and capitalise things appropriately.
- When adding `:`, don't add space beforehand.
- Remove unnecessary and `filler` words.
	- If something can be said in less words, say it in less words.
	- Always try to be clear and to the point.
- Both in code and prose. Say more with less and never repeat myself.
- Minimise using word `here`.
	- When wanting to reference a certain URL, I don't need to mention `and you can see it here`. I can most likely just change the thing I want to show and wrap that in URL. It looks cleaner this way and saves space.

## Social Networks
- Only add/follow people I am really interested in or know.

## Sharing files
- Temporary shares with [Transfer](https://transfer.sh).
- More permanent shares are shared with Dropbox. Don't delete items from `shares` folder unless I delete the file somewhere myself.
- If the files are stored on GitHub somewhere, I can use [Rawgit](http://rawgit.com/) to get a downloadable link of the raw file.

## Releasing workflows with OneUpdater
- Don't forget to update version number and then also commit the workflow so info.plist on GitHub has the latest version.

## SnippetsLab
- Upload all non personal snippets as Gists.
- Create smart group for each new prefix I add.
- Add correct language syntax to every snippet.
- Prefer to use comments in the code itself rather than notes (easier to follow).
- All snippet titles are lowercased. Notes are uppercased.
- Map command line tools in the curated list. Only add commands where I do want to add some optional commands as snippet to SnippetsLab.

## Planning
- Plan the next day fully in advance in the calendar. Follow through with the plan but adapt where necessary.

## 2Do
- Only contain things I really need to get done in `Today`.
- All 2Do tasks need to be actionable.
- Work on tasks based on priority (starred completed first).

## Ideas
- If the idea is not private. Share it in Trello (hopefully soon in Crafting Ideas).

## Capitalising things
- Capitalise everything where appropriate (Sentences / Headings / Alfred workflows / GitHub descriptions / GitHub issues and PRs / Comments in SnippetsLab / Main dashed points / Tweets / ..).

## Ship app
- Prioritise repos in the sidebar based on priority/urgency.

## Stack Overflow CV
- Only add repositories that are meaningful to the list.

## Reading books
- Make notes as I read books.
- Review each book I read on Goodreads and add it to an appropriate category. Add the review to [books](../books/books.md) after.

## Productivity
- Work to complete the task I have at hand.
- Set deadlines for all things. Adjust plan according to the deadlines.
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

## File naming
- Prefer to use `-` instead of `_` in naming script files. And lowercased.

## Prefixing
- Prefer to have the prefix reference the meaning in some way for easier mental mapping of prefix to meaning.
- Name and prefix things with focus on clarity. The semantics of prefixes should ideally be obvious. And naming should be appropriate.

## Sleep
- Prefer to sleep and wake up at the same time.
	- 5:30AM - Wake up.
	- 10:00PM - Sleep.
- Sleep for ~ 7 hours each day.

## Timing
- Minimise `Distracting` time and aim at `< 1 hour` `Distracted` time a day. Ideally less. Block everything that is `Distracting` when working.
- Mark all time past in Timing. Adjust filters and projects to what kind of things I am working on.
- Only mark non distracting blocks with tasks. Those blocks are ones that matter.

## Day One
- Mark all memorable moments and experiences in Day One.
- Mark memorable moments throughout the day and put them in `Photo a Day` and `Travel` journals.

## Tools
- Only use what is necessary. Experiment with things but never keep things you are not actually using. That mostly includes plugins and modifications.
	 - Try to establish a good mental model of what you need and what you use as you apprach modifying your workflow and tools.

## Images
- Keep screenshots and interesting images saved in Pixave.

## Knowledge indexing
- Index knowledge that is actionable in case I might need to reference it later.

## Email
- Do emails in burst. Bring my inbox to 0 every day. Action on each email.
- Archive emails if I want them to be searchable later. Delete emails otherwise.
- Star emails that are important and I may need to reference late.
- Sort newsletters automatically into groups for later viewings.
- Only show relevant emails to me in Inbox.
- Label all emails accordingly (if it is appropriate).

## Finance
- Track every purchase made with YNAB. Putting it in its respective category.
- Adhere to my budget that I have set for each of the categories. If there is no budget left for eating out. Don't eat out.

## SnippetsLab
- Attach the URL to snippets as note as a first thing always if URL is needed. Add other notes underneath the URL.

## Information
- Don't save unnecessary information. This mostly concerns tracking data and the like. More data is not always better. Quality and purposeful data is.
- Make the data actionable in some way. Either as a way that you can query it later and read or look back upon in the future (as is case with Day One) or as a way you can immediately take action on in form of analysis of it.
	- Knowing what goals you have and the reasons why you are collecting or using certain kind of data is important and saves a lot of time in the future when you want to `minimise` and `simplify` things in life.

## Trello
- No need to create `Doing now` and `Do next` blocks. I can use `watching` feature for that and seperate cards instead by topic. Use `tags` to give priority to tasks. (In my case `Important`, `Interesting` and `Next` are always present tags in most boards).
- Share boards publically that contain no sensetive or private information.
- For boards like `Learning` and `Books`, no need to add links to cards as the links to these things should already be added in LA curated lists which are easily queriable.
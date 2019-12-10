# Git

I love Git and version control. And I use version control over any project I do. I follow a [series of rules](../focusing/rules.md#git) when dealing with Git.

## Notes

- Commit as often as you can. Ideally after each micro-iteration, when something new is working.
  - This way, at the end of the day you can just rebase the whole branch and squash all of the micro-commits in a whole commit implementing the whole new features.
- Good git workflow to make changes to new projects: clone, fork (`hub fork`), 'git checkout -b my-feature', work, commit, 'git push -u nikitavoloboev my-feature', work, commit, 'git push'.
- If you’re doing things right, there’s only two kinds of branches anyways, master and feature branches. Feature branches can be squashed and rebased off master (minimizing the issue of merge conflicts and making for easier management of the commit history) and merged to master from there without requiring further conflict resolution. ([Trunk-Based Development](https://paulhammant.com/2013/04/05/what-is-trunk-based-development/))
- A Git branch is just a pointer to a commit. Git commits, however, also contain the hash of the parent commit(s), so by referring to that commit you also refer too all ancestors.
- Squash + rebase (for feature branches) are good for PRs. No one cares that it took you 20 tries to get the feature right, what matters is what went into the pull request, which is usually one commit.
- [To me the beauty of git stems from the fact that it is an implementation of a functional data structure. It‘s a tree index that is read-only, and updating it involves creating a complete copy of the tree and giving it a new name. Then the only challenge is to make that copy as cheap as possible - for which the tree lends itself, as only the nodes on the path to the root need to get updated. As a result, you get lock-free transactions (branches) and minimal overhead. It is so beautiful in fact that when I think about systems that need to maintain long-running state in concurrent environments, my first reaction is ”split up the state into files, and maintain it through git“.](https://news.ycombinator.com/item?id=21418033)

## Links

- [Git from the Bottom Up](https://jwiegley.github.io/git-from-the-bottom-up/)
- [Flight rules for git](https://github.com/k88hudson/git-flight-rules#readme)
- [Great Git workflow instructions](https://github.com/rvolosatovs/turtlitto/blob/master/CONTRIBUTING.md#readme)
- [GIT Conventions](https://medium.com/@tjholowaychuk/git-conventions-a940ee20862d)
- [Learn Git branching](https://learngitbranching.js.org/)
- [Gitbase](https://github.com/src-d/gitbase) - SQL interface to Git repositories.
- [Gitea](https://github.com/go-gitea/gitea) - Easiest, fastest, and most painless way of setting up a self-hosted Git service.
- [How to Write a Git Commit Message](https://chris.beams.io/posts/git-commit/)
- [How I Use Git](https://hugogiraudel.com/2018/02/17/how-i-use-git/)
- [clog-cli](https://github.com/clog-tool/clog-cli#readme) - Generate beautiful changelogs from your Git commit history.
- [git rebase for fame and power](http://www.charlesetc.com/rebase-for-fame/)
- [gitbatch](https://github.com/isacikgoz/gitbatch) - Manage your git repositories in one place.
- [Conventional Commits](https://www.conventionalcommits.org/en/v1.0.0-beta.2/) - Specification for adding human and machine readable meaning to commit messages.
- [git absorb](https://github.com/tummychow/git-absorb) - Git commit --fixup, but automatic.
- [ghq](https://github.com/motemen/ghq) - Manage remote repository clones.
- [git-flow](https://github.com/nvie/gitflow) - Collection of Git extensions to provide high-level repository operations for Vincent Driessen's [branching model](http://nvie.com/git-model).
- [gitin](https://github.com/isacikgoz/gitin) - Commit/branch/status explorer for git.
- [Lab](https://github.com/zaquestion/lab) - Wraps Git or Hub, making it simple to clone, fork, and interact with repositories on GitLab.
- [Tips for a disciplined git workflow (2019)](https://drewdevault.com/2019/02/25/Using-git-with-discipline.html) ([Reddit](https://www.reddit.com/r/programming/comments/bfbar2/tips_for_a_disciplined_git_workflow/))
- [Git Town](https://github.com/Originate/git-town) - Generic, high-level Git workflow support.
- [libgit2](https://libgit2.org/) - Portable, pure C implementation of the Git core methods provided as a re-entrant linkable library.
- [git-delete-squashed](https://github.com/not-an-aardvark/git-delete-squashed) - Delete branches that have been squashed and merged into master.
- [rebase-editor](https://github.com/sjurba/rebase-editor) - Simple terminal based sequence editor for git interactive rebase.
- [git-standup](https://github.com/kamranahmedse/git-standup) - Recall what you did on the last working day. Psst! or be nosy and find what someone else in your team did.
- [git-rs](https://github.com/chrisdickinson/git-rs) - Implementing git in rust for fun and education.
- [Commit messages guide](https://github.com/RomuloOliveira/commit-messages-guide#readme)
- [The Many Benefits of Using a Monorepo (2019)](https://pspdfkit.com/blog/2019/benefits-of-a-monorepo/) ([HN](https://news.ycombinator.com/item?id=19795482))
- [Lefthook](https://github.com/Arkweid/lefthook) - Git hooks manager.
- [git rebase in depth](https://git-rebase.io/) ([HN](https://news.ycombinator.com/item?id=19877811))
- [Highlights from Git 2.22 (2019)](https://github.blog/2019-06-07-highlights-from-git-2-22/)
- [What is a fork, really, and how GitHub changed its meaning (2019)](https://drewdevault.com/2019/05/24/What-is-a-fork.html) ([Lobsters](https://lobste.rs/s/txx8vu/what_is_fork_really_how_github_changed_its)) ([HN](https://news.ycombinator.com/item?id=20001570))
- [More productive Git](https://increment.com/open-source/more-productive-git/) ([HN](https://news.ycombinator.com/item?id=20004224))
- [List of commonly used Git commands](https://github.com/joshnh/Git-Commands#readme)
- [Hercules](https://github.com/src-d/hercules) - Fast, insightful and highly customizable Git history analysis.
- [Git Tower Guides](https://www.git-tower.com/learn/)
- [git-stashout](https://github.com/aviaviavi/git-stashout) - Custom git checkout command to automatically manage a per-branch stash.
- [git-bug](https://github.com/MichaelMure/git-bug) - Distributed, offline-first bug tracker embedded in git.
- [Git Standards (2018)](https://blog.carlmjohnson.net/post/2018/git-gud/)
- [BFG Repo-Cleaner](https://github.com/rtyley/bfg-repo-cleaner) - Removes large or troublesome blobs like git-filter-branch does, but faster.
- [gitstatus](https://github.com/romkatv/gitstatus) - 10x faster implementation of `git status` command.
- [git-revise](https://github.com/mystor/git-revise) - Handy tool for doing efficient in-memory commit rebases & fixups.
- [git-secret](https://github.com/sobolevn/git-secret) - Bash-tool to store your private data inside a git repository.
- [pre-commit](https://github.com/pre-commit/pre-commit) - Framework for managing and maintaining multi-language pre-commit hooks.
- [git-o-matic](https://github.com/muesli/gitomatic) - Tool to monitor git repositories and automatically pull & push changes.
- [GitGraph.js](https://github.com/nicoespeon/gitgraph.js) - JavaScript library to draw pretty git graphs in the browser.
- [Git First-Parent-- Have your messy history and eat it too (2018)](http://www.davidchudzicki.com/posts/first-parent/)
- [Git Alias](https://github.com/GitAlias/gitalias) - Git alias commands for faster easier version control.
- [Git Interactive Rebase Tool](https://github.com/MitMaro/git-interactive-rebase-tool) - Native cross platform full feature terminal based sequence editor for git interactive rebase. Written in Rust using ncurses.
- [Apache Arrow Git Tips](https://andygrove.io/apache_arrow_git_tips/)
- [GitSheet](https://gitsheet.wtf/) - Dead simple git cheatsheet.
- [Git Commit naming](https://pbs.twimg.com/media/DBPQbTrXkAA4v-H.jpg)
- [HN: My Favourite Git Commit (2019)](https://news.ycombinator.com/item?id=21289827)
- [Onefetch](https://github.com/o2sh/onefetch) - Git repository summary on your terminal.
- [Delta](https://github.com/dandavison/delta) - Syntax-highlighting pager for git.
- [Building Git](https://shop.jcoglan.com/building-git/) - Deep dive into the internals of the Git version control system.
- [gitrs](https://github.com/haltode/gitrs) - Re-implementation of the git version control system in Rust.
- [Sourcehut's year in alpha (2019)](https://sourcehut.org/blog/2019-11-15-sourcehut-1-year-alpha/) ([HN](https://news.ycombinator.com/item?id=21545145))
- [Awesome Monorepo](https://github.com/korfuri/awesome-monorepo#readme) - Curated list of awesome Monorepo tools, software and architectures.
- [rug](https://github.com/samrat/rug) - Stripped-down version of Git, implemented in Rust.
- [Git from Beginner to Advanced](https://www.madebymike.com.au/writing/how-to-git/)
- [GitUp](https://gitup.co/) - Git client with nice interface.

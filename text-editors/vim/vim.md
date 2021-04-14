# [Vim](https://www.vim.org/)

The best text editing experience you will ever get. It is worth learning it as you will start thinking about text editing differently.

I use some kind of vim bindings in any app I use and if there is a vim plugin for the app, I often use it. I use vim editor often when connecting to remote sessions and when I need a console editor.

I use [Neovim](https://github.com/neovim/neovim) instead of vim and my vimrc can be found [here](https://github.com/nikitavoloboev/dotfiles/blob/master/nvim/init.vim). And [here](vim-plugins.md) are all the plugins I use with it.

I switch between Ayu Light and Ayu Mirage themes as I change between macOS appearances.

![](https://i.imgur.com/LKSOmms.png)

> Ayu Light

![](https://i.imgur.com/f0wFbdH.png)

> Ayu Mirage

I also use a little trick where I change the cursor from thin block in insert mode to underline when in normal mode. This lets me quickly know which mode I am in.

## Notes

- I can profile vim startup time with [startuptime](https://github.com/tweekmonster/startuptime.vim).
- I can run `:CheckHealth` to see if there is anything wrong with my installation (only Neovim).
- Get plugins that meet your needs.
- You're more likely to find useful plugins if you approach it the other way around: I want to do X, vim can't do it nicely by itself, which plugin can help me do it.
- If something doesn't work in vimrc, put the binding in the bottom of vimrc (perhaps something overrides it).
- [I care about reducing friction, and Vim's editing language is very good at that.](https://www.reddit.com/r/vim/comments/hg37kq/question_how_much_time_do_you_save_using_vim/)

## Links

- [You don't grok vim](https://stackoverflow.com/questions/1218390/what-is-your-most-productive-shortcut-with-vim/1220118#1220118)
- [How to start learning Vim](https://www.reddit.com/r/vim/comments/8fqsju/learning_vim_and_switching_to_it/dy5qkpe/?context=1)
- [Vim Koans](https://sanctum.geek.nz/arabesque/vim-koans/) - Quite funny.
- [Macros](http://vim.wikia.com/wiki/Macros)
- [Seven habits of effective text editing](http://www.moolenaar.net/habits.html)
- [Go client for neovim](https://github.com/neovim/go-client)
- [Vim Golf](http://www.vimgolf.com/)
- [Vim Is Not About hjkl](http://sandymaguire.me/blog/vim-is-not-about-hjkl/)
- [Ask HN: What is one Vim trick most people don't know? (2018)](https://news.ycombinator.com/item?id=17422516)
- [Vim Proverbs](https://www.vi-improved.org/vim-proverbs/)
- [tpope's Vim Config and Plugins (2014)](https://www.youtube.com/watch?v=MGmIJyTf8pg)
- [veonim](https://github.com/veonim/veonim) - Simple modal IDE built on neovim.
- [The Last Statusline For Vim (2018)](https://kadekillary.work/post/statusline-vim/)
- [How I revamped my Vim setup (2019)](https://alex.dzyoba.com/blog/vim-revamp/)
- [Intermediate vim](https://dn.ht/intermediate-vim/) ([Lobsters](https://lobste.rs/s/78yjp6/intermediate_vim)) ([HN](https://news.ycombinator.com/item?id=19794558))
- [vi Complete Key Binding List](https://hea-www.harvard.edu/~fine/Tech/vi.html)
- [Effective VimScript (2019)](https://arp242.net/effective-vimscript.html) ([Reddit](https://www.reddit.com/r/vim/comments/dpmnd1/effective_vimscript/))
- [Lobsters: Whats your vim setup like? (2019)](https://lobste.rs/s/ffhwse/whats_your_vim_setup_like)
- [nvim-lsp](https://github.com/neovim/nvim-lsp) - Common configurations for Neovim Language Servers.
- [We can have nice things (2019)](https://www.youtube.com/watch?v=Bt-vmPC_-Ho) ([Slides](https://vimconf.org/2019/slides/justin.pdf))
- [Vim like apps/plugins](https://vim.reversed.top/) ([HN](https://news.ycombinator.com/item?id=21692418#21693299))
- [pack](https://github.com/maralla/pack) - Missing vim8 package manager.
- [Neovide](https://github.com/Kethku/neovide) - No Nonsense Neovim Client in Rust.
- [Lisp in Vim with Slimv or Vlime (2019)](https://susam.in/blog/lisp-in-vim-with-slimv-or-vlime/)
- [Build Your Own Vim Emulation for VS Code (2020)](https://johtela.github.io/vscode-modaledit/docs/.vscode/settings.html) ([HN](https://news.ycombinator.com/item?id=22383841))
- [130+ Essential Vim Commands (2020)](https://catswhocode.com/vim-commands/)
- [libvim](https://github.com/onivim/libvim) - Core Vim editing engine as a minimal C library.
- [diagnostic-nvim](https://github.com/haorenW1025/diagnostic-nvim) - Wrapper for neovim built in LSP diagnosis config.
- [Mastering Vim Quickly](https://jovicailic.org/mastering-vim-quickly/)
- [Missing Semester VIM lecture (2020)](https://www.youtube.com/watch?v=a6Q8Na575qc) ([HN](https://news.ycombinator.com/item?id=23436392))
- [Make Vim Amazing videos](https://www.youtube.com/playlist?list=PLm323Lc7iSW9kRCuzB3J_h7vPjIDedplM)
- [Vim Life videos](https://www.youtube.com/playlist?list=PLm323Lc7iSW9CtibHhhQErDh167XfL4EU)
- [General purpose Language Server that integrate with linter to support diagnostic features](https://github.com/iamcco/diagnostic-languageserver)
- [Neovim Async Tutorial (2020)](https://ms-jpq.github.io/neovim-async-tutorial/)
- [How Did Vim Become So Popular (2020)](https://pragmaticpineapple.com/how-did-vim-become-so-popular/) ([HN](https://news.ycombinator.com/item?id=23689091)) ([Lobsters](https://lobste.rs/s/smszig/how_did_vim_become_so_popular)) ([Reddit](https://www.reddit.com/r/vim/comments/hilqrw/how_did_vim_become_so_popular/))
- [Vim's Versatile CLI (2020)](https://www.youtube.com/watch?v=pt36X1OJRG4) ([Lobsters](https://lobste.rs/s/v0vfcy/vim_s_versatile_cli))
- [GNvim](https://github.com/vhakulinen/gnvim) - Rich Neovim GUI without any web bloat.
- [Why Should You Learn Vim in 2020](https://pragmaticpineapple.com/why-should-you-learn-vim-in-2020/) ([Lobsters](https://lobste.rs/s/07tbby/why_should_you_learn_vim_2020)) ([HN](https://news.ycombinator.com/item?id=24064809))
- [Vim Search, Find and Replace: a Detailed Guide (2020)](https://thevaluable.dev/vim-search-find-replace/)
- [Neovim 0.5.0: Language Server Protocol (2020)](https://nathansmith.io/posts/neovim-lsp)
- [Neovim for macOS](https://github.com/JaySandhu/neovim-mac) - Fast, minimal, Neovim GUI for macOS.
- [VimTricks](https://vimtricks.com/) - Weekly Vim tips, tricks, guides, screencasts, and more.
- [Learn Vim (the Smart Way)](https://github.com/iggredible/Learn-Vim) ([HN](https://news.ycombinator.com/item?id=24287566))
- [VIM quick-start, cheat-sheet, and links (2013)](https://jkirchartz.com/2013/04/vim_quick_start_cheatsheet_and_links.html)
- [VimConf](https://www.vimconf.live/) - Virtual vim conference.
- [Vim: Sharpening the Axe (2020)](https://www.youtube.com/watch?v=iEShYRRVZOE)
- [Using LibUV in Neovim](https://teukka.tech/vimloop.html)
- [Beyond basic modal editing. Using vim's command-line mode (2020)](https://thoughtbot.com/blog/beyond-basic-modal-editing-using-vims-command-line-mode)
- [Vimscript Language Server](https://github.com/google/vimscript-language-server)
- [Vim Registers: What, How, and Macros (2020)](https://www.youtube.com/watch?v=Q5eDxR7bU2k)
- [GoNeovim](https://github.com/akiyosi/goneovim) - Neovim GUI written in Go, using a Qt binding for Go.
- [pyvim](https://github.com/prompt-toolkit/pyvim) - Pure Python Vim clone.
- [NeoVims built-in Language Server Client and why you should use it (2020)](https://expectationmax.github.io/2020/NeoVims-Language-Server-Client/)
- [Getting started using Lua in Neovim](https://github.com/nanotee/nvim-lua-guide)
- [Firenvim](https://github.com/glacambre/firenvim) - Embed Neovim in your browser.
- [Why I teach vim (2020)](https://blog.ceos.io/2020/11/14/why-i-teach-vim/) ([HN](https://news.ycombinator.com/item?id=25097788))
- [Improving Vim Workflow With fzf (2020)](https://pragmaticpineapple.com/improving-vim-workflow-with-fzf/) ([Lobsters](https://lobste.rs/s/buaxj3/improving_vim_workflow_with_fzf))
- [Vim Script Parser written in Go](https://github.com/haya14busa/go-vimlparser)
- [glrnvim](https://github.com/beeender/glrnvim) - Fast & stable neovim GUI could be accelated by GPU.
- [Five lines I put in a blank .vimrc (2020)](https://swordandsignals.com/2020/12/13/5-lines-in-vimrc.html) ([HN](https://news.ycombinator.com/item?id=25410390)) ([Lobsters](https://lobste.rs/s/du8i6z/5_lines_i_put_blank_vimrc))
- [neovim-editor](https://github.com/rhysd/neovim-component) - WebComponent to embed Neovim to your app with great ease.
- [How do you index code in your projects? (2020)](https://lobste.rs/s/ujr9mg/how_do_you_index_code_your_projects)
- [At least one Vim trick you might not know (2019)](https://www.hillelwayne.com/post/intermediate-vim/) ([HN](https://news.ycombinator.com/item?id=20077580))
- [A Vim Guide for Intermediate Users (2020)](https://thevaluable.dev/vim-intermediate/) ([Lobsters](https://lobste.rs/s/mr8ynv/vim_guide_for_intermediate_users))
- [The ABCs of vim (2020)](https://www.kovach.me/The_ABCs_of_vim.html)
- [Customizing Neovim LSP (nvim-lsp) (2020)](https://rishabhrd.github.io/jekyll/update/2020/09/19/nvim_lsp_config.html)
- [popfix](https://github.com/RishabhRD/popfix) - Neovim lua API for highly extensible quickfix window.
- [vim9jit](https://github.com/tjdevries/vim9jit) - Vimscript but in Lua and fast.
- [sql.nvim](https://github.com/tami5/sql.nvim) - SQLite/LuaJIT binding for lua and neovim.
- [Neovim client for Julia](https://github.com/bfredl/Neovim.jl)
- [neovim-lib](https://github.com/daa84/neovim-lib) - Rust library for Neovim clients.
- [Your First VimRC 2021](https://www.youtube.com/watch?v=DogKdiRx7ls)
- [vim.so](https://www.vim.so/) - Learn and Master Vim faster with interactive exercises. ([HN](https://news.ycombinator.com/item?id=25846347))
- [lspsaga.nvim](https://github.com/glepnir/lspsaga.nvim) - Light-weight lsp plugin based on neovim built-in lsp with highly performance UI.
- [nlua.nvim](https://github.com/tjdevries/nlua.nvim) - Lua Development for Neovim.
- [Configuring Neovim using Lua (2021)](https://icyphox.sh/blog/nvim-lua/)
- [A Vim Guide for Advanced Users (2021)](https://thevaluable.dev/vim-advanced/) ([HN](https://news.ycombinator.com/item?id=26284618)) ([Reddit](https://www.reddit.com/r/programming/comments/ltm4vh/a_vim_guide_for_advanced_users/))
- [Extreme Vim Macros for Traditionalist Catholics (2021)](https://www.youtube.com/watch?v=FXCitlsA7eQ) ([Reddit](https://www.reddit.com/r/vim/comments/lw71ax/extreme_vim_macros_for_traditionalist_catholics/))
- [Vim is Turing-Complete (2021)](https://buttondown.email/hillelwayne/archive/vim-is-turing-complete/)
- [nvcode](https://github.com/ChristianChiarulli/nvcode) - Ultimate Neovim Config NVCode.
- [Vim normal mode grammar](https://glts.github.io/2013/04/28/vim-normal-mode-grammar.html)

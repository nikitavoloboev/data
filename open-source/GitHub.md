# GitHub

I 💚 GitHub. It is really one of the best communities out there that exists with so many different projects and ideas being worked on at any given time.

I publish most of my work and make it open source for other people to use and reuse and I build my own system to how I can be productive using the GitHub website and all the tooling and services it provides.

[Here](https://my.mindnode.com/ZKGETDkUaQUsL3q8q9z788CxG84oEHgDiT79GuzX#-143.5,-902.6,0) you can find all the repositories I have shared. 

# Issues

For managing issues I have, I use [Ship macOS app](https://www.realartists.com).

If I can't work on a task straight away, I assign `future` tag on it. 

Ship app also allows you to put certain issues into Up Next group as can be seen below : 

![](https://i.imgur.com/DgMzQa1.png)

And then I can go through these issues one by one.

# Tips and Tricks

- selecting some text in a github comment and pressing `r` will make a reply of that text
- [GitHub jump](https://github.com/lox/alfred-github-jump) workflow is one of the best things you can install to navigate GitHub like a God
	- I [slightly modified it](https://www.dropbox.com/s/t3iyjt3pyuz8mup/github%20jumps-.alfredworkflow?dl=1) to jump to different parts of the repo with modifiers like going to issues, releases, opening new issue or even cloning the repo instantly
- [this](http://www.somsubhra.com/github-release-stats/) is really cool website where you can find how people people have downloaded various GitHub releases of people


## go to first commit of currently open repo bookmarklet

\`\`\`Javascript
javascript:(b=\>fetch('https://api.github.com/repos/'+b[1]+'/commits?sha='+(b[2]||'')).then(c=\>Promise.all([c.headers.get('link'),c.json()])).then(c=\>{if(c[0]){var d=c[0].split(',')[1].split(';')[0].slice(2,-1);return fetch(d).then(e=\>e.json())}return c[1]}).then(c=\>c.pop().html\_url).then(c=\>window.location=c))(window.location.pathname.match(/\/([^\/]+\/[^\/]+)(?:\/tree\/([^\/]+))?/));
\`\`\`

I run it from Keyboard Maestro like so : 

![](https://i.imgur.com/UIQNrjr.png)
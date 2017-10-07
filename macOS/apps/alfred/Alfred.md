# Alfred

[Alfred](https://www.alfredapp.com/) is a macOS launcher that is extremely powerful as you can program it to show what you want.

## clean prompt

Sometimes when passing things from one object into another, you want to clear your {query} being passed to get a clean prompt. To achieve this you can set argument utility and set it as empty like here : 

![](https://i.imgur.com/seduWW7.png)

# Links

- [LA](https://learn-anything.xyz/software/tooling/productivity/alfred)
- [how to: workflow / environment variables](https://www.alfredforum.com/topic/9070-how-to-workflowenvironment-variables/?tab=comments#comment-45177)

# Notes

- you can't exclude paths in Alfred file filters but you can exclude certain other things to get around this 
	- i.e. I wanted to search for folders in my [knowledge repo](https://github.com/nikitavoloboev/knowledge) but not search node\_modules
	- I added a finder comment to node\_modules like so 
![](https://i.imgur.com/J8Co1dt.png)
- and then in my Alfred File filter I said to skip files where the comment is _skip_

![](https://i.imgur.com/qRIYTZi.png)
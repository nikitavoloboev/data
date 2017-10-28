# Wiki Structure
Below you will find all the notes I have made. I continuously update them and I use them myself for my own reference. However I try to write things and structure them so everyone can benefit from reading them or even reference what I know easily and quickly.

Each folder is a topic that can include other topics under it, ideally related to the parent topic.

Gitbook provides a pretty awesome search feature that lets you search across this entire knowledge base, so feel free to use it.

I also really love using night mode. You can switch to it from here if you haven't :
![](https://i.imgur.com/k6EDR4K.png)

If this is your first time visiting this wiki, you can just start reading from the top entry down and see what sparks your interest. The top folders include content that is often changing which includes [projects](../projects/Projects.md) I am working on, [things I am working on right now](../working-on/working-on.md), things I am [learning](../working-on/Learning.md) and [reading](../working-on/Reading.md) and [some ideas](../working-on/Ideas.md) I have that I don't have time or skill to work on. 

I will be updating this wiki quite often as I use it myself daily both to keep an account of things I know as well as things I want to know and everything in-between.

For super fast access of this wiki, I do advise you to try out [Alfred My Mind](https://github.com/nikitavoloboev/alfred-my-mind).

If you find some mistake, especially if something that I say is plain wrong, please fork [this repository](https://github.com/nikitavoloboev/knowledge) and make a PR with correct changes. Or [open an issue](https://github.com/nikitavoloboev/knowledge/issues/new) saying what you think is wrong.

## Content Structure
Each topic will have a title, some description of it, usually my own thoughts and knowledge on it as well as referencing some resources or links I have liked or used that helped me either understand the topic or gain appreciation of it.

The structure of each of the posts will often look roughly like this : 
1. Title
2. Description
	- my thoughts on the topic
3. Subtopics 
	- various subtopics related to the main topic
4. Snippets
	- if the topic is programming related, I will attach a few Gists related to the topic I personally use
5. Links
	- links related to the topic
6. Notes
	- my own personal notes on the matter as well as things I found interesting on the internet regarding the topic
		- I often give a link of where I got it from

In links I very often link to a  [Learn Anything](https://learn-anything.xyz/) map as that in my opinion is the perfect place to start learning any topic as it should ideally show you the most efficient path for learning anything. Aside from the link to LA, I may also include some of my own things I wrote or really liked as well as my comments on these links.

As my focus at current time is writing code to solve various problems, I will also be including a lot of code and covering various technical and interesting to me subjects. The code I will be linking to will be embedded as Gists and you can use and reuse it as you see fit. 

## My workflow in writing and maintaining this wiki
I write all these entries in [Ulysses](../macOS/apps/Ulysses.md) app and since all of the entries are simply markdown files sorted into folders and [hosted on GitHub](https://github.com/nikitavoloboev/knowledge). I hook this git _knowledge_ folder to Ulysses as external folder : 

![](https://i.imgur.com/EgReXPY.png)

And then I have all the power of Ulysses search across my entire wiki base + my personal entries : 

![](https://i.imgur.com/M7vXSjl.png)

I also use two specialised Alfred workflows. One that specifically searches through folders in this _knowledge_ main folder : 

![](https://i.imgur.com/FdZB0Aj.png)

And then on activation will open these folders in Ulysses so I can quickly add more notes to the folder. I then have another workflow that searches for all the markdown files in this _knowledge_ folder : 

![](https://i.imgur.com/g7riHqR.png)

And then it opens these markdown files in Ulysses for quickly editing them.

Aside from these two workflows I heavily automate and customise Ulysses to my liking. And I use this Keyboard Maestro macro to quickly start editing my [SUMMARY file](https://github.com/nikitavoloboev/knowledge/blob/master/SUMMARY.md) which defines the order in which Gitbook actually renders my markdowns. Here is how this simple macro looks like : 

![](https://i.imgur.com/2GLv5A4.png)

I also use this macro to quickly create folders in this wiki : 

![](https://i.imgur.com/aGywogI.png)

And finally, the most used macro I use to fully automate updating this wiki is simply committing all the changes in this wiki with a single hotkey : 

![](https://i.imgur.com/BuKbfkn.png)

And that is pretty much it. This lets me be extremely fast with editing and maintaining this wiki. I also heavily make use of this wiki by searching through Gitbook book itself : 

![](https://i.imgur.com/UY8B6gc.png)
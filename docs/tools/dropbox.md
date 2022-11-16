---
title: Dropbox
---

# [Dropbox](https://www.dropbox.com)

I now use [iCloud](https://www.icloud.com/) to sync my files.

I used Dropbox previously to share and keep files on the cloud. Don't use it any more as I try to minimize unneccesary running processes on my macOS. 

I currently only use it for sharing some files to download for people although plan to write my own service for this with [Cloudflare Workers R2](../cloud-computing/serverless-computing/cloudflare-workers.md). I have a folder `Shares` and inside it is where I put various files that I want to publicly share with people. I then use a modified version of [Directory Watches workflow](https://github.com/nikitavoloboev/small-workflows/blob/master/augmentations/Directory%20watches.alfredworkflow?raw=true) to scan through this folder from Alfred. I then modified one of the actions so that on pressing `return` it gives me a shareable link of the file. Here is how that Alfred filter looks like for me:

![](https://i.imgur.com/ipbEhil.png)

It's pretty amazing as I can very easily query all the links I shared with anyone. I want to move this workflow to something I own myself.

For sharing files temporarily I use [transfer.sh](https://transfer.sh) and [Vitor's awesome Alfred workflow](https://www.alfredforum.com/topic/5233-uploadfile-%E2%80%94-upload-files-and-directories-for-easy-sharing/) for it.

I also used Dropbox for sharing configuration of some apps like [Alfred](../macOS/apps/alfred/alfred.md).

![](https://i.imgur.com/F9nsqBn.png)

## Notes

- [Dropbox on mac can slow down Swift compilation](https://twitter.com/macguru17/status/1458037982435418119)

## Links

- [Replacing Dropbox in favor of DigitalOcean spaces (2021)](https://mitjafelicijan.com/replacing-dropbox-in-favor-of-digitalocean-spaces.html) ([HN](https://news.ycombinator.com/item?id=25909336))
- [Maestral](https://maestral.app/) - Open source Dropbox client. ([Code](https://github.com/samschott/maestral)) ([HN](https://news.ycombinator.com/item?id=30831214))
- [Dropbox Sync does not natively support Apple Silicon](https://twitter.com/mitchellh/status/1453394500848537605) ([Reddit](https://www.reddit.com/r/apple/comments/qh6or2/dropbox_doesnt_support_apple_silicon_natively_yet/)) ([HN](https://news.ycombinator.com/item?id=29026304)) ([Tweet](https://twitter.com/marcoarment/status/1453735403626766341))
- [git-remote-dropbox](https://github.com/anishathalye/git-remote-dropbox) - Transparent bidirectional bridge between Git and Dropbox.
- [Dropbox Rust SDK](https://github.com/dropbox/dropbox-sdk-rust)
- [dboxpaper](https://github.com/mattn/dboxpaper) - Client for Dropbox Paper.

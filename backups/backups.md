# Backups

Currently I use [Arq](https://www.arqbackup.com) to make scheduled (every 120 hours & after 21:00) encrypted cloud backups of macOS to [BackBlaze](https://www.backblaze.com/). iOS as well as `~/Documents`, `~/Desktop` on macOS are automatically backed up to iCloud. Other documents are backed up with [Dropbox](https://www.dropbox.com/).

I love the fact that I can set up any new mac or phone to use my `exact` and perfect setup without any cruft in seconds. All my [dotfiles](https://github.com/nikitavoloboev/dotfiles) are online and can be setup on any new mac via few commands, thanks to [Nix](../package-managers/nix/nix.md).

## Backup tools

- [Arq](https://www.arqbackup.com/) - Automatically backs up Macs and Windows PCs.
- [Restic](https://github.com/restic/restic) - Fast, secure, efficient backup program. ([HN](https://news.ycombinator.com/item?id=21410833))
- [BorgBackup](https://github.com/borgbackup/borg) - Deduplicating archiver with compression and authenticated encryption.
  - [Vorta Backup Client](https://github.com/borgbase/vorta) - Backup client for macOS and Linux desktops. It integrates the mighty BorgBackup with your desktop environment to protect your data from disk failure, ransomware and theft.
  - [BorgBase](https://www.borgbase.com/) - Specialized Hosting Service for BorgBackup.
- [HashBackup](http://www.hashbackup.com/) - Unix command-line backup program to create a local backup, remote offsite backup, or both, in your own storage accounts.
- [Kopia](https://github.com/kopia/kopia) - Simple, cross-platform tool for managing encrypted backups in the cloud.
- [knoxite](https://github.com/knoxite/knoxite) - Data storage & backup system.
- [Rclone](https://github.com/rclone/rclone) - Command line program to sync files and directories to and from different cloud storage providers. ([HN](https://news.ycombinator.com/item?id=22791036))

## Cloud Storage

- [BackBlaze](https://www.backblaze.com/)
- [Wasabi](https://wasabi.com/)

## Notes

- Data that is not backed up is lost data.
- Even the smallest file can take days to recreate.

## Links

- [Tao backup](http://taobackup.com/) - Pretty funny take on backups.
- [Rsync time backup](https://github.com/laurent22/rsync-time-backup)
- [Cloud Backup Comparison. Which one is the fastest? (2019)](https://www.arqbackup.com/cloud-backup-comparison.html)
- [Lessons learned when my SSD died (2018)](https://bsago.me/blog/lessons-learned-when-my-ssd-died)
- [Lobsters: Where do you host your back ups? (2019)](https://lobste.rs/s/c8long/where_do_you_host_your_back_ups)
- [go-backblaze](https://github.com/kothar/go-backblaze) - Go client for Backblaze's B2 storage.

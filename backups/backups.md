# Backups

Currently I use [Arq](https://www.arqbackup.com) to make scheduled (every 120 hours & after 21:00) encrypted cloud backups of macOS to [BackBlaze](https://www.backblaze.com/). iOS as well as `~/Documents`, `~/Desktop` on macOS are automatically backed up to iCloud. Other documents are backed up with [Dropbox](https://www.dropbox.com/).

I love the fact that I can set up any new mac or phone to use my `exact` and perfect setup without any cruft in seconds. All my [dotfiles](https://github.com/nikitavoloboev/dotfiles) are online and can be setup on any new mac via few commands, thanks to [Nix](../package-managers/nix/nix.md).

## Backup tools

- [Arq](https://www.arqbackup.com/) - Automatically backs up Macs and Windows PCs.
- [Restic](https://github.com/restic/restic) - Fast, secure, efficient backup program. ([HN](https://news.ycombinator.com/item?id=21410833)) ([Web](https://restic.net/))
- [BorgBackup](https://github.com/borgbackup/borg) - Deduplicating archiver with compression and authenticated encryption.
  - [Vorta Backup Client](https://github.com/borgbase/vorta) - Backup client for macOS and Linux desktops. It integrates the mighty BorgBackup with your desktop environment to protect your data from disk failure, ransomware and theft.
  - [BorgBase](https://www.borgbase.com/) - Specialized Hosting Service for BorgBackup.
- [HashBackup](http://www.hashbackup.com/) - Unix command-line backup program to create a local backup, remote offsite backup, or both, in your own storage accounts.
- [Kopia](https://github.com/kopia/kopia) - Simple, cross-platform tool for managing encrypted backups in the cloud.
- [knoxite](https://github.com/knoxite/knoxite) - Data storage & backup system.
- [Rclone](https://github.com/rclone/rclone) - Command line program to sync files and directories to and from different cloud storage providers. ([HN](https://news.ycombinator.com/item?id=22791036)) ([Web](https://rclone.org/))
- [bup](https://github.com/bup/bup) - Very efficient backup system based on the git packfile format, providing fast incremental saves and global deduplication.
- [encbup](https://github.com/skorokithakis/encbup) - Companion to bup, the backup program. encbup adds encryption to bup, while still allowing per-file deduplication.
- [Conserve](https://github.com/sourcefrog/conserve) - Robust portable backup tool in Rust.
- [tm](https://github.com/erica/tm) - Time Machine recovery utility for the six-or-so people who still use Time Machine for version control.
- [Unison](https://github.com/bcpierce00/unison) - File-synchronization tool for OSX, Unix, and Windows. ([Web](https://www.cis.upenn.edu/~bcpierce/unison/))
- [Tarsnap](https://www.tarsnap.com/) - Online backups for the truly paranoid. ([HN](https://news.ycombinator.com/item?id=24535046))
- [Bupstash](https://bupstash.io/) - Encrypted backups made easy. ([Code](https://github.com/andrewchambers/bupstash)) ([Introducing Bupstash](https://acha.ninja/blog/introducing_bupstash/)) ([Lobsters](https://lobste.rs/s/k5opww/introducing_bupstash))
- [BlobBackup](https://blobbackup.com/) - Simple Backups to Any Storage.

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
- [How to store data forever (2020)](https://drewdevault.com/2020/04/22/How-to-store-data-forever.html) ([Lobsters](https://lobste.rs/s/il9z0e/how_store_data_forever))
- [Backblaze B2 Cloud Storage Now Has S3 Compatible APIs (2020)](https://www.backblaze.com/blog/backblaze-b2-s3-compatible-api/) ([HN](https://news.ycombinator.com/item?id=23069114))
- [tmignore](https://github.com/samuelmeuli/tmignore) - Exclude development files from Time Machine backups.
- [Setting up your backup service (2020)](https://www.williamjbowman.com/blog/2020/06/30/setting-up-your-backup-service/)
- [Organizing Data Through the Lens of Deduplication (2020)](https://www.anishathalye.com/2020/08/03/periscope/) ([Lobsters](https://lobste.rs/s/udqu02/organizing_data_through_lens))
- [Backing up data like the adult I supposedly am (2020)](https://magnusson.io/post/backups/) ([Lobsters](https://lobste.rs/s/bmqi6l/backing_up_data_like_adult_i_supposedly_am)) ([HN](https://news.ycombinator.com/item?id=24526706))
- [Automatic restic backups using systemd services and timers](https://github.com/erikw/restic-systemd-automatic-backup)
- [Encrypted Backup Shootout (2021)](https://acha.ninja/blog/encrypted_backup_shootout/) ([HN](https://news.ycombinator.com/item?id=25618346))

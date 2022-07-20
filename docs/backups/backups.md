# Backups

Currently I use [Arq](https://www.arqbackup.com) to make scheduled (every 120 hours & after 21:00) encrypted cloud backups of macOS to [BackBlaze](https://www.backblaze.com/). iOS as well as `~/Documents`, `~/Desktop` on macOS are automatically backed up to iCloud. Other documents are backed up with [Dropbox](../tools/dropbox.md).

I love the fact that I can set up any new mac or phone to use my `exact` and perfect setup without any cruft in seconds. All my [dotfiles](https://github.com/nikitavoloboev/dotfiles) are online and can be setup on any new mac via few commands, thanks to [Nix](../package-managers/nix/nix.md).

For programming, [Walrus](https://github.com/Clivern/Walrus) seems nice.

## Backup tools

- [Arq](https://www.arqbackup.com/) - Automatically backs up Macs and Windows PCs.
- [Restic](https://github.com/restic/restic) - Fast, secure, efficient backup program. ([HN](https://news.ycombinator.com/item?id=21410833)) ([Web](https://restic.net/)) ([Rest Server](https://github.com/restic/rest-server)) ([GitHub](https://github.com/restic)) ([Home backups using Restic](https://w.hutson.gy/projects/home-backups-using-restic)) ([HN](https://news.ycombinator.com/item?id=29209455)) ([Restic by Example](https://github.com/rubiojr/rapi)) ([HN](https://news.ycombinator.com/item?id=30822631)) ([Awesome](https://github.com/rubiojr/awesome-restic))
- [BorgBackup](https://github.com/borgbackup/borg) - Deduplicating archiver with compression and authenticated encryption.
  - [Vorta Backup Client](https://github.com/borgbase/vorta) - Backup client for macOS and Linux desktops. It integrates the mighty BorgBackup with your desktop environment to protect your data from disk failure, ransomware and theft.
  - [BorgBase](https://www.borgbase.com/) - Specialized Hosting Service for BorgBackup.
- [HashBackup](http://www.hashbackup.com/) - Unix command-line backup program to create a local backup, remote offsite backup, or both, in your own storage accounts.
- [Kopia](https://github.com/kopia/kopia) - Simple, cross-platform tool for managing encrypted backups in the cloud. ([Web](https://kopia.io/)) ([HN](https://news.ycombinator.com/item?id=27471945))
- [knoxite](https://github.com/knoxite/knoxite) - Data storage & backup system.
- [Rclone](https://github.com/rclone/rclone) - Command line program to sync files and directories to and from different cloud storage providers. ([HN](https://news.ycombinator.com/item?id=22791036)) ([Web](https://rclone.org/)) ([HN](https://news.ycombinator.com/item?id=29435760)) ([syncrclone](https://github.com/Jwink3101/syncrclone))
- [bup](https://github.com/bup/bup) - Very efficient backup system based on the git packfile format, providing fast incremental saves and global deduplication.
- [encbup](https://github.com/skorokithakis/encbup) - Companion to bup, the backup program. encbup adds encryption to bup, while still allowing per-file deduplication.
- [Conserve](https://github.com/sourcefrog/conserve) - Robust portable backup tool in Rust.
- [tm](https://github.com/erica/tm) - Time Machine recovery utility for the six-or-so people who still use Time Machine for version control.
- [Unison](https://github.com/bcpierce00/unison) - File-synchronization tool for OSX, Unix, and Windows. ([Web](https://www.cis.upenn.edu/~bcpierce/unison/))
- [Tarsnap](https://www.tarsnap.com/) - Online backups for the truly paranoid. ([CLI client](https://github.com/Tarsnap/tarsnap)) ([HN](https://news.ycombinator.com/item?id=24535046))
- [Bupstash](https://bupstash.io/) - Encrypted backups made easy. ([Code](https://github.com/andrewchambers/bupstash)) ([Introducing Bupstash](https://acha.ninja/blog/introducing_bupstash/)) ([Lobsters](https://lobste.rs/s/k5opww/introducing_bupstash)) ([The Bupstash Garbage Collector](https://acha.ninja/blog/the_bupstash_garbage_collector/))
- [BlobBackup](https://blobbackup.com/) - Simple Backups to Any Storage.
- [zrepl](https://github.com/zrepl/zrepl) - One-stop ZFS backup & replication solution. ([Docs](https://zrepl.github.io/))
- [Linux Time Machine](https://github.com/cytopia/linux-timemachine) - Rsync-based OSX-like time machine for Linux, MacOS and BSD for atomic and resumable local and remote backups.
- [Autorestic](https://github.com/cupcakearmy/autorestic) - High level CLI utility for restic. ([Docs](https://autorestic.vercel.app/))
- [Walrus](https://github.com/Clivern/Walrus) - Fast, Secure and Reliable System Backup, Set up in Minutes. ([Web](https://clivern.github.io/Walrus/))
- [awsbackup.sh](https://github.com/keisentraut/awsbackup) - Bash script for personal backup with Amazon Web Services S3 Glacier Deep Archive.
- [nFreezer](https://github.com/josephernest/nfreezer) - Encrypted-at-rest backup tool, designed specifically for the case when the destination server is untrusted.
- [rsync](https://www.rsync.net/index.html) - Cloud Storage for Offsite Backup. ([Code](https://github.com/WayneD/rsync)) ([rsyncparse](https://github.com/stapelberg/rsyncparse)) ([rsync-go](https://github.com/gokrazy/rsync)) ([rsync: Scenarios](https://michael.stapelberg.ch/posts/2022-06-18-rsync-article-1-scenarios/)) ([HN](https://news.ycombinator.com/item?id=31958536))
- [Bakelite](https://github.com/richfelker/bakelite) - Incremental backup with strong cryptographic confidentiality baked into the data model.
- [shallow-backup](https://github.com/alichtman/shallow-backup) - Git-integrated backup tool for macOS and Linux devs.
- [paperback](https://github.com/cyphar/paperback) - Paper-based backup scheme that is secure and easy-to-use.
- [Duplicati](https://github.com/duplicati/duplicati) - Store securely encrypted backups on cloud storage services.
- [Back In Time](https://github.com/bit-team/backintime) - Simple backup tool for Linux.

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
- [Back up your MacOS (2021)](https://dzx.cz/2021/05/23/back_up_your_macos/)
- [Personal data backup plan (2020)](https://jacobbednarz.com/personal-data-backup-plan)
- [How I store my files (2021)](https://www.unixsheikh.com/articles/how-i-store-my-files-and-why-you-should-not-rely-on-fancy-tools-for-backup.html) ([HN](https://news.ycombinator.com/item?id=28003119))
- [Saving a restic backup the hard way (2021)](http://blog.pkh.me/p/30-saving-a-restic-backup-the-hard-way.html)
- [Ask HN: How should I back up data on devices if I'm not smart? (2021)](https://news.ycombinator.com/item?id=28758415)
- [Asimov](https://github.com/stevegrunwell/asimov) - Automatically exclude development dependencies from Apple Time Machine backups.
- [Backups and Corruption (2021)](https://www.collicutt.co.uk/notebook/backups.html)
- [What is the best backup solution for self-hosted services? (2021)](https://www.reddit.com/r/selfhosted/comments/qq1zpv/what_is_the_best_backup_solution_for_selfhosted/)
- [Btrbk](httpvs://github.com/digint/btrbk) - Tool for creating snapshots and remote backups of btrfs subvolumes. ([Web](https://digint.ch/btrbk/))
- [Frictionless external backups with systemd (2021)](https://jmtd.net/log/systemd_ext_backups/) ([Lobsters](https://lobste.rs/s/wjxxin/frictionless_external_backups_with))
- [The 3-2-1 Backup Rule – Why Your Data Will Always Survive (2021)](https://www.vmwareblog.org/3-2-1-backup-rule-data-will-always-survive/)
- [How I back up all my data](https://github.com/geerlingguy/my-backup-plan)
- [Backblaze restore for Personal Backup is awful (2021)](https://news.ycombinator.com/item?id=29533753)
- [Ask HN: What is your system for backing up family photos and video? (2022)](https://news.ycombinator.com/item?id=29978099)
- [Bleanser](https://github.com/karlicoss/bleanser) - Tool for cleaning old and redundant backups.
- [qmpbackup](https://github.com/abbbi/qmpbackup) - Live Qemu Incremental backup using dirty-bitmaps.
- [Tape](https://github.com/vitorgalvao/tape) - Backup and restore software settings on macOS.
- [Ask HN: If Apple integrated Time Machine with iCloud, would you use it? (2022)](https://news.ycombinator.com/item?id=30223422)
- [A almost perfect rsync over ssh backup script (2022)](https://blog.zazu.berlin/software/a-almost-perfect-rsync-over-ssh-backup-script.html)
- [Blobbackup](https://blobbackup.com/) - Private, Secure Computer Backups. ([Code](https://github.com/blobbackup/blobbackup)) ([HN](https://news.ycombinator.com/item?id=30577625))
- [Backblaze B2 JavaScript Client](https://github.com/benaubin/b2-js)
- [Blackbox](https://github.com/lemonsaurus/blackbox) - Magically save your database backups and critical logs in your favorite cloud storage provider.
- [bacup](https://github.com/galeone/bacup) - Easy-to-use backup tool designed for servers - written in Rust.
- [Correct Backups Require Filesystem Snapshots (2022)](https://cyounkins.medium.com/correct-backups-require-filesystem-snapshots-23062e2e7a15) ([HN](https://news.ycombinator.com/item?id=31401151))
- [Hoard](https://github.com/Shadow53/hoard) - Program for backing up files from across a filesystem into a single directory and restoring them later. ([Docs](https://hoard.rs/))
- [Immutable backups so simple that unborkable](https://github.com/nathants/backup) ([Lobsters](https://lobste.rs/s/qfva1e/immutable_backups_so_simple_unborkable))
- [Advice needed for backing up and hosting large amount of files (2022)](https://news.ycombinator.com/item?id=31996612)
- [Restic Backup Docker Container](https://github.com/lobaro/restic-backup-docker)

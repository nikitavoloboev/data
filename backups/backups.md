# Backups

Currently only use encrypted Time Machine backups.

But in future, I plan to start using [Arq](https://www.arqbackup.com) to make encrypted cloud backups in addition to Time Machine backups.

I love the fact that I can set up any new machine or phone to use my `exact` and perfect setup without any cruft in seconds.

## Backup tools

- [Arq](https://www.arqbackup.com/) - Automatically backs up Macs and Windows PCs.
- [Restic](https://github.com/restic/restic) - Fast, secure, efficient backup program.
- [BorgBackup](https://github.com/borgbackup/borg) - Deduplicating archiver with compression and authenticated encryption.
  - [Vorta Backup Client](https://github.com/borgbase/vorta) - Backup client for macOS and Linux desktops. It integrates the mighty BorgBackup with your desktop environment to protect your data from disk failure, ransomware and theft.
  - [BorgBase](https://www.borgbase.com/) - Specialized Hosting Service for BorgBackup.

## Notes

- Data that is not backed up is lost data.
- Even the smallest file can take days to recreate.

## Links

- [Tao backup](http://taobackup.com/) - Pretty funny take on backups.
- [Rsync time backup](https://github.com/laurent22/rsync-time-backup)
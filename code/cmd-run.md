# CMD Run

```bash
# Get IP address
ipconfig getifaddr en0 | pbcopy

# Get macOS version
defaults read loginwindow SystemVersionStampAsString | pbcopy

# Flash DNS Cache
sudo dscacheutil -flushcache;sudo killall -HUP mDNSResponder

# Rebuild launch services
/System/Library/Frameworks/CoreServices.framework/Versions/A/Frameworks/LaunchServices.framework/Versions/A/Support/lsregister﻿ -kill -r -domain local -domain user﻿﻿﻿﻿﻿﻿﻿﻿﻿

# List all commands installed in $PATH
print -l $commands

# List all commands/aliases installed in $PATH
whence -m \*

# List all commands/aliases/functions installed in $PATH
print -rl - ${(k)builtins} ${(k)functions} ${(k)commands}

# Switch to root user on Linux
sudo su -

# Get current process PID
echo $$

# Reset launchpad
defaults write com.apple.dock ResetLaunchPad -bool true; killall Dock

# List GPG secret keys
gpg --list-secret-keys --keyid-format LONG

# List GPG keys
gpg --list-keys

# Restart nginx with systemctl
sudo systemctl restart nginx

# Symlink source dir of Alfred workflow (https://gist.github.com/deanishe/35faae3e7f89f629a94e)
workflow-install -s -source

# List all processes running
ps aux

# Print exit code of last command ran
echo $?

# See all aliases assigned
alias

# List VSCode installed packages
code --list-extensions

# convert video to GIF
ffmpeg -i http://myVideo.mov myVideo.gif
```

# CMDs

### [screencapture](https://ss64.com/osx/screencapture.html) (capture images from screen)

```bash
# Take screenshot & save it to Desktop with current date as name
screencapture -ixoa -t jpg ~/"Desktop/$(date).jpg"
```

### [docker](https://github.com/docker/cli) (run processes in isolated containers)

```bash
# Delete all containers
for i in $(docker ps -a -q); do docker rm $i; done

# List all containers
docker ps -a

# Remove a container after it’s stopped
docker run --rm [...]
```

### [kubectl](https://kubernetes.io/docs/reference/kubectl/overview/) (run commands against Kubernetes clusters)

```bash
# Get info on pods in use. Has info on why they failed if they did.
kubectl describe pods

# Get services across all namespaces
kubectl get svc --all-namespaces

# Port forward the <pod> from 5432 to localhost 5300 port
kubectl port-forward <pod> 5300:5432
```

[nix-env](https://nixos.wiki/wiki/Nix-env) (manipulate or query Nix user environments)

```bash
# See installed packages
nix-env -q

# Install packages
nix-env -iA
```

[basename](https://www.explainshell.com/explain/1/basename) (strip directory and suffix from filenames)

```bash
basename "/Users/nikivi/Documents/books/Thinking, fast and slow.pdf" # => Thinking, fast and slow.pdf
```

[chmod](https://www.explainshell.com/explain/1/chmod) (change file mode bits)

```bash
# All users can read and write but cannot execute
chmod 666

# All actions for all users
chmod 777

# Only owner can do all actions; group and other users are allowed only to read
chmod 744
```

[tail](https://www.explainshell.com/explain/1/tail) (output the last part of files) https://www.explainshell.com/explain/tail

```bash
# Shows the last 10 lines of file
tail /etc/passwd

# -n option allows to change the number of lines to display
# where n is the number of lines you want to see
# ie
tail -5 /etc/passwd

# Use tail +n to print lines starting at line n
```

[playgo](https://github.com/plutov/playgo) (send .go file to the Go Playground)

```bash
# Open Go playground of file in browser
playgo <file>
```

[tr](https://www.explainshell.com/explain/tr) (translate or delete characters)

```bash
# Convert all input to upper case
ls | tr a-z a-z

# Take the output and put into a single line
ls | tr  "\n" " "

# Get rid of all numbers
ls -lt | tr -d 0-9
```

[diff](https://www.explainshell.com/explain/diff) (compare files line by line)

```bash
# Compare file1 with file2
diff file1 file2
```

[find](https://www.explainshell.com/explain/find) (walk a file hierarchy)

```bash
# Will search in current directory (.) for the file 'hello_world.py'
# and will return the path to the file
find . -name 'hello_world.py'

# You can also search multiple directories
# will search both Documents and Desktop folders for the file
find Documents Desktop -name 'hello_world.py'
```

[grep](https://www.explainshell.com/explain/grep) (file pattern searcher)

```bash
# Print lines from a file or input stream that match an expression

# -i = case insensitive search

# -v = return all lines that do not contain {}
grep -v {} story.txt
```

[man](https://www.explainshell.com/explain/man) (open manual pages)

```bash
# Search for manual page by keyword
man -k keyword

# ie if you are looking for command to sort something, run
# output will include man page name, man section and quick description
man -k sort

# **Online manual sections**

# (1) = user commands
# (2) = system calls
# (3) = higher-level unix programming library documentation
# (4) = device interface and driver information
# (5) = file descriptions (system configuration files)
# (6) = games
# (7) = file formats, conventions, and encodings (ASCII, suffixes, and so on)
# (8) = system commands and servers

# Open manual page of passwd on section 5
man 5 passwd

# I can often get information about options of some command using either --help or -h flags
# ie
vim --help # vim -h would also work
```

[sort](https://www.explainshell.com/explain/sort) (put the lines of a text file in alphanumeric order)

```bash
# Will process the results of ps aux command with grep
# and will then sort the output with 'sort' command
ps aux | grep bash | sort

# -n option sorts in numerical order
# -r option reverses the order of the sort
```

PlistBuddy (read and write values to plists)

```bash
# Change version of Alfred workflow info.plist
/usr/libexec/PlistBuddy -c "Set :version \"X.Y.Z\"" info.plist
```

[kill](https://www.explainshell.com/explain/kill) (send a signal to a process)

```bash
# Force quit the process with id 1456
kill -9 1456

# Stop process 1456
kill -STOP 1456
```

[head](https://www.explainshell.com/explain/head) (output the first part of files)

```bash
# Shows the first 10 lines of file
head /etc/passwd

# -n option allows to change the number of lines to display
# where n is the number of lines you want to see
# ie
head -5 /etc/passwd
```

tar (manipulate tape archives)

```bash
# Extract tar files. -x = 'extract'. -v = verbose. -f = point to tar fle
tar -xvf some_file.tar.gz
```

env (set environment and execute command, or print environment)

```bash
# View enviroment variables
env

# Variables to know
$HOME # Expands to the path of my home folder
$PS1 # Represents my command prompt line

# I can thus change the way my command prompt looks like this
PS1="\w >"

$PATH # Lists all the directories that can be executable with commands

# Add directory /Users/nikivi/bin to the path
export PATH=/Users/nikivi/bin:$PATH

# If you put executables there, they will be available

# Exported variables get passed on to child processes. Not-exported variables do not.
```

cat (concatenate and print files)

```bash
# Print what is in file1 and file2 to screen
cat file1 file2
```

[ngrok](https://github.com/bubenshchykov/ngrok) (expose your localhost to the web)

```bash
# Creates a shareable link of my local server on port 3000
ngrok http 3000
```

rmdir (remove empty directories)

```bash
# Remove directory
rmdir dir

# Remove non empty directores
rm -rf dir

# Don't use -rf flags with globs such as star (*)
# Best to double check commands before running
```

[mediumexporter](https://github.com/xdamman/mediumexporter) (export medium.com articles to markdown)

```bash
# i.e.
mediumexporter https://medium.com/@nikitavoloboev/karabiner-god-mode-7407a5ddc8f6 > medium_post.md
```

file (determine file type)

```bash
# See type of file
file <file>
```

time (time command execution)

```bash
# Time how long command took to run
time <cmd>
```

# Definitions

```bash
# $PATH

# Add to beginning of $PATH
PATH=~/opt/bin:$PATH

# Add to end of $PATH
PATH=$PATH:~/opt/bin
```

```bash
# Standard input & output

# Send output of command to a file instead of terminal
# Shell will create a file if it does not exist, if it does, the shell erases (clobbers) the original file first
command > file

# Append to file
command >> file

# Match every line that I type after that has 'this' in it and put the results in to hello_grep.txt file
grep this > hello_grep.txt

# Put all the lines matching 'line' into grep_output.txt file
grep line hello.txt > grep_output.txt

# >> Redirect output to append to FILE
grep line hello.txt >> grep_output.txt

# Redirect standard error stream to FILE
# where 2> means redirect output stream 2 to write results to the
# file and not to the terminal
find / -name 'story.txt' 2> error_log.txt

# 1> specifies stream ID 1 (standard output) (default)
# 2> specifies stream ID 2 (standard error)

# Sends both standard error and output to some location
ls /fffff > f 2>&1
```

```md
- Accept - Specifies the file format the requester wants.
- Accept-Language - Specifies the human-readable language, like English, Spanish, or Russian.
- Cache-Control - Specifies whether the response can be generated from a cache.
```

```bash
# Globbing

# Shell can match simple patterns to file and directory names
# this process is known as globbing

# Simplest of these is the glob character * which tells the shell to match any number of arbitrary characters

# Print a list of files in current directory
echo *

# Shell matches arguments containing globs to filenames
# substititutes the filenames for those arguments, and then runs the revised command line
# the substitution is called expansion because the shell substitutes all matching filenames


at* # Expands to all filenames that start with at

*at # Expands to all filenames that end with at

*at* # Expands to all filenames that contain at


# Another shell glob character is (?) - it instructs the shell to match exactly one arbitrary character

# Match both boat and brat
b?at

# If you don't want the shell to expand a glob in a command, enclose the glob in single quotes ('')

# It is important to remember that the shell performs expansions before running commands
```

```md
ASCII code

- ASCII is an encoding representing each typed letter by a number
- each number is stored in one byte (so the number is in 0..255)
  - A is 65
  - B is 66
  - space is 32
- "Unicode" is an encoding for mandarin, greek, arabic, etc. languages
  - typically 2-bytes per "letter"

![](https://i.imgur.com/qbFb4Ik.png)
```

```md
HTTP status codes

- 200-299 are good
- 300-399 understood (located elsewhere)
- 400-499 error (404 = resource no longer exists)
- 400-599 server end (500 = some error on the server)
```

```md
HTTP REST methods

- GET - Used for fetching either a collection of resources or a single resource.
- POST - Used to add a new resource to a collection.
  - For example, we wouldn't POST to /players/567 or /games/1234 because they aren't collections.
  - We would, however, POST to /players or /games to create a new player or a new game.
- PUT - Use when we want to update a record. We wouldn't use PUT on collection or list URLs.
- DELETE - Used for sending a DELETE request to a detail record, a URL for a single record, should delete just that record.
  - Sending DELETE to an entire collection would delete the whole collection but that's usually not implemented, with good reason.
```

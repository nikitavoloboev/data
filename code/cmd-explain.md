# CMD Explain

```bash
# Create symbolic link
ln -s /path/to/original /path/to/symlink
      └───────┬───────┘
              └── the path to the original file/folder
                  can use . or ~ or other relative paths
```

```bash
# Copy directory & contents to another place (-r for recursive)
cp -r /path/to/dir /path/to/wheretocopy
```

```bash
# Preview output of script.
# Piping things into less will preview the content
curl https://21.co | less
```

```bash
# Move everything from one dir to another.

# i.e. move all files in Downloads to Desktop. * after folder means select everything
mv -v ~/Downloads/* ~/Desktop/
```

```bash
# Get full path to file
realpath <file>
```

```bash
# Do regex replace (with perl) on a file
perl -pi -e "<regex>" <file>

# i.e.
perl -pi -e "s/^\s*\{[^\n']*'([^\n']*)'[^\n']*'([^\n#']*)#[^\n']*'[^\n'}]*\},?/'\1', '\2'\n/gm" triggers.js
```

```bash
# Download URL contents
curl -0 <url>
```

```bash
# Give executable permission to file.
chmod a+x <file>
```

```bash
# Discard output of command.
# Will direct ln cmd output to null device that will delete anything written to it.
ln > /dev/null
```

```bash
# See how unicode string is encoded internally.
# https://wiki.soimort.org/unix/cli/
echo hello | hexdump -C

# See how it is encoded in UTF-16: (assume UTF-8 is default encoding)
echo hello | iconv -f utf-8 -t utf-16 | hexdump -C
```

```bash
# Search for occurrence of word in a file.
# i.e. show all occurrences of word fox in file story.txt
grep 'fox' story.txt
```

```bash
# Search whether a string occurs in a directory
# -l = print filenames of matching files
grep -rl "string" /path
```

# Bash
Not too experienced with Bash. But I do know that you have to __quote your variables__. 

Also it's really useful to use [SpellCheck](https://www.shellcheck.net).

# Snippets
### removes some string from a string
```Bash
url="https://github.com/learn-anything/maps"
# removes https:// 
url="${url#https://}"
```

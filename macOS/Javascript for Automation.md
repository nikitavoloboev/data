# Javascript for Automation
Still learning it.

# useful snippets
## get the URL of the frontmost tab
```Javascript
Application('Google Chrome').windows[0].activeTab.url()
```

## get selected text in Chrome
```Javascript
Application('Google Chrome').windows[0].activeTab.execute({javascript:'window.getSelection().toString()'})
```


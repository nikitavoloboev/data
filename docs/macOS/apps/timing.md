---
title: Timing
---

# [Timing](https://timingapp.com/)

I once used Timing app to fully automate tracking time I spend on my computer.

![](https://i.imgur.com/tj0nmih.png)

I setup projects in Timing with rules assigned for each project. I then had Timing infer time I spent on these projects automatically based on the rules.

No longer use any tracking apps as part of my minimization of tools I use.

## Notes

- Can press ⌥ + → in `Review` tab to expand all items fully (assuming I selected the items I want to expand). Can `CMD+A` to select all projects beforehand.
- Running `defaults write info.eurocomp.Timing2 letTaskSuggestionsOverrideExistingProjects -bool true` would make task suggestions independent of current project.
- `defaults write info.eurocomp.Timing2 TaskActivitySuggestionsTableViewController.selectFirstSuggestionByDefault -bool true` makes the first suggestion on by default.

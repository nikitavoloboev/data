# [Timing](https://timingapp.com/?lang=en)

I use Timing app to fully automate tracking time I spend on my computer.

![](https://i.imgur.com/K2wrkjS.png)

I setup projects in Timing with rules assigned for each project. I then have Timing infer time I spent on these projects automatically based on the rules.

![](https://i.imgur.com/Kcvxqjh.png)

## Notes

- Can press ⌥ + → in `Review` tab to expand all items fully (assuming I selected the items I want to expand). Can `CMD+A` to select all projects beforehand.
- Running `defaults write info.eurocomp.Timing2 letTaskSuggestionsOverrideExistingProjects -bool true` would make task suggestions independent of current project.
- `defaults write info.eurocomp.Timing2 TaskActivitySuggestionsTableViewController.selectFirstSuggestionByDefault -bool true` makes the first suggestion on by default.

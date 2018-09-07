# [Fantastical](https://flexibits.com/fantastical)
I use the app to manage events in my life. I sync all events with Google Calendar and I sync many of [my Trello boards](../../sharing/my-trello.md) deadlines with iCal.

I use Fantastical's natural language input together with [many Typinator expansions](https://medium.com/@nikitavoloboev/fantastical-natural-input-text-expansions-3ea8cf7ccac3#.pv5937ncr) to make new events in seconds.

I always view my events from `Week` view. And show 5 days only with all 24h shown for all days. This lets me have a perspective over what I have to do know. What deadlines I have to complete soon. And gives my the freedom to adjust my schedule in light of upcoming deadlines and events.

![](https://i.imgur.com/wjZqdnk.png)

As I heavily use [Trello](../../sharing/my-trello.md) as a way to stay on top of deadlines for various projects, I have it sync with my calendar and I often view Fantastical in `Month` view to know of the upcoming deadlines I have. I trigger a [KM macro](keyboard-maestro/km-macros.md) that will switch to `Month` view with my `Routine` calendar hidden. It looks like this:

![](https://i.imgur.com/fw42gez.png)

## Notes
- To mass delete events, make empty search query (single space) on events I want to remove and delete them. Can also use stock Calendar app and search for `""` and cmd+shift+arrows events I want to delete.
- [Google calendar sync select page](https://calendar.google.com/calendar/syncselect) allows you to select which calendars you want to sync to external apps like Fantastical. Useful for letting Fantastical know about `Other calendars` which by default don't sync.
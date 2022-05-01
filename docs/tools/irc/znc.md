---
title: ZNC
---

# [ZNC](https://wiki.znc.in/ZNC)

## Interesting Modules

- [Advanced ZNC playback](http://wiki.znc.in/Playback)
- [otr](https://wiki.znc.in/Otr) - Allows you to encrypt private conversations using the OTR protocol.

## ZNC commands

- `/quote detach ..` - Detach a channel from ZNC
  - i.e. `/quote detach ##graphics`
  - Pay attention to the slashes as they matter.
  - Connecting back is then simply `/join ..` (or connecting through Textual GUI click).
- `/znc jump` - Reconnect to the network. Useful for activating changes you made to the ZNC itself. Such as changing real name or nick.
- `/msg *sasl set NICK PASSWORD` where NICK is IRC nick and PASSWORD is IRC password to make ZNC authenticate with SASL where you are identified before being fully connected to the network.

## Links

- [ZNC commands](https://wiki.znc.in/Using_commands)
- [Modules](https://wiki.znc.in/Modules)
- [LunarBNC](https://lunarbnc.net/) - Free BNC service.

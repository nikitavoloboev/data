---
title: Tor Project
---

# [Tor Project](https://www.torproject.org/download/)

## Notes

- [For users wanting to prevent their ISP from sniffing around then tor works as intended. Against advertisers it also work decently as a self cleaning browsers that constantly change its IP address.](https://news.ycombinator.com/item?id=24564415)
  - For developers and sysadmins that want to get an outside look at their own services or investigate third party websites (like fraudulent lookalike) it work pretty effective with some caveats.
  - It also works mostly fine against national and ISP firewalls that is intended to censor citizens and lead people away from places which the state has declared unsuited for its population.
  - Against police force it seem to mostly work as a free tool that get used by criminals as something better than nothing, but with some larger caveats and the police have cases from time to time where they have identified criminals (from either good investigations or parallel constructions depending on who you ask). The tor browsers has also not been immune to malware.
  - Against national-level intelligence agency, "citizen scores", and whistleblowers employed within such agencies, the protection granted by tor may be very far from 100%. It is not recommended by anyone to depend on tor against that threat model.

## Links

- [TinyTor](https://github.com/Marten4n6/TinyTor) - Tiny Tor client implementation (in pure python).
- [Awesome Tor](https://github.com/ajvb/awesome-tor)
- [Exitmap](https://github.com/NullHypothesis/exitmap) - Fast and modular Python-based scanner for Tor exit relays.
- [New Release: Tor Browser 9.5 (2020)](https://blog.torproject.org/new-release-tor-browser-95) ([HN](https://news.ycombinator.com/item?id=23392814))
- [Operating a Tor Relay (2020)](https://birb007.github.io/blog/2020/06/06/operating-a-tor-relay.html) ([Lobsters](https://lobste.rs/s/is2li1/operating_tor_relay)) ([HN](https://news.ycombinator.com/item?id=23536949))
- [MoreOnionsPorfavor: Onionize your website and take back the internet (2020)](https://blog.torproject.org/more-onions-porfavor) ([HN](https://news.ycombinator.com/item?id=23775704))
- [Snowflake Proxy on Mobile (2020)](https://blog.torproject.org/gsoc-2020-snowflake-proxy-mobile)
- [OnionSearch](https://github.com/megadose/OnionSearch) - Script that scrapes urls on different .onion search engines.
- [tor-controller](https://github.com/kragniz/tor-controller) - Run Tor onion services on Kubernetes.
- [Paper review: Statistical and Combinatorial Analysis of the TOR Routing Protocol (2021)](https://dustri.org/b/paper-review-statistical-and-combinatorial-analysis-of-the-tor-routing-protocol.html)
- [Public TOR IPv6 Only Gateway (2021)](https://blog.shamm.as/posts/2021-01-18-public_tor_ipv6_gateway/)
- [Bine](https://github.com/cretz/bine) - Go library for accessing and embedding Tor clients and servers.
- [Arti](https://gitlab.torproject.org/tpo/core/arti/) - Pure-Rust Tor implementation. ([Article](https://blog.torproject.org/announcing-arti)) ([Reddit](https://www.reddit.com/r/programming/comments/ogw5o7/the_tor_project_announces_arti_a_tor/)) ([HN](https://news.ycombinator.com/item?id=30683879))
- [onion2web](https://github.com/starius/onion2web) - Access .onion sites without Tor Browser.
- [Tor is a Great SysAdmin Tool (2020)](https://www.jamieweb.net/blog/tor-is-a-great-sysadmin-tool/)
- [Lightnion](https://github.com/spring-epfl/lightnion) - JavaScript library that you can include on your webpage to let any modern browser make anonymous requests. Light version of Tor.
- [Real-World Onion Sites](https://github.com/alecmuffett/real-world-onion-sites) - List of substantial, commercial-or-social-good mainstream websites which provide onion services.
- [DoHoT](https://github.com/alecmuffett/dohot) - Making practical use of DNS over HTTPS over Tor.
- [finshir](https://github.com/isgasho/finshir) - High-performant, coroutines-driven, and fully customisable implementation of Low & Slow load generator designed for real-world pentesting. Complete undetectability is achieved by connecting through Tor.
- [Emerald Onion](https://emeraldonion.org/) - Tor host with a mission to protect privacy, anonymity, access to information, and free speech online.
- [mkp224o](https://github.com/cathugger/mkp224o) - Vanity address generator for ed25519 onion services.
- [Tor Browser 11.0 (2021)](https://blog.torproject.org/new-release-tor-browser-11-0) ([HN](https://news.ycombinator.com/item?id=29165747))
- [Tor blocked in Russia: how to circumvent censorship (2021)](https://forum.torproject.net/t/tor-blocked-in-russia-how-to-circumvent-censorship/982)
- [Tor Project Forum](https://forum.torproject.net/)
- [The Tor Project · GitLab](https://gitlab.torproject.org/tpo)
- [Using ephemeral Onion Services for quick NAT traversal (2021)](https://www.trickster.dev/post/using-ephemeral-onion-services-for-quick-nat-traversal/)
- [tor-dam](https://github.com/parazyd/tordam) - Library for peer discovery inside the Tor network.
- [Tor Snowflake Proxy](https://snowflake.torproject.org/) ([HN](https://news.ycombinator.com/item?id=29634636)) ([Tweet](https://twitter.com/genderjokes/status/1497284560811225095))
- [Tor in 2022](https://blog.torproject.org/tor-in-2022/) ([HN](https://news.ycombinator.com/item?id=29635213))
- [Expose server behind NAT with Tor (2019)](https://golb.hplar.ch/2019/01/expose-server-tor.html) ([HN](https://news.ycombinator.com/item?id=29929399))
- [Tor Browser: a legacy of advancing private browsing innovation (2022)](https://blog.torproject.org/tor-browser-advancing-privacy-innovation/) ([HN](https://news.ycombinator.com/item?id=30123982))
- [oniongrok](https://github.com/cmars/oniongrok) - Onion addresses for anything.
- [go-libtor](https://github.com/berty/go-libtor) - Self-contained, fully statically linked Tor library for Go.
- [Help Censored Users – Run a Tor Bridge (2022)](https://blog.torproject.org/run-a-bridge-campaign/) ([HN](https://news.ycombinator.com/item?id=30566093))
- [Shrek](https://github.com/innix/shrek) - Vanity .onion address generator written in Go.
- [ExitGather](https://github.com/uforia/exitgather) - Python script to generate a list of IP addresses being used as TOR and VPN exit nodes.
- [Tor Relay Availability Checker](https://github.com/ValdikSS/tor-relay-scanner) - For using it as a bridge in countries with censorship.
- [libtor-sys](https://github.com/MagicalBitcoin/libtor-sys) - Rust crate that builds a static version of Tor.
- [garlicshare](https://github.com/R4yGM/garlicshare) - Private and self-hosted file sharing over the Tor network written in Go.
- [Big Dark Web Links List](https://www.webhostingsecretrevealed.net/blog/security/dark-web-websites-onion-links/)
- [Tor Metrics](https://metrics.torproject.org/) ([HN](https://news.ycombinator.com/item?id=32240037))
- [OnionScan](https://github.com/s-rah/onionscan) - Free and open source tool for investigating the Dark Web. ([Twitter](https://twitter.com/OnionScan))
- [Tor is fighting and beating Russian censorship (2022)](https://www.wired.com/story/tor-browser-russia-blocks/) ([HN](https://news.ycombinator.com/item?id=32288058))

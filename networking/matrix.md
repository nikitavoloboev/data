# [Matrix](https://matrix.org)

Matrix is kind of like a federated IRC system and [Riot](https://about.riot.im) is a web and mobile interface to the Matrix system.

[Dendrite](https://github.com/matrix-org/dendrite) is Matrix homeserver written in Go. It is a rewrite of [synapse](https://github.com/matrix-org/synapse) that was written in Python which could not scale.

## Bridges

- [IRC](https://github.com/matrix-org/matrix-ircd)
- [WhatsApp](https://github.com/tulir/mautrix-whatsapp)

## Notes

- Dendrite: You don't need to run Kafka unless working on distributed stuff. It is composed of several microservices (room server, sync api, media api...). Kafka is used to transfer messages between those microservices.

## Links

- [What is Matrix?](https://matrix.org/docs/guides/faq.html#what-is-matrix)
- [Matrix doc](https://github.com/matrix-org/matrix-doc)
- [sytest](https://github.com/matrix-org/sytest) - Black-box integration testing for Matrix homeservers.
- [Unofficial list of public Matrix servers](https://www.hello-matrix.net/public_servers.php)
- [Matrix GSOC ideas](https://github.com/matrix-org/GSoC/blob/master/IDEAS.md#what-is-matrix)
- [Matrix IRCd](https://github.com/matrix-org/matrix-ircd)
- [Riot Web](https://github.com/vector-im/riot-web) - Web client for Matrix.
- [Ruma](https://github.com/ruma/homeserver) - Matrix homeserver written in Rust. Discontinued.
- [Client-Server API](https://matrix.org/docs/spec/client_server/r0.3.0.html)
- [Matrix Blog](https://matrix.org/blog/posts/)
- [The 2018 Matrix Holiday Special](https://matrix.org/blog/2018/12/25/the-2018-matrix-holiday-special/)
- [HN: The Matrix.org 2018 year in review](https://news.ycombinator.com/item?id=18756787)
- [Nio](https://github.com/niochat/nio) - Upcoming matrix client for iOS.
- [Dendrite design](https://github.com/matrix-org/dendrite/blob/master/docs/DESIGN.md)
- [Dendrite checklist](https://docs.google.com/spreadsheets/d/1tkMNpIpPjvuDJWjPFbw_xzNzOHBA-Hp50Rkpcr43xTw) - API to finish for release.
- [Running your own secure communication service with Matrix and Jitsi (2020)](https://matrix.org/blog/2020/04/06/running-your-own-secure-communication-service-with-matrix-and-jitsi) ([HN](https://news.ycombinator.com/item?id=22802645))
- [matrix-rust-sdk](https://github.com/matrix-org/matrix-rust-sdk) - Matrix Client-Server SDK for Rust.
- [Declarative, Decentralised, and Secure communication via Matrix, Jitsi, & NixOS (2019)](https://kaushikc.org/posts/matrix-jitsi-nixos.html)
- [mautrix-telegram](https://github.com/tulir/mautrix-telegram) - Matrix-Telegram hybrid puppeting/relaybot bridge.
- [Maelstrom](https://github.com/maelstrom-rs/maelstrom) - High performance Matrix Home Server in rust.
- [Matrix iOS SDK](https://github.com/matrix-org/matrix-ios-sdk)
- [Conduit](https://git.koesters.xyz/timo/conduit) - Matrix homeserver written in Rust.
- [Introducing P2P Matrix (2020)](https://matrix.org/blog/2020/06/02/introducing-p-2-p-matrix) ([HN](https://news.ycombinator.com/item?id=23393935))
- [Thoughts on Peer-to-Peer Matrix (2020)](https://neilalexander.dev/2020/06/02/thoughts-p2p-matrix.html)
- [Seaglass](https://github.com/neilalexander/seaglass) - Truly native Matrix client for macOS - written in Swift/Cocoa, with E2E encryption support.
- [NovaChat](https://nova.chat/) - Multi-Network Chat. ([HN](https://news.ycombinator.com/item?id=23693371))
- [HN: Matrix (2020)](https://news.ycombinator.com/item?id=24239564)
- [Gitter is joining Matrix (2020)](https://matrix.org/blog/2020/09/30/welcoming-gitter-to-matrix) ([HN](https://news.ycombinator.com/item?id=24638438))
- [Dendrite is entering Beta (2020)](https://matrix.org/blog/2020/10/08/dendrite-is-entering-beta) ([HN](https://news.ycombinator.com/item?id=24721160))
- [Free Small Matrix Server](https://matrix.org/docs/guides/free-small-matrix-server)
- [Nheko](https://github.com/Nheko-Reborn/nheko) - Desktop client for Matrix using Qt and C++17.
- [Combating abuse in Matrix – without E2EE backdoors (2020)](https://matrix.org/blog/2020/10/19/combating-abuse-in-matrix-without-backdoors) ([Lobsters](https://lobste.rs/s/ntyvtw/combating_abuse_matrix_without)) ([HN](https://news.ycombinator.com/item?id=24836987))
- [How we fixed Synapse's scalability (2020)](https://matrix.org/blog/2020/11/03/how-we-fixed-synapses-scalability) ([Lobsters](https://lobste.rs/s/gwwnei/how_we_fixed_synapse_s_scalability_matrix))
- [Using Matrix to replace proprietary and centralized chat apps (2020)](https://jae.moe/blog/2020/11/using-matrix-to-replace-proprietary-and-centralized-chat-apps/) ([HN](https://news.ycombinator.com/item?id=25091614))
- [Free Small Matrix Server](https://github.com/ptman/matrix-docs/tree/master/free-matrix-server)
- [libQuotient](https://github.com/quotient-im/libQuotient/) - Qt5 library to write cross-platform clients for Matrix.
- [Quaternion](https://github.com/quotient-im/Quaternion/) - Qt5-based IM client for Matrix.

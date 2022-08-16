---
title: Matrix
---

# [Matrix](https://matrix.org)

Matrix is kind of like a federated IRC system and [Element](https://element.io) is a web and mobile interface to the Matrix system.

[Dendrite](https://github.com/matrix-org/dendrite) is Matrix homeserver written in Go. It is a rewrite of [synapse](https://github.com/matrix-org/synapse) that is written in Python that has issues scaling.

## Bridges

- [IRC](https://github.com/matrix-org/matrix-ircd)
- [WhatsApp](https://github.com/tulir/mautrix-whatsapp)
- [matrix-hookshot](https://github.com/Half-Shot/matrix-hookshot) - Bridge between Matrix and multiple project management services, such as GitHub, GitLab and JIRA.

## Bots

- [hebbot](https://github.com/haecker-felix/hebbot) - Matrix bot which can generate "This Week in X" like blog posts.
- [maubot](https://github.com/maubot/maubot) - Plugin-based Matrix bot system.

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
- [mautrix-telegram](https://github.com/mautrix/telegram) - Matrix-Telegram hybrid puppeting/relaybot bridge.
- [Maelstrom](https://github.com/maelstrom-rs/maelstrom) - High performance Matrix Home Server in rust.
- [Matrix iOS SDK](https://github.com/matrix-org/matrix-ios-sdk)
- [Conduit](https://conduit.rs/) - Simple, fast and reliable chat server powered by Matrix. Written in Rust. ([Code](https://gitlab.com/famedly/conduit))
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
- [The Matrix Holiday Special 2020](https://matrix.org/blog/2020/12/25/the-matrix-holiday-special-2020) ([Lobsters](https://lobste.rs/s/sxt5zc/matrix_holiday_special_2020))
- [Cerulean](https://github.com/matrix-org/cerulean) - Experimental Matrix client intended to demonstrate the viability of freestyle public threaded conversations a la Twitter.
- [retrix](https://github.com/agraven/retrix) - Lightweight matrix client.
- [Building a virtual coffee bot for Mattermost and Matrix (2021)](https://n8n.io/blog/how-to-host-virtual-coffee-breaks-with-n8n/)
- [Element Matrix Client Code](https://github.com/vector-im/element-web)
- [Matrix (An open network for secure, decentralized communication) server setup using Ansible and Docker](https://github.com/spantaleev/matrix-docker-ansible-deploy)
- [Message Retention Policies in Matrix (2021)](https://brendan.abolivier.bzh/matrix-retention-policies/)
- [FOSDEM 2021: Building massive virtual communities in Matrix](https://www.youtube.com/watch?v=TzUfS08lMek)
- [How we hosted FOSDEM 2021 on Matrix (2021)](https://matrix.org/blog/2021/02/15/how-we-hosted-fosdem-2021-on-matrix) ([HN](https://news.ycombinator.com/item?id=26142654))
- [Hummingbard](https://hummingbard.com/hummingbard/introducing-hummingbard) - Decentralized communities built on Matrix. ([HN](https://news.ycombinator.com/item?id=26277602)) ([Code](https://github.com/hummingbard/hummingbard))
- [Element Home](https://element.io/element-home) - Element app, but faster, personalized and under your control. ([Article](https://element.io/blog/element-home/)) ([HN](https://news.ycombinator.com/item?id=26311801))
- [Pinecone](https://github.com/matrix-org/pinecone) - Peer-to-peer overlay routing for the Matrix ecosystem. ([Article](https://matrix.org/blog/2021/05/06/introducing-the-pinecone-overlay-network)) ([HN](https://news.ycombinator.com/item?id=27077660))
- [Element raises $30M to boost Matrix (2021)](https://matrix.org/blog/2021/07/27/element-raises-30-m-to-boost-matrix) ([HN](https://news.ycombinator.com/item?id=27969624))
- [Matrix JavaScript SDK](https://github.com/matrix-org/matrix-js-sdk)
- [Cinny](https://cinny.in/) - Web-based matrix client. ([Code](https://github.com/ajbura/cinny)) ([HN](https://news.ycombinator.com/item?id=27986376)) ([Web Code](https://github.com/cinnyapp/cinny-site))
- [TheBoard](https://github.com/toger5/TheBoard) - Collaborative Whiteboard powered by the [matrix] protocol and infrastructure.
- [matrix-media-repo](https://github.com/turt2live/matrix-media-repo) - Matrix media repository with multi-domain in mind.
- [Sygnal](https://github.com/matrix-org/sygnal) - Reference Push Gateway for Matrix.
- [Public Rooms - Matrix Static](https://view.matrix.org/) - Static Go generated preview of public world readable Matrix rooms. ([Code](https://github.com/matrix-org/matrix-static))
- [Gotrix](https://github.com/chanbakjsd/gotrix) - Implementation of the client portion of Matrix's client-server API.
- [Element One](https://ems.element.io/element-one) - All of Matrix, WhatsApp, Signal and Telegram in one place. ([Article](https://element.io/blog/element-one-all-of-matrix-whatsapp-signal-and-telegram-in-one-place/)) ([HN](https://news.ycombinator.com/item?id=28997898))
- [Is Matrix less private than other E2EE messaging protocols? (2019)](https://gitlab.com/libremonde-org/papers/research/privacy-matrix.org/-/blob/master/part1/README.md) ([HN](https://news.ycombinator.com/item?id=29105098))
- [Hemppa](https://github.com/vranki/hemppa) - Generic modular bot for Matrix.
- [Matrix Discord Bridge](https://github.com/Half-Shot/matrix-appservice-discord)
- [Sign in with Matrix](https://github.com/mishushakov/signin-with-matrix) - Federated sign-in component for your web app (using Matrix). ([HN](https://news.ycombinator.com/item?id=29240154))
- [OCaml matrix](https://github.com/clecat/ocaml-matrix) - Implementation of a matrix server in OCaml for MirageOS.
- [Awesome Matrix](https://github.com/jryans/awesome-matrix)
- [Matrix Corporal](https://github.com/devture/matrix-corporal) - Reconciliator and gateway for a managed Matrix server.
- [Matrix CRDT](https://github.com/YousefED/Matrix-CRDT) - Enables you to use Matrix as a backend for distributed, real-time collaborative web applications that sync automatically. ([HN](https://news.ycombinator.com/item?id=29978659))
- [Third Room](https://github.com/matrix-org/thirdroom) - Decentralized metaverse platform built on top of the Matrix protocol. ([Intro](https://github.com/matrix-org/thirdroom/discussions/20))
- [mautrix-go](https://github.com/mautrix/go) - Go Matrix framework.
- [matrix-oauth](https://github.com/turt2live/matrix-oauth) - OAuth 2.0 server for Matrix accounts.
- [Native Matrix VoIP with Element Call (2022)](https://element.io/blog/introducing-native-matrix-voip-with-element-call/) ([HN](https://news.ycombinator.com/item?id=30568164))
- [mx-puppet-slack](https://github.com/Sorunome/mx-puppet-slack) - Slack puppeting bridge for Matrix.
- [Matrix Video Chat](https://call.element.io/)
- [Element Call for macOS](https://github.com/pixlwave/Element-Call-macOS) - Wrapper app for Element Call that adds support for Camera & Microphone toggling via keyboard.
- [Matrix Highlight](https://github.com/DanilaFe/matrix-highlight) - Decentralized and federated way of annotating the web based on Matrix.
- [Playing with Matrix: Conduit and Synapse (2022)](https://akselmo.dev/2022/04/01/Playing-with-Matrix-Conduit-and-Synapse.html)
- [Matrix.to](https://matrix.to/) - Simple stateless privacy-protecting URL redirecting service for Matrix. ([Code](https://github.com/matrix-org/matrix.to))
- [Element X iOS](https://github.com/vector-im/element-x-ios) - Matrix client provided by Element and based on the Matrix Rust SDK.
- [Matrix React SDK](https://github.com/matrix-org/matrix-react-sdk)
- [Mautrix](https://github.com/mautrix/python) - Python 3 asyncio Matrix framework.
- [OAuth 2.0 authentication server for Matrix](https://github.com/matrix-org/matrix-authentication-service)
- [Matrix protocol spec](https://github.com/matrix-org/matrix-spec)
- [Quadrix](https://github.com/alariej/quadrix) - Minimal, simple, multi-platform chat client for the Matrix protocol.
- [Matrix notes (2022)](https://anarc.at/blog/2022-06-17-matrix-notes/)
- [matrix-commander](https://github.com/8go/matrix-commander) - Simple but convenient CLI-based Matrix client app for sending and receiving.
- [Threematrix](https://github.com/bitbetterde/Threematrix) - Messenger bridge between Threema and Matrix.
- [HG HomeServer](https://github.com/heusalagroup/hghs) - Minimal Matrix HomeServer written in TypeScript.
- [Matrix public archive](https://github.com/matrix-org/matrix-public-archive) - View the history of world readable Matrix rooms.
- [The Matrix Summer Special 2022](https://matrix.org/blog/2022/08/15/the-matrix-summer-special-2022)

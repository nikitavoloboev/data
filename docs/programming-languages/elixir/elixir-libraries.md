# Elixir libraries

## Web

- [Mint](https://github.com/elixir-mint/mint) - Functional HTTP client for Elixir with support for HTTP/1 and HTTP/2.
- [Mint.WebSocket](https://github.com/elixir-mint/mint_web_socket) - HTTP/1 and HTTP/2 WebSocket support for Mint.
- [Req](https://github.com/wojtekmach/req) - HTTP client with a focus on ease of use and composability, built on top of Finch.
- [Finch](https://github.com/keathley/finch) - HTTP client with a focus on performance, built on top of Mint and NimblePool.
- [Ash](https://github.com/ash-project/ash) - Resource based framework for building Elixir applications.
- [SiteEncrypt](https://github.com/sasa1977/site_encrypt) - Integrated certification via Let's encrypt for Elixir-powered sites.
- [ExVCR](https://github.com/parroty/exvcr) - Record and replay HTTP interactions library for Elixir.
- [WebSockex](https://github.com/Azolo/websockex) - Elixir Websocket Client.
- [Tesla](https://github.com/teamon/tesla) - Flexible HTTP client library for Elixir, with support for middleware and multiple adapters.
- [Mojito](https://github.com/appcues/mojito) - Easy-to-use Elixir HTTP client, built on the low-level Mint library.
- [elixir-auth-google](https://github.com/dwyl/elixir-auth-google) - Minimalist Google OAuth Authentication for Elixir Apps.
- [Ace](https://github.com/CrowdHailer/Ace) - HTTP web server and client, supports http1 and http2.
- [Raxx](https://github.com/CrowdHailer/raxx) - Interface for HTTP webservers, frameworks and clients.
- [Corsica](https://github.com/whatyouhide/corsica) - Elixir library for dealing with CORS requests.
- [Plug](https://github.com/elixir-plug/plug) - Specification and conveniences for composable modules between web applications.
- [Open API Spex](https://github.com/open-api-spex/open_api_spex) - Open API Specifications for Elixir Plug applications.
- [Quenya](https://github.com/tyrchen/quenya) - Framework to build high-quality REST API applications based on extended OpenAPI spec.
- [Bandit](https://github.com/mtrudel/bandit) - HTTP server for Plug apps. ([Article](https://mat.geeky.net/2021/06/24/bandit-on-the-loose.html)) ([Talk](https://github.com/mtrudel/talks/blob/main/ElixirConf%202021.pdf))
- [ElixirST](https://github.com/gertab/ElixirST) - Session Types in Elixir.
- [PardallMarkdown](https://github.com/alfredbaudisch/pardall_markdown) - Reactive publishing framework, filesystem-based with support for Markdown, nested hierarchies, and instant content rebuilding. Written in Elixir.
- [OAuth2](https://github.com/scrogson/oauth2) - Elixir OAuth 2.0 Client Library.
- [Joken](https://github.com/joken-elixir/joken) - Elixir JWT library.
- [Ueberauth](https://github.com/ueberauth/ueberauth) - Elixir Authentication System for Plug-based Web Applications.
- [Temple](https://github.com/mhanberg/temple) - HTML DSL for Elixir and Phoenix.
- [NimbleTOTP](https://github.com/dashbitco/nimble_totp) - Tiny Elixir library for time-based one time passwords (TOTP).
- [Canary](https://github.com/cpjk/canary) - Elixir authorization and resource-loading library for Plug applications.
- [Riverside](https://github.com/lyokato/riverside) - Plain WebSocket Server Framework for Elixir.
- [ElixirRetry](https://github.com/safwank/ElixirRetry) - Simple Elixir macros for linear retry, exponential backoff and wait with composable delays.
- [Webhoox](https://github.com/cpursley/webhoox) - Deal with inbound webhooks.
- [Hologram](https://github.com/bartblast/hologram) - Full stack isomorphic Elixir web framework that can be used on top of Phoenix.
- [EasyHTML](https://github.com/wojtekmach/easyhtml) - Makes it easy to work with HTML in Elixir.
- [Riptide](https://github.com/ironbay/riptide) - Elixir framework for creating real time applications.
- [LetMe](https://github.com/woylie/let_me) - Authorization library for Elixir.

## DB

- [Ecto](https://github.com/elixir-ecto/ecto) - Toolkit for data mapping and language integrated query for Elixir. ([Brief Guide to Ecto](https://serokell.io/blog/ecto-guide-for-beginners)) ([Tweet](https://twitter.com/josevalim/status/1489522409426325504))
- [TypedEctoSchema](https://github.com/bamorim/typed_ecto_schema) - Library to define Ecto schemas with typespecs without all the boilerplate code.
- [Rummage.Ecto](https://github.com/annkissam/rummage_ecto) - Search, Sort and Pagination for ecto queries.
- [Paper Trail](https://github.com/izelnakri/paper_trail) - Track and record all the changes in your database with Ecto. Revert back to anytime in history.
- [EctoMorph](https://github.com/Adzz/ecto_morph) - Morphs your Ecto capabilities.
- [DBConnection](https://github.com/elixir-ecto/db_connection) - Database connection behavior.
- [Hierarch](https://github.com/Byzanteam-Labs/hierarch) - Tree structure & hierarchy for ecto models with ltree (Postgres).
- [Excellent Migrations](https://github.com/Artur-Sulej/excellent_migrations) - Detect potentially dangerous or destructive operations in your database migrations.
- [Params](https://github.com/vic/params) - Easily define parameter structure and validate/cast with Ecto.Schema.
- [Postgrex](https://github.com/elixir-ecto/postgrex) - PostgreSQL driver for Elixir.
- [CubDB](https://github.com/lucaong/cubdb) - Embedded key-value database written in the Elixir language.
- [Elixir driver for EdgeDB](https://github.com/nsidnev/edgedb-elixir)
- [Khepri](https://github.com/rabbitmq/khepri) - Tree-like replicated on-disk database library for Erlang and Elixir.
- [Ecto SQLite3 adapter](https://github.com/elixir-sqlite/ecto_sqlite3)
- [Honeydew](https://github.com/koudelka/honeydew) - Job Queue for Elixir. Clustered or Local. Straight BEAM. Optional Ecto.
- [EctoJob](https://github.com/mbuhot/ecto_job) - Transactional job queue built with Ecto and GenStage.
- [TaskBunny](https://github.com/shinyscorpion/task_bunny) - Background processing application written in Elixir and uses RabbitMQ as a messaging backend.
- [litestream](https://github.com/akoutmos/litestream) - Wraps Litestream so that SQLite databases can be easily backed up.
- [Paper Trail](https://github.com/izelnakri/paper_trail) - Track and record all the changes in your database with Ecto. Revert back to anytime in history.
- [ExAudit](https://github.com/ZennerIoT/ex_audit) - Ecto auditing library that transparently tracks changes and can revert them.
- [DataSchema](https://github.com/Adzz/data_schema) - Declarative schemas for data transformations.
- [Fields](https://github.com/dwyl/fields) - Collection of commonly used fields implemented as custom Ecto types with validation, sanitizing and encryption/hashing.
- [Spear](https://github.com/NFIBrokerage/spear) - Sharp EventStoreDB 20+ client backed by Mint.
- [Yesql](https://github.com/lpil/yesql) - Elixir library for using SQL.
- [Etso](https://github.com/evadne/etso) - Ecto 3 adapter allowing use of Ecto schemas held in ETS tables.
- [Triplex](https://github.com/ateliware/triplex) - Simple and effective way to build multitenant applications on top of Ecto.
- [Xandra](https://github.com/lexhide/xandra) - Fast, simple, and robust Cassandra driver for Elixir.
- [Pointers](https://github.com/bonfire-networks/pointers) - One foreign key to rule them all and in the darkness, bind them.

## Test

- [Mox](https://github.com/dashbitco/mox) - Library for defining concurrent mocks in Elixir.
- [Hammox](https://github.com/msz/hammox) - Automated contract testing for Elixir functions and mocks.
- [PropCheck](https://github.com/alfert/propcheck) - Property based testing for Elixir.
- [Mock](https://github.com/jjh42/mock) - Mocking library for Elixir language.
- [ExMachina](https://github.com/thoughtbot/ex_machina) - Create test data for Elixir applications.
- [Patch](https://github.com/ihumanable/patch) - Ergonomic Mocking for Elixir.
- [StreamData](https://github.com/whatyouhide/stream_data) - Elixir library for data generation and property-based testing.
- [Machete](https://github.com/mtrudel/machete) - Literate test matchers for ExUnit.

## ML

- [Axon](https://github.com/elixir-nx/axon) - Nx-powered Neural Networks. ([HN](https://news.ycombinator.com/item?id=26737956))
- [Scholar](https://github.com/elixir-nx/scholar) - Machine learning tools built on top of Nx.

## CLI

- [Owl](https://github.com/fuelen/owl) - Toolkit for writing command-line user interfaces.
- [Vexil](https://github.com/Ovyerus/vexil) - Elixir flag parser that does just enough.
- [Artificery](https://github.com/bitwalker/artificery) - Toolkit for creating terminal user interfaces in Elixir.
- [ExTermbox](https://github.com/ndreynolds/ex_termbox) - Low-level termbox bindings for Elixir.
- [ExCLI](https://github.com/danhper/ex_cli) - User friendly CLI apps for Elixir.
- [CliSpinners](https://github.com/blackode/elixir_cli_spinners) - Spinnig Animations for Command Line Applications.

## Math

- [Statistics](https://github.com/msharp/elixir-statistics) - Statistical functions and distributions for Elixir.

## Livebook

- [Kino](https://github.com/livebook-dev/kino) - Client-driven interactive widgets for Livebook.

## Crypto

- [Cloak](https://github.com/danielberkompas/cloak) - Elixir encryption library designed for Ecto.

## GraphQL

- [BridgeEx](https://github.com/primait/bridge_ex) - Library to build bridges to GraphQL services.
- [AbsintheCache](https://github.com/IvanIvanoff/absinthe_cache) - Caching Absinthe GraphQL queries.
- [CommonGraphqlClient](https://github.com/annkissam/common_graphql_client) - Elixir GraphQL Client with HTTP and WebSocket Support.
- [GQL](https://github.com/balexand/gql) - Simple GraphQL client for Elixir.

## Other

- [GraphQL toolkit for Elixir](https://github.com/absinthe-graphql/absinthe)
- [Sage](https://github.com/Nebo15/sage) - Dependency-free tool to run distributed transactions in Elixir, inspired by Sagas pattern.
- [Timex](https://github.com/bitwalker/timex) - Complete date/time library for Elixir projects.
- [Core Scenic Library](https://github.com/boydm/scenic) - Client application library written directly on the Elixir/Erlang/OTP stack. With it you can build applications that operate identically across all supported operating systems.
- [Benchee](https://github.com/bencheeorg/benchee) - Easy and extensible benchmarking in Elixir providing you with lots of statistics.
- [Opus](https://github.com/zorbash/opus) - Framework for pluggable business logic components.
- [Pow](https://github.com/danschultzer/pow) - Robust, modular, and extendable user authentication system.
- [Rihanna](https://github.com/samphilipd/rihanna) - High performance postgres-backed job queue for Elixir.
- [Broadway](https://github.com/plataformatec/broadway) - Concurrent and multi-stage data ingestion and data processing with Elixir.
- [Goth](https://github.com/peburrows/goth) - Elixir package for Oauth authentication via Google Cloud APIs.
- [Timber](https://github.com/timberio/timber-elixir) - Great Elixir logging made easy.
- [Parent](https://github.com/sasa1977/parent) - Support for custom parenting of processes.
- [Nebulex](https://github.com/cabol/nebulex) - In-Process and Distributed Cache Toolkit for Elixir.
- [FunWithFlags](https://github.com/tompave/fun_with_flags) - Elixir feature flag library.
- [Scientist](https://github.com/cwbriones/scientist) - Library for carefully refactoring critical paths in your elixir application.
- [libgraph](https://github.com/bitwalker/libgraph) - Graph data structure library for Elixir projects.
- [GenStage](https://github.com/elixir-lang/gen_stage) - Producer and consumer actors with back-pressure for Elixir.
- [Elixir RakNet](https://github.com/X-Plane/elixir-raknet) - Elixir implementation of the RakNet/RakLib networking communication protocol.
- [CI](https://github.com/sasa1977/ci) - CI/CD toolkit as an Elixir library.
- [Waffle](https://github.com/elixir-waffle/waffle) - Flexible file upload and attachment library for Elixir.
- [Oban](https://github.com/sorentwo/oban) - Robust job processing in Elixir, backed by modern PostgreSQL. ([Web](https://getoban.pro/))
- [Nx](https://github.com/elixir-nx/nx) - Multi-dimensional tensors Elixir lib with multi-staged compilation (CPU/GPU). ([Article](https://dashbit.co/blog/nx-numerical-elixir-is-now-publicly-available)) ([HN](https://news.ycombinator.com/item?id=26168304)) ([Lobsters](https://lobste.rs/s/qzwqdu/nx_numerical_elixir_is_now_publicly)) ([Introducing Nx Talk](https://www.youtube.com/watch?v=fPKMmJpAGWc)) ([HN](https://news.ycombinator.com/item?id=29822849))
- [Boundary](https://github.com/sasa1977/boundary) - Manage and restrain cross-module dependencies in Elixir projects.
- [Poison](https://github.com/devinus/poison) - Incredibly fast, pure Elixir JSON library.
- [Exq](https://github.com/akira/exq) - Job processing library for Elixir - compatible with Resque / Sidekiq.
- [Twirp](https://github.com/keathley/twirp-elixir) - Elixir implementation of the twirp RPC framework.
- [libcluster](https://github.com/bitwalker/libcluster) - Automatic cluster formation/healing for Elixir applications.
- [Horde](https://github.com/derekkraan/horde) - Distributed Supervisor and Registry backed by DeltaCrdt.
- [libring](https://github.com/bitwalker/libring) - Fast consistent hash ring for Elixir.
- [Aja](https://github.com/sabiwara/aja) - Extension of the Elixir standard library focused on data structures, data manipulation and performance.
- [Schoolbus](https://github.com/rupurt/schoolbus) - Manage one or more PubSub instances using the Elixir registry.
- [Msgpax](https://github.com/lexmag/msgpax) - Fast and comprehensive library for serializing and deserializing Elixir terms using the MessagePack format.
- [Manifold](https://github.com/discord/manifold) - Fast batch message passing between nodes for Erlang/Elixir.
- [FastGlobal](https://github.com/discord/fastglobal) - Fast no copy globals for Elixir & Erlang.
- [Wasmex](https://github.com/tessi/wasmex) - Execute WebAssembly / WASM from Elixir.
- [Recase](https://github.com/sobolevn/recase) - Convert strings to any case.
- [Regulator](https://github.com/keathley/regulator) - Adaptive concurrency limits around external resources.
- [Delta](https://github.com/slab/delta-elixir) - Simple yet expressive format to describe contents and changes.
- [Instruments](https://github.com/discord/instruments) - Simple and Fast metrics for Elixir.
- [Para](https://github.com/syamilmj/para) - Declarative way to parse and validate parameters.
- [Norm](https://github.com/keathley/norm) - Data specification and generation.
- [Algae](https://github.com/witchcrafters/algae) - Bootstrapped algebraic data types for Elixir.
- [VegaLite](https://github.com/livebook-dev/vega_lite) - Elixir bindings for Vega-Lite.
- [Elixir Thrift](https://github.com/pinterest/elixir-thrift)
- [Uncharted](https://github.com/unchartedelixir/uncharted) - Simple Elixir charting library that generates easy to customize charts.
- [Rustic Result](https://github.com/linkdd/rustic_result) - Result monad for Elixir inspired by Rust Result type. ([HN](https://news.ycombinator.com/item?id=28913884))
- [Explorer](https://github.com/elixir-nx/explorer) - Dataframes for Elixir. ([Intro](https://www.cigrainger.com/posts/introducing-explorer/))
- [Burrito](https://github.com/burrito-elixir/burrito) - Cross-Platform Elixir Deployments.
- [ex_cldr](https://github.com/elixir-cldr/cldr) - Elixir implementation of CLDR/ICU.
- [TypedStruct](https://github.com/ejpcmac/typed_struct) - Library for defining structs with a type without writing boilerplate code.
- [CSV](https://github.com/beatrichartz/csv) - CSV Decoding and Encoding for Elixir.
- [Daat](https://github.com/QuinnWilton/daat) - Parameterized modules for Elixir.
- [Expat](https://github.com/vic/expat) - Reusable, composable patterns across Elixir libraries.
- [Vex](https://github.com/CargoSense/vex) - Data Validation for Elixir.
- [NimbleParsec](https://github.com/dashbitco/nimble_parsec) - Simple and fast library for text-based parser combinators.
- [Elasticlunr](https://atandarash.me/blog/introduction-to-elasticlunr) - Full-text search library for Elixir. ([HN](https://news.ycombinator.com/item?id=29857712))
- [Witchcraft](https://github.com/witchcrafters/witchcraft) - Monads and other dark magic for Elixir. ([Article](https://blog.appsignal.com/2022/02/08/functional-programming-in-elixir-with-witchcraft.html)) ([HN](https://news.ycombinator.com/item?id=30257542)) ([Lobsters](https://lobste.rs/s/iaulch/functional_programming_elixir_with))
- [Domo](https://github.com/IvanRublev/Domo) - Library to validate values of nested structs with their type spec t() and associated precondition functions.
- [Ink](https://github.com/ivx/ink) - Logger backend that logs JSON.
- [MuonTrap](https://github.com/fhunleth/muontrap) - Keep programs, deamons, and applications launched from Erlang and Elixir contained and well-behaved.
- [StreamSplit](https://github.com/tallakt/stream_split) - Split a stream into a head and tail, without iterating the tail.
- [Unicode Set](https://github.com/elixir-unicode/unicode_set) - Unicode set parsing, expansion, macros and guards for Elixir.
- [CAStore](https://github.com/elixir-mint/castore) - Up-to-date certificate store for Elixir.
- [NimblePool](https://github.com/dashbitco/nimble_pool) - Tiny resource-pool implementation for Elixir.
- [SSHKit](https://github.com/bitcrowd/sshkit.ex) - Elixir toolkit for performing tasks on one or more servers, built on top of Erlang’s SSH application.
- [Quark](https://github.com/witchcrafters/quark) - Common combinators for Elixir.
- [Protox](https://github.com/ahamez/protox) - Fast, easy to use and 100% conformant Elixir library for Google Protocol Buffers.
- [Flow](https://github.com/dashbitco/flow) - Computational parallel flows on top of GenStage. ([Intro to Elixir Flow](https://blog.joelabshier.com/%F0%9F%8C%8A-a-realworld-introduction-to-elixir-flow/))
- [Jsonrs](https://github.com/benhaney/Jsonrs) - Rust powered JSON library for Elixir.
- [Semaphore](https://github.com/discord/semaphore) - Fast semaphore using ETS.
- [MJML](https://github.com/adoptoposs/mjml_nif) - Elixir NIF bindings for the MJML Rust implementation.
- [gRPC Elixir](https://github.com/elixir-grpc/grpc) - Elixir implementation of gRPC.
- [Focus](https://github.com/smpoulsen/focus) - Lightweight, pure Elixir functional optics.
- [Envy](https://github.com/BlakeWilliams/envy) - Elixir library to load environment variables from .env and environment specific env files.
- [ExWaiter](https://github.com/baldwindavid/ex_waiter) - Helper for waiting on asynchronous conditions to be met.
- [Rollbax](https://github.com/ForzaElixir/rollbax) - Exception tracking and logging from Elixir to Rollbar.
- [Xray](https://github.com/fireproofsocks/xray) - Academic exploration of strings and encodings in Elixir.
- [PromEx](https://github.com/akoutmos/prom_ex) - Elixir Prometheus metrics collection library built on top of Telemetry with accompanying Grafana dashboards.
- [Grizzly](https://github.com/smartrent/grizzly) - Elixir library for Z-Wave.
- [Brex.Result](https://github.com/brexhq/result) - Provides tools to handle three common return values in Elixir.
- [Combination](https://github.com/seantanly/elixir-combination) - Simple combinatorics library providing combination and permutation.
- [Mogrify](https://github.com/elixir-mogrify/mogrify) - Image processing in Elixir.
- [Jetstream](https://github.com/mmmries/jetstream) - Elixir client for NATS Jetstream.
- [ExImageInfo](https://github.com/Group4Layers/ex_image_info) - Elixir library to parse images (binaries) and get the dimensions (size), detected mime-type.
- [NifIo](https://github.com/rusterlium/NifIo) - Rustler example. Implements file IO in a NIF.
- [Hush](https://github.com/gordalina/hush) - Runtime configuration loader for Elixir applications.
- [Stripe for Elixir](https://github.com/beam-community/stripity_stripe)
- [Solid](https://github.com/edgurgel/solid) - Liquid template engine in Elixir.
- [MerkleTree](https://github.com/yosriady/merkle_tree) - Merkle Tree implementation in pure Elixir.
- [Entity Component System](https://github.com/yosriady/ecs)
- [Blex](https://github.com/gyson/blex) - Fast Bloom filter with concurrent accessibility, powered by :atomics module.
- [Talos](https://github.com/balance-platform/talos) - Elixir parameter validation library. Simple and extensible.
- [Faker](https://github.com/elixirs/faker) - Pure Elixir library for generating fake data.
- [Tarearbol](https://github.com/am-kantox/tarearbol) - Lightweight task manager, allowing retries, callbacks, assurance that the task succeeded, and more.
- [Telemetria](https://github.com/am-kantox/telemetria) - Helper application that simplifies and standardizes telemetry usage.
- [Swarm](https://github.com/bitwalker/swarm) - Easy clustering, registration, and distribution of worker processes for Erlang/Elixir.
- [Decimal](https://github.com/ericmj/decimal) - Arbitrary precision decimal arithmetic.
- [Ravenx](https://github.com/acutario/ravenx) - Notification dispatch library for Elixir applications.
- [Notifiex](https://github.com/burntcarrot/notifiex) - Simple Elixir library for sending notifications to various messaging services.
- [Qrusty](https://github.com/nbw/qrusty) - Elixir QR Code library that leverages precompiled Rust.
- [Lens](https://github.com/obrok/lens) - Utility for working with nested data structures.
- [MixInstallExamples](https://github.com/wojtekmach/mix_install_examples) - Collection of simple Elixir scripts that are using Mix.install/2.
- [JsonDiffEx](https://github.com/olafura/json_diff_ex) - Diff and patch for JSON in Elixir, works really well and is really fast.
- [Pexelmatch](https://github.com/user-docs/pexelmatch) - Pure Elixir port of Pixelmatch.
- [Honey Potion](https://github.com/lac-dcc/honey-potion) - Writing eBPF with Elixir.
- [Sippet](https://github.com/balena/elixir-sippet) - Elixir library designed to be used as SIP protocol middleware.
- [Vapor](https://github.com/elixir-toniq/vapor) - Runtime configuration system for Elixir.
- [NimbleOptions](https://github.com/dashbitco/nimble_options) - Tiny library for validating and documenting high-level options.
- [Scribe](https://github.com/codedge-llc/scribe) - Pretty print tables of Elixir structs and maps.
- [Queuetopia](https://github.com/annatel/queuetopia) - Persistent blocking job queue built with Ecto.
- [GenRegistry](https://github.com/discord/gen_registry) - Simple and efficient local Process Registry.
- [Rambo](https://github.com/jayjun/rambo) - Run your command. Send input. Get output.
- [Rexbug](https://github.com/nietaki/rexbug) - Thin Elixir wrapper for the redbug Erlang tracing debugger.
- [Mux Elixir](https://github.com/muxinc/mux-elixir) - Official Mux API wrapper for Elixir projects, supporting both Mux Data and Mux Video.
- [StreamGzip](https://github.com/ne-sachirou/stream_gzip) - Gzip or gunzip an Elixir stream.

## Links

- [Hex packages](https://hex.pm/packages) - Package manager for the Erlang ecosystem.
- [ExDoc](https://github.com/elixir-lang/ex_doc) - Tool to generate documentation for your Elixir projects. ([Tweet](https://twitter.com/fhunleth/status/1460430228271112196))
- [Log all the things - features of Elixir 1.11 Logger application (2022)](https://hauleth.dev/post/log-all-the-things/) ([Lobsters](https://lobste.rs/s/i29sqw/log_all_things_features_elixir_1_11_logger))
- [Introduction to Telemetry in Elixir (2022)](https://blog.miguelcoba.com/introduction-to-telemetry-in-elixir) ([Lobsters](https://lobste.rs/s/ywwxdj/introduction_telemetry_elixir))
- [A Guide to Secure Elixir Package Updates (2022)](https://blog.appsignal.com/2022/03/15/a-guide-to-secure-elixir-package-updates.html)
- [How EEx Turns Your Template Into HTML (2022)](https://www.mitchellhanberg.com/how-eex-turns-your-template-into-html/)
- [Elixir runtime configuration trap (2022)](https://michaldolata.appunite.com/post/elixir-runtime-configuration-trap)
- [Lettuce](https://github.com/gilacost/lettuce) - Code watcher for elixir projects.
- [The Curious Case of String.slice/3 (2022)](https://adworse.codes/the-curious-case-of-string-slice-3/)

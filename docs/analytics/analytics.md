# Analytics

Analytics if implemented should be actionable and in no way impact the performance of the app. i.e [Netlify](https://www.netlify.com/products/analytics/) or [Google Analytics](https://developers.google.com/tag-platform/tag-manager/server-side/intro) but done on server side and not on the client.

I try to not waste time and even block pages like Insights Traffic pane on GitHub so as to not waste time viewing useless vanity stats.

Well implemented analytics together with proper logging/tracing should let you let you view user journeys of your app, which parts of the app are barely used and see how you can improve.

[PostHog](https://github.com/PostHog/posthog) & [umami](https://umami.is/) seem nice for analytics. [Product Analytics from Scratch](https://tmfarrell.github.io/writing/2022/04/26/product_analytics_from_scratch/) is a nice read. [Databend](https://github.com/datafuselabs/databend) seems great for more serious data processing needs.

## Notes

- [Just because you can measure quantitative things doesn't mean you can ignore qualitative things.](https://twitter.com/_kamlesh_/status/1451907162867847179)

## Links

- [Fathom](https://usefathom.com/) - Simple, trustworthy website analytics. Built with Go & Preact. ([Code](https://github.com/usefathom/fathom)) ([Best option for GDPR compliance](https://usefathom.com/features/eu-isolation))
- [Matamo](https://matomo.org/) - Open source alternative to Google Analytics.
- [LocustDB](https://github.com/cswinter/LocustDB) - Massively parallel, high performance analytics database that will rapidly devour all of your data.
- [Grafana](https://github.com/grafana/grafana) - Tool for beautiful monitoring and metric analytics & dashboards for Graphite, InfluxDB & Prometheus & More. ([Grafana Backup Tool](https://github.com/ysde/grafana-backup-tool))
- [Cube.js](https://github.com/statsbotco/cube.js) - Open source modular framework to build analytical web applications.
- [Analytics.js](https://github.com/segmentio/analytics.js) - Hassle-free way to integrate analytics into any web application.
- [Kedro](https://github.com/kedro-org/kedro) - Python library for building robust production-ready data and analytics pipelines. ([Article](https://medium.com/@QuantumBlack/introducing-kedro-the-open-source-library-for-production-ready-machine-learning-code-d1c6d26ce2cf))
- [Plausible Analytics](https://plausible.io/) - Simple, privacy-friendly alternative to Google Analytics. ([Code](https://github.com/plausible/analytics)) ([HN](https://news.ycombinator.com/item?id=24696145)) ([HN 2](https://news.ycombinator.com/item?id=24868012)) ([Docs](https://docs.plausible.io/)) ([Lobsters](https://lobste.rs/s/xksjh5/plausible_analytics_self_hosted_privacy))
- [Metabase](https://www.metabase.com/) - Easy, open source way for everyone in your company to ask questions and learn from data. ([Code](https://github.com/metabase/metabase)) ([Binary deployment](https://github.com/metabase/metabase-deploy))
- [OmniSci](https://www.omnisci.com/) - Interactively query, visualize, and power data science workflows over billions of records. ([1.1 Billion Taxi Rides using OmniSciDB and a MacBook Pro](https://tech.marksblogg.com/omnisci-macos-macbookpro-mbp.html))
- [Freshlytics](https://github.com/sheshbabu/freshlytics) - Open source privacy-friendly analytics software. It aims to be reliable, friendly to use and easy to deploy.
- [GoatCounter](https://www.goatcounter.com/) - Simple web statistics. No tracking of personal data. ([Code](https://github.com/zgoat/goatcounter))
- [MixPanel](https://mixpanel.com/) - Analyze user behavior across your sites and apps. Then send messages and run experiments from what you learned–all in Mixpanel.
- [What is your preferred web traffic analytics platform? (2020)](https://lobste.rs/s/gzkue1/what_is_your_preferred_web_traffic)
- [blackrock](https://github.com/rekki/blackrock) - Events & Analytics.
- [RudderStack](https://rudderstack.com/) - Platform for collecting, storing and routing customer event data to dozens of tools. ([Code](https://github.com/rudderlabs/rudder-server))
- [OpenTelemetry API for JavaScript](https://github.com/open-telemetry/opentelemetry-js) - Framework for collecting traces and metrics from applications.
- [PostHog](https://posthog.com/) - Developer-friendly, open-source product analytics. ([Code](https://github.com/PostHog/posthog))
- [A simple way to get more value from metrics](https://danluu.com/metrics-analytics/)
- [We ditched Google Analytics for good (2020)](https://missiveapp.com/blog/privacy-first-analytics) ([HN](https://news.ycombinator.com/item?id=23378524))
- [Netlify Analytics](https://www.netlify.com/products/analytics/)
- [Analytics](https://github.com/DavidWells/analytics) - Lightweight analytics abstraction layer for tracking page views, custom events, & identifying visitors. ([Docs](https://getanalytics.io/))
- [Roll Your Own Analytics](https://www.pcmaffey.com/roll-your-own-analytics)
- [Lightweight alternatives to Google Analytics (2020)](https://lwn.net/SubscriberLink/822568/b8f0709a45910e49/) ([Lobsters](https://lobste.rs/s/lvdj3w/lightweight_alternatives_google))
- [More alternatives to Google Analytics (2020)](https://lwn.net/SubscriberLink/824294/fe8f9331eca8b9ee/) ([Lobsters](https://lobste.rs/s/rwl3cp/more_alternatives_google_analytics))
- [Fathom Analytics: Count on it (2020)](https://brycewray.com/posts/2020/06/fathom-analytics-count-on-it/)
- [Time-based analytics, longitudinal clustering of usage data](http://www.feltpresence.com/analytics.html)
- [Deep Dive Into My Stats Page (2020)](https://sld.codes/articles/Deep-Dive-Into-My-Stats-Page)
- [Using event naming conventions to keep analytics data clean](https://davidwells.io/blog/clean-analytics)
- [Goodhart's Law and how systems are shaped by the metrics you chase (2020)](https://whyisthisinteresting.substack.com/p/why-is-this-interesting-the-goodharts) ([HN](https://news.ycombinator.com/item?id=23762526))
- [Simple Analytics](https://simpleanalytics.com/) - Simple, clean, and privacy-friendly analytics. ([Telegram Widget](https://github.com/simpleanalytics/chat))
- [GoAccess](https://goaccess.io/) - Visual Web Log Analyzer.
- [Monitoring your own infrastructure using Grafana, InfluxDB, and CollectD (2020)](https://serhack.me/articles/monitoring-infrastructure-grafana-influxdb-connectd/) ([HN](https://news.ycombinator.com/item?id=23906165))
- [kSense](https://ksense.io/) - Insanely Fast Analytics.
- [Umami](https://umami.is/) - Self-hosted open-source alternative to Google Analytics. ([HN](https://news.ycombinator.com/item?id=24198329))
- [Privacy Focused Analytics From Scratch (2020)](https://healeycodes.com/privacy-focused-analytics-from-scratch/)
- [Cloudflare Web Analytics](https://www.cloudflare.com/web-analytics/) ([Article](https://blog.cloudflare.com/free-privacy-first-analytics-for-a-better-web/)) ([HN](https://news.ycombinator.com/item?id=24627204)) ([HN 2](https://news.ycombinator.com/item?id=24628628)) ([Review](https://markosaric.com/cloudflare-analytics-review/)) ([HN 3](https://news.ycombinator.com/item?id=24846300))
- [Big tech finally challenges Fathom Analytics (2020)](https://usefathom.com/blog/big-tech-vs-fathom)
- [Umami](https://umami.is/) - Simple, fast, website analytics alternative to Google Analytics. ([Code](https://github.com/mikecao/umami)) ([HN](https://news.ycombinator.com/item?id=27181622))
- [The Analytics That Matter (2020)](https://css-tricks.com/the-analytics-that-matter/)
- [The Startup Guide to Analytics](https://windsor.io/guide)
- [userTrack](https://www.usertrack.net/) - Self-Hosted Analytics. ([HN](https://news.ycombinator.com/item?id=24746921))
- [Ackee](https://github.com/electerious/Ackee) - Self-hosted, Node.js based analytics tool for those who care about privacy. ([Web](https://ackee.electerious.com/))
- [Shynet](https://github.com/milesmcc/shynet/) - Modern, privacy-friendly, and detailed web analytics that works without cookies or JS.
- [Frovedis](https://github.com/frovedis/frovedis) - Framework of vectorized and distributed data analytics.
- [AWS Web Analytics](https://github.com/goatandsheep/aws-web-analytics) - Privacy-focused alternative to Google Analytics on AWS Pinpoint.
- [Microsoft Clarity](https://clarity.microsoft.com/) - Analytics for Websites.
- [Splitbee](https://splitbee.io/) - Friendly all-in-one analytics & conversion tool.
- [Real-time Security Insights: Apache Pinot at Confluera (2020)](https://medium.com/confluera-engineering/real-time-security-insights-apache-pinot-at-confluera-a6e5f401ff02)
- [Ask HN: How to run analytics on data without access to the data? (2020)](https://news.ycombinator.com/item?id=25429749)
- [Analytics Zoo](https://github.com/intel-analytics/analytics-zoo) - Distributed Tensorflow, Keras and PyTorch on Apache Spark/Flink & Ray.
- [How I replaced google analytics on my website (2020)](https://tnickel.de/2020/12/24/2020-12-How-I-replaced-google-analytics-on-my-website/)
- [Datadog](https://www.datadoghq.com/) - Cloud Monitoring as a Service. ([GitHub](https://github.com/DataDog))
- [Datadog Agent](https://github.com/DataDog/datadog-agent)
- [Awesome Analytics](https://github.com/onurakpolat/awesome-analytics)
- [Panelbear](https://panelbear.com/) - Fast and privacy-friendly website analytics.
- [Squzy](https://github.com/squzy/squzy) - High-performance open-source monitoring, incident and alert system written in Go with Bazel.
- [Webmention Analytics](https://github.com/maxboeck/webmention-analytics) - Analytics dashboard for webmention.io data. ([Article](https://mxb.dev/blog/webmention-analytics/))
- [Why Databricks Is Winning In The Data & Analytics Market (2021)](https://cloudnativeenterprise.substack.com/p/why-databricks-winning-market) ([HN](https://news.ycombinator.com/item?id=26135144))
- [Minimally Invasive (and More Accurate) Analytics: GoAccess and Athena/SQL (2021)](https://brandur.org/minimal-analytics) ([HN](https://news.ycombinator.com/item?id=26155361))
- [Google Analytics: Stop feeding the beast (2021)](https://casparwre.de/blog/stop-using-google-analytics/) ([HN](https://news.ycombinator.com/item?id=26263149))
- [Counter](https://counter.dev/) - Simple and Free Web Analytics. ([HN](https://news.ycombinator.com/item?id=26379569)) ([Code](https://github.com/ihucos/counter.dev))
- [Redata](https://www.redata.team/) - Better monitoring for data teams. ([Code](https://github.com/redata-team/redata))
- [SaaS Metrics](https://www.causal.app/saas-metrics) - Collection of articles and interactive models, to help you understand the metrics that matter to your SaaS business.
- [Volument](https://volument.com/) - Smarter take on website analytics. ([HN](https://news.ycombinator.com/item?id=27186249))
- [June](https://june.so/) - Instant analytics reports, on top of Segment.
- [A list of privacy-friendly Google Analytics alternatives (2021)](https://creativerly.com/google-analytics-alternatives/) ([HN](https://news.ycombinator.com/item?id=27604673))
- [ganalytics](https://github.com/lukeed/ganalytics) - Tiny (312B) client-side module for tracking with Google Analytics.
- [Load-testing my Web Analytics Tool](https://johnmathews.eu/load-testing-web-analytics-tool.html)
- [Tech-savvy audiences block Google Analytics (2021)](https://plausible.io/blog/google-analytics-adblockers-missing-data) ([HN](https://news.ycombinator.com/item?id=28365163))
- [Monitor ClickHouse with Prometheus & Grafana (2021)](https://tech.marksblogg.com/clickhouse-prometheus-grafana.html)
- [Pirsch Analytics](https://pirsch.io/) - Cookie-Free Alternative to Google Analytics.
- [Dune Analytics](https://dune.xyz/home) - Free crypto analytics by and for the community. ([Dune Snippets](https://github.com/sambacha/dune-snippets))
- [Bast](https://github.com/kooparse/bast) - Web analytics focusing on privacy and simplicity.
- [Sensor Tower](https://sensortower.com/) - Mobile App Store Marketing Intelligence.
- [Analytics Next](https://github.com/segmentio/analytics-next) - Latest version of Segment’s JavaScript SDK - enabling you to send your data to any tool.
- [Ask HN: Best alternatives to Google Analytics in 2021?](https://news.ycombinator.com/item?id=29662859)
- [Ask HN: Good open source alternatives to Google Analytics? (2022)](https://news.ycombinator.com/item?id=29888599)
- [Census](https://www.getcensus.com/) - Operational Analytics. Sync data from your warehouse into all your business tools. ([Article](https://blog.getcensus.com/series-b-the-future-of-operational-analytics/))
- [Chaos Genius](https://github.com/chaos-genius/chaos_genius) - ML powered analytics engine for outlier detection and root cause analysis.
- [consent-manager](https://github.com/segmentio/consent-manager) - Drop-in consent management plugin for analytics.js.
- [AppSignal](https://www.appsignal.com/) - Application Monitoring for Ruby on Rails, Elixir & Node.js.
- [Offen](https://www.offen.dev/) - Fair and lightweight alternative to common web analytics tools. ([Code](https://github.com/offen/offen))
- [Ballcone](https://github.com/dustalov/ballcone) - Fast and lightweight server-side Web analytics solution. It requires no JavaScript on your website.
- [Fugu](https://fugu.lol/) - Simple and privacy-friendly product analytics. ([Code](https://github.com/shafy/fugu))
- [Product Analytics from Scratch (2022)](https://tmfarrell.github.io/writing/2022/04/26/product_analytics_from_scratch/)
- [In-Memory Analytics with Apache Arrow (2022)](https://www.packtpub.com/product/in-memory-analytics-with-apache-arrow/9781801071031) ([Code](https://github.com/PacktPublishing/In-Memory-Analytics-with-Apache-Arrow-))
- [Objectiv](https://github.com/objectiv/objectiv-analytics) - Open-source product analytics infrastructure with a generic event taxonomy. ([Web](https://objectiv.io/)) ([HN](https://news.ycombinator.com/item?id=31432859))
- [Snowplow](https://github.com/snowplow/snowplow) - Enterprise-grade behavioral data engine (web, mobile, server-side, webhooks), running cloud-natively on AWS and GCP.
- [dbt Product Analytics](https://github.com/mjirv/dbt_product_analytics) ([HN](https://news.ycombinator.com/item?id=31877959))

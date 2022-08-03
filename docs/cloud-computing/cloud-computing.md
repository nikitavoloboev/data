# Cloud computing

Like using [fly.io](https://fly.io/) for serverless workloads. For personal servers I use [Hetzner](https://www.hetzner.com/) with [unlimited traffic and gigabit uplink](https://news.ycombinator.com/item?id=30055054). [OVHcloud](https://www.ovhcloud.com/en-ie/) is nice too. Had horrendous experience with [GCP](gcp/gcp.md)/[AWS](aws/aws.md) so avoid using it.

## ML

- [JarvisLabs.ai](https://cloud.jarvislabs.ai/) - One of the best pricing for renting compute. ([Tweet](https://twitter.com/bhutanisanyam1/status/1448971253654556674))
- [Paperspace](https://www.paperspace.com/) - Effortless infrastructure for Machine Learning and Data Science.
- [Google Cloud](https://cloud.google.com/) - Has generous [free tier](https://cloud.google.com/free/).
- [Amazon SageMaker](https://aws.amazon.com/sagemaker/) - Build, train, and deploy machine learning models at scale.
  - [Tutorial](https://docs.aws.amazon.com/sagemaker/latest/dg/whatis.html)
  - [Sagemaker examples](https://github.com/awslabs/amazon-sagemaker-examples)
- [Google Cloud Developer Cheat Sheet](https://medium.com/google-cloud/the-google-cloud-developer-cheat-sheet-429775bd6d11)
- [HN: Cloud Computing Without Containers](https://news.ycombinator.com/item?id=18415708)
- [Google Cloud AI Platform Products](https://cloud.google.com/ai-platform/docs) ([Code](https://github.com/GoogleCloudPlatform/ai-platform-samples))

## Servers

- [Digital Ocean](https://www.digitalocean.com/) - Cloud infrastructure provider.
- [Vultr](https://www.vultr.com/) - SSD VPS Servers, Cloud Servers and Cloud Hosting. ([Vultr CLI](https://github.com/vultr/vultr-cli)) ([GitHub](https://github.com/vultr))
- [OVHcloud](https://www.ovh.com/world/) - Web hosting, cloud computing and dedicated servers.
- [Hetzner](https://www.hetzner.com/) - Dedicated Root Server, VPS & Hosting. ([Moving To Hetzner Cloud from DigitalOcean](https://figbert.com/posts/moving-to-hetzner-from-digitalocean/)) ([Lobsters](https://lobste.rs/s/anzwio/moving_hetzner_cloud_from_digitalocean)) ([CLI](https://github.com/hetznercloud/cli)) ([Terraform](https://github.com/hetznercloud/terraform-provider-hcloud)) ([GitHub](https://github.com/hetznercloud)) ([Awesome](https://github.com/hetznercloud/awesome-hcloud)) ([HCloud provider](https://github.com/pulumi/pulumi-hcloud))
- [Prgmr](https://prgmr.com/xen/) - Linux & BSD virtual private servers.
- [Scaleway](https://www.scaleway.com/en/) - Cloud instances delivered in seconds with backup, network and security options. ([Scaleway cheaper than AWS for VMs](https://twitter.com/jpetazzo/status/1294561330805059585))
- [Feral Hosting](https://www.feralhosting.com/pricing)
- [Contabo](https://contabo.com/en/) - Quality VPS & Dedicated Servers At Incredible Prices.
- [Equinix Metal](https://metal.equinix.com/) - Bare Metal Server Provider - Global, Automated, Interconnected. ([GitHub](https://github.com/packethost))
- [GRAVITL](https://gravitl.com/) - Build the distributed cloud. ([GitHub](https://github.com/gravitl))
- [Exoscale Cloud Hosting](https://www.exoscale.com/) - Data centers in Germany, Austria & Switzerland. ([GitHub](https://github.com/exoscale))

## Web

- [Netlify](https://www.netlify.com/) - Build, test, and deploy globally with Netlify’s all-in-one platform for modern web projects. ([CLI](https://github.com/netlify/cli))
- [Vercel](https://vercel.com) - Optimal workflow for frontend teams. All-in-one: Static and JAMstack deployment, Serverless Functions, and Global CDN. ([HN](https://news.ycombinator.com/item?id=25442072)) ([Discussions](https://github.com/vercel/community/discussions))
- [Render](https://render.com/) - Unified platform to build and run all your apps and websites with free SSL, a global CDN, private networks and auto deploys from Git.
- [Mutable](https://mutable.io/) - Public Edge Cloud.
- [Zeet](https://zeet.co/) - Run your code.
- [Railway](https://railway.app/) - Develop code in a cloud that feels local. When you're ready, deploy from anywhere. ([GitHub](https://github.com/railwayapp)) ([Starters](https://github.com/railwayapp/starters)) ([CLI](https://github.com/railwayapp/cli)) ([Jake Cooper - Railway](https://www.youtube.com/watch?v=0YLNrvy_F70)) ([Awesome](https://github.com/railwayapp/awesome-railway))
- [Flightcontrol](https://flightcontrol.dev/) - Full stack deploy platform that runs on your own cloud.
- [Deta](https://www.deta.sh/) - Build your apps in hours, deploy them in seconds. ([Deta Space](https://deta.space/)) ([GitHub](https://github.com/deta)) ([Awesome](https://github.com/deta/awesome-deta))

## Other

- [Digital Ocean App Platform](https://www.digitalocean.com/products/app-platform/) ([Announcement](https://www.digitalocean.com/blog/introducing-digitalocean-app-platform-reimagining-paas-to-make-it-simpler-for-you-to-build-deploy-and-scale-apps/)) ([HN](https://news.ycombinator.com/item?id=24698334)) ([Reddit](https://www.reddit.com/r/webdev/comments/j699a0/digitalocean_launches_app_platform_a_fully/))
- [Nanobox](https://nanobox.io/) - Run any app on any cloud.
- [Hadean](https://hadean.com/) - Distributed Computing at Massive Scale.
- [Renderro](https://renderro.com/) - Powerful Cloud Computer For Creatives.
- [Yandex Cloud](https://cloud.yandex.com/en/) ([GitHub](https://github.com/yandex-cloud))
- [Akash Network](https://akash.network/) - Secure, transparent, and peer-to-peer cloud computing network. ([Code](https://github.com/ovrclk/akash)) ([GitHub](https://github.com/ovrclk)) ([Awesome](https://github.com/ovrclk/awesome-akash))
- [SetOps](https://www.setops.co/) - Cloud Deployment made for Developers.

## Notes

- [My goto hosting stack for the hyperbeast stack when I'm on a budget: 1. Stick everything inside a single VPS. 2. Use Dokku to: Config Nginx for me. Manage docker containers. Run cron job to backup the db.](https://twitter.com/benawad/status/1366423507555536896)
- [The PaaS (e.g. Heroku) vs. IaaS (e.g. AWS) debate needs to be revisited. For most teams, the most surefire way to blow deadlines and run over budget is to go it alone on AWS instead of sticking with a batteries-included PaaS that lets you forget about infra and focus on the app.](https://twitter.com/searls/status/1379461145799618564)
- [Hetzner has nice cheap servers](https://twitter.com/shipilev/status/1444618666826424322)
- [Nothing beats a good VPS server. No need for complex architectures/serverless in most projects.](https://twitter.com/pierregillesl/status/1445339116405604359)
- [The greatest trick the cloud vendors + Snowflake ever played was convincing orgs that by spending more on storage (of raw data) and compute (data modeling), teams can transform garbage data into valuable insights.](https://twitter.com/sarahcat21/status/1499863693277876225)

## Links

- [What does Unsplash cost in 2019?](https://medium.com/unsplash/what-does-unsplash-cost-in-2019-f499620a14d0) ([HN](https://news.ycombinator.com/item?id=19827521))
- [pack](https://github.com/buildpack/pack) - Local CLI for building apps using Cloud Native Buildpacks.
- [Common solutions and tools developed by Google Cloud's Professional Services team](https://github.com/GoogleCloudPlatform/professional-services)
- [Google Cloud Developer's Cheat Sheet](https://github.com/gregsramblings/google-cloud-4-words)
- [Awesome Cloud Build](https://github.com/Timtech4u/awesome-cloudbuild) - Curated list of resources about all things Google Cloud Build.
- [How to burn the most money with a single click in Azure (2020)](https://mijailovic.net/2020/03/28/azure-money-burning/) ([HN](https://news.ycombinator.com/item?id=22718330))
- [Awesome Cloud Native](https://github.com/rootsongjc/awesome-cloud-native)
- [How to deploy side projects as web services for free (2020)](https://ashishb.net/tech/how-to-deploy-side-projects-as-web-services-for-free/) ([Lobsters](https://lobste.rs/s/nn0kpt/how_deploy_side_projects_as_web_services))
- [DigitalOcean VPC](https://blog.digitalocean.com/vpc-trust-platform/) ([HN](https://news.ycombinator.com/item?id=23007860))
- [Finala](https://github.com/similarweb/finala) - Resource cloud scanner that analyzes and reports about wasteful and unused resources to cut unwanted expenses.
- [Dear Google Cloud: Your Deprecation Policy is Killing You (2020)](https://medium.com/@steve.yegge/dear-google-cloud-your-deprecation-policy-is-killing-you-ee7525dc05dc) ([HN](https://news.ycombinator.com/item?id=24165445))
- [env0](https://www.env0.com/) - Self-service Cloud Environments. Your team manage their own environments in AWS, Azure and Google.
- [Cloud Native Landscape](https://landscape.cncf.io/) - Map of uncharted terrain of cloud native technologies. ([Code](https://github.com/cncf/landscape)) ([HN](https://news.ycombinator.com/item?id=29935767))
- [How I Setup My Own Personal CDN (2020)](https://joel.net/how-i-setup-my-own-personal-cdn)
- [Optimyze.cloud](https://optimyze.cloud/) - Hyperscaler software efficiency. For everybody.
- [The Cloud Christmas](https://thecloud.christmas/)
- [Ask HN: Where to get cheap VPS with big storage? (2021)](https://news.ycombinator.com/item?id=25713160)
- [New Directions in Cloud Programming (2021)](https://arxiv.org/abs/2101.01159)
- [The Big Little Guide to Running Code in the Cloud(s) (2021)](https://sudhir.io/the-big-little-guide-to-running-code-in-the-clouds/)
- [New Directions in Cloud Programming (2021)](http://cidrdb.org/cidr2021/papers/cidr2021_paper16.pdf)
- [A crash course on running your own servers with a shoestring budget (2021)](https://blog.alexgleason.me/run-your-own-server/) ([Lobsters](https://lobste.rs/s/77orlj/crash_course_on_running_your_own_servers))
- [Cloudlist](https://github.com/projectdiscovery/cloudlist) - Multi-cloud tool for getting Assets (Hostnames, IP Addresses) from Cloud Providers.
- [Scout Suite](https://github.com/nccgroup/ScoutSuite) - Multi-Cloud Security Auditing Tool.
- [Cloud Native Glossary](https://github.com/cncf/glossary/blob/main/cloudnative-glossary.pdf)
- [Cloudflare on the Edge (2021)](https://stratechery.com/2021/cloudflare-on-the-edge/) ([HN](https://news.ycombinator.com/item?id=27120677))
- [Ask HN: Where is a nice place to host which is not AWS / GCP types (2021)](https://news.ycombinator.com/item?id=27150689)
- [Northflank](https://northflank.com/) - Full-stack cloud platform.
- [Spot instances for personal servers (2021)](https://pitr.ca/2021-05-23-personal-spot) - Save 70% on personal servers.
- [Hosting: Render (2021)](https://blog.placemark.io/2021/05/14/render.html)
- [Self taught guide to cloud computing](https://github.com/madebygps/self-taught-guide-to-cloud-computing)
- [Cloud Cost Handbook](https://handbook.vantage.sh/) ([Code](https://github.com/vantage-sh/handbook))
- [Alicorn](https://alicorncloud.io/) - Easily move between AWS, GCP and Azure. ([HN](https://news.ycombinator.com/item?id=28167349))
- [When AWS, Azure, or GCP Becomes the Competition (2019)](https://www.gkogan.co/blog/big-cloud/) ([HN](https://news.ycombinator.com/item?id=28351951))
- [cloud-init](https://github.com/canonical/cloud-init) - Industry standard multi-distribution method for cross-platform cloud instance initialization. ([Web](https://cloud-init.io/))
- [Cloudflare’s Disruption (2021)](https://stratechery.com/2021/cloudflares-disruption/) ([HN](https://news.ycombinator.com/item?id=28707317))
- [Cost-effective way to publish many GBs of data on internet with HTTPS access (2021)](https://twitter.com/shipilev/status/1444292088254943234)
- [AWS is playing Chess. Cloudflare is playing Go. (2021)](https://www.swyx.io/cloudflare-go/) ([HN](https://news.ycombinator.com/item?id=28903982))
- [Awesome PaaS](https://github.com/debarshibasak/awesome-paas) - Curated list of PaaS, developer platforms tools to emulate PaaS on cloud, Cloud IDEs and ADNs.
- [Cloudflare announced four major new features that could disrupt industries (2021)](https://www.indiehackers.com/post/cloudflare-just-disrupted-3-industries-in-1-week-907e44a8f5) ([HN](https://news.ycombinator.com/item?id=28746880))
- [Cloud Native Computing Foundation](https://www.cncf.io/) ([Technical Oversight Committee (TOC)](https://github.com/cncf/toc))
- [Cross-Cloud Access to SaaS Application Books](https://www.manning.com/liveprojectseries/cross-cloud-access-to-SaaS-application)
- [What is the next "cloud computing"? (2021)](https://www.reddit.com/r/investing/comments/qw71tf/what_is_the_next_cloud_computing/)
- [Interview: Matthew Prince with Eric Goldman (2021)](https://www.youtube.com/watch?v=30QNAMBdkbc) ([Transcript](https://knightfoundation.org/interview-matthew-prince-with-eric-goldman/))
- [The Impending Cloud Reshuffle (2021)](https://erikbern.com/2021/11/30/storm-in-the-stratosphere-how-the-cloud-will-be-reshuffled.html) ([HN](https://news.ycombinator.com/item?id=29411566))
- [We handle 80TB and 5M page views a month for under $400 (2022)](https://blog.polyhaven.com/how-we-handle-80tb-and-5m-page-views-a-month-for-under-400/) ([HN](https://news.ycombinator.com/item?id=29816504))
- [I got pwned by my cloud costs (2022)](https://www.troyhunt.com/how-i-got-pwned-by-my-cloud-costs/) ([HN](https://news.ycombinator.com/item?id=30054739)) ([Reddit](https://www.reddit.com/r/programming/comments/sbhgup/how_i_got_pwned_by_my_cloud_costs/))
- [Stratus Red Team](https://github.com/DataDog/stratus-red-team) - Granular, Actionable Adversary Emulation for the Cloud.
- [Tell HN: I got 10x Hetzner storage at the same price (2022)](https://news.ycombinator.com/item?id=30398534)
- [I'm so sorry everyone. Or: why I'm switching to Cloudflare (2022)](https://ersei.saggis.com/en/blog/im-sorry-everyone-cloudflare) ([HN](https://news.ycombinator.com/item?id=30443178))
- [Cloudflare Provider](https://github.com/pulumi/pulumi-cloudflare) - Cloudflare resource provider for Pulumi lets you use Cloudflare resources in your cloud programs.
- [Cloud Pricing API](https://github.com/infracost/cloud-pricing-api) - GraphQL-based API that includes all public prices from AWS, Azure and Google.
- [Mutagen](https://mutagen.io/) - Cloud-based development using your local tools. ([HN](https://news.ycombinator.com/item?id=30957156)) ([Mutagen Compose](https://github.com/mutagen-io/mutagen-compose))
- [NooBaa](https://www.noobaa.io/) - Data federation across multiple private and public clouds. ([Code](https://github.com/noobaa/noobaa-core))
- [Open Source PaaS List](https://github.com/guettli/open-source-paas)
- [Cloud Cost Control: Best Practices (2022)](https://bestcloudplatform.com/cloud-cost-control-best-practices/) ([HN](https://news.ycombinator.com/item?id=31271340))
- [Heroku: Core Impact (2022)](https://brandur.org/nanoglyphs/033-heroku) ([HN](https://news.ycombinator.com/item?id=31391272))
- [Nice bare cloud hardware providers (2022)](https://news.ycombinator.com/item?id=31496680)
- [The Story of Heroku (2022)](https://leerob.io/blog/heroku) ([HN](https://news.ycombinator.com/item?id=31559270)) ([Tweet](https://twitter.com/leeerob/status/1531279732213485568))
- [The Co-op Cloud](https://coopcloud.tech/) - Public interest infrastructure. Alternative to corporate clouds built by tech co-ops.
- [Komiser](https://github.com/mlabouardy/komiser) - Cloud Environment Inspector.
- [Ask HN: Why are there no big cloud vendors based in Europe? (2022)](https://news.ycombinator.com/item?id=31768109)
- [Ask HN: What do you use VMs for regularly? (2022)](https://news.ycombinator.com/item?id=31822932)
- [Hop](https://hop.io/) - Cloud Platform for Real time. Low-latency products built for the real time era. ([Docs](https://docs.hop.io/)) ([Docs Code](https://github.com/hopinc/hop-client-js))
- [Awesome Cloud Native Trainings](https://github.com/joseadanof/awesome-cloudnative-trainings)
- [Which cloud provider are your favorite and why? (2022)](https://www.reddit.com/r/devops/comments/w7vhvj/which_cloud_provider_are_your_favorite_and_why/)
- [Cloud Security Wiki](https://www.secwiki.cloud/) ([Code](https://github.com/withsecurelabs/cloud-wiki))
- [Use one big server (2022)](https://specbranch.com/posts/one-big-server/) ([HN](https://news.ycombinator.com/item?id=32319147)) ([Lobsters](https://lobste.rs/s/yimhoy/use_one_big_server))

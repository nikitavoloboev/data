---
title: Google Cloud
---

# [Google Cloud](https://cloud.google.com/)

Avoid GCP now because it's pricing is insane. I got billed 100 euro for empty Postgres database that I held for 2 weeks only and made 0 hits to as I was setting things up. I get that you [rent a full virtual machine](https://www.reddit.com/r/googlecloud/comments/9uramg/cloud_sql_cost_too_high_how_to_optimize/) but it's still a rip off. Actually of opinion now that both AWS and GCP are trash and should be avoided.

Looking into [Railway](https://railway.app/) & [Cloudflare](https://www.cloudflare.com/en-gb/) now for any cloud compute I want to do.

## Connecting to CloudSQL from Go

1. [Connecting to Cloud SQL (Postgres) from a Go web app](https://github.com/GoogleCloudPlatform/golang-samples/blob/main/cloudsql/postgres/database-sql/README.md)
2. [Connect using the Cloud SQL Auth proxy](https://cloud.google.com/sql/docs/postgres/connect-admin-proxy)

## Links

- [Cloud Tasks](https://cloud.google.com/tasks) - Asynchronous task execution.
- [Awesome Google Cloud Platform](https://github.com/GoogleCloudPlatform/awesome-google-cloud)
- [Google Cloud Buildpacks](https://github.com/GoogleCloudPlatform/buildpacks) - Builders and buildpacks designed to run on Google Cloud's container platforms.
- [How to find and delete idle GCP Projects (2020)](https://cloudberry.engineering/article/find-and-delete-idle-gcp-projects/)
- [Stricter Access Control to Google Cloud Registry (2020)](https://cloudberry.engineering/article/stricter-access-control-to-gcr/)
- [Introduction to Google Cloud Functions (2020)](https://ncona.com/2020/11/introduction-to-google-cloud-functions/) ([HN](https://news.ycombinator.com/item?id=24987197))
- [Terraform module for configuring GKE clusters](https://github.com/terraform-google-modules/terraform-google-kubernetes-engine)
- [google-cloud-rs](https://github.com/google-apis-rs/google-cloud-rs) - Asynchronous Rust bindings for Google Cloud Platform gRPC APIs.
- [2020: The year in databases at Google Cloud](https://cloud.google.com/blog/products/databases/managed-cloud-database-services-what-was-new-in-2020)
- [GCPSketchnote](https://github.com/priyankavergadia/GCPSketchnote) - Google Cloud Developer's Visual Notes.
- [Google Cloud Workflows](https://cloud.google.com/workflows) - Orchestrate and automate Google Cloud and HTTP-based API services with serverless workflows.
- [Google Cloud vs. AWS Onboarding Comparison (2021)](https://www.kevinslin.com/notes/ebd7fd65-988f-422a-93f5-b1fe5c3f29ce.html) ([HN](https://news.ycombinator.com/item?id=26252010))
- [Google API Improvement Proposals](https://google.aip.dev/)
- [WebSockets, HTTP/2 and gRPC bidirectional streams for Cloud Run (2021)](https://cloud.google.com/blog/products/serverless/cloud-run-gets-websockets-http-2-and-grpc-bidirectional-streams)
- [fake-gcs-server](https://github.com/fsouza/fake-gcs-server) - Google Cloud Storage emulator & testing library.
- [Terraform Google GKE Cluster](https://github.com/jetstack/terraform-google-gke-cluster) - Terraform module to create a best-practice Google Kubernetes Engine (GKE) cluster.
- [Building a high-scale chat server on Cloud Run (2021)](https://ahmet.im/blog/cloud-run-chat-server/) ([HN](https://news.ycombinator.com/item?id=26406206))
- [Compare AWS and Azure Services to Google Cloud](https://cloud.google.com/free/docs/aws-azure-gcp-service-comparison) ([HN](https://news.ycombinator.com/item?id=26562230))
- [setup-gcloud](https://github.com/google-github-actions/setup-gcloud) - Collection of GitHub Actions for interfacing with Google Cloud Platform.
- [runsd](https://github.com/ahmetb/runsd) - Drop-in Service Discovery capabilities for Google Cloud Run.
- [Google Cloud vs AWS in 2021 (Comparing the Giants)](https://kinsta.com/blog/google-cloud-vs-aws/)
- [Why I distrust Google Cloud more than than AWS or Azure (2021)](http://www.iasylum.net/writings/2021-04-21-why-I-distrust-google-cloud-more-than-AWS-or-Azure.html) ([HN](https://news.ycombinator.com/item?id=26897106))
- [First look at GKE Autopilot (2021)](https://ahmet.im/blog/gke-autopilot/)
- [Cloud Storage Rust](https://github.com/ThouCheese/cloud-storage-rs) - Crate for uploading files to Google cloud storage, and for generating download URLs.
- [Google APIs](https://github.com/googleapis/googleapis) - Public interface definitions of Google APIs.
- [Google API Linter](https://github.com/googleapis/api-linter) - API linter provides real-time checks for compliance with many of Google's API standards.
- [Berglas](https://github.com/GoogleCloudPlatform/berglas) - Tool for managing secrets on Google Cloud.
- [Google Cloud Rust Client](https://github.com/mozilla-services/google-cloud-rust)
- [gRPC and gRPC Web on Google Cloud Run (serverless) (2021)](https://blog.gendocu.com/posts/grpc-on-google-cloud/)
- [Google Cloud Digital Leader Certification Course (2021)](https://www.youtube.com/watch?v=UGRDM86MBIQ)
- [provider-gcp](https://github.com/crossplane/provider-gcp) - Crossplane Google Cloud Platform (GCP) infrastructure provider.
- [oidc-auth-google-cloud](https://github.com/sethvargo/oidc-auth-google-cloud) - GitHub Action for authenticating to Google Cloud with GitHub Actions OIDC tokens and Workload Identity Federation.
- [Google Cloud Client Libraries Go](https://github.com/googleapis/google-cloud-go)
- [Sample blueprints for Google Cloud](https://github.com/GoogleCloudPlatform/blueprints)
- [Google Cloud Deploy: Managed continuous delivery to GKE (2021)](https://cloud.google.com/blog/products/devops-sre/google-cloud-deploy-automates-deploys-to-gke)
- [GCP ping](https://gcping.com/) - Measure your latency to GCP regions.
- [Rowy](https://www.rowy.io/) - GCP as easy as ABC. ([Article](https://dev.to/harinilabs/gcp-as-easy-as-abc-a-low-code-platform-for-firestore-cloud-functions-deo)) ([Twitter](https://twitter.com/RowyIO)) ([HN](https://news.ycombinator.com/item?id=28758598))
- [Listing of Rust crates for use with Google Cloud](https://github.com/paulgb/are-we-google-cloud-yet)
- [gcpdiag](https://github.com/GoogleCloudPlatform/gcpdiag) - Command-line diagnostics tool for GCP customers.
- [Native Google Cloud Pulumi Provider](https://github.com/pulumi/pulumi-google-native)
- [Google Auth Library: Node.js Client](https://github.com/googleapis/google-auth-library-nodejs)
- [GCP Tau VMs](https://cloud.google.com/tau-vm) - Compute Engine virtual machines optimized for scale-out workloads. ([Tweet](https://twitter.com/uhoelzle/status/1450881369114963973))
- [Local Emulator for Google Cloud Storage](https://github.com/oittaa/gcp-storage-emulator)
- [csi-gcs](https://github.com/ofek/csi-gcs) - Kubernetes CSI driver for Google Cloud Storage.
- [Google Cloud Build Local Builder](https://github.com/GoogleCloudPlatform/cloud-build-local)
- [cloud-provider-gcp](https://github.com/kubernetes/cloud-provider-gcp)
- [Push to GCR GitHub Action](https://github.com/RafikFarhad/push-to-gcr-github-action)
- [GCP Developer Center](https://cloud.google.com/developers) ([Code](https://github.com/GoogleCloudPlatform/community))
- [gcloud](https://github.com/actions-hub/gcloud) - GitHub Action for interacting with Google Cloud Platform (GCP).
- [Cloud Run: multiple processes in a container (the lazy way) (2019)](https://ahmet.im/blog/cloud-run-multiple-processes-easy-way/index.html) ([Code](https://github.com/ahmetb/multi-process-container-lazy-solution))
- [Serving gRPC+HTTP from a Go app on Cloud Run (and elsewhere) (2021)](https://ahmet.im/blog/grpc-http-mux-go/)
- [Deploying to Cloud Run with Go (2021)](https://ahmet.im/blog/cloud-run-deploy-api/)
- [Bootstrap your Google Cloud Foundation with Terraform and Gitlab CI (2021)](https://johansiebens.dev/posts/2021/05/bootstrap-your-google-cloud-foundation-with-terraform-and-gitlab-ci/)
- [Microservices Architecture on Google Cloud (2021)](https://cloud.google.com/blog/topics/developers-practitioners/microservices-architecture-google-cloud) ([HN](https://news.ycombinator.com/item?id=29294267))
- [Cloud Run Proxy](https://github.com/GoogleCloudPlatform/cloud-run-proxy) - Local proxy for authenticating requests to Cloud Run.
- [Awesome Cloud Run](https://github.com/steren/awesome-cloudrun)
- [Get up and running with Go on Google Cloud](https://github.com/einride/cloudrunner-go)
- [Python Client for Google Cloud Storage](https://github.com/googleapis/python-storage)
- [Google APIs Node.js Client](https://github.com/googleapis/google-api-nodejs-client)
- [GCP Auth](https://github.com/hrvolapeter/gcp_auth) - Minimal authentication library for Google Cloud Platform (GCP) in Rust.
- [Google Cloud API Client Libraries for Rust](https://github.com/googleapis/google-cloud-rust)
- [Google Kubernetes Engine](https://cloud.google.com/kubernetes-engine/)
- [Google Cloud Common: Node.js Client](https://github.com/googleapis/nodejs-common)
- [Google Cloud Platform API hello world samples](https://github.com/salrashid123/gcpsamples)
- [GCSB](https://github.com/cloudspannerecosystem/gcsb) - Cloud Spanner load generator to load test your application and pre-warm the database before launch.
- [Google Cloud Developer Cheat Sheet](https://googlecloudcheatsheet.withgoogle.com/) ([HN](https://news.ycombinator.com/item?id=30357355))
- [Google Cloud architecture diagramming tool](https://googlecloudcheatsheet.withgoogle.com/architecture) ([Article](https://cloud.google.com/blog/topics/developers-practitioners/introducing-google-cloud-architecture-diagramming-tool))
- [Terraform Provider for Google Cloud Platform](https://github.com/hashicorp/terraform-provider-google-beta)
- [Google Rust APIs](https://github.com/Byron/google-apis-rs) - Binding and CLI generator for all Google APIs.
- [Google APIs Client Library for Go](https://github.com/googleapis/google-api-go-client) - Auto-generated Google APIs for Go.
- [gcloudx](https://github.com/emicklei/gcloudx) - Extra features for accessing the Google Cloud Platform.
- [Updates to Google Cloud’s infrastructure capabilities and pricing (2022)](https://cloud.google.com/blog/products/infrastructure/updates-to-google-clouds-infrastructure-pricing) ([HN](https://news.ycombinator.com/item?id=30671997))
- [Google Auth Python Library](https://github.com/googleapis/google-auth-library-python)
- [Penny Wise and Cloud Foolish (2022)](https://danielcompton.net/2022/03/21/penny-wise-cloud-foolish) ([Lobsters](https://lobste.rs/s/yz7yjv/penny_wise_cloud_foolish))
- [Google Compute Engine Persistent Disk CSI Driver](https://github.com/kubernetes-sigs/gcp-compute-persistent-disk-csi-driver)
- [HarbourBridge](https://github.com/cloudspannerecosystem/harbourbridge) - Stand-alone open source tool for Cloud Spanner evaluation and migration, using data from an existing PostgreSQL, MySQL, SQL Server, Oracle or DynamoDB database.
- [Google Kubernetes Cluster review tool](https://github.com/google/gke-policy-automation) - Validates selected GKE cluster against set of REGO policies.
- [TPU Starter](https://github.com/ayaka14732/tpu-starter) - Everything you want to know about Google Cloud TPUs.
- [Time limited, auto-expiring group memberships for users on Google Cloud](https://github.com/salrashid123/iam_autorevoke)
- [Configure Google production keys for Auth0](https://github.com/natbat/pillarpointstewards/issues/52) ([Tweet](https://twitter.com/simonw/status/1510642840149250055))
- [Google API Extensions for Go](https://github.com/googleapis/gax-go)
- [log](https://github.com/apsystole/log) - Go logging library for GCP App Engine, Cloud Run, Cloud Functions.
- [gcsfs](https://github.com/fsspec/gcsfs) - Pythonic file-system interface for Google Cloud Storage.
- [Litestream & Cloud Run example](https://github.com/steren/litestream-cloud-run-example)
- [Awesome Bigtable](https://github.com/zrosenbauer/awesome-bigtable)
- [Google Cloud service emulators for local development stacks](https://github.com/fullstorydev/emulators)
- [docker-credential-gcr](https://github.com/GoogleCloudPlatform/docker-credential-gcr) - Docker credential helper for GCR users.
- [Setup Google Cloud SQL Proxy on macOS (2022)](https://estebanborai.com/notes/setup-google-cloud-sql-proxy-on-macos)
- [IAM Go](https://github.com/einride/iam-go) - Opinionated Open Source implementation of the google.iam APIs on top of Cloud Spanner.
- [Protobuf + BigQuery + Go](https://github.com/einride/protobuf-bigquery-go) - Seamlessly save and load protocol buffers to and from BigQuery using Go.
- [Google Cloud Project Factory Terraform Module](https://github.com/terraform-google-modules/terraform-google-project-factory)
- [google-cloud-rust](https://github.com/yoshidan/google-cloud-rust) - Google Cloud Client Libraries for Rust.
- [Google Cloud Vertex AI Samples](https://github.com/GoogleCloudPlatform/vertex-ai-samples)
- [GroupSync: Cloud SQL IAM Database Authentication for Groups](https://github.com/GoogleCloudPlatform/cloud-sql-iam-db-authn-groups)
- [Cloud SQL Connector for Python Drivers](https://github.com/GoogleCloudPlatform/cloud-sql-python-connector)
- [Google Cloud Platform Pricing and Cost Calculator](https://github.com/Cyclenerd/google-cloud-pricing-cost-calculator) - Calculate estimated monthly costs of Google Cloud Platform products and resources via YAML files and Linux CLI program.
- [CI for Data in BigQuery CLI utility](https://github.com/GoogleCloudPlatform/ci-for-data-in-bigquery)
- [Security Response Automation](https://github.com/GoogleCloudPlatform/security-response-automation) - Take automated actions against threats and vulnerabilities.
- [Google Cloud Run, Satisfaction, and Scalability with Steren Giannini (2022)](https://overcast.fm/+RWDXAOH-s)
- [Building Serverless Applications with Google Cloud Run (2020)](https://wietsevenema.eu/cloud-run-book/)
- [BigQuery Emulator](https://github.com/goccy/bigquery-emulator) - Provides a way to launch a BigQuery server on your local machine for testing and development.
- [Lifecycle of a container on Cloud Run (2021)](https://cloud.google.com/blog/topics/developers-practitioners/lifecycle-container-cloud-run)
- [Database migration via Cloud SQL Proxy for Cloud SQL in Google Compute Engine VM (2022)](https://www.hairizuan.com/database-migration-via-cloud-sql-proxy-for-cloud-sql-in-google-compute-engine-vm/)
- [Securely connecting external tools to your GCP SQL database (2020)](https://trevor.io/blog/securely-connecting-external-tools-to-your-gcp-sql-database/)
- [Google Datastore](https://github.com/kitsonk/google-datastore) - APIs that allow interfacing to Google Datastore on GCP from Deno.
- [GCP Project Operator](https://github.com/openshift/gcp-project-operator) - Responsible for creating and destroying projects and service accounts in GCP.
- [pubsub_sendmail](https://github.com/GoogleCloudPlatform/cloud-pubsub-sendmail) - Send emails from Google Cloud Pub/Sub events.
- [bqtail](https://github.com/viant/bqtail) - BigQuery Google Storage Based Data Loader.
- [Networking 101 GCP sheet](https://github.com/jesuispy/networking-101-gcp-sheet)
- [Proof of concept use Cloud Run to deploy a SQL database on demand](https://github.com/guillaumeblaquiere/serverless-sql)
- [spanner-truncate](https://github.com/cloudspannerecosystem/spanner-truncate) - Tool to delete all rows from the tables in a Cloud Spanner database without deleting tables themselves.
- [Google API Client Library for JavaScript](https://github.com/google/google-api-javascript-client)
- [Google Cloud Platform ESPv2](https://github.com/GoogleCloudPlatform/esp-v2) - General-purpose L7 service proxy that enables API management capabilities for JSON/REST or gRPC API services.

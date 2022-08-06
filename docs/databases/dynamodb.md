---
title: DynamoDB
---

# [DynamoDB](https://aws.amazon.com/dynamodb)

[NoSQL Data Modeling with Amazon DynamoDB](https://www.youtube.com/watch?v=h7mH2Bxkc6k) is good intro talk. [DonutDB](https://github.com/psanford/donutdb) looks nice.

## OSS DynamoDB apps

- [DynamoDB Instagram](https://github.com/alexdebrie/dynamodb-instagram)
- [Note Service (Next Generation)](https://github.com/deeheber/note-service-next-generation) - GraphQL CRUD API built using AWS AppSync, Lambda, DynamoDB, CDK, Typescript.

## Notes

- Each item has a unique primary key and any number of attributes.
- Scanning for an item means looking through **every item in a table** thus it is less efficient than query search.
- Table is a collection of items, and each item is a collection of attributes.
- [DynamoDB will burn you in so many ways. It’s only fine if you have very predictably sized objects and non-spiky workloads.](https://twitter.com/danielrhodes/status/1466645623814328324)

## Links

- [Dynamo: Amazon’s Highly Available Key-value Store (2007)](https://www.allthingsdistributed.com/files/amazon-dynamo-sosp2007.pdf)
- [Intro to Dynamo](https://gist.github.com/jlafon/d8f91086e3d00c4bff3b)
- [Awesome DynamoDB](https://github.com/alexdebrie/awesome-dynamodb) - List of resources for learning about modeling, operating, and using Amazon DynamoDB.
- [SQL, NoSQL, and Scale: How DynamoDB scales where relational databases don't (2020)](https://www.alexdebrie.com/posts/dynamodb-no-bad-queries/)
- [dynamo](https://github.com/glassechidna/dynamo) - Dead-simple AWS DynamoDB CLI.
- [Using (and Ignoring) DynamoDB Best Practices with Serverless | Alex DeBrie (2019)](https://acloud.guru/series/serverlessconf-nyc-2019/view/dynamodb-best-practices)
- [Build with DynamoDB - Single-Table Design Pros and Cons (2020)](https://www.twitch.tv/videos/544223958)
- [DQL](https://github.com/stevearc/dql) - SQL-ish language for DynamoDB.
- [The DynamoDB Book: Data Modeling with NoSQL and DynamoDB](https://www.dynamodbbook.com/) ([HN](https://news.ycombinator.com/item?id=23193093))
- [5 Things I Learned from The DynamoDB Book (2020)](https://www.swyx.io/writing/dynamodb-book/) ([Twitter](https://twitter.com/swyx/status/1247585165766832128))
- [Rules for Data Modeling with DynamoDB (2020)](https://www.trek10.com/blog/the-ten-rules-for-data-modeling-with-dynamodb) ([HN](https://news.ycombinator.com/item?id=22813908))
- [Dynobase](https://dynobase.dev/) - Professional GUI Client for DynamoDB.
- [dynomite](https://github.com/softprops/dynomite) - Make your rust types fit DynamoDB and visa versa.
- [PynamoDB](https://github.com/pynamodb/PynamoDB) - Pythonic interface to Amazon's DynamoDB.
- [Introducing the GoLD Stack: Go + Lambda + DynamoDB (2020)](https://dev.to/prozz/introduction-to-the-gold-stack-5b66)
- [Modeling Graph Relationships in DynamoDB (2020)](https://medium.com/developing-koan/modeling-graph-relationships-in-dynamodb-c06141612a70)
- [Amazon DynamoDB Deep Dive: Advanced Design Patterns for DynamoDB (2018)](https://www.youtube.com/watch?v=HaEPXoXVf2k) ([Notes on NoSQL patterns](https://github.com/dideler/notes/blob/master/nosql-patterns.md))
- [How to understand DynamoDB](https://consulting.0x4447.com/articles/how_to/how-to-understand-dynamodb.html)
- [DynamoDB, explained](https://www.dynamodbguide.com/) - Primer on the DynamoDB NoSQL database.
- [Patterns: Serverless Rust + GraphQL + DynamoDB on AWS Lambda](https://github.com/codetalkio/patterns-serverless-rust)
- [Live Migration of DynamoDB Tables (2020)](https://codetalk.io/posts/2020-03-19-Live-Migration-of-DynamoDB-Tables.html)
- [dynamit-cli](https://github.com/floydspace/dynamodb-migrations-tool) - DynamoDB Migrations Tool CLI.
- [AsyncIO DynamoDB](https://github.com/HENNGE/aiodynamo) - Asynchronous, fast, pythonic DynamoDB Client. ([Docs](https://aiodynamo.readthedocs.io/en/latest/))
- [Best Practices for Designing and Architecting with DynamoDB](https://docs.aws.amazon.com/amazondynamodb/latest/developerguide/best-practices.html)
- [Rotary](https://github.com/akkoro/rotary) - ORM-like query library for DynamoDB.
- [dynein](https://github.com/awslabs/dynein) - DynamoDB CLI written in Rust.
- [graphql-ttl-transformer](https://github.com/flogy/graphql-ttl-transformer) - Enable DynamoDB's time-to-live feature to auto-delete old entries in your AWS Amplify API.
- [Comparing Fauna and DynamoDB: Architecture and Pricing (2020)](https://fauna.com/blog/comparing-fauna-and-dynamodb) ([HN](https://news.ycombinator.com/item?id=25363056))
- [Next.js + AWS DynamoDB](https://github.com/leerob/nextjs-aws-dynamodb)
- [jest-dynalite](https://github.com/freshollie/jest-dynalite) - Jest preset to run Dynalite (DynamoDB local) per test runner.
- [DynamoDB OneTable](https://github.com/sensedeep/dynamodb-onetable) - Access and management for one table designs with NodeJS.
- [DynamoDB JavaScript DocumentClient cheat sheet](https://github.com/dabit3/dynamodb-documentclient-cheat-sheet)
- [DynamoDB sessions at AWS re:Invent 2020](https://www.youtube.com/playlist?list=PL_EDAAla3DXWshFxx1R5P5MNaER84zHsU)
- [Serverless DynamoDB Local Plugin](https://github.com/99x/serverless-dynamodb-local) - Allows to run DynamoDB locally for serverless.
- [Fundamentals of Amazon DynamoDB Single Table Design (2020)](https://www.youtube.com/watch?app=desktop&v=KYy8X8t4MB8)
- [DynamoDB Best Practices](https://dynobase.dev/dynamodb-best-practices/)
- [DynamoDB Checklist (2021)](https://www.sensedeep.com/blog/posts/2021/dynamodb-checklist.html)
- [DynamoDB with Go](https://dev.to/jbszczepaniak/dynamodb-with-go-1-setup-1nnm)
- [Amazon DynamoDB Developer Guide](https://github.com/awsdocs/amazon-dynamodb-developer-guide)
- [Amazon DynamoDB data modeling with NoSQL Workbench with Gunnar and Samaneh (2021)](https://www.twitch.tv/videos/984733547)
- [You should always use DynamoDB global tables now (2021)](https://acloudguru.com/blog/engineering/you-should-always-use-dynamodb-global-tables-now)
- [Asynchronous API with DynamoDB Streams (2021)](https://medium.com/nerd-for-tech/asynchronous-api-with-dynamodb-streams-4117776f2fa4) ([HN](https://news.ycombinator.com/item?id=27232637))
- [LucidDyanamoDB](https://github.com/dineshsonachalam/Lucid-Dynamodb) - Simple Python wrapper to AWS DynamoDB. ([HN](https://news.ycombinator.com/item?id=27334430))
- [ddbimport](https://github.com/a-h/ddbimport) - Import CSV data into DynamoDB.
- [Event Sourced Database in DynamoDB / TypeScript](https://github.com/a-h/hde)
- [The What, Why and How of DynamoDB (2021)](https://adamrackis.dev/dynamo-introduction/)
- [Best Practices for Modeling Relational Data in DynamoDB](https://docs.aws.amazon.com/amazondynamodb/latest/developerguide/bp-relational-modeling.html)
- [Amazon DynamoDB Under the Hood: How We Built a Hyper-Scale Database (2018)](https://www.youtube.com/watch?v=yvBR71D0nAQ)
- [AWS DynamoDB Component](https://github.com/serverless-components/aws-dynamodb) - Easily provision AWS DynamoDB tables using Serverless Components.
- [AWS DynamoDBtoCSV](https://github.com/edasque/DynamoDBtoCSV) - Dump DynamoDB data into a CSV file.
- [dynm](https://github.com/kocisov/dynm) - Utility for easier interaction with DynamoDB.
- [raiden-dynamo](https://github.com/raiden-rs/raiden-dynamo) - DynamoDB library for Rust.
- [DynamoDB + Lambda Workshop](https://github.com/alexdebrie/lambda-dynamodb-workshop)
- [DynamoDB --> Stream --> Elasticsearch](https://github.com/matrus2/dynamodb-stream-elasticsearch) - Missing blueprint for AWS Lambda. Reads stream from AWS DynamoDB and writes it to ElasticSearch.
- [DynamoDB Examples](https://github.com/aws-samples/aws-dynamodb-examples)
- [DonutDB](https://github.com/psanford/donutdb) - SQL database implemented on DynamoDB and SQLite.
- [DynamoDB Lock Client for Go](https://github.com/cirello-io/dynamolock)
- [DynamoDB Foreign Data Wrapper for PostgreSQL](https://github.com/pgspider/dynamodb_fdw)
- [New DynamoDB Table Class – Save Up To 60% in Your DynamoDB Costs (2021)](https://aws.amazon.com/blogs/aws/new-dynamodb-table-class-save-up-to-60-in-your-dynamodb-costs/)
- [dynamodump](https://github.com/mifi/dynamodump) - Node CLI for backing up and restoring schema+data from DynamoDB tables.
- [Starter project with Go, Gin and DynamoDB](https://github.com/vsouza/go-gin-boilerplate)
- [Go DynamoDB Web App Starter](https://github.com/kaihendry/go-web-dynamo-starter)
- [How to develop an AWS hosted DynamoDB Web application locally](https://github.com/kaihendry/local-audio)
- [TypeDORM](https://github.com/typedorm/typedorm) - Strongly typed ORM for DynamoDB - Built with the single-table-design pattern in mind.
- [typesafe-dynamodb](https://github.com/sam-goodwin/typesafe-dynamodb) - TypeSafe type definitions for the AWS DynamoDB API.
- [serde_dynamodb](https://github.com/mockersf/serde_dynamodb) - Talk with dynamodb using your existing structs thanks to serde.
- [Happy 10th Birthday, DynamoDB (2022)](https://aws.amazon.com/blogs/aws/happy-birthday-dynamodb/)
- [Serverless Typescript Demo](https://github.com/aws-samples/serverless-typescript-demo) - Simple serverless application built in Typescript and uses Node.js 14 runtime.
- [Serverless search for DynamoDB](https://github.com/jakejscott/dynamodb-stream-indexer)
- [Alex DeBrie: NoSQL Data Modeling with Amazon DynamoDB (2022)](https://www.youtube.com/watch?v=h7mH2Bxkc6k)
- [AWS Lambda in Private Subnet connecting to DynamoDB with VPC Endpoint](https://github.com/oieduardorabelo/cdk-lambda-private-subnet-dynamodb-vpc-endpoint)
- [Dynosaur](https://github.com/SystemFw/dynosaur) - Declarative bidirectional codecs for DynamoDb AttributeValue, SDK 2.
- [Dynomite](https://github.com/Netflix/dynomite) - Inspired by Dynamo whitepaper, is a thin, distributed dynamo layer for different storage engines and protocols.
- [dynamodb_lock](https://crates.io/crates/dynamodb_lock) - Distributed lock backed by DynamoDB.
- [Amazon DynamoDB Design Patterns](https://github.com/aws-samples/amazon-dynamodb-design-patterns)
- [How to ensure cross-region data integrity with Amazon DynamoDB global tables (2022)](https://www.stedi.com/blog/how-to-ensure-cross-region-data-integrity-with-amazon-dynamodb-global-tables)
- [Serverless Go Demo](https://github.com/aws-samples/serverless-go-demo) - Consists of an API Gateway backed by four Lambda functions and a DynamoDB table for storage.
- [The DynamoDB paper (2022)](https://brooker.co.za/blog/2022/07/12/dynamodb.html) ([HN](https://news.ycombinator.com/item?id=32094046)) ([Summary](http://muratbuffalo.blogspot.com/2022/07/amazon-dynamodb-scalable-predictably.html))
- [Single Table DynamoDB](https://github.com/jeffreyyoung/single-table-dynamo) - DynamoDB client built for using one table.
- [Some notes on DynamoDB 2022 paper](http://_.0xffff.me/dynamodb2022.html) ([HN](https://news.ycombinator.com/item?id=32361558))

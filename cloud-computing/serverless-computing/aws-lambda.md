# [AWS Lambda](https://aws.amazon.com/lambda/)

## Notes

- [I just reduced latency by \~ 10x by bumping memory from default to \~1GB. w/ provisioned concurrency brought the request time down to around 30ms w DDB](https://twitter.com/dabit3/status/1299846109448282112)

## Links

- [Writing AWS Lambda Functions in Rust](https://github.com/SilentByte/rust-lambda) ([Article](https://silentbyte.com/writing-aws-lambda-functions-in-rust))
- [sqs-lambda](https://github.com/insanitybit/sqs-lambda) - Rust crate for writing AWS Lambdas that are triggered by SQS.
- [WTF is AWS Lambda course (2020)](https://egghead.io/lessons/aws-wtf-is-aws-lambda)
- [Create an AWS Lambda function from scratch (2020)](https://dev.to/tlakomy/create-an-aws-lambda-function-from-scratch-3fdd)
- [AWSome Lambda Layers](https://github.com/mthenw/awesome-layers)
- [5 reasons why you might use AWS Lambda for your next project (2020)](https://dev.to/tlakomy/5-reasons-why-you-might-use-aws-lambda-for-your-next-project-ik8)
- [AWS SAM Local](https://aws.amazon.com/about-aws/whats-new/2017/08/introducing-aws-sam-local-a-cli-tool-to-test-aws-lambda-functions-locally/) - Allows you to run your serverless application locally for quick development and testing.
- [python-lambda-layer-builder](https://github.com/tobilg/python-lambda-layer-builder) - Build tool for creating optimized Python AWS Lambda layers.
- [Working with AWS Lambda and Chalice (2020)](https://papercup.dev/posts/working-with-aws-lambda-and-chalice/)
- [Notes on AWS Lambda (2020)](https://twitter.com/scottdomes/status/1249787042797465600)
- [AWS Lambda for Go](https://github.com/aws/aws-lambda-go) - Libraries, samples and tools to help Go developers develop AWS Lambda functions.
- [Apollo AWS Lambda with GraphQL subscriptions](https://github.com/michalkvasnicak/aws-lambda-graphql)
- [Lambda Powertools](https://github.com/awslabs/aws-lambda-powertools) - Suite of utilities for AWS Lambda Functions that makes tracing with AWS X-Ray, structured logging and creating custom metrics asynchronously easier.
- [Bash in AWS Lambda](https://github.com/gkrizek/bash-lambda-layer)
- [AWS Lambda Go Api Proxy](https://github.com/awslabs/aws-lambda-go-api-proxy) - Makes it easy to run Golang APIs written with frameworks such as Gin with AWS Lambda and Amazon API Gateway.
- [AWS Lambda Power Tuning](https://github.com/alexcasalboni/aws-lambda-power-tuning) - AWS Step Functions state machine that helps you optimize your Lambda functions in a data-driven way.
- [Swift AWS Lambda Runtime](https://github.com/swift-server/swift-aws-lambda-runtime) ([HN](https://news.ycombinator.com/item?id=23352501))
- [A Shared File System for Your Lambda Functions (2020)](https://aws.amazon.com/blogs/aws/new-a-shared-file-system-for-your-lambda-functions/) ([HN](https://news.ycombinator.com/item?id=23543554))
- [Chalice](https://github.com/aws/chalice) - Framework for writing serverless apps in python. It allows you to quickly create and deploy applications that use AWS Lambda.
- [Chaos Lambda](https://github.com/artilleryio/chaos-lambda) - Serverless chaos monkey for AWS (runs on AWS Lambda).
- [AWS CDK Made Simple: Run a Lambda function locally (2020)](https://tlakomy.com/run-cdk-lambda-function-locally)
- [Building Your First Serverless Service With AWS Lambda Functions (2020)](https://css-tricks.com/building-your-first-serverless-service-with-aws-lambda-functions/)
- [Introducing Direct Lambda Resolvers: AWS AppSync GraphQL APIs without VTL](https://aws.amazon.com/blogs/mobile/appsync-direct-lambda/)
- [Cache AWS Lambda responses with Cloudflare (2020)](https://kylebarron.dev/blog/caching-lambda-functions-cloudflare)
- [Debugging AWS Lambda Timeouts (2020)](https://lumigo.io/blog/debugging-aws-lambda-timeouts/)
- [Show HN: AWS Lambda TypeScript Middleware](https://github.com/dbartholomae/lambda-middleware) ([Docs](https://dbartholomae.github.io/lambda-middleware/)) ([HN](https://news.ycombinator.com/item?id=24280237))
- [Serverless ML Inference with AWS Lambda + Amazon EFS (2020)](https://medium.com/faun/setup-serverless-ml-inference-with-aws-lambda-efs-738546fa2e03)
- [Lambda calculus and Graham’s number (2012)](https://mindsarentmagic.org/2012/11/22/lambda-graham/)
- [Rust on AWS Lambda Using AWS CDK for Deployment](https://github.com/codetalkio/patterns-serverless-rust-minimal)
- [The Complete AWS Lambda Handbook for Beginners (2020)](https://dashbird.io/blog/complete-aws-lambda-handbook-beginners-part-1/)
- [AWS Lambda Extensions – In preview (2020)](https://aws.amazon.com/blogs/compute/introducing-aws-lambda-extensions-in-preview/) ([Tweet](https://twitter.com/dhruvsood/status/1314236371570024449))
- [Building Bad Lambda (2020)](https://kohidave.dev/posts/building-bad-lambda/)
- [golambda](https://github.com/rakyll/golambda) - AWS Lambda Go functions made easy.
- [Determine prominent colors in a picture, your first AWS Lambda in Go (2020)](https://buddy.works/tutorials/determine-prominent-colors-in-a-picture-your-first-aws-lambda-in-go)
- [Using AWS Lambda to sync an Elastic Load Balancer's Target Groups (2020)](https://emilenijssen.nl/6-aws-lambda-sync-elastic-load-balancer-target-groups/)
- [AWS Lambda Runtime Interface Emulator](https://github.com/aws/aws-lambda-runtime-interface-emulator)
- [Integration testing for AWS Lambda in Go with Docker-compose (2020)](https://buddy.works/tutorials/integration-testing-for-aws-lambda-in-go-with-docker-compose)
- [Optimizing Lambda Cost with Multi-Threading (2020)](https://www.sentiatechblog.com/aws-re-invent-2020-day-3-optimizing-lambda-cost-with-multi-threading)
- [AWS Lambda Extensions](https://github.com/aws-samples/aws-lambda-extensions) - Collection of sample extensions to help you get started with AWS Lambda Extensions.
- [AWS Lambda Python Runtime Interface Client](https://github.com/aws/aws-lambda-python-runtime-interface-client)
- [aws-lambda-nodejs-esbuild](https://github.com/floydspace/aws-lambda-nodejs-esbuild) - AWS CDK Construct to build Node.js AWS lambdas using esbuild.
- [aws-lambda-deploy](https://github.com/aws-samples/aws-lambda-deploy) - Collection of sample tools to enable canary deployments of AWS Lambda functions.
- [Mangum](https://github.com/jordaneremieff/mangum) - AWS Lambda & API Gateway support for ASGI.
- [AWS Lambda Terraform Cookbook with working examples](https://github.com/nsriram/lambda-the-terraform-way) ([HN](https://news.ycombinator.com/item?id=25588898))
- [lambda_decorators](https://github.com/dschep/lambda-decorators) - Collection of useful decorators for making AWS Lambda handlers.
- [Optimizing Lambda functions packaged as container images (2021)](https://aws.amazon.com/blogs/compute/optimizing-lambda-functions-packaged-as-container-images/)
- [AWS Lambda Power tools (Python)](https://github.com/awslabs/aws-lambda-powertools-python)
- [Rust Runtime for AWS Lambda](https://github.com/lamedh-dev/aws-lambda-rust-runtime)
- [Serverless TypeScript: A Complete Setup for AWS SAM Lambdas (2021)](https://evilmartians.com/chronicles/serverless-typescript-a-complete-setup-for-aws-sam-lambda)
- [cdk-watch](https://github.com/teamplanes/cdk-watch) - CLI to watch and live-update your CDK Stack's Lambdas.
- [re:Web](https://github.com/apparentorder/reweb) - Enables classic web applications to run on AWS Lambda.

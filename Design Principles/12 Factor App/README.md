# The Twelve-Factor App

The Twelve-Factor App is a set of best practices for building modern software-as-a-service (SaaS) applications.

## Codebase

A codebase is any single repo or any set of repos that share a root commit (in a decentralised revision control system like Git). A twelve-factor app is always tracked in a version control system.

There is only one codebase per app, but there will be many app deployments. Deploy is a running instance of the app. This is typically a production site and one or more staging sites.

## Dependencies

**Explicitly declare and isolate dependencies**

The app should explicitly declare all dependencies, such as libraries and packages, to ensure consistency across environments. This prevents "it works on my machine" problems.

One benefit of explicit dependency declaration is that it simplifies setup for developers new to the app. The new developer can check out the app’s codebase on their development machine, requiring only the language runtime and dependency manager installed as prerequisites. They will be able to set up everything needed to run the app’s code with a deterministic build command.

## Config

**Store config in the environment**

Apps sometimes store config as constants in the code. This is a violation of the twelve-factor, which requires strict separation of config from code. Config varies substantially across deploys, code does not.

The twelve-factor app stores config in environment variables. Env vars are easy to change between deploys without changing any code; unlike config files, there is little chance of them being checked into the code repo accidentally; and unlike custom config files, or other config mechanisms such as Java System Properties, they are a language- and OS-agnostic standard.

## Backing Services

**Treat backing services as attached resources**

A backing service is any service the app consumes over the network as part of its normal operation. Examples include data stores (such as MySQL or CouchDB), messaging/queueing systems (such as RabbitMQ or Beanstalkd), SMTP services for outbound email (such as Postfix), and caching systems (such as Memcached).

The code for a twelve-factor app makes no distinction between local and third-party services. To the app, both are attached resources, accessed via a URL or other locator/credentials stored in the config. Deployment of the twelve-factor app should be able to swap out a local MySQL database with one managed by a third party (such as Amazon RDS) without any changes to the app’s code. Likewise, a local SMTP server could be swapped with a third-party SMTP service (such as Postmark) without code changes. In both cases, only the resource handle in the config needs to change.

## Build, Release, Run

**Strictly separate build and run stages**

The twelve-factor app uses strict separation between the build, release, and run stages.

The build stage is a transform that converts a code repo into an executable bundle known as a build. Using a version of the code at a commit specified by the deployment process, the build stage fetches vendors' dependencies and compiles binaries and assets.

The release stage takes the build produced by the build stage and combines it with the deployment’s current config. The resulting release contains both the build and the config and is ready for immediate execution in the execution environment.

The run stage (also known as “runtime”) runs the app in the execution environment, by launching some set of the app’s processes against a selected release.

## Processes

**Execute the app as one or more stateless processes**

The app is executed in the execution environment as one or more processes. Twelve-factor processes are stateless and share nothing. Any data that needs to persist must be stored in a stateful backing service, typically a database.

The twelve-factor app never assumes that anything cached in memory or on disk will be available on a future request or job – with many processes of each type running, chances are high that a future request will be served by a different process.

## Port Binding

**Export services via port binding**

Port Binding means that the application itself is responsible for running its web server and exposing itself to the internet or a network.

## Concurrency

**Scale out via the process model**

The app should be designed to scale out by running multiple processes or instances, each handling a part of the workload. This is often achieved through horizontal scaling, where more instances are added to handle the increased load.

## Disposability

The app should be designed for quick startup and graceful shutdown to maximize robustness.

## Dev/Prod Parity

**Keep development, staging, and production as similar as possible**

The twelve-factor app is designed for continuous deployment by keeping the gap between development and production small.

## Logs

**Treat logs as event streams**

Logs should be directed to standard output and captured by the environment, rather than being managed by the app itself. This allows logs to be centralized and processed in different ways depending on the needs of the environment.

## Admin Processes

**Run admin/management tasks as one-off processes**

Administrative or management tasks, like database migrations or running one-off scripts, should be run as one-off (once or occasionally) processes against a deployed release. This ensures these tasks run in the same environment as the app, reducing the chance of environment-specific issues.
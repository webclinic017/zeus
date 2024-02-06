---
sidebar_position: 1
displayed_sidebar: zK8s
---

# Introduction

You can use the Zeus API to programmatically interact with our cloud platform's core infrastructure apis and services,
or you can use our web interface to interact with the platform.

## 1-on-1 Tailored Onboarding

Want one-on-one help getting started tailored to your use case? Schedule a Google meet with an expert
at [https://calendly.com/zeusfyi/solutions-engineering](https://calendly.com/zeusfyi/solutions-engineering).

## Sign Up

You can sign up for a free account at [https://cloud.zeus.fyi/login](https://cloud.zeus.fyi/login)

![Screenshot 2023-10-01 at 7 03 17 PM](https://github.com/zeus-fyi/zeus/assets/17446735/e76133f9-8fdd-49b0-9652-be8066587f86)

You can use your Google account to sign up, or you can sign up with your email address and verify your email.

## Quick Start

See our Cookbooks/Ethereum section on our docs site for a quick start guide on how to deploy an Ethereum beacon and you'll
get familiar with the platform in no time.

## Start with Reading the Platform Overview

This section covers the system elements for the platform, and how they work together. It also shows you the code
that applies these concepts to the platform.

## Then Learn zK8s Client & APIs

Learn how to use the zK8s client and APIs to build your own infrastructure using only Go code and YAML templates.

## Generate an API Key

Prerequisites: You'll need to generate an API key from the access panel if you don't have an existing one.

![APIkeypage](https://github.com/zeus-fyi/zeus/assets/17446735/7352892d-49ad-4a72-add1-5b212a90b914)

## Want to see a video?

### Sui Node Deployment using DigitalOcean Local NVMe on Kubernetes

<iframe width="560" height="315" src="https://www.youtube.com/embed/G8JBECjC6fc?si=I2ZADqNzS6Fh11WW&amp;start=1" title="YouTube video player" frameborder="0" allow="accelerometer; autoplay; clipboard-write; encrypted-media; gyroscope; picture-in-picture; web-share" allowfullscreen></iframe>

### 1-Click Upgrading a Fleet of zK8s Apps on AWS, GCP, & DigitalOcean

<iframe width="560" height="315" src="https://www.youtube.com/embed/H7sMsK_dnj4?si=RJ9PpuJ8AfXf_Wai&amp;start=1" title="YouTube video player" frameborder="0" allow="accelerometer; autoplay; clipboard-write; encrypted-media; gyroscope; picture-in-picture; web-share" allowfullscreen></iframe>

Checkout our YouTube channel for more videos on how to use the power of Zeus to build your own infrastructure.
https://www.youtube.com/@Zeusfyi

## Want to follow a tutorial?

### Option 1: Easy - Build an Ethereum Beacon on zK8s

This design process covers how to build an Ethereum Beacon using Go code for zK8s, and how to deploy it on Zeus. Even
a complete beginner can follow along.

https://medium.com/zeusfyi/zeus-k8s-infra-as-code-concepts-47e690c6e3c5

### Option 2: Medium - Learn Best Practices: Rapid Infra Development on zK8s

This design process covers how I packaged a Cosmos node on Zeus
within one day.

https://medium.com/zeusfyi/zeus-k8s-infra-as-code-concepts-47e690c6e3c5

### Option 3: Easy-Medium - Follow a Cookbook Recipe

...and build infra using only Go code and a YAML template base.

https://github.com/zeus-fyi/zeus/tree/main/cookbooks

https://medium.com/zeusfyi/zeus-k8s-infra-as-code-concepts-47e690c6e3c5

### Option 4: Easiest - Tailored one on one help by schedule.

Feel free to email us at support@zeus.fyi or schedule an expert and talk to us directly!

https://calendly.com/zeusfyi/solutions-engineering

## Deploy a Free Trial Cluster

You can deploy any* one app for evaluation purposes without a credit card. It will delete itself after one hour,
unless you add billing before then.

\* that costs $500/mo or less without a credit card

## Top-k questions about the platform

### What is the difference between a cluster and an app?

The app is the naming convention for the underlying codebase, e.g. they're Postgres apps, but we still need disk, ram, a
place to deploy it and so on. Some postgres apps need sharded configs, etc.

In the context of configuration

    App == Postgres Family & Derivative Configs
    Cluster == Specific Postgres Version + Disk + RAM + CPU + Ingress + DNS + Load Balancer + etc.

A Cluster topology is a complete underlying infrastructure hierarchy that binds at least one fully deployable zK8s app.

But in the context of the platform dashboard

    Clusters == Where you can view your live Kubernetes namespaces, locations, app details, and can manage some DevOps tasks.

And in the context of Kubernetes

    Cluster == meaning it is is a set of nodes that run containerized applications.
    Node == VM or physical machine that Kubernetes runs containers on.

### What is the difference between a cluster and a topology?

    Class Topologies:  Base, Cluster, Matrix, Network-Orchestration 

A topology is a class specified and well defined unit of infrastructure that can be anywhere from a single standalone
microservice to a fully orchestrated network topology. These are all covered in the platform overview in detail with
code examples, check them out there!

### What's the difference between zK8s apps and regular Kubernetes workloads?

None, they're completely equivalent. We just call them zK8s apps since we also build hierarchical rules on top of them.
We also let you code and unit test them with strongly typed Go code as an alternative to messy templated Helm charts in
YAML, that everyone is afraid to touch, and for good reason.

### On-Demand Pricing vs Reserved Pricing vs Spot

On-demand pricing is the price you pay for guaranteed resources you can provision on demand. These are more ideal for
short term workload spikes, development purposes, or low traffic apps.

Spot pricing offers significant discounts up to 70% on compute resources, but they are not guaranteed. If you're running
a workload that isn't mission critical, and can be interrupted, then spot pricing is a great way to save money. We can
help you figure out when it makes sense to use spot instances, and when it doesn't.

If you can reasonably forecast spending for at least 1-month reach out to us, and we can get you a better deal on
reserved compute through a variety of options

### I currently use GitOps, how do I keep a GitOps flow with Zeus?

#### The Easy Way

    Add the Upload Chart Call as CI/CD Stage

Like in this test case

```go
func (t *RedisCookbookTestSuite) TestUploadRedis() {
    _, rerr := redisClusterDefinition.UploadChartsFromClusterDefinition(ctx, t.ZeusTestClient, true)
    t.Require().Nil(rerr)
}
```

    Add the Deployment Call as the following CI/CD Stage

Full example here

```go
func (t *RedisCookbookTestSuite) TestDeployRedis() {
    t.TestUploadRedis()
    cdep := redisClusterDefinition.GenerateDeploymentRequest()

    _, err := t.ZeusTestClient.DeployCluster(ctx, cdep)
	t.Require().Nil(err)
}
```

It will then replace it with your most recent upload, and deploy it to your cluster. All infrastructure definitions are
immutable upon creation and each has a unique id, so you can always reference it by that id and use that for version
control, it is also unix timestamped. You can also hash the workload definition and use that as a version control
and for integrity checks.

### I currently use ...X Cloud or DevOps platform, how would I migrate to Zeus, I don't have time to learn a new tool?

We'll help you migrate your existing infrastructure to Zeus, and we'll help migrate it for you as part of an
enterprise contract. We'll also help you otherwise, but we'll prioritize enterprise customers first, then by FIFO.
You'll be surprised at how easy & fast it is.

### Do Enterprise & Private Clusters work with the UI?

Yes, and only your org's authorized users will be able to see, access, or modify any of your private clusters.

### Can I use my startup credits on AWS, GCP, Azure, DigitalOcean, etc.?

Yes, but only for Enterprise & Private Clusters.

### How do I get Prometheus metrics for my app?

Send us an email. There's some on-boarding steps required to get access to Prometheus metrics via PromQL that we'll need
to send you. We're working on automating this process, but for now it's semi-manual.

### Can I use this to build a serverless technology?

Yes. Contact us for more details on our ultra low latency serverless Thunderbolts program, and how you can use them to
build your own serverless technology such as GPU accelerated serverless functions, or serverless databases. We'll be releasing
more details about these publicly when they've been iterated on enough to package a release meant for a wider audience.

### I don't see many On-Demand choices for GPUs, why?

You'll need to contact us for full GPU access via our search API and app deployment dashboard. There's some additional steps
required to get access to GPUs. Don't worry though the process is easy, it requires a short manual review due to high
cost and availability constraints.

### Do I need to understand Kubernetes?

No, you don't need to understand Kubernetes in much detail to use Zeus, but we do recommend reading up on some basics
though for the best experience and scheduling an expert on-boarding session with us and we'll to get you up to speed, we even
have free cloud credits dedicated for educational purposes for those who simply want to learn Kubernetes you should
definitely ask us about if you're interested!

We've abstracted away the complexity of Kubernetes to the point where you really only need to know a few keywords
and concepts to be competent in nearly every aspect of managing your DevOps & Infra via API and Dashboard on the platform.
However, unlike other platforms, we don't make up terminology to hide the underlying Kubernetes system, we use the same terminology as Kubernetes,
so that you can use well understood terminology if you ever transition to needing more powerful Kubernetes features,
without having to re-haul your entire system.

I'd recommend this technical reference book, which contains just the right amount of detail to get you up to speed while
leaving room for you to learn more if you want to.

#### Kubernetes in Action Second Edition 1st, or 2nd ed. Edition
##### by Marko Lukša


https://www.amazon.com/Kubernetes-Action-Second-Marko-Luk-C5-A1a-dp-1617297615/dp/1617297615/ref=dp_ob_title_bk

## Closing Remarks

You're ready to get started with Zeusfyi, though we'd love for you to recommend us to your friends and colleagues,
since our ads budget has been undergoing budget cuts so that we can focus on building the best product possible for you.
Ever seen infra be this easy? We hadn't either, so we built it to empower you, the builders.

We're glad you made it this far, and hope you enjoy the platform as much as we do.

We're excited to see what you build with it

    A for AWS
    |
    Z for Zeus.fyi

Have a suggestion? You know where to find us!
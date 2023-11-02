---
sidebar_position: 6
---

# Serverless EVM

### Foundry's Anvil as an EVM Sim Environment on Demand

## Free Beta Access

During beta testing we're offering free access to our serverless on-demand anvil execution environments,
which you can use to run Tx simulations or smart contract CI/CD. Not to mention, while also even load balancing
your production traffic through our Adaptive RPC Load Balancer on QuickNode Marketplace, truly all in one.

#### Requires Adaptive RPC Load Balancer Sign Up on QuickNode Marketplace.

## Overview

Serverless EVM environments that you can use your QuickNode, other providers, or even your own self-managed endpoints
to run Tx simulations, smart contract CI/CD, while also even load balancing your production traffic through our Adaptive
RPC Load Balancer on QuickNode Marketplace.

- Each serverless execution environment lasts for up to 10 minutes max, but you can end early to release.
- Each environment contains an anvil service you can use for simulations/forks + automated smart contract CI/CD.
- Each user during beta can run up to 5 concurrent serverless execution environments for free.
- It will automatically convert `hardhat_` prefixed rpc methods to the equivalent `anvil_` method.
- Integrated Go web3 client under `pkg/artemis/web3/client` on our GitHub repo.
- We disabled the embedded Anvil rate limiter, so you can use at full speed. `NO_RATE_LIMITS=true`

```mermaid
flowchart TD
    subgraph ServerlessEnvironment ["Serverless EVM"]
        User[("User\n(up to 5 concurrent environments)")] -->|Request with Header| Session
        Session[("Serverless Execution Environment\n(up to 10 minutes)")] --> AnvilService
        AnvilService[("Anvil Service\n(Simulations/Forks + CI/CD)")]
    end

    subgraph LoadBalancer ["Adaptive RPC Load Balancer"]
        Session -->|Production traffic| QuickNodeMarketplace[("QuickNode Marketplace")]
    end

    subgraph Networks
        LocalHardhat[("Local Hardhat\nNo X-Route-Group Header")] --> Session
        Ethereum[("Ethereum\nX-Route-Group: Table-Name")] --> Session
    end

    LoadBalancer --> ServerlessEnvironment

```

## Serverless Sessions

You need to add a header with a unique session-id name of your choice, per each concurrent serverless environment.

    SERVERLESS_HEADER
        KEY: X-Anvil-Session-Lock-ID
        VALUE: 0 characters < ANY string < 256 characters

    NETWORKS:
        LOCAL_HARDHAT: Don't add X-Route-Group
        
        ETHEREUM: 
            KEY: X-Route-Group
            VALUE: Table-Name (eg. ethereum-mainnet)

## What we show in under 60 seconds...

### Creating a local hardhat serverless session (no X-Route-Group header)

- Calls `eth_chainId` to confirm local hardhat session

```curl
curl --location 'https://iris.zeus.fyi/v1/router' \
--header 'X-Anvil-Session-Lock-ID: local-hardhat-env' \
--header 'Content-Type: application/json' \
--header 'Authorization: Bearer BEARER' \
--data '{
    "jsonrpc": "2.0",
    "method": "eth_chainId",
    "params": [],
    "id": 1
}' 
```

### Creating a mainnet serverless session (with X-Route-Group header)

- Adds `X-Route-Group: ethereum-mainnet` header to the request
- Creates a second serverless session with a different `X-Anvil-Session-Lock-ID` header

```curl
curl --location 'https://iris.zeus.fyi/v1/router' \
--header 'X-Anvil-Session-Lock-ID: test-fork-mainnet' \
--header 'X-Route-Group: ethereum-mainnet' \
--header 'Content-Type: application/json' \
--header 'Authorization: Bearer BEARER' \
--data '{
    "jsonrpc": "2.0",
    "method": "eth_chainId",
    "params": [],
    "id": 1
}' 
```

### Creating an ethereum serverless session (with X-Route-Group header)

#### Using a pinned block number

```json
{
  "jsonrpc": "2.0",
  "method": "anvil_reset",
  "params": [
    {
      "forking": {
        "jsonRpcUrl": "http://localhost:8545",
        "blockNumber": 14390000
      }
    }
  ],
  "id": 1
}
```

#### Using latest block number

```json
{
  "jsonrpc": "2.0",
  "method": "anvil_reset",
  "params": [
    {
      "forking": {
        "jsonRpcUrl": "http://localhost:8545"
      }
    }
  ],
  "id": 1
}
```

#### Calls get block number to confirm session

```json

{
  "jsonrpc": "2.0",
  "method": "eth_blockNumber",
  "params": [],
  "id": 1
}
```

#### How to use the load balancer with your serverless session

- Using your pre-existing routing table header (eg. ethereum-mainnet)

#### How to end your serverless session

- By using a `X-End-Session-Lock-ID` header with your `X-Anvil-Session-Lock-ID` value

## YouTube Walkthrough

<iframe width="1000" height="700" src="https://www.youtube.com/embed/KXkFGW4DGPU?si=ESiYQWXlCj0g4Oqe" title="YouTube video player" frameborder="0" allow="accelerometer; autoplay; clipboard-write; encrypted-media; gyroscope; picture-in-picture; web-share" allowfullscreen></iframe>

## Top Reason to Use the Adaptive RPC Load Balancer?

There's many, but the top one? People are sick of hitting 429 rate limiting and 5xx errors. What's double worse is
paying for the request that failed. So we set out to solve this problem.

Now there's a solution, backed by extensive studies. Turn many endpoints into one super endpoint that handles the scale
you need without the errors. It can handle Nx more requests with N being the number of requests/sec than any single
endpoint can handle. That's in addition to the other significant proven performance gains and error rate reductions you
can expect from our adaptive load balancing technology, see our benchmarking section for details.
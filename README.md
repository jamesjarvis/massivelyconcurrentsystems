# Massively Concurrent Systems Design

[![Go Reference](https://pkg.go.dev/badge/github.com/jamesjarvis/massivelyconcurrentsystems.svg)](https://pkg.go.dev/github.com/jamesjarvis/massivelyconcurrentsystems)
![GitHub Workflow Status (branch)](https://img.shields.io/github/workflow/status/jamesjarvis/massivelyconcurrentsystems/go%20testing/master)

---

## What is this?

This repo is an exercise to see what the limitations of go's concurrency model is, and to see if it is realistic to be able to build systems based on the Massively Concurrent design.

We require 2 things:

1. A method of concurrently executing individual units of work (requests), that are able to depend on other units of work.
   This should allow for more readable code, reduced latency (each unit of work now focuses on just itself, rather than others within a unnecessary batch), and throughput, provided the machine running the program is equipped with enough resources.
2. A method of synchronising these units of work into batchable pipelines at required I/O bottlenecks of the system.
  e.g. If there are 10 requests executing concurrently, and they all need to make the same DB call.
  They should be able to be optimistically batched together into a single DB call, and then split back out into their own units of work.
  This should happen more or less invisibly to the user, and should be a safe, common operation.
  Given this ability, batch sizes to downstream I/O bound resources can be greatly optimised, improving throughput dramatically.
  Another benefit is this allows us to implement rate limiting of downstream RPC calls easily, and if we are able to monitor the saturation of each I/O stage's workers, then we can also scale up/down the number of workers on the fly to maintain an optimal balance of batch sizes and latency.

However, this all falls apart if the system cannot efficiently schedule thousands of concurrently running requests on a single instance.

For reference, we should be looking for throughput of around 50k/s and library latency of less than 2ms per req on a machine with fewer than 20 cores.

---
## Benchmarks

You are able to view the latest benchmark results at [jamesjarvis.github.io/massivelyconcurrentsystems](https://jamesjarvis.github.io/massivelyconcurrentsystems/)

- [deps](https://jamesjarvis.github.io/massivelyconcurrentsystems/bench/deps)
- [pool](https://jamesjarvis.github.io/massivelyconcurrentsystems/bench/pool)

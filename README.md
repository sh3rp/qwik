# Qwik - Filesystem event agent
Qwik is a daemon agent that listens to changes on a configured
file paths and publishes those changes to a message bus.  The
only message bus implementation at the moment is NATS.
## Qwik Start

Prerequisites:
* Install dep (https://github.com/golang/dep)

1. Run 'make'.
2. Modify your qwik.json configuration appropriately.
3. Run 'qwik'.
4. [Optional] Run 'qwiklog' to subscribe to the NATS channel qwik is publishing to.
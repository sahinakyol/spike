## Quickstart

In this guide, you will learn how to build, deploy, and test **SPIKE** from the 
source.

## Prerequisites

This quickstart guide assumes you are using an [Ubuntu Linux][ubuntu] operating system.
The steps may slightly differ if you are using a different operating system.

[ubuntu]: https://ubuntu.com/

Here's the OS details that we are testing this guide on:

```text
DISTRIB_ID=Ubuntu
DISTRIB_RELEASE=24.04
DISTRIB_CODENAME=noble
DISTRIB_DESCRIPTION="Ubuntu 24.04 LTS"
```

In addition, you will need the usual suspects:

* [`git`](https://git-scm.com/)
* [`make`](https://www.gnu.org/software/make/)
* [`node`](https://nodejs.org) (*at least [GitHub Copilot][copilot] requires it on Linux*)
* [`build-essential`](https://packages.ubuntu.com/hirsute/build-essential)
  (*i.e., `sudo apt install build-essential`*)

[copilot]: https://copilot.github.com/ "GitHub Copilot"

## Building SPIRE

To get started let's create a development version of SPIRE. Note that this is
not a production-ready setup. For production, you should follow the 
[official SPIRE documentation][spire-prod].

[spire-prod]: https://spiffe.io/docs/latest/deploying/configuring/

Let's first build SPIRE from the source:

```bash
cd $WORKSPACE
git clone https://github.com/spiffe/spire && cd spire
make build
````

Add the SPIRE binaries to your PATH:

```bash
# ~/.profile
export PATH=$PATH:$WORKSPACE/spire/bin
```

Verify installation:

```bash 
source ~/.profile

spire-server -h
```

Output:

```text
Usage: spire-server [--version] [--help] <command> [<args>]

Available commands are:
    agent                
    bundle               
    entry                
    federation           
    healthcheck          Determines server health status
    jwt                  
    localauthority       
    logger               
    run                  Runs the server
    token                
    upstreamauthority    
    validate             Validates a SPIRE server configuration file
    x509  
```

## Building SPIKE

Next, build **SPIKE binaries:

```bash
cd $WORKSPACE/spike
./hack/build-spike.sh

# Created files:
#   keeper*
#   nexus*
#   spike*
```

## Initialize SPIRE Server

```bash
cd $WORKSPACE/spike
./hack/spire-server-starter.sh
```

## Initialize SPIRE Agent

Once SPIRE Server is running, start the SPIRE Agent:

```bash
./hack/spire-agent-starter.sh
```

## Start SPIKE Components

Then start **SPIKE** components:

Make sure you started the following binaries each run on a specific terminal 
window.

Start the workloads:

```bash
cd $WORKSPACE/spike
./nexus  # Nexus
./keeper # Keeper
```

## Using SPIKE Pilot

Define an alias to **SPIKE** Pilot:

```bash
# ~/.bashrc

# path to the SPIKE Pilot binary (`spike`)
alias spike=$USER/WORKSPACE/spike/spike
```

Run **SPIKE** Pilot and explore the CLI:

```bash
spike

# Sample Output:
# SPIKE v0.1.0
# >> Secure your secrets with SPIFFE: https://spike.ist/ #
# Usage: spike <command> [args...]
# Commands:
#   init
#   put <path> <key=value>...
#   get <path> [-version=<n>]
#   delete <path> [-versions=<n1,n2,...>]
#   undelete <path> [-versions=<n1,n2,...>]
#   list
```

## Testing Out SPIKE

Let test **SPIKE** by creating a secret

```text
# you need to initialize the SPIKE Pilot before you can use it:
spike init

# Register a secret:
spike put /secrets/db-secret username=postgres password=postgres

spike get /secrets/db/secret 
# Wil return:
# username=postgres 
# password=postgres

spike delete /secrets/db-secret # Deleting the current secret

spike get /secrets/db-secret 
# WIll be empty.
```

That's about it.

Enjoy.

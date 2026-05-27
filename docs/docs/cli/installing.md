---
title: Installing
description: Learn how to install the Vydon CLI onto your operating system of choice.
id: installing
hide_title: false
slug: /cli/installing
---

# Installing the CLI

Instructions on how to get the Vydon CLI installed on your local machine.

Vydon is delivered through a Command Line Interface (CLI) to make it easy for developers to use in their native workflows.
The Vydon CLI lets you view accounts, jobs and sync data locally. To get started with Vydon, follow the instructions below to download the CLI.

## MacOS

Homebrew is the simplest way to install vydon CLI on the Mac. This can also be used on Linux, as well as on Windows 10 with Windows Subsystem for Linux.

### Homebrew

Once the project cuts its first tagged release, a Homebrew formula is
published to the [vydon-io/homebrew-tap](https://github.com/vydon-io/homebrew-tap)
repository. If you don't have Homebrew installed, follow these
[instructions](https://docs.brew.sh/Installation). Then run:

```console
brew install vydon-io/tap/vydon
```

If the tap install fails with `404`, no tagged release is published
yet — fall back to the direct download or the Docker image below.
Once installed, keep Vydon up to date with:

```console
brew upgrade
```

## MacOS/Linux Direct Download

Navigate to the Vydon [releases](https://github.com/vydon-io/vydon/releases)
page and pick the binary that matches your machine. If the releases
page is empty, the project has not cut its first tagged release yet —
build the CLI from source in the meantime with `make build/cli` from a
clone of the repo.

After you've downloaded and untarred the tarball, move it into your local bin to make it easy to run. If you're using Windows 10/11, see the Windows section below for more details.

Replace `<VERSION>` with the tag you downloaded (for example `0.4.0`) and pick the archive matching your machine (`darwin_arm64`, `darwin_amd64`, `linux_amd64`, ...).

```console
tar xzf vydon_<VERSION>_darwin_arm64.tar.gz vydon
mv vydon /usr/local/bin/vydon
```

### Verifying your installation

Once you've successfully installed the CLI, verify your installation by following these steps:

1. Open a new terminal window.
2. Type in `vydon help` into your terminal and press enter.
3. If installed successfully, you will see something similar to this help menu

```console
Terminal UI that interfaces with the Vydon system.

Usage:
  vydon [command]

Available Commands:
  accounts    Parent command for account
  completion  Generate the autocompletion script for the specified shell
  connections Parent command for connections
  help        Help about any command
  jobs        Parent command for jobs
  login       Login to Vydon
  sync        One off sync job to local resource
  version     Print the client version information
  whoami      Find out who you are

Flags:
      --api-key string   Vydon API Key. Takes precedence over $VYDON_API_KEY
      --config string    config file (default is $HOME/.vydon/config.yaml)
      --debug            Run in debug mode
  -h, --help             help for vydon
  -v, --version          version for vydon

Use "vydon [command] --help" for more information about a command.
```

Now that you've successfully downloaded the Vydon CLI, you're ready to start building and deploying services. Check out the next section to get familiar with the Vydon CLI commands.

## Windows 10/11 Direct Download

Navigate to Vydon [releases](https://github.com/vydon-io/vydon/releases) page of the CLI repository in the Vydon Github. From there you can choose which binary to download based on your machine's architecture for Windows. Some examples are listed below.

After the download has completed, unzip the contents into a new folder. The most important file is vydon.exe. This can be left here, but a more appropriate place to move it would be to a folder such as: `C:\Vydon` or `C:\Apps\Vydon`

Afterwards, this location needs to be added to the system path. This can be done by going into Settings, and searching for "environment variables" in the search bar. Click "Edit Environment Variables for your Account".

The "Path" variable should be edited. User or System is dependent on your preferences.

### Verifying your installation

Once you've successfully installed the CLI, verify your installation by following these steps:

1. Open a new terminal window.
   - Note: the examples below are using Powershell 5.1.x on Windows 11
2. Type in `vydon help` into your terminal and press enter.
3. If installed successfully, you will see something similar to this help menu

```console
vydon
Terminal UI that interfaces with the Vydon system.

Usage:
  vydon [command]

Available Commands:
  accounts    Parent command for account
  completion  Generate the autocompletion script for the specified shell
  connections Parent command for connections
  help        Help about any command
  jobs        Parent command for jobs
  login       Login to Vydon
  sync        One off sync job to local resource
  version     Print the client version information
  whoami      Find out who you are

Flags:
      --api-key string   Vydon API Key. Takes precedence over $VYDON_API_KEY
      --config string    config file (default is $HOME/.vydon/config.yaml)
      --debug            Run in debug mode
  -h, --help             help for vydon
  -v, --version          version for vydon

Use "vydon [command] --help" for more information about a command.
```

## Docker

Once a release is cut, a matching Docker image is published to GitHub Container Registry. Each versioned image includes the Vydon CLI release with the same version number.

These images wrap the Vydon executable, allowing you to run Vydon subcommands by passing in their names and arguments as part of `docker run`.

The list of published images can be found on [Github](https://github.com/vydon-io/vydon/pkgs/container/vydon%2Fcli). If the package does not yet appear, the project has not cut its first tagged release yet — install via Homebrew or direct download in the meantime.

### Configuration

The container will need further configuration so that Vydon can access configuration files, and possibly source code if the plan is to issue source-code deployments with Vydon CLI.

See the example below for how to login to the CLI, and then view a list of environments in a Vydon account.

```console
docker run -it --rm -p 4242:4242 --mount source=vydoncfg,target=/root/.vydon ghcr.io/vydon-io/vydon/cli:latest login
```

```console
docker run -it --rm --mount source=vydoncfg,target=/root/.vydon ghcr.io/vydon-io/vydon/cli:latest accounts ls
```

The command above will print out a list of environments that are in the account associated with the logged in credentials. Note that the port mapping isn't required here, as that is only necessary during the login flow.

The docker volume is necessary in order to persist the Vydon CLI configuration data between runs. This today is namely used to persist auth data used during the login process so that it can be used with the other CLI commands.

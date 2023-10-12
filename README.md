# bosh-cli-completion

Automatically generated shell autocomplete for [bosh-cli](https://github.com/cloudfoundry/bosh-cli) (bash, zsh, fish, powershell):

- Utilizes [Cobra](https://github.com/spf13/cobra/) for autocompletion, loading the Cobra command tree from `bosh-cli` via reflection.
- Session is compatible with `bosh-cli` (it uses environment variables and flags in the same manner as `bosh-cli`).
- Offers **cacheable suggestions** for deployment names and other parameters, as querying the `Bosh Director API` can be time-consuming and the results don't update frequently.

> [Cobra](https://github.com/spf13/cobra/) is used in many Go projects such as Kubernetes, Hugo, and GitHub CLI to name a few. This list contains a more extensive list of projects using Cobra.

Check out this solution and consider voting for it to be [merged into the main bosh-cli repo](https://github.com/cloudfoundry/bosh-cli/pull/629)

## Installation
 
### Bash

1. Add these lines to `.bashrc`:
   ```shell
   bosh_switch() {
      if [ "$1" == "completion" ] || [ "$1" == "__complete" ]; then
        \bosh-cli-completion "$@"
      else
        \bosh "$@"
      fi
   }
   
   source <(\bosh-cli-completion completion bash)
   alias bosh=bosh_switch
   alias b=bosh_switch
   complete -o default -F __start_bosh b
   ```
2. Restart shell

## Test it

1. Type `bosh -d ` and `TAB key`
2. Wait a moment to give time to query the Bosh Director API, the response will be cached for 15s. 
3. Type `TAB key` again.
4. For example
   ```shell
   b -d cf ssh <TAB>
   api                                                          log-api
   api/eb477fea-77dd-4833-bb62-e9025595f020                     log-api/22074b77-5a65-47c4-add5-fab4f8165b2c
   api/c03f3d40-e878-42d4-85d7-1aeeac2bc96a                     log-api/1b7b4364-0cb5-48b0-bb76-9f96ace90d28
   ...

### Other Shells

```shell
bosh-cli-completion completion -help
```

Generate the autocompletion script for bosh for the specified shell. See each sub-command's help for details on how to use the generated script.

```shell
bosh-cli-completion completion [shell]
```

Available Shells:
- `bash` Generate the autocompletion script for bash
- `fish` Generate the autocompletion script for fish
- `powershell` Generate the autocompletion script for powershell
- `zsh` Generate the autocompletion script for zsh


```shell
bosh-cli-completion completion bash/fish/powershell/zsh --help
```

## Development

### Test

```shell
go test bosh-cli-completion/cmd/completion
```

### Build

```shell
export BBC_VER="v0.0.1-beta-1"
```
```shell
GOOS=linux GOARCH=amd64; go build -o "bosh-cli-completion-$BBC_VER-$GOOS-$GOARCH"
GOOS=linux GOARCH=arm64; go build -o "bosh-cli-completion-$BBC_VER-$GOOS-$GOARCH"

GOOS=darwin GOARCH=amd64; go build -o "bosh-cli-completion-$BBC_VER-$GOOS-$GOARCH"
GOOS=darwin GOARCH=arm64; go build -o "bosh-cli-completion-$BBC_VER-$GOOS-$GOARCH"

GOOS=windows GOARCH=amd64; go build -o "bosh-cli-completion-$BBC_VER-$GOOS-$GOARCH"
GOOS=windows GOARCH=arm64; go build -o "bosh-cli-completion-$BBC_VER-$GOOS-$GOARCH"
```

Use these commands for building a standalone, statically-linked binary using an external linker:
```shell
GOOS=linux GOARCH=amd64; go build -ldflags '-linkmode external -extldflags "-static"' -o "bosh-cli-completion-$BBC_VER-$GOOS-$GOARCH-standalone"
GOOS=linux GOARCH=arm64; go build -ldflags '-linkmode external -extldflags "-static"' -o "bosh-cli-completion-$BBC_VER-$GOOS-$GOARCH-standalone"

GOOS=darwin GOARCH=amd64; go build -ldflags '-linkmode external -extldflags "-static"' -o "bosh-cli-completion-$BBC_VER-$GOOS-$GOARCH-standalone"
GOOS=darwin GOARCH=arm64; go build -ldflags '-linkmode external -extldflags "-static"' -o "bosh-cli-completion-$BBC_VER-$GOOS-$GOARCH-standalone"

GOOS=windows GOARCH=amd64; go build -ldflags '-linkmode external -extldflags "-static"' -o "bosh-cli-completion-$BBC_VER-$GOOS-$GOARCH-standalone"
GOOS=windows GOARCH=arm64; go build -ldflags '-linkmode external -extldflags "-static"' -o "bosh-cli-completion-$BBC_VER-$GOOS-$GOARCH-standalone"
```


### Debug

1. Add these environment variables:
    ```shell
    export BOSH_LOG_LEVEL=debug
    export BOSH_LOG_PATH=~/.bosh/log/bosh-cli-debug.log
    ```
2. Restart shell
3. Watch logs `tail -f ~/.bosh/log/bosh-cli-debug.log`
4. Type `bosh -d ` and `TAB key`

## mycli completion bash

Generate the autocompletion script for bash

### Synopsis

Generate the autocompletion script for the bash shell.

This script depends on the 'bash-completion' package.
If it is not installed already, you can install it via your OS's package manager.

To load completions in your current shell session:

	source <(mycli completion bash)

To load completions for every new session, execute once:

#### Linux:

	mycli completion bash > /etc/bash_completion.d/mycli

#### macOS:

	mycli completion bash > /usr/local/etc/bash_completion.d/mycli

You will need to start a new shell for this setup to take effect.


```
mycli completion bash
```

### Options

```
  -h, --help              help for bash
      --no-descriptions   disable completion descriptions
```

### SEE ALSO

* [mycli completion](mycli_completion.md)	 - Generate the autocompletion script for the specified shell

###### Auto generated by spf13/cobra on 1-Feb-2022

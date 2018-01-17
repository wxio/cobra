# Test 

To demonstrate issue fixed by https://github.com/spf13/cobra/pull/615

The output of the follow two command should be the same

```
go run main.go -t=true sub abc
```

and

```
go run main.go -t sub abc
```

# Output

```
>go run main.go -t=t sub abc
sub called with args [abc]
```

```
>go run main.go -t sub abc
Usage:
   [command]

Available Commands:
  help        Help about any command
  sub

Flags:
  -h, --help     help for this command
  -t, --toggle   Help message for toggle

Use " [command] --help" for more information about a command.
```

The reason for the difference is that in `-t sub`, the `sub` is treated as a param of the `-t`. 

# Discussion

As pointed out in the comments to the PR by @eparis 

> If think you are saying that you have: flag.Value.Type() == "bool" but that flag.NoOptDefVal == "" I believe that shouldn't happen... Can you write a test how that happened.

`flag.NoOptDefVal` should be empty for bools, but at this point in the code it looks like it it.

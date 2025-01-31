## kubemqctl queries attach

Attach to 'queries' channels command

### Synopsis

Attach command allows to display 'queries' channel content for debugging proposes

```
kubemqctl queries attach [flags]
```

### Examples

```

	# attach to all 'queries' channels and output running messages
	kubemqctl queries attach *
	
	# attach to some-query 'queries' channel and output running messages
	kubemqctl queries attach some-query

	# attach to some-queries1 and some-queries2 'queries' channels and output running messages
	kubemqctl queries attach some-queries1 some-queries2 

	# attach to some-queries 'queries' channel and output running messages filter by include regex (some*)
	kubemqctl queries attach some-queries -i some*

	# attach to some-queries 'queries' channel and output running messages filter by exclude regex (not-some*)
	kubemqctl queries attach some-queries -e not-some*

```

### Options

```
  -e, --exclude stringArray   set (regex) strings to exclude
  -h, --help                  help for attach
  -i, --include stringArray   set (regex) strings to include
```

### SEE ALSO

* [kubemqctl queries](kubemqctl_queries.md)	 - Execute KubeMQ 'queries' RPC based commands



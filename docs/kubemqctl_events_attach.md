## kubemqctl events attach

Attach to 'events' channels command

### Synopsis

Attach command allows to display 'events' channel content for debugging proposes

```
kubemqctl events attach [flags]
```

### Examples

```

	# attach to all 'events' channels and output running messages
	kubemqctl events attach *
	
	# attach to some-events 'events' channel and output running messages
	kubemqctl events attach some-events

	# attach to some-events1 and some-events2 'events' channels and output running messages
	kubemqctl events attach some-events1 some-events2 

	# attach to some-events 'events' channel and output running messages filter by include regex (some*)
	kubemqctl events attach some-events -i some*

	# attach to some-events 'events' channel and output running messages filter by exclude regex (not-some*)
	kubemqctl events attach some-events -e not-some*

```

### Options

```
  -e, --exclude stringArray   set (regex) strings to exclude
  -h, --help                  help for attach
  -i, --include stringArray   set (regex) strings to include
```

### SEE ALSO

* [kubemqctl events](kubemqctl_events.md)	 - Execute KubeMQ 'events' Pub/Sub commands



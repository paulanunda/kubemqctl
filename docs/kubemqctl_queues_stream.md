## kubemqctl queues stream

Stream a message from a queue command

### Synopsis

Stream command allows to receive message from a queue in push mode response an appropriate action

```
kubemqctl queues stream [flags]
```

### Examples

```

	# Stream 'queues' message in transaction mode
	kubemqctl queue stream q1

	# Stream 'queues' message in transaction mode with visibility set to 120 seconds and wait time of 180 seconds
	kubemqctl queue stream q1 -v 120 -w 180

```

### Options

```
  -h, --help             help for stream
  -v, --visibility int   set initial visibility seconds (default 30)
  -w, --wait int         set how many seconds to wait for 'queues' messages (default 60)
```

### SEE ALSO

* [kubemqctl queues](kubemqctl_queues.md)	 - Execute KubeMQ 'queues' commands



## kubemqctl queues peek

Peek a messages from a queue channel command

### Synopsis

Peek command allows to peak one or many messages from a queue channel without removing them from the queue

```
kubemqctl queues peek [flags]
```

### Examples

```

	# Peek 1 messages from a queue and wait for 2 seconds (default)
	kubemqctl queue peek some-channel

	# Peek 3 messages from a queue and wait for 5 seconds
	kubemqctl queue peek some-channel -m 3 -w 5

```

### Options

```
  -h, --help           help for peek
  -m, --messages int   set how many messages we want to peek from queue (default 1)
  -w, --wait int       set how many seconds to wait for peeking queue messages (default 2)
```

### SEE ALSO

* [kubemqctl queues](kubemqctl_queues.md)	 - Execute KubeMQ 'queues' commands



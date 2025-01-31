## kubemqctl queues receive

Receive a messages from a queue channel command

### Synopsis

Receive command allows to receive one or many messages from a queue channel

```
kubemqctl queues receive [flags]
```

### Examples

```

	# Receive 1 messages from a queue channel q1 and wait for 2 seconds (default)
	kubemqctl queue receive q1

	# Receive 3 messages from a queue channel and wait for 5 seconds
	kubemqctl queue receive q1 -m 3 -t 5

	# Watching 'queues' channel messages
	kubemqctl queue receive q1 -w

```

### Options

```
  -h, --help               help for receive
  -m, --messages int       set how many messages we want to get from a queue (default 1)
  -t, --wait-timeout int   set how many seconds to wait for 'queues' messages (default 2)
  -w, --watch              set watch on 'queues' channel
```

### SEE ALSO

* [kubemqctl queues](kubemqctl_queues.md)	 - Execute KubeMQ 'queues' commands



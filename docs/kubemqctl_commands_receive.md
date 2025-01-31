## kubemqctl commands receive

Receive a message from 'commands' channel command

### Synopsis

Receive (Subscribe) command allows to consume a message from 'commands' channel and response with appropriate reply

```
kubemqctl commands receive [flags]
```

### Examples

```

	# Receive commands from a 'commands' channel (blocks until next message)
	kubemqctl commands receive some-channel

	# Receive commands from a 'commands' channel with group (blocks until next message)
	kubemqctl commands receive some-channel -g G1

```

### Options

```
  -a, --auto-response   set auto response executed command for each command received
  -g, --group string    set 'commands' channel consumer group (load balancing)
  -h, --help            help for receive
```

### SEE ALSO

* [kubemqctl commands](kubemqctl_commands.md)	 - Execute KubeMQ 'commands' RPC commands



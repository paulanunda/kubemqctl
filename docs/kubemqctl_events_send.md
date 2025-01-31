## kubemqctl events send

Send messages to an 'events' channel command

### Synopsis

Send command allows to send (publish) one or many messages to an 'events' channel

```
kubemqctl events send [flags]
```

### Examples

```

	# Send (Publish) message to a 'events' channel
	kubemqctl events send some-channel some-message
	
	# Send (Publish) message to a 'events' channel with metadata
	kubemqctl events send some-channel some-message --metadata some-metadata
	
	# Send (Publish) batch of 10 messages to a 'events' channel
	kubemqctl events send some-channel some-message -m 10

```

### Options

```
  -h, --help              help for send
  -m, --messages int      set how many 'events' messages to send (default 1)
      --metadata string   set message metadata field
```

### SEE ALSO

* [kubemqctl events](kubemqctl_events.md)	 - Execute KubeMQ 'events' Pub/Sub commands



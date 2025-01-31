## kubemqctl events_store send

Send messages to an 'events store' channel command

### Synopsis

Send command allows to send (publish) one or many messages to an 'events store' channel

```
kubemqctl events_store send [flags]
```

### Examples

```

	# Send (Publish) message to an 'events store' channel
	kubemqctl events_store send some-channel some-message
	
	# Send (Publish) message to an 'events store' channel with metadata
	kubemqctl events_store send some-channel some-message --metadata some-metadata

	# Send 10 messages to an 'events store' channel
	kubemqctl events_store send some-channel some-message -m 10

```

### Options

```
  -h, --help              help for send
  -m, --messages int      set how many 'events store' messages to send (default 1)
      --metadata string   set message metadata field
```

### SEE ALSO

* [kubemqctl events_store](kubemqctl_events_store.md)	 - Execute KubeMQ 'events_store' Pub/Sub commands



## kubemqctl queues

Execute KubeMQ 'queues' commands

### Synopsis

Execute KubeMQ 'queues' commands

```
kubemqctl queues [flags]
```

### Examples

```

	# Execute send 'queues' command
	kubemqctl queues send

	# Execute attached to 'queues' command
	kubemqctl queues attach

	# Execute receive 'queues' command
	kubemqctl queues receive
	
	# Execute list 'queues' command
	kubemqctl queues list

	# Execute peek 'queues' command
	kubemqctl queues peak

	# Execute ack 'queues' command
	 kubemqctl queues ack

	# Execute stream 'queues' command
	kubemqctl queues stream

```

### Options

```
  -h, --help   help for queues
```

### SEE ALSO

* [kubemqctl](kubemqctl.md)	 - 
* [kubemqctl queues ack](kubemqctl_queues_ack.md)	 - Ack all messages in a 'queues' channel
* [kubemqctl queues attach](kubemqctl_queues_attach.md)	 - Attach to 'queues' channels command
* [kubemqctl queues list](kubemqctl_queues_list.md)	 - Get a list of 'queues' channels / clients command
* [kubemqctl queues peek](kubemqctl_queues_peek.md)	 - Peek a messages from a queue channel command
* [kubemqctl queues receive](kubemqctl_queues_receive.md)	 - Receive a messages from a queue channel command
* [kubemqctl queues send](kubemqctl_queues_send.md)	 - Send a message to a queue channel command
* [kubemqctl queues stream](kubemqctl_queues_stream.md)	 - Stream a message from a queue command



package events

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	"github.com/kubemq-io/kubemqctl/pkg/config"
	"github.com/kubemq-io/kubemqctl/pkg/k8s"
	"github.com/kubemq-io/kubemqctl/pkg/kubemq"
	"github.com/kubemq-io/kubemqctl/pkg/utils"
	"github.com/spf13/cobra"
)

type EventsSendOptions struct {
	cfg       *config.Config
	transport string
	channel   string
	message   string
	metadata  string
	messages  int
}

var eventsSendExamples = `
	# Send (Publish) message to a 'events' channel
	kubemqctl events send some-channel some-message
	
	# Send (Publish) message to a 'events' channel with metadata
	kubemqctl events send some-channel some-message --metadata some-metadata
	
	# Send (Publish) batch of 10 messages to a 'events' channel
	kubemqctl events send some-channel some-message -m 10
`
var eventsSendLong = `Send command allows to send (publish) one or many messages to an 'events' channel`
var eventsSendShort = `Send messages to an 'events' channel command`

func NewCmdEventsSend(ctx context.Context, cfg *config.Config) *cobra.Command {
	o := &EventsSendOptions{
		cfg: cfg,
	}
	cmd := &cobra.Command{

		Use:     "send",
		Aliases: []string{"s"},
		Short:   eventsSendShort,
		Long:    eventsSendLong,
		Example: eventsSendExamples,
		Run: func(cmd *cobra.Command, args []string) {
			ctx, cancel := context.WithCancel(ctx)
			defer cancel()
			utils.CheckErr(o.Complete(args, cfg.ConnectionType), cmd)
			utils.CheckErr(o.Validate())
			utils.CheckErr(k8s.SetTransport(ctx, cfg))
			utils.CheckErr(o.Run(ctx))
		},
	}
	cmd.PersistentFlags().StringVarP(&o.metadata, "metadata", "", "", "set message metadata field")
	cmd.PersistentFlags().IntVarP(&o.messages, "messages", "m", 1, "set how many 'events' messages to send")

	return cmd
}

func (o *EventsSendOptions) Complete(args []string, transport string) error {
	o.transport = transport
	if len(args) >= 2 {
		o.channel = args[0]
		o.message = args[1]
		return nil
	}
	return fmt.Errorf("missing arguments, must be 2 arguments, channel and a message")
}

func (o *EventsSendOptions) Validate() error {
	return nil
}

func (o *EventsSendOptions) Run(ctx context.Context) error {
	client, err := kubemq.GetKubeMQClient(ctx, o.transport, o.cfg)
	if err != nil {
		return fmt.Errorf("create kubemq client, %s", err.Error())
	}

	defer func() {
		client.Close()
	}()
	for i := 1; i <= o.messages; i++ {
		msg := client.E().
			SetChannel(o.channel).
			SetId(uuid.New().String()).
			SetBody([]byte(fmt.Sprintf("%s - (%d)", o.message, i))).
			SetMetadata(o.metadata)
		err = msg.Send(ctx)
		if err != nil {
			return fmt.Errorf("sending 'events' message, %s", err.Error())
		}
		utils.Printlnf("[message: %d] [channel: %s] [client id: %s] -> {id: %s, metadata: %s, body: %s}", i, msg.Channel, msg.ClientId, msg.Id, msg.Metadata, msg.Body)
	}
	return nil
}

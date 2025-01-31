package commands

import (
	"context"
	"fmt"
	"github.com/AlecAivazis/survey/v2"
	"github.com/kubemq-io/kubemqctl/pkg/config"
	"github.com/kubemq-io/kubemqctl/pkg/k8s"
	"github.com/kubemq-io/kubemqctl/pkg/kubemq"
	"github.com/kubemq-io/kubemqctl/pkg/utils"
	"github.com/spf13/cobra"
	"os"
	"text/tabwriter"
	"time"
)

type CommandsReceiveOptions struct {
	cfg          *config.Config
	transport    string
	channel      string
	group        string
	autoResponse bool
}

var commandsReceiveExamples = `
	# Receive commands from a 'commands' channel (blocks until next message)
	kubemqctl commands receive some-channel

	# Receive commands from a 'commands' channel with group (blocks until next message)
	kubemqctl commands receive some-channel -g G1
`
var commandsReceiveLong = `Receive (Subscribe) command allows to consume a message from 'commands' channel and response with appropriate reply`
var commandsReceiveShort = `Receive a message from 'commands' channel command`

func NewCmdCommandsReceive(ctx context.Context, cfg *config.Config) *cobra.Command {
	o := &CommandsReceiveOptions{
		cfg: cfg,
	}
	cmd := &cobra.Command{

		Use:     "receive",
		Aliases: []string{"r", "rec", "subscribe", "sub"},
		Short:   commandsReceiveShort,
		Long:    commandsReceiveLong,
		Example: commandsReceiveExamples,
		Run: func(cmd *cobra.Command, args []string) {
			ctx, cancel := context.WithCancel(ctx)
			defer cancel()
			utils.CheckErr(o.Complete(args, cfg.ConnectionType), cmd)
			utils.CheckErr(o.Validate())
			utils.CheckErr(k8s.SetTransport(ctx, cfg))
			utils.CheckErr(o.Run(ctx))
		},
	}

	cmd.PersistentFlags().StringVarP(&o.group, "group", "g", "", "set 'commands' channel consumer group (load balancing)")
	cmd.PersistentFlags().BoolVarP(&o.autoResponse, "auto-response", "a", false, "set auto response executed command for each command received")
	return cmd
}

func (o *CommandsReceiveOptions) Complete(args []string, transport string) error {
	o.transport = transport
	if len(args) >= 1 {
		o.channel = args[0]
		return nil
	}
	return fmt.Errorf("missing channel argument")
}

func (o *CommandsReceiveOptions) Validate() error {
	return nil
}

func (o *CommandsReceiveOptions) Run(ctx context.Context) error {
	client, err := kubemq.GetKubeMQClient(ctx, o.transport, o.cfg)
	if err != nil {
		return fmt.Errorf("create kubemq client, %s", err.Error())

	}
	defer func() {
		client.Close()
	}()
	w := tabwriter.NewWriter(os.Stdout, 0, 0, 1, ' ', tabwriter.TabIndent)

	errChan := make(chan error, 1)
	commandsChan, err := client.SubscribeToCommands(ctx, o.channel, o.group, errChan)

	if err != nil {
		utils.Println(fmt.Errorf("receive commands messages, %s", err.Error()).Error())
	}
	for {
		utils.Println("waiting for the next command message...")
		select {
		case command, opened := <-commandsChan:
			if !opened {
				utils.Println("server disconnected")
				return nil
			}
			fmt.Fprintf(w, "[channel: %s]\t[id: %s]\t[metadata: %s]\t[body: %s]\n", command.Channel, command.Id, command.Metadata, command.Body)
			w.Flush()
			if o.autoResponse {
				err = client.R().SetRequestId(command.Id).SetExecutedAt(time.Now()).SetResponseTo(command.ResponseTo).Send(ctx)
				if err != nil {
					return err
				}
				utils.Println("auto execution sent executed response ")
				continue
			}
			var isExecuted bool
			prompt := &survey.Confirm{
				Renderer: survey.Renderer{},
				Message:  "Set executed ?",
				Help:     "",
			}
			err := survey.AskOne(prompt, &isExecuted)

			if err != nil {
				return err
			}
			if isExecuted {
				err = client.R().SetRequestId(command.Id).SetExecutedAt(time.Now()).SetResponseTo(command.ResponseTo).Send(ctx)
				if err != nil {
					return err
				}
				continue
			}
			err = client.R().SetRequestId(command.Id).SetError(fmt.Errorf("commnad not executed")).SetResponseTo(command.ResponseTo).Send(ctx)
			if err != nil {
				return err
			}
		case <-ctx.Done():
			return nil
		}
	}

}

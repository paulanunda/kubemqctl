package events_store

import (
	"context"
	"github.com/kubemq-io/kubemqctl/pkg/config"
	"github.com/spf13/cobra"
)

var eventsExamples = `
	# Execute send 'events store' command 
	kubemqctl events_store send

	# Execute receive 'events store' command
	kubemqctl events_store receive

	# Execute attach to 'events store' command
	 kubemqctl events_store attach

	# Execute list of 'events store' channels command
 	kubemqctl events_store list
`
var eventsLong = `Execute KubeMQ 'events_store' Pub/Sub commands`
var eventsShort = `Execute KubeMQ 'events_store' Pub/Sub commands`

// NewCmdCreate returns new initialized instance of create sub command
func NewCmdEventsStore(ctx context.Context, cfg *config.Config) *cobra.Command {
	cmd := &cobra.Command{
		Use:       "events_store",
		Aliases:   []string{"es"},
		Short:     eventsLong,
		Long:      eventsShort,
		Example:   eventsExamples,
		ValidArgs: []string{"send", "receive", "attach", "list"},
		Run: func(cmd *cobra.Command, args []string) {
			cmd.Help()
		},
	}
	cmd.AddCommand(NewCmdEventsStoreSend(ctx, cfg))
	cmd.AddCommand(NewCmdEventsStoreReceive(ctx, cfg))
	cmd.AddCommand(NewCmdEventsStoreAttach(ctx, cfg))
	cmd.AddCommand(NewCmdEventsStoreList(ctx, cfg))

	return cmd
}

package events_store

import (
	"context"
	"fmt"
	"github.com/kubemq-io/kubemqctl/pkg/attach"
	"github.com/kubemq-io/kubemqctl/pkg/config"
	"github.com/kubemq-io/kubemqctl/pkg/k8s"
	"github.com/kubemq-io/kubemqctl/pkg/utils"
	"github.com/spf13/cobra"
)

type EventsStoreAttachOptions struct {
	cfg       *config.Config
	transport string
	include   []string
	exclude   []string
	resources []string
}

var eventsAttachExamples = `
	# Attach to all events store channels and output running messages
	kubemqctl events attach *
	
	# Attach to some-events-store 'events store' channel and output running messages
	kubemqctl events_store attach some-events-store

	# Attach to some-events-store1 and some-events-store2 'events store' channels and output running messages
	kubemqctl events attach some-events-store1 some-events-store2 

	# Attach to some-events-store 'events store' channel and output running messages filter by include regex (some*)
	kubemqctl events attach some-events -i some*

	# Attach to some-events-store 'events store' channel and output running messages filter by exclude regex (not-some*)
	kubemqctl events attach some-events -e not-some*
`
var eventsAttachLong = `Attach command allows to display 'events store' channel content for debugging proposes`
var eventsAttachShort = `Attach to events store channels command`

func NewCmdEventsStoreAttach(ctx context.Context, cfg *config.Config) *cobra.Command {
	o := &EventsStoreAttachOptions{
		cfg: cfg,
	}
	cmd := &cobra.Command{

		Use:     "attach",
		Aliases: []string{"a", "att", "at"},
		Short:   eventsAttachShort,
		Long:    eventsAttachLong,
		Example: eventsAttachExamples,
		Run: func(cmd *cobra.Command, args []string) {
			ctx, cancel := context.WithCancel(ctx)
			defer cancel()
			utils.CheckErr(k8s.SetTransport(ctx, cfg))
			utils.CheckErr(o.Complete(args, cfg.ConnectionType), cmd)
			utils.CheckErr(o.Validate())
			utils.CheckErr(o.Run(ctx))
		},
	}
	cmd.PersistentFlags().StringArrayVarP(&o.include, "include", "i", []string{}, "set (regex) strings to include")
	cmd.PersistentFlags().StringArrayVarP(&o.exclude, "exclude", "e", []string{}, "set (regex) strings to exclude")
	return cmd
}

func (o *EventsStoreAttachOptions) Complete(args []string, transport string) error {
	o.transport = transport
	if len(args) == 0 {
		return fmt.Errorf("missing channel argument")

	}

	for _, a := range args {
		rsc := fmt.Sprintf("events_store/%s", a)
		o.resources = append(o.resources, rsc)
		utils.Printlnf("adding '%s' to attach list", a)
	}
	return nil
}

func (o *EventsStoreAttachOptions) Validate() error {
	return nil
}

func (o *EventsStoreAttachOptions) Run(ctx context.Context) error {
	err := attach.Run(ctx, o.cfg, o.resources, o.include, o.exclude)
	if err != nil {
		return err
	}
	<-ctx.Done()
	return nil
}

package logstream

import (
	"code.cloudfoundry.org/go-loggregator"
	"code.cloudfoundry.org/go-loggregator/conversion"
	"code.cloudfoundry.org/go-loggregator/rpc/loggregator_v2"
	"context"
	"github.com/cloudfoundry/sonde-go/events"
)

type Streamer interface {
	Stream(ctx context.Context, req *loggregator_v2.EgressBatchRequest) loggregator.EnvelopeStream
}

type V2Adapter struct {
	streamer Streamer
	cancel context.CancelFunc
}

func NewV2Adapter(s Streamer) *V2Adapter {
	return &V2Adapter{
		streamer: s,
	}
}

func (a *V2Adapter) Firehose(subscriptionID string) chan *events.Envelope {
	ctx, cancel := context.WithCancel(context.Background())
	a.cancel = cancel

	es := a.streamer.Stream(ctx, &loggregator_v2.EgressBatchRequest{
		ShardId: subscriptionID,
		Selectors: []*loggregator_v2.Selector{
			{
				Message: &loggregator_v2.Selector_Counter{
					Counter: &loggregator_v2.CounterSelector{},
				},
			},
			{
				Message: &loggregator_v2.Selector_Gauge{
					Gauge: &loggregator_v2.GaugeSelector{},
				},
			},
			{
				Message: &loggregator_v2.Selector_Timer{
					Timer: &loggregator_v2.TimerSelector{},
				},
			},
		},
	})

	var msgs = make(chan *events.Envelope, 100)
	go func() {
		for ctx.Err() == nil {
			for _, e := range es() {
				for _, v1e := range conversion.ToV1(e) {
					msgs <- v1e
				}
			}
		}
		close(msgs)
	}()

	return msgs
}

func (a *V2Adapter) Close() {
	a.cancel()
}

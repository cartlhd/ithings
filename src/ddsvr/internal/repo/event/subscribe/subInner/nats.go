package subInner

import (
	"context"
	"github.com/i-Things/things/shared/clients"
	"github.com/i-Things/things/shared/conf"
	"github.com/i-Things/things/shared/devices"
	"github.com/i-Things/things/shared/events"
	"github.com/i-Things/things/shared/events/topics"
	"github.com/i-Things/things/shared/traces"
	"github.com/nats-io/nats.go"
	"github.com/zeromicro/go-zero/core/logx"
)

type (
	NatsClient struct {
		client *nats.Conn
	}
)

const (
	ThingsDDDeliverGroup = "things_dd_group"
)

func newNatsClient(conf conf.NatsConf) (SubInner, error) {
	nc, err := clients.NewNatsClient(conf)
	if err != nil {
		return nil, err
	}
	return &NatsClient{client: nc}, nil
}

func (n *NatsClient) SubToDevMsg(handle Handle) error {
	_, err := n.client.QueueSubscribe(topics.DeviceDownAll, ThingsDDDeliverGroup,
		events.NatsSubscription(func(ctx context.Context, msg []byte, natsMsg *nats.Msg) error {
			topic, payload := devices.GetPublish(msg)
			//给设备回包之前，将链路信息span推送至jaeger
			_, span := traces.StartSpan(ctx, topics.DeviceDownAll, "")
			logx.WithContext(ctx).Infof("ddsvr.mqtt.SubDevMsg topic:%s payload:%v",
				topic, string(payload))
			defer span.End()
			return handle(ctx).PublishToDev(topic, payload)
		}))
	return err
}

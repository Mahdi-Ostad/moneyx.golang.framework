package outbox

import (
	"reflect"

	"gofr.dev/pkg/gofr"
	"google.golang.org/protobuf/proto"
)

const (
	outboxRepository = "StorageOutboxMessage"
)

type outboxMessage struct {
	messageName string `bson:"messageName"`
	data        []byte `bson:"data"`
	ID          string `bson:"_id,omitempty" json:"id"`
	lockTime    int64  `bson:"lockTime"`
}

func Publish(ctx *gofr.Context, message proto.Message) error {
	protoMessage, err := proto.Marshal(message)
	if err != nil {
		return err
	}
	messageType := reflect.TypeOf(message)

	_, err = ctx.Mongo.InsertOne(ctx, outboxRepository, outboxMessage{
		data:        protoMessage,
		messageName: messageType.String(),
		lockTime:    0,
	})
	return err
}

func PublishMany(ctx *gofr.Context, messages []proto.Message) error {
	protoMessages := make([]any, len(messages))
	for _, message := range messages {
		protoMessage, err := proto.Marshal(message)
		if err != nil {
			return err
		}
		messageType := reflect.TypeOf(message)
		protoMessages = append(protoMessages, outboxMessage{
			data:        protoMessage,
			messageName: messageType.String(),
			lockTime:    0,
		})
	}

	_, err := ctx.Mongo.InsertMany(ctx, outboxRepository, protoMessages)
	return err
}

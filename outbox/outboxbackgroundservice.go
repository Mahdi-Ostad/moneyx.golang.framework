package outbox

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"gofr.dev/pkg/gofr"
	"moneyx.golang.framework/config"
	"moneyx.golang.framework/logger"
	RoutineManagement "moneyx.golang.framework/routinemanagement"
	"moneyx.golang.framework/utils"
)

// type OutboxManager struct {
// 	ch chan

// }

const (
	outboxProccessorDelay = time.Millisecond * 200
)

func OutboxBackgroundService(ctx context.Context, logger *logger.MoneyxLog, contextMaker func(context.Context) *gofr.Context) {
	microserviceName := config.GetConfig("MICROSERVICE_NAME")
	logger.Infof("Outbox Background Service is starting on client '%s' ...", microserviceName)
	retryPolicy := utils.NewRetryPolicy(time.Second*2, func(retryCount int, err error) {
		logger.Errorf("Retrying in 2 seconds (Reason: %v) (Retry count: %d)", err, retryCount)
	})
	retryPolicy.RetryForever(func() error {
		for {
			select {
			case <-ctx.Done():
				return nil
			default:
				newContext := contextMaker(ctx)
				var messages []outboxMessage
				now := time.Now().UnixMilli()
				count, err := newContext.Mongo.UpdateMany(newContext.Context, outboxRepository, bson.D{
					{Key: "lockTime", Value: bson.D{{Key: "$lt", Value: now}}},
				}, bson.D{
					{Key: "$set", Value: bson.D{{Key: "lockTime", Value: now}}},
				})
				if count == 0 {
					break
				}
				if err != nil {
					return err
				}
				err = newContext.Mongo.Find(newContext.Context, outboxRepository, bson.D{
					{Key: "lockTime", Value: now},
				}, &messages)
				if err != nil {
					return err
				}
				publishMessageFunc := func(msg outboxMessage) error {
					err := newContext.PubSub.Publish(ctx, msg.messageName, msg.data)
					return err
				}
				publishMessageOnFail := func(err error) {
					logger.Errorf("an error has occurred while processing Saga State outbox: %v", err)
				}
				publishMessage := RoutineManagement.NewFanModel(publishMessageFunc, publishMessageOnFail)
				cleanOutboxFunc := func(msg outboxMessage) error {
					_, err := newContext.Mongo.DeleteOne(newContext.Context, outboxRepository, bson.D{
						{Key: "_id", Value: msg.ID},
					})
					return err
				}
				cleanOutboxOnFail := func(err error) {
					logger.Errorf("an error has occurred while processing Saga State outbox: %v", err)
				}
				cleanOutboxMessage := RoutineManagement.NewFanModel(cleanOutboxFunc, cleanOutboxOnFail)
				fans := []RoutineManagement.FanModel[outboxMessage]{
					*publishMessage,
					*cleanOutboxMessage,
				}
				RoutineManagement.FanInFanOut(messages, fans)
			}
			time.Sleep(outboxProccessorDelay)
		}
	})
}

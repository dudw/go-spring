/*
 * Copyright 2012-2019 the original author or authors.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *      https://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package StarterRabbitMQConsumer

import (
	"context"

	"github.com/go-spring/spring-boost/log"
	"github.com/go-spring/spring-boost/util"
	"github.com/go-spring/spring-core/gs"
	"github.com/go-spring/spring-core/mq"
	"github.com/go-spring/starter-rabbitmq/server"
)

func init() {
	gs.Object(new(Starter)).Export((*gs.AppEvent)(nil))
}

type Starter struct {
	Server *StarterRabbitMQServer.AMQPServer `autowire:""`
}

func (starter *Starter) OnStartApp(ctx gs.Environment) {

	cMap := map[string][]mq.Consumer{}
	{
		var consumers []mq.Consumer
		err := ctx.GetBean(&consumers)
		util.Panic(err).When(err != nil)

		var bindConsumers *gs.Consumers
		err = ctx.GetBean(&bindConsumers)
		util.Panic(err).When(err != nil)

		bindConsumers.ForEach(func(c mq.Consumer) {
			consumers = append(consumers, c)
		})

		for _, consumer := range consumers {
			for _, topic := range consumer.Topics() {
				cMap[topic] = append(cMap[topic], consumer)
			}
		}
	}

	go func() {
		// TODO 使用 goroutine 池提高消费速率
		for topic, consumers := range cMap {
			delivery, err := starter.Server.Channel.Consume(
				topic, // queue
				"",    // consumer
				true,  // auto-ack
				false, // exclusive
				false, // no-local
				false, // no-wait
				nil,   // args
			)
			if err != nil {
				log.Error(err)
				continue
			}
			d := <-delivery
			msg := mq.NewMessage().WithBody(d.Body).WithTopic(topic)
			for _, c := range consumers {
				c.Consume(context.TODO(), msg)
			}
		}
	}()
}

func (starter *Starter) OnStopApp(ctx context.Context) {

}

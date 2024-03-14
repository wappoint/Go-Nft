/*
 * Licensed to the Apache Software Foundation (ASF) under one or more
 * contributor license agreements.  See the NOTICE file distributed with
 * this work for additional information regarding copyright ownership.
 * The ASF licenses this file to You under the Apache License, Version 2.0
 * (the "License"); you may not use this file except in compliance with
 * the License.  You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package mq

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/apache/rocketmq-clients/golang"
	"github.com/apache/rocketmq-clients/golang/credentials"
)

const (
	Topic     = "xxxxxx"
	GroupName = "xxxxxx"
	Endpoint  = "xxxxxx"
	NameSpace = "xxxxxx"
	AccessKey = "xxxxxx"
	SecretKey = "xxxxxx"
)

var (
	// maximum waiting time for receive func
	awaitDuration = time.Second * 5
	// maximum number of messages received at one time
	maxMessageNum int32 = 16
	// invisibleDuration should > 20s
	invisibleDuration = time.Second * 20
	// receive messages in a loop
)

func main() {
	// log to console
	err := os.Setenv("mq.consoleAppender.enabled", "true")
	if err != nil {
		return
	}
	golang.ResetLogger()
	// new simpleConsumer instance
	simpleConsumer, err := golang.NewSimpleConsumer(&golang.Config{
		Endpoint:      Endpoint,
		ConsumerGroup: GroupName,
		NameSpace:     NameSpace,
		Credentials: &credentials.SessionCredentials{
			AccessKey:    AccessKey,
			AccessSecret: SecretKey,
		},
	},
		golang.WithAwaitDuration(awaitDuration),
		golang.WithSubscriptionExpressions(map[string]*golang.FilterExpression{
			Topic: golang.SUB_ALL,
		}),
	)
	if err != nil {
		log.Fatal(err)
	}
	// start simpleConsumer
	err = simpleConsumer.Start()
	if err != nil {
		log.Fatal(err)
	}
	// gracefule stop simpleConsumer
	defer func(simpleConsumer golang.SimpleConsumer) {
		err := simpleConsumer.GracefulStop()
		if err != nil {
			log.Fatal(err)
		}
	}(simpleConsumer)

	go func() {
		for {
			fmt.Println("start recevie message")
			mvs, err := simpleConsumer.Receive(context.TODO(), maxMessageNum, invisibleDuration)
			if err != nil {
				fmt.Println(err)
			}
			// ack message
			for _, mv := range mvs {
				//TODO 处理消息

				tag := *mv.GetTag()
				if tag == "xxxxxx" {

				} else if tag == "xxxxxx" {

				}

				simpleConsumer.Ack(context.TODO(), mv)
			}
		}
	}()
	// run for a while
	time.Sleep(time.Minute)
}

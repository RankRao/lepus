/*
Copyright 2014-2022 The Lepus Team Group, website: https://www.lepus.cc
Licensed under the GNU General Public License, Version 3.0 (the "GPLv3 License");
You may not use this file except in compliance with the License.
You may obtain a copy of the License at
    https://www.gnu.org/licenses/gpl-3.0.html
Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
Special note:
Please do not use this source code for any commercial purpose,
or use it for commercial purposes after secondary development, otherwise you may bear legal risks.
*/

package kafka

import (
	"fmt"
	"github.com/Shopify/sarama"
	"lepus/src/libary/conf"
	"time"
)

/*
init kafka client
*/

func InitClient() sarama.Client {
	servers := conf.Option["kafka_server"]
	version := conf.Option["kafka_version"]
	config := sarama.NewConfig()
	config.Net.DialTimeout = 3 * time.Second
	config.Net.ReadTimeout = 5 * time.Second
	config.Net.WriteTimeout = 3 * time.Second
	config.Producer.Timeout = 3 * time.Second
	config.Producer.Return.Errors = true
	config.Producer.Partitioner = sarama.NewRandomPartitioner
	configKafkaVersion, err := sarama.ParseKafkaVersion(version)
	if err != nil {
		panic(fmt.Sprintln("Parse kafka version err:", err))
	}
	config.Version = configKafkaVersion
	client, err := sarama.NewClient([]string{servers}, config)
	if err != nil {
		panic(fmt.Sprintln("Init kafka client err:", err))
	}
	return client
}

/*
create new client
*/

func NewClient(servers, version string) sarama.Client {
	config := sarama.NewConfig()
	config.Net.DialTimeout = 3 * time.Second
	config.Net.ReadTimeout = 5 * time.Second
	config.Net.WriteTimeout = 3 * time.Second
	config.Producer.Timeout = 3 * time.Second
	config.Producer.Return.Errors = true
	config.Producer.Partitioner = sarama.NewRandomPartitioner
	configKafkaVersion, err := sarama.ParseKafkaVersion(version)
	if err != nil {
		panic(fmt.Sprintln("Parse kafka version err:", err))
	}
	config.Version = configKafkaVersion
	client, err := sarama.NewClient([]string{servers}, config)
	if err != nil {
		panic(fmt.Sprintln("Init kafka client err:", err))
	}
	return client
}

/*
create new producer from kafka client
*/
func NewProducer(client sarama.Client) (sarama.AsyncProducer, error) {
	producer, err := sarama.NewAsyncProducerFromClient(client)
	if err != nil {
		return nil, err
	}
	return producer, nil
}

/*
create new consumer from kafka client
*/
func NewConsumer(client sarama.Client) (sarama.Consumer, error) {
	consumer, err := sarama.NewConsumerFromClient(client)
	if err != nil {
		return nil, err
	}
	return consumer, nil
}

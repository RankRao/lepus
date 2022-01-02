/*
Copyright 2014-2021 The Lepus Team Group, website: https://www.lepus.cc
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

package redis

import (
	"fmt"
	"github.com/go-redis/redis"
	"lepus/src/libary/conf"
	"time"
)

var redisClient *redis.Client

func InitClient() *redis.Client {
	redisClient = redis.NewClient(&redis.Options{
		Addr:         fmt.Sprintf("%s:%s", conf.Option["redis_host"], conf.Option["redis_port"]),
		Password:     conf.Option["redis_pass"], // no password set
		DB:           0,  // use default DB
		PoolSize:     1000,
		ReadTimeout:  time.Millisecond * time.Duration(200),
		WriteTimeout: time.Millisecond * time.Duration(200),
		IdleTimeout:  time.Second * time.Duration(600),
	})
	_, err := redisClient.Ping().Result()
	if err != nil {
		panic(fmt.Sprintln("Init redis client err:", err))
	}

	return redisClient
}

func CreateClient(host, port, password string) (*redis.Client, error) {
	redisClient = redis.NewClient(&redis.Options{
		Addr:         fmt.Sprintf("%s:%s", host, port),
		Password:     password, // no password set
		DB:           0,        // use default DB
		PoolSize:     1000,
		ReadTimeout:  time.Millisecond * time.Duration(200),
		WriteTimeout: time.Millisecond * time.Duration(200),
		IdleTimeout:  time.Second * time.Duration(600),
	})
	_, err := redisClient.Ping().Result()
	if err != nil {
		return nil, err
	}
	return redisClient, nil
}
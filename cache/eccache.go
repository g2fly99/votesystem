package cache

import (
	"fmt"
	"time"
	"votesystem/models"

	"github.com/beego/beego/v2/core/logs"
	"github.com/go-redis/redis"
)

var (
	gRedisClient *redis.Client
)

func Init(addr, passwd string) {

	if gRedisClient != nil {
		return
	}

	gRedisClient = redis.NewClient(&redis.Options{
		Addr:         addr,
		Password:     passwd,
		PoolSize:     800,
		MinIdleConns: 5,
		IdleTimeout:  180 * time.Second,
		DB:           8,
	})

	res := gRedisClient.Ping()
	if res.Err() != nil {
		logs.Error("redis connect failed:%v", res.String())
		gRedisClient = nil
	} else {
		logs.Info("redis connect success......")
	}
}

type esCacheT struct {
	EsId      int
	Finishd   bool
	StartTime string
}

func ecKeyCreate(ecId int) string {
	return fmt.Sprintf("ecId:%v", ecId)
}

func Save(ec *models.ElectionCampaignT, data string) error {

	esKey := ecKeyCreate(ec.EcId)

	expire := ec.Expire.Sub(time.Now())
	result := gRedisClient.Set(esKey, data, expire)
	if result.Err() != nil {
		logs.Debug("save [%v] to redis:%v,err:%v", ec.EcId, result.Val(), result.Err())
	}

	return nil
}

func Get(ecId int) (string, error) {

	esKey := ecKeyCreate(ecId)

	result := gRedisClient.Get(esKey)
	return result.Result()

}

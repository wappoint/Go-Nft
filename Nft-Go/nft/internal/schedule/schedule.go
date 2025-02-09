package schedule

import (
	"Nft-Go/common/db"
	"Nft-Go/common/util"
	"context"
	"github.com/dubbogo/gost/log/logger"
	"github.com/duke-git/lancet/v2/convertor"
	"github.com/robfig/cron/v3"
	"sync"
	"time"
)

func InitListeningRankAdd() {
	red := db.GetRedis()
	ctx := context.Background()
	c := cron.New(cron.WithSeconds())
	//10秒钟遍历一遍redis的有序集合
	_, err := c.AddFunc("*/10 * * * * *", func() {
		// 获取 rankAdd 哈希表数据
		result, err := red.HGetAll(ctx, "rankAdd").Result()
		if err != nil {
			logger.Error("获取rankAdd失败", err)
			return
		}
		// 设置初始化该键
		today := util.FormatDateForDay(time.Now())
		var wg sync.WaitGroup
		for key, value := range result {
			// 开启协程处理
			wg.Add(1)
			go func(key, value string) {
				defer wg.Done()
				// 获取分布式锁
				locked, err := red.SetNX(ctx, "lock:"+key, 1, 10*time.Second).Result()
				if err != nil || !locked {
					// 获取分布式锁失败
					logger.Error("获取分布式锁失败", err)
					return
				}
				defer red.Del(ctx, "lock:"+key)

				float, err := convertor.ToFloat(value)
				if err != nil {
					logger.Error("转换失败", err)
					return
				}
				_, err = red.ZIncrBy(ctx, today, float, key).Result()
				if err != nil {
					logger.Error("排行榜热度添加失败", err)
					return
				}
				err = red.HSet(ctx, "rankAdd", key, "0").Err()
				if err != nil {
					logger.Error("操作 rankAdd 失败", err)
				}
			}(key, value)
		}
		wg.Wait()
	})
	if err != nil {
		logger.Error("定时任务添加失败")
	}
}

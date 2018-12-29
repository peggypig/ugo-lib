package snowflake

import (
	"errors"
	"fmt"
	"sync"
	"time"
)

/**
41 bit 作为毫秒数 - 41位的长度可以使用69年
10 bit 作为机器编号 （5个bit是数据中心，5个bit的机器ID） - 10位的长度最多支持部署1024个节点
12 bit 作为毫秒内序列号 - 12位的计数顺序号支持每个节点每毫秒产生4096个ID序号
*/

type SnowFlakeConfig struct {
	StartTime        int64  // 系统开始时间截
	WorkerIdBits     uint64 // 机器id所占的位数
	DataCenterIdBits uint64 // 数据中心id所占的位数
	SequenceBits     uint64 // 序列在id中占的位数
	WorkerId         int64  // 工作机器ID
	DataCenterId     int64  // 数据中心ID
}

type SnowFlake struct {
	startTime            int64      // 系统开始时间截
	workerIdBits         uint64     // 机器id所占的位数
	dataCenterIdBits     uint64     // 数据中心id所占的位数
	sequenceBits         uint64     // 序列在id中占的位数
	workerId             int64      // 工作机器ID
	dataCenterId         int64      // 数据中心ID
	maxWorkerId          int64      // 每个数据中心支持的最大机器数目(十进制) 1 << workerIdBits
	maxDataCenterId      int64      // 支持的最大数据中心数目 1 << dataCenterIdBits
	workerIdMoveBits     uint64     // 机器id左移位数 sequenceBits
	dataCenterIdMoveBits uint64     // 数据中心标识id 左移位数 sequenceBits + workerIdBits;
	timestampMoveBits    uint64     // 时间戳左移位数 sequenceBits + workerIdBits + dataCenterIdBits;
	sequenceMask         int64      // 生成序列的掩码(sequenceBits位所对应的最大整数值)
	lastTimestamp        int64      // 上次生成ID的时间
	sequence             int64      // 当前序列
	valid                bool       // 配置参数是否合法
	mutex                sync.Mutex // 锁
	dict                 string
}

// 设置config
func (s *SnowFlake) SetConfig(config SnowFlakeConfig) (err error) {
	if err == nil && config.WorkerIdBits+config.DataCenterIdBits+config.SequenceBits != 22 {
		err = errors.New("sum(WorkerIdBits,DataCenterIdBits,SequenceBits) should be 22")
	}
	if err == nil && config.WorkerId > 1<<config.WorkerIdBits {
		err = errors.New("WorkerId too large")
	}
	if err == nil && config.DataCenterId > 1<<config.DataCenterIdBits {
		err = errors.New("WorkerId too large")
	}
	if err == nil {
		// 参数配置
		s.sequenceBits = config.SequenceBits
		s.dataCenterId = config.DataCenterId
		s.workerId = config.WorkerId
		s.dataCenterIdBits = config.DataCenterIdBits
		s.workerId = config.WorkerId
		s.startTime = config.StartTime
		s.workerIdBits = config.WorkerIdBits
		s.maxWorkerId = -1 ^ (-1 << config.WorkerIdBits)
		s.maxDataCenterId = -1 ^ (-1 << config.DataCenterIdBits)
		s.sequenceMask = -1 ^ (-1 << config.SequenceBits)
		s.workerIdMoveBits = config.SequenceBits
		s.dataCenterIdMoveBits = config.SequenceBits + config.WorkerIdBits
		s.timestampMoveBits = config.SequenceBits + config.WorkerIdBits + config.DataCenterIdBits
		s.lastTimestamp = -1
		s.valid = true
		s.dict = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"
	}
	return
}

// 获得以毫秒为单位的当前时间
func (s *SnowFlake) getCurrentTime() (t int64) {
	return time.Now().UnixNano() / 1e6
}

// 阻塞到下一个毫秒 即 直到获得新的时间戳
func (s *SnowFlake) blockTillNextMillis(lastTimestamp int64) int64 {
	timestamp := s.getCurrentTime()
	for timestamp <= lastTimestamp {
		timestamp = s.getCurrentTime()
	}
	return timestamp
}

func (s *SnowFlake) Snow() (sfs string) {
	var sf int64
	if !s.valid {
		sf = 0
	} else {
		s.mutex.Lock()
		defer s.mutex.Unlock()
		currentTime := s.getCurrentTime()
		//如果当前时间小于上一次ID生成的时间戳: 说明系统时钟回退过 - 这个时候应当抛出异常
		if currentTime < s.lastTimestamp {
			panic(fmt.Sprintf("Clock moved backwards.  Refusing to generate id for %d milliseconds", s.lastTimestamp-currentTime))
		}
		//如果是同一时间生成的，则进行毫秒内序列
		if s.lastTimestamp == currentTime {
			s.sequence = (s.sequence + 1) & s.sequenceMask
			//毫秒内序列溢出 即 序列 > sequenceMask
			if s.sequence == 0 {
				//阻塞到下一个毫秒,获得新的时间戳
				currentTime = s.blockTillNextMillis(s.lastTimestamp)
			}
		} else {
			//时间戳改变，毫秒内序列重置
			s.sequence = 0
		}
		//上次生成ID的时间截
		s.lastTimestamp = currentTime

		//移位并通过或运算拼到一起组成64位的ID
		sf = ((s.lastTimestamp - s.startTime) << s.timestampMoveBits) |
			(s.dataCenterId << s.dataCenterIdMoveBits) |
			(s.workerId << s.workerIdMoveBits) |
			s.sequence
	}
	return s.format(sf)
}

func (s *SnowFlake) format(num int64) string {
	var str []byte
	for {
		var result byte
		number := num % int64(len(s.dict))
		result = s.dict[number]
		str = append([]byte{result},str...)
		num = num / int64(len(s.dict))
		if num == 0 {
			break
		}
	}
	return string(str)
}

package snowflake


import (
	"blog/core/setting"
	"errors"
"sync"
"time"
)

const (
	epoch            = int64(1577808000000) // 初始化时间戳
	workerIdBits     = int64(5)             // workerId 占用的位数
	datacenterIdBits = int64(5)             // dataCenter id 占用的位数
	sequenceBits     = int64(12)            // 序号占用的位数
	timestampBits    = int64(41)            // 时间戳占用的位数

	maxWorkerId     = int64(-1 ^ (-1 << workerIdBits))     // workerId 最大值
	maxDataCenterId = int64(-1 ^ (-1 << datacenterIdBits)) // dataCenterID最大值
	maxSequence     = int64(-1 ^ (-1 << sequenceBits))     // sequence 最大值

	workerIdShift     int64 = sequenceBits                                   // workerId 左移位数 为序列号的位数
	dataCenterIdShift int64 = workerIdBits + sequenceBits                    // dataCenterId左移位数 为序列号位数加workerId位数
	timestampShift    int64 = datacenterIdBits + workerIdBits + sequenceBits // 时间戳左移位数为 dataCenterId位数+worker位数+序列号位数
)

/**
 *  0     | 0000 0000 0000 0000 0000 0000 0000 0000 0000 00000 |  00000     00000      | 0000 0000 0000 |
 *  符号位 ｜         时间戳                                     | workerid datacenterId | sequence    |
 *   1                41                                          5            5           12
 */
type Snowflake struct {
	mutex         sync.Mutex
	workerID      int64
	dataCenterID  int64
	sequence      int64
	lastTimestamp int64
}

func New(setting *setting.Snowflake) (*Snowflake, error) {
	if setting.WorkID < 0 || setting.WorkID > maxWorkerId {
		return nil, errors.New("Worker ID excess of quantity: " + string(setting.WorkID))
	}
	if setting.DataCenterID < 0 || setting.DataCenterID > maxDataCenterId {
		return nil, errors.New("DataCenter ID excess of quantity: " + string(setting.DataCenterID))
	}
	return &Snowflake{
		workerID:      setting.DataCenterID,
		dataCenterID:  setting.DataCenterID,
		sequence:      0,
		lastTimestamp: 0,
	}, nil
}

func (s *Snowflake) NextID() (int64, error) {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	timeStamp := time.Now().UnixNano() / 1000000
	if timeStamp < s.lastTimestamp {
		return 0, errors.New("Clock moved backwards. Refusing for generate id for " + string(timeStamp) + " seconds")
	}
	if timeStamp == s.lastTimestamp {
		// 时间戳相等，看序列号有没有用完，没有序列号加1
		s.sequence = (s.sequence + 1) & maxSequence
		if s.sequence == 0 {
			timeStamp = tileNextMillis(s.lastTimestamp)
		}
	} else {
		s.sequence = 0
	}
	s.lastTimestamp = timeStamp
	return int64((timeStamp-epoch)<<timestampShift | s.dataCenterID<<dataCenterIdShift | s.workerID<<workerIdShift | s.sequence), nil
}

func tileNextMillis(lastTimestamp int64) int64 {
	timesTamp := time.Now().UnixNano()
	for timesTamp <= lastTimestamp {
		timesTamp = time.Now().UnixNano()/ 1000000
	}
	return timesTamp
}


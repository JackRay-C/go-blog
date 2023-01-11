package utils

import (
	"errors"
	"strconv"
	"sync"
	"time"
)

const (
	epoch          = int64(1577808000)                // 时间戳初始化
	sequenceBits   = int64(22)                        // 文件序号占用的位数
	timestampBits  = int64(41)                        // 时间戳占用的位数
	maxSequence    = int64(-1 ^ (-1 << sequenceBits)) // 序号的最大值
	timestampShift = sequenceBits                     // 时间戳左移位数
)

// 0     ｜ 00000000000000000000000000000000000000000 ｜ 0000000000000000000000 ｜
// 符号位1｜ 时间戳位数41  							  ｜ 文件序号位 22
type Generator struct {
	mutex         sync.Mutex
	sequence      int64
	lastTimestamp int64
}

func NewFilenameGenerator() *Generator {
	return &Generator{sequence: 0, lastTimestamp: 0}
}

type Info struct {
	Name      string
	Sequence  int64
	Timestamp int64
	Year      string
	Month     string
	Day       string
}

// NewName return a new id
func (g *Generator) NewName() (*Info, error) {
	g.mutex.Lock()
	defer g.mutex.Unlock()

	// 生成时间戳
	timestamp := time.Now().Unix()
	if timestamp < g.lastTimestamp {
		return nil, errors.New("Clock moved backwards. Refusing for generate id for " + strconv.FormatInt(timestamp, 10) + " seconds")
	}

	// 判断时间戳是否过期，同一个时间戳可以生成2^22个序列
	if timestamp == g.lastTimestamp {
		g.sequence = (g.sequence + 1) & maxSequence
		if g.sequence == 0 {
			timestamp = tileNextMillis(g.lastTimestamp)
		}
	} else {
		g.sequence = 0
	}

	// 将新的时间戳设置为最新的
	g.lastTimestamp = timestamp

	// 解析时间戳的年月日，生产文件名
	unix := time.Unix(timestamp, 0)

	return &Info{
		Name:      strconv.FormatInt(int64((timestamp-epoch)<<timestampShift|g.sequence), 10),
		Sequence:  g.sequence,
		Timestamp: (timestamp - epoch) >> timestampShift,
		Year:      strconv.FormatInt(int64(unix.Year()), 10),
		Month:     strconv.FormatInt(int64(unix.Month()), 10),
		Day:       strconv.FormatInt(int64(unix.Day()), 10),
	}, nil
}

// ParseName parse file name return timestamp and sequence, or return error
func (g *Generator) ParseName(name string) (timestamp int64, sequence int64, err error) {
	// 解析文件名
	parseInt, err := strconv.ParseInt(name, 10, 64)

	if err != nil {
		return 0, 0, err
	}

	// 向右移22位得到时间戳间隔，加上epoch初始时间得到生成文件名时的时间戳，这样就可以判断文件路径 例如: 2022/6/13
	timestamp = int64((parseInt >> timestampShift) + epoch)

	// timestamp左移22位得到新值, 再与文件名"与"运算得到序列号
	sequence = int64(timestamp<<timestampShift) ^ parseInt

	// 返回解析的时间戳和序列号
	return timestamp, sequence, nil
}

func tileNextMillis(lastTimestamp int64) int64 {
	timesTamp := time.Now().Unix()
	for timesTamp <= lastTimestamp {
		timesTamp = time.Now().Unix()
	}
	return timesTamp
}

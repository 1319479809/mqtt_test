package utils

import (
	"errors"
	"sync"
	"time"

	"github.com/1319479809/mqtt_test/utils/slog"
)

var (
	//Userid 交易模块的雪花ID
	FastId   *id
	NormalId *id
)

func init() {
	initIdGenerator()
}

func initIdGenerator() {
	NormalId, _ = NewNormalid(1)
}

// A id struct holds the basic information needed for a FastId generator worker
type id struct {
	sync.Mutex
	timestamp int64
	workerid  int64
	sequence  int64
	timeendmp int64
}

const (
	normalepoch          = int64(1609430400 >> 6)                  //开始时间截 (2020-01-01)
	normalepoch6         = int64(1609430400)                       //开始时间截 (2020-01-01)
	normalworkeridBits   = uint(3)                                 //机器id所占的位数
	normalsequenceBits   = uint(10)                                //序列所占的位数
	normalworkeridMax    = int64(1<<normalworkeridBits) - 1        //支持的最大机器id数量
	normalsequenceMask   = int64(1<<normalsequenceBits) - 1        //
	normalworkeridShift  = normalsequenceBits                      //机器id左移位数
	normaltimestampShift = normalsequenceBits + normalworkeridBits //时间戳左移位数
)

// Newnormalid NewNode returns a new normalId worker that can be used to generate normalId IDs
func NewNormalid(workerid int64) (*id, error) {
	if workerid < 0 || workerid > normalworkeridMax {
		return nil, errors.New("workerid must be between 0 and 1023")
	}
	return &id{
		timestamp: 0,
		workerid:  workerid,
		sequence:  0,
	}, nil
}

// Generate creates and returns a unique normalId ID
// Generate creates and returns a unique normalId ID
func GenerateNormalId() int64 {
	NormalId.Lock()
	defer NormalId.Unlock()
	now := time.Now().Unix() >> 6
	if NormalId.timestamp == now {
		NormalId.sequence = NormalId.sequence + 1
		if NormalId.sequence > normalsequenceMask {
			for now <= NormalId.timestamp {
				time.Sleep(10 * time.Millisecond)
				now = time.Now().Unix() >> 6
			}
			NormalId.sequence = 1
			slog.Cp.Debug().Msgf("GenerateNormalId=================now=%d  timestamp=%d time=%d timeendmp=%d time2=%d", now, NormalId.timestamp, time.Now().Unix(), NormalId.timeendmp, time.Now().Unix()-NormalId.timeendmp)

			slog.Cp.Debug().Msgf("GenerateNormalId==1================== nows=%b", (now-normalepoch)<<normaltimestampShift)
			slog.Cp.Debug().Msgf("GenerateNormalId==2================== workerid=%b", (NormalId.workerid << normalworkeridShift))
			slog.Cp.Debug().Msgf("GenerateNormalId==3================== sequence=%b", (NormalId.sequence))
			slog.Cp.Debug().Msgf("5 nows=%d %d workerid=%d %d sequence=%d", (now - normalepoch), (now-normalepoch)<<normaltimestampShift, NormalId.workerid, (NormalId.workerid << normalworkeridShift), (NormalId.sequence))
		}
	} else {
		NormalId.sequence = 1
	}
	NormalId.timestamp = now
	NormalId.timeendmp = time.Now().Unix()
	r := int64((now-normalepoch)<<normaltimestampShift | (NormalId.workerid << normalworkeridShift) | (NormalId.sequence))
	return r
}

// Generate creates and returns a unique normalId ID
func GenerateNormalId6() int64 {
	NormalId.Lock()
	defer NormalId.Unlock()
	now := time.Now().Unix()
	if NormalId.timestamp == now {
		NormalId.sequence = NormalId.sequence + 1
		if NormalId.sequence > normalsequenceMask {
			for now <= NormalId.timestamp {
				//time.Sleep(time.Second)
				time.Sleep(10 * time.Millisecond)
				now = time.Now().Unix()
			}
			end := int64(2684397885)
			slog.Cp.Debug().Msgf("GenerateNormalId=================now=%d  timestamp=%d ", now, NormalId.timestamp)
			slog.Cp.Debug().Msgf("GenerateNormalId==================== nows=%b", (now-normalepoch6)<<normaltimestampShift)
			slog.Cp.Debug().Msgf("GenerateNormalId=======20055============= nows=%b", (end-normalepoch6)<<normaltimestampShift)
			slog.Cp.Debug().Msgf("GenerateNormalId==2================== workerid=%b", (NormalId.workerid << normalworkeridShift))
			slog.Cp.Debug().Msgf("GenerateNormalId==3================== sequence=%b", (NormalId.sequence))

			slog.Cp.Debug().Msgf("GenerateNormalId==3================== r=%b", int64((now-normalepoch6)<<normaltimestampShift|(NormalId.workerid<<normalworkeridShift)|(NormalId.sequence)))
			slog.Cp.Debug().Msgf("5 nows=%d %d workerid=%d %d sequence=%d", (now - normalepoch6), (now-normalepoch6)<<normaltimestampShift, NormalId.workerid, (NormalId.workerid << normalworkeridShift), (NormalId.sequence))

			NormalId.sequence = 1
		}
	} else {
		NormalId.sequence = 1
	}
	NormalId.timestamp = now
	r := int64((now-normalepoch6)<<normaltimestampShift | (NormalId.workerid << normalworkeridShift) | (NormalId.sequence))
	return r
}

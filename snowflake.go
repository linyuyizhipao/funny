//算法描述：
//
//最高位是符号位，始终为0，不可用。
//41位的时间序列，精确到毫秒级，41位的长度可以使用69年。时间位还有一个很重要的作用是可以根据时间进行排序。
//10位的机器标识，10位的长度最多支持部署1024个节点。
//12位的计数序列号，序列号即一系列的自增id，可以支持同一节点同一毫秒生成多个ID序号，12位的计数序列号支持每个节点每毫秒产生4096个ID序号。

package funny

import (
	"errors"
	"sync"
	"time"
)

//你需要先去了解下 golang中   & |  <<  >> 几个符号产生的运算意义
const (
	workerBits  uint8 = 10   //10bit工作机器的id，如果你发现1024台机器不够那就调大次值
	numberBits  uint8 = 12  //12bit 工作序号，如果你发现1毫秒并发生成4096个唯一id不够请调大次值
	workerMax   int64 = -1 ^ (-1 << workerBits)
	numberMax   int64 = -1 ^ (-1 << numberBits)
	timeShift   = workerBits + numberBits
	workerShift = numberBits
	startTime   int64 = 1525705533000   //雪花算法有效期以该时间为起点往后推69年，有效
)

type Worker struct {
	mu        sync.Mutex
	timestamp int64
	workerId  int64
	number    int64
}

func NewWorker(workerId int64) (*Worker, error) {
	if workerId < 0 || workerId > workerMax {
		return nil, errors.New("Worker ID excess of quantity")
	}
	// 生成一个新节点
	return &Worker{
		timestamp: 0,
		workerId:  workerId,
		number:    0,
	}, nil
}

func (w *Worker) GetId() int64 {
	w.mu.Lock()
	defer w.mu.Unlock()
	now := time.Now().UnixNano() / 1e6
	if w.timestamp == now {
		w.number++
		if w.number > numberMax {
			for now <= w.timestamp {
				now = time.Now().UnixNano() / 1e6
			}
			w.number = 0
			w.timestamp = now
		}
	} else {
		w.number = 0
		w.timestamp = now
	}
	ID := (now-startTime)<<timeShift | (w.workerId << workerShift) | (w.number)
	return ID
}

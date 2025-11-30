package routine

import (
	"log"
	"time"

	"github.com/panjf2000/ants/v2"
)

// GlobalPool 是一个默认的 ants 池，用于处理通用的异步任务
var GlobalPool *ants.Pool

func init() {
	// 初始化一个默认容量的池，防止无限创建 goroutine
	// WithNonblocking(true) 表示如果池满了，Submit 会返回错误而不是阻塞
	var err error
	GlobalPool, err = ants.NewPool(50000, ants.WithNonblocking(true))
	if err != nil {
		panic(err)
	}
}

// Submit 提交一个任务到全局池
func Submit(task func()) error {
	return GlobalPool.Submit(task)
}

// BatchProcessor 通用批量处理器
// T: 处理的数据类型
type BatchProcessor[T any] struct {
	inputChan  chan T
	batchSize  int
	interval   time.Duration
	handleFunc func([]T) // 实际执行批量处理的函数
}

// NewBatchProcessor 创建并启动一个新的批量处理器
// batchSize: 触发处理的批次大小
// interval: 触发处理的时间间隔
// handleFunc: 真正的处理逻辑（例如批量写入数据库）
func NewBatchProcessor[T any](batchSize int, interval time.Duration, handleFunc func([]T)) *BatchProcessor[T] {
	bp := &BatchProcessor[T]{
		inputChan:  make(chan T, batchSize*10), // 缓冲区给大一点，防止突发流量阻塞
		batchSize:  batchSize,
		interval:   interval,
		handleFunc: handleFunc,
	}
	bp.start()
	return bp
}

// Add 添加一条数据到处理队列 (非阻塞)
func (bp *BatchProcessor[T]) Add(item T) {
	select {
	case bp.inputChan <- item:
	default:
		// 队列满了，为了不阻塞主流程，这里选择丢弃
		// 实际生产中应该记录 metrics
		log.Println("[BatchProcessor] Queue full, item dropped")
	}
}

func (bp *BatchProcessor[T]) start() {
	// 使用 ants 提交后台 worker
	// 注意：这个 worker 是长驻的 loop，占用池里的 1 个坑位
	_ = Submit(func() {
		buffer := make([]T, 0, bp.batchSize)
		ticker := time.NewTicker(bp.interval)
		defer ticker.Stop()

		for {
			select {
			case item := <-bp.inputChan:
				buffer = append(buffer, item)
				if len(buffer) >= bp.batchSize {
					bp.flush(buffer)
					buffer = make([]T, 0, bp.batchSize)
				}
			case <-ticker.C:
				if len(buffer) > 0 {
					bp.flush(buffer)
					buffer = make([]T, 0, bp.batchSize)
				}
			}
		}
	})
}

func (bp *BatchProcessor[T]) flush(data []T) {
	defer func() {
		if r := recover(); r != nil {
			log.Printf("[BatchProcessor] Panic in flush: %v\n", r)
		}
	}()
	// 执行回调
	bp.handleFunc(data)
}

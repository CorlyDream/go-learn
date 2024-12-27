package context

import (
	"context"
	"fmt"
	"testing"
	"time"
)

func Test_cancel(t *testing.T) {
	// 创建一个可取消的context
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel() // 确保函数结束时取消context

	// 启动一个goroutine，监听cancel信号
	go func(ctx context.Context) {
		// 判断是否有设置Deadline，如果没有则为设置了cancel函数
		if deadline, ok := ctx.Deadline(); ok {
			fmt.Println("Deadline set for:", deadline)
		} else {
			fmt.Println("No deadline set.")
		}
		for {
			select {
			case <-ctx.Done():
				fmt.Println("Goroutine被取消")
				return
			default:
				fmt.Println("Goroutine运行中")
				time.Sleep(1 * time.Second)
			}
		}
	}(ctx)

	// 运行3秒后取消Context
	time.Sleep(3 * time.Second)
	cancel()

	// 等待一会儿以便观察输出
	time.Sleep(1 * time.Second)
	fmt.Println("主程序结束")
}

func Test_cacel1(t *testing.T) {
    ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
    defer func() {
        fmt.Println("即将执行cancel")
        cancel()
        fmt.Println("已经执行cancel")
    }() // 走到这里或其他地方调用了cancel。

    // 如果注释掉这个，则输出"timeout not reached"
    go func() {
        time.Sleep(time.Second)
        cancel()
    }()

    select {
    case <-time.After(3 * time.Second):
        fmt.Println("timeout not reached")
    case <-ctx.Done():// 调用了cancel相当于这个chan里读到了值，于是进这里执行
        fmt.Println("timeout reached")
    }
}

func doWork(ctx context.Context) {
	select {
	case <-time.After(5 * time.Second): // 模拟这个任务需要10s执行结束。
		fmt.Println("Done with work")
	case <-ctx.Done():
		fmt.Println("Canceled work:", ctx.Err())
	}
}

func Test_cancelSelect(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	go doWork(ctx)

	// 比如用户手动触发终止任务执行，即调用cancel函数，此时异步执行的work需要终止，
	cancel()
	fmt.Println("After cancel")
	time.Sleep(10 * time.Second)
}

func Test_multiChannelSelect(t *testing.T) {
    ch1 := make(chan string)
    ch2 := make(chan string)
    ch3 := make(chan string)
    
    go func() { ch1 <- "one" }()
    go func() { ch2 <- "two" }()
    go func() { ch3 <- "three" }()
    
    select {
    case msg1 := <-ch1:
        fmt.Println("ch1:", msg1)
    case msg2 := <-ch2:
        fmt.Println("ch2:", msg2)
    case msg3 := <-ch3:
        fmt.Println("ch3:", msg3)
    }
}

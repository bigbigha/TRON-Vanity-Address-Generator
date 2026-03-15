package vanity

import (
	"tron-vanity-go/internal/tron"
)

// WorkerLoop 工作协程的主循环
func WorkerLoop(m *Matcher, out chan<- Result, stop <-chan struct{}) {
	for {
		// 先检查是否需要停止
		select {
		case <-stop:
			return
		default:
		}

		// 生成TRON地址
		privHex, addr, err := tron.GenerateTronAddress()
		if err != nil {
			continue
		}

		// 检查是否匹配
		if m.Match(addr) {
			// 在发送结果前再次检查stop信号
			select {
			case <-stop:
				return
			case out <- Result{
				PrivateHex:    privHex,
				AddressBase58: addr,
			}:
			}
		}
	}
}
package main

import (
	"flag"
	"fmt"
	"os"
	"runtime"
	"sync"
	"time"
	"tron-vanity-go/internal/vanity"
)

func main() {
	// 解析命令行参数
	prefix := flag.String("prefix", "", "prefix for TRON address (Base58, optional)")
	suffix := flag.String("suffix", "", "suffix for TRON address (Base58, optional)")
	workers := flag.Int("workers", runtime.NumCPU(), "number of workers")
	count := flag.Int("count", 1, "number of matches to find")
	flag.Parse()

	// 验证参数
	if *prefix == "" && *suffix == "" {
		fmt.Fprintln(os.Stderr, "must set at least --prefix or --suffix")
		flag.Usage()
		os.Exit(1)
	}

	// 检查前缀/后缀长度并给出警告
	if len(*prefix) > 6 || len(*suffix) > 6 {
		fmt.Printf("Warning: Long prefix/suffix (%d/%d characters) may take significant time to find matches\n", len(*prefix), len(*suffix))
	}

	// 创建匹配器
	matcher := vanity.NewMatcher(*prefix, *suffix)
	
	// 创建通道
	results := make(chan vanity.Result)
	stop := make(chan struct{})

	// 启动工作协程
	var wg sync.WaitGroup
	fmt.Printf("Starting %d workers...\n", *workers)
	fmt.Printf("Searching for addresses with prefix '%s' and suffix '%s'...\n", *prefix, *suffix)
	
	startTime := time.Now()
	
	for i := 0; i < *workers; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			vanity.WorkerLoop(matcher, results, stop)
		}()
	}

	// 在单独的goroutine中等待所有worker完成后关闭results通道
	go func() {
		wg.Wait()
		close(results)
	}()

	// 收集结果
	found := 0
	for r := range results {
		found++
		fmt.Printf("\n=== Found Match #%d ===\n", found)
		fmt.Printf("Address: %s\n", r.AddressBase58)
		fmt.Printf("Private Key: %s\n", r.PrivateHex)
		
		if found >= *count {
			fmt.Printf("\nFound %d matches. Stopping...\n", found)
			close(stop)
			break
		}
	}

	duration := time.Since(startTime)
	fmt.Printf("Total time: %v\n", duration)
	fmt.Println("Done!")
}
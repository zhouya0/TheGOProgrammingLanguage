package main

// 使用time.tick来报出进度 [∆]
// 使用sync.WaitGroup来调控什么时候要关闭channeel [∆]
// 一个问题： 多个goroutine真的会
// 使用带缓存的channel来实现信号量的控制(并发的控制) [∆]

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"sync"
	"time"
)

// sema will limit the concurrency
var sema = make(chan struct{}, 20)

func dirents(dir string) []os.FileInfo {
	sema <- struct{}{}
	defer func() {
		<-sema
	}()
	entries, err := ioutil.ReadDir(dir)
	if err != nil {
		fmt.Fprintf(os.Stderr, "du1: %v\n", err)
		return nil
	}
	return entries
}

func walkDir(dir string, fileSizes chan<- int64, wg *sync.WaitGroup) {
	defer wg.Done()
	for _, entry := range dirents(dir) {
		if entry.IsDir() {
			wg.Add(1)
			subdir := filepath.Join(dir, entry.Name())
			go walkDir(subdir, fileSizes, wg)
		} else {
			fileSizes <- entry.Size()
		}
	}
}

func main() {
	counts := time.Tick(1 * time.Second)
	var wg sync.WaitGroup
	// Determine the initial directories
	flag.Parse()
	roots := flag.Args()
	if len(roots) == 0 {
		roots = []string{"."}
	}
	// Traverse the file tree
	fileSizes := make(chan int64)
	// 这个goroutine的作用并不是为了并发，只是说这个channnel必须要两个goroutine来做。
	go func() {
		for _, root := range roots {
			wg.Add(1)
			walkDir(root, fileSizes, &wg)
		}

		wg.Wait()
		close(fileSizes)

	}()

	// Print the results
	var nfiles, nbytes int64
loop:
	for {
		select {
		case size, ok := <-fileSizes:
			if !ok {
				break loop
			}
			nfiles++
			nbytes += size
		case <-counts:
			printDiskUsage(nfiles, nbytes)
		}
	}
	printDiskUsage(nfiles, nbytes)
}

func printDiskUsage(nfiles, nbytes int64) {
	fmt.Printf("%d files %.1f MB\n", nfiles, float64(nbytes)/1e6)
}

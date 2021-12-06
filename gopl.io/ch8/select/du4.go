//go:build ignore
// +build ignore

/*
 * @Author: damao
 * @Date: 2021-12-06 11:50:51
 */

package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"sync"
)

var done = make(chan int)

func cancled() bool {
	select {
	case <-done:
		return true
	default:
		return false
	}
}

var verbose = flag.Bool("v", false, "show verbose progress messages")

// sema是计数信号量用来限制dirents同时打开太多文件
var sema = make(chan int, 20)

func main() {
	roots := []string{"F:\\"}
	fileSizes := make(chan int64)
	var n sync.WaitGroup
	for _, root := range roots {
		n.Add(1)
		go walkDir(root, &n, fileSizes)
	}
	go func() {
		n.Wait()
		close(fileSizes)
	}()

	go func() {
		os.Stdin.Read(make([]byte, 1))
		close(done)
	}()

	var nfiles, nsizes int64

loop:
	for {
		select {
		case <-done:
			for range fileSizes {
				// do nothing
			}
		case size, ok := <-fileSizes:
			if !ok {
				break loop
			}
			nfiles++
			nsizes += size
		}
	}

	printDiskUsage(nfiles, nsizes)
}
func walkDir(dir string, n *sync.WaitGroup, fileSize chan<- int64) {
	defer n.Done()
	if cancled() {
		return
	}
	for _, entry := range dirents(dir) {
		if entry.IsDir() {
			n.Add(1)
			subDir := filepath.Join(dir, entry.Name())
			go walkDir(subDir, n, fileSize)
		} else {
			fileSize <- entry.Size()
		}
	}
}

func dirents(dir string) []os.FileInfo {
	defer func() { <-sema }()

	select {
	case <-done:
		return nil
	case sema <- 1:
	}

	entries, err := ioutil.ReadDir(dir)
	if err != nil {
		fmt.Fprintf(os.Stdout, "du1: %v\n", err)
		return nil
	}
	return entries
}

func printDiskUsage(nfiles, nbytes int64) {
	fmt.Printf("%d files  %.1f GB\n", nfiles, float64(nbytes)/1e9)
}

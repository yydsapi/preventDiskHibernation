// Copyright (C) 2018 Betalo AB - All Rights Reserved

package main

import (
	"fmt"
	"github.com/go-co-op/gocron"
	"os"
	"strings"
	"time"
)

func main() {
	gocron.Every(1).Minutes().Do(WriteDiskPartition, "your text file position")
	gocron.Start()
}
func WriteDiskPartition(p string) {
	go WriteText(strings.TrimSpace(p))
}
func WriteText(fp string) {
	fmt.Println("start write" + fp)
	defer func() {
		if err := recover(); err != nil {
			fmt.Printf("get sub panic，err:%v\n", err)
		}
	}()
	err := os.WriteFile(fp, []byte("prevent_sleep"), 0644)
	if err != nil {
		fmt.Println(err)
		return
	}
	time.Sleep(800 * time.Millisecond)
	err = os.Remove(fp)
	if err != nil {
		fmt.Println(err)
		return
	}
}
